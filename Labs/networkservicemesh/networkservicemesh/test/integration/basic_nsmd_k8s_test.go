// +build basic

package nsmd_integration_tests

import (
	"context"
	"fmt"
	. "github.com/onsi/gomega"
	"k8s.io/api/core/v1"
	"net"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/registry"
	nsmd2 "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/nsmd"
	"github.com/networkservicemesh/networkservicemesh/test/kubetest"
	"github.com/networkservicemesh/networkservicemesh/test/kubetest/pods"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNSMDDRegistryNSE(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kubetest.NewK8s(true)
	defer k8s.Cleanup()

	Expect(err).To(BeNil())

	nsmd := k8s.CreatePod(pods.NSMgrPod("nsmgr-1", nil, k8s.GetK8sNamespace()))

	k8s.WaitLogsContains(nsmd, "nsmd", "NSMD: Restore of NSE/Clients Complete...", defaultTimeout)

	fwd, err := k8s.NewPortForwarder(nsmd, 5000)
	Expect(err).To(BeNil())
	defer fwd.Stop()

	// We need to wait unti it is started

	k8s.WaitLogsContains(nsmd, "nsmd-k8s", "nsmd-k8s initialized and waiting for connection", fastTimeout)
	logs, err := k8s.GetLogs(nsmd, "nsmd-k8s")
	logrus.Printf("%s", logs)

	e := fwd.Start()
	if e != nil {
		logrus.Printf("Error on forward: %v retrying", e)
	}

	serviceRegistry := nsmd2.NewServiceRegistryAt(fmt.Sprintf("localhost:%d", fwd.ListenPort))

	discovery, err := serviceRegistry.DiscoveryClient()
	Expect(err).To(BeNil())
	req := &registry.FindNetworkServiceRequest{
		NetworkServiceName: "my_service",
	}
	response, err := discovery.FindNetworkService(context.Background(), req)
	Expect(response).To(BeNil())
	logrus.Print(err.Error())

	Expect(err.Error()).To(Equal("rpc error: code = Unknown desc = no NetworkService with name: my_service"))

	// Lets register few hundred NSEs and check how it works.

	registryClient, err := serviceRegistry.NseRegistryClient()
	Expect(err).To(BeNil())

	nses := []string{}
	nme := "my-network-service"

	t1 := time.Now()
	//nme := fmt.Sprintf("my-network-service-%d-%d", i, t1.Second())

	reg := &registry.NSERegistration{
		NetworkService: &registry.NetworkService{
			Name:    nme,
			Payload: "tcp",
		},
		NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
			NetworkServiceName: nme,
		},
	}
	resp, err := registryClient.RegisterNSE(context.Background(), reg)
	Expect(err).To(BeNil())
	logrus.Printf("Register time: time: %v %v", time.Since(t1), resp)

	t2 := time.Now()
	response, err = discovery.FindNetworkService(context.Background(), &registry.FindNetworkServiceRequest{
		NetworkServiceName: nme,
	})
	Expect(err).To(BeNil())
	logrus.Printf("Find time: time: %v %d", time.Since(t2), len(response.NetworkServiceEndpoints))

	for _, lnse := range response.NetworkServiceEndpoints {
		nses = append(nses, lnse.EndpointName)
	}

	logs, err = k8s.GetLogs(nsmd, "nsmd-k8s")
	logrus.Printf("%s", logs)
	// Remove all added NSEs

	for _, lnse := range nses {
		t2 := time.Now()
		removeNSE := &registry.RemoveNSERequest{
			EndpointName: lnse,
		}
		_, err = registryClient.RemoveNSE(context.Background(), removeNSE)
		response, err2 := discovery.FindNetworkService(context.Background(), &registry.FindNetworkServiceRequest{
			NetworkServiceName: nme,
		})
		logrus.Printf("Find time: time: %v %d", time.Since(t2), len(response.NetworkServiceEndpoints))
		Expect(err).To(BeNil())
		if err != nil {
			logrus.Printf("Bad")
		}
		Expect(err2).To(BeNil())
	}

	logs, err = k8s.GetLogs(nsmd, "nsmd-k8s")
	logrus.Printf("%s", logs)

}

func TestUpdateNSM(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kubetest.NewK8s(true)
	defer k8s.Cleanup()
	Expect(err).To(BeNil())

	nsmd := k8s.CreatePod(pods.NSMgrPod("nsmgr-1", nil, k8s.GetK8sNamespace()))

	fwd, err := k8s.NewPortForwarder(nsmd, 5000)
	Expect(err).To(BeNil())
	defer fwd.Stop()

	// We need to wait until it is started
	k8s.WaitLogsContains(nsmd, "nsmd-k8s", "nsmd-k8s initialized and waiting for connection", fastTimeout)
	// To be sure NSMD is already called for register.
	k8s.WaitLogsContains(nsmd, "nsmd", "Waiting for dataplane available...", defaultTimeout)

	e := fwd.Start()
	if e != nil {
		logrus.Printf("Error on forward: %v retrying", e)
	}

	serviceRegistry := nsmd2.NewServiceRegistryAt(fmt.Sprintf("localhost:%d", fwd.ListenPort))

	discovery, err := serviceRegistry.DiscoveryClient()
	Expect(err).To(BeNil())

	nseRegistryClient, err := serviceRegistry.NseRegistryClient()
	Expect(err).To(BeNil())

	nsmRegistryClient, err := serviceRegistry.NsmRegistryClient()
	Expect(err).To(BeNil())

	networkService := "icmp-responder"
	nsmName := "master"
	url1 := "1.1.1.1:1"
	url2 := "2.2.2.2:2"
	failures := InterceptGomegaFailures(func() {
		response, err := nsmRegistryClient.RegisterNSM(context.Background(), &registry.NetworkServiceManager{
			Name: nsmName,
			Url:  url1,
		})
		logrus.Info(response)
		Expect(err).To(BeNil())

		nseResp, err := nseRegistryClient.RegisterNSE(context.Background(), &registry.NSERegistration{
			NetworkService: &registry.NetworkService{
				Name:    networkService,
				Payload: "tcp",
			},
			NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
				NetworkServiceName: networkService,
			},
		})
		Expect(err).To(BeNil())
		logrus.Info(nseResp)
		Expect(getNsmUrl(discovery)).To(Equal(url1))

		updNsm, err := nsmRegistryClient.RegisterNSM(context.Background(), &registry.NetworkServiceManager{
			Name: nsmName,
			Url:  url2,
		})
		Expect(err).To(BeNil())
		Expect(updNsm.GetUrl()).To(Equal(url2))
		Expect(getNsmUrl(discovery)).To(Equal(url2))
	})

	if len(failures) > 0 {
		logrus.Errorf("Failures: %v", failures)

		nsmdLogs, _ := k8s.GetLogs(nsmd, "nsmd")
		logrus.Errorf("===================== NSMD output since test is failing %v\n=====================", nsmdLogs)

		nsmdk8sLogs, _ := k8s.GetLogs(nsmd, "nsmd-k8s")
		logrus.Errorf("===================== NSMD K8S output since test is failing %v\n=====================", nsmdk8sLogs)

		nsmdpLogs, _ := k8s.GetLogs(nsmd, "nsmdp")
		logrus.Errorf("===================== NSMDP output since test is failing %v\n=====================", nsmdpLogs)

		t.Fail()
	}
}

func TestGetEndpoints(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kubetest.NewK8s(true)
	defer k8s.Cleanup()
	Expect(err).To(BeNil())

	nsmd, err := kubetest.SetupNodes(k8s, 1, defaultTimeout)
	Expect(err).To(BeNil())

	k8s.WaitLogsContains(nsmd[0].Nsmd, "nsmd", "NSMD: Restore of NSE/Clients Complete...", defaultTimeout)

	nseRegistryClient, nsmRegistryClient, fwdClose := kubetest.PrepareRegistryClients(k8s, nsmd[0].Nsmd)
	defer fwdClose()

	url := "1.1.1.1:1"

	responseNsm, err := nsmRegistryClient.RegisterNSM(context.Background(), &registry.NetworkServiceManager{
		Url: url,
	})
	Expect(err).To(BeNil())
	Expect(responseNsm.Url).To(Equal(url))
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	for i := 0; i < len(letters); i++ {
		nsName := "icmp-responder-" + letters[i]
		_, err := nseRegistryClient.RegisterNSE(context.Background(), &registry.NSERegistration{
			NetworkService: &registry.NetworkService{
				Name:    nsName,
				Payload: "IP",
			},
			NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
				NetworkServiceName: nsName,
			},
		})
		Expect(err).To(BeNil())
	}

	nseList, err := nsmRegistryClient.GetEndpoints(context.Background(), &empty.Empty{})
	Expect(err).To(BeNil())
	Expect(len(nseList.NetworkServiceEndpoints)).To(Equal(len(letters)))

	_, err = nseRegistryClient.RegisterNSE(context.Background(), &registry.NSERegistration{
		NetworkService: &registry.NetworkService{
			Payload: "IP",
			Name:    "icmp-responder",
		},
		NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
			EndpointName: nseList.NetworkServiceEndpoints[0].EndpointName,
		},
	})
	Expect(err).NotTo(BeNil())
	Expect(err.Error()).To(ContainSubstring("already exists"))
}

func TestDuplicateEndpoint(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kubetest.NewK8s(true)
	defer k8s.Cleanup()
	Expect(err).To(BeNil())

	nsmd, err := kubetest.SetupNodes(k8s, 1, defaultTimeout)
	Expect(err).To(BeNil())

	// We need to wait unti it is started
	k8s.WaitLogsContains(nsmd[0].Nsmd, "nsmd-k8s", "nsmd-k8s initialized and waiting for connection", defaultTimeout)

	k8s.WaitLogsContains(nsmd[0].Nsmd, "nsmd", "NSMD: Restore of NSE/Clients Complete...", defaultTimeout)

	nseRegistryClient, nsmRegistryClient, fwdClose := kubetest.PrepareRegistryClients(k8s, nsmd[0].Nsmd)
	defer fwdClose()

	url := "1.1.1.1:1"

	responseNsm, err := nsmRegistryClient.RegisterNSM(context.Background(), &registry.NetworkServiceManager{
		Url: url,
	})
	Expect(err).To(BeNil())
	Expect(responseNsm.Url).To(Equal(url))

	nsName := "icmp-responder"
	_, err = nseRegistryClient.RegisterNSE(context.Background(), &registry.NSERegistration{
		NetworkService: &registry.NetworkService{
			Payload: "IP",
			Name:    "icmp-responder",
		},
		NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
			EndpointName: nsName,
		},
	})
	Expect(err).To(BeNil())

	nseList, err := nsmRegistryClient.GetEndpoints(context.Background(), &empty.Empty{})
	Expect(err).To(BeNil())
	Expect(len(nseList.NetworkServiceEndpoints)).To(Equal(1))

	_, err = nseRegistryClient.RegisterNSE(context.Background(), &registry.NSERegistration{
		NetworkService: &registry.NetworkService{
			Payload: "IP",
			Name:    "icmp-responder",
		},
		NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
			EndpointName: nsName,
		},
	})
	Expect(err).NotTo(BeNil())
	Expect(err.Error()).To(ContainSubstring("already exists"))
}

func getNsmUrl(discovery registry.NetworkServiceDiscoveryClient) string {
	logrus.Infof("Finding NSE...")
	response, err := discovery.FindNetworkService(context.Background(), &registry.FindNetworkServiceRequest{
		NetworkServiceName: "icmp-responder",
	})
	Expect(err).To(BeNil())
	Expect(len(response.GetNetworkServiceEndpoints()) >= 1).To(Equal(true))

	endpoint := response.GetNetworkServiceEndpoints()[0]
	logrus.Infof("Endpoint: %v", endpoint)
	return response.NetworkServiceManagers[endpoint.NetworkServiceManagerName].GetUrl()
}

func TestClusterInfo(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kubetest.NewK8s(true)
	defer k8s.Cleanup()
	Expect(err).To(BeNil())

	clientset, err := k8s.GetClientSet()
	Expect(err).To(BeNil())
	cm, err := clientset.CoreV1().ConfigMaps("kube-system").Get("kubeadm-config", metav1.GetOptions{})

	if cm == nil || err != nil {
		t.Skip("Skip, no kubeadm-config")
		return
	}

	k8s.Prepare("nsmgr")
	nsmd, err := kubetest.SetupNodes(k8s, 1, defaultTimeout)
	Expect(err).To(BeNil())

	k8s.WaitLogsContains(nsmd[0].Nsmd, "nsmd", "NSMD: Restore of NSE/Clients Complete...", defaultTimeout)

	fwd, err := k8s.NewPortForwarder(nsmd[0].Nsmd, 5000)
	Expect(err).To(BeNil())
	defer fwd.Stop()

	e := fwd.Start()
	if e != nil {
		logrus.Printf("Error on forward: %v retrying", e)
	}

	serviceRegistry := nsmd2.NewServiceRegistryAt(fmt.Sprintf("localhost:%d", fwd.ListenPort))
	clusterInfo, err := serviceRegistry.ClusterInfoClient()
	Expect(err).To(BeNil())

	config, err := clusterInfo.GetClusterConfiguration(context.Background(), &empty.Empty{})
	Expect(err).To(BeNil())
	Expect(config.PodSubnet).ToNot(BeNil())
	Expect(config.ServiceSubnet).ToNot(BeNil())

	_, _, err = net.ParseCIDR(config.PodSubnet)
	Expect(err).To(BeNil())

	_, _, err = net.ParseCIDR(config.ServiceSubnet)
	Expect(err).To(BeNil())
}

func createSingleNsmgr(k8s *kubetest.K8s, name string) *v1.Pod {
	nsmgr := k8s.CreatePod(pods.NSMgrPod(name, nil, k8s.GetK8sNamespace()))

	// We need to wait until it is started
	k8s.WaitLogsContains(nsmgr, "nsmd-k8s", "nsmd-k8s initialized and waiting for connection", fastTimeout)
	// To be sure NSMD is already called for register.
	k8s.WaitLogsContains(nsmgr, "nsmd", "Waiting for dataplane available...", defaultTimeout)

	return nsmgr
}

func TestLostUpdate(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kubetest.NewK8s(true)
	defer k8s.Cleanup()
	Expect(err).To(BeNil())

	nsmgr1 := createSingleNsmgr(k8s, "nsmgr-1")

	sr1, closeFunc := kubetest.ServiceRegistryAt(k8s, nsmgr1)

	nsmRegistryClient, err := sr1.NsmRegistryClient()
	Expect(err).To(BeNil())

	discovery, err := sr1.DiscoveryClient()
	Expect(err).To(BeNil())

	nseRegistryClient, err := sr1.NseRegistryClient()
	Expect(err).To(BeNil())

	networkService := "icmp-responder"
	nsmName := "master"
	url1 := "1.1.1.1:1"
	response, err := nsmRegistryClient.RegisterNSM(context.Background(), &registry.NetworkServiceManager{
		Name: nsmName,
		Url:  url1,
	})
	logrus.Info(response)
	Expect(err).To(BeNil())

	nseResp, err := nseRegistryClient.RegisterNSE(context.Background(), &registry.NSERegistration{
		NetworkService: &registry.NetworkService{
			Name:    networkService,
			Payload: "tcp",
		},
		NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
			NetworkServiceName: networkService,
		},
	})
	Expect(err).To(BeNil())
	logrus.Info(nseResp)
	Expect(getNsmUrl(discovery)).To(Equal(url1))
	closeFunc()
	k8s.DeletePods(nsmgr1)

	nsmgr2 := createSingleNsmgr(k8s, "nsmgr-2")

	sr2, closeFunc2 := kubetest.ServiceRegistryAt(k8s, nsmgr2)
	defer closeFunc2()

	discovery2, err := sr2.DiscoveryClient()
	Expect(err).To(BeNil())

	nseRegistryClient2, err := sr2.NseRegistryClient()
	Expect(err).To(BeNil())

	nseResp, err = nseRegistryClient2.RegisterNSE(context.Background(), &registry.NSERegistration{
		NetworkService: &registry.NetworkService{
			Name:    networkService,
			Payload: "tcp",
		},
		NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
			NetworkServiceName: networkService,
		},
	})
	Expect(err).To(BeNil())
	logrus.Info(nseResp)
	Expect(getNsmUrl(discovery2)).ToNot(Equal(url1))
}

func TestRegistryConcurrentModification(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kubetest.NewK8s(true)
	defer k8s.Cleanup()
	Expect(err).To(BeNil())

	nsmgr1 := createSingleNsmgr(k8s, "nsmgr-1")

	sr1, closeFunc := kubetest.ServiceRegistryAt(k8s, nsmgr1)
	defer closeFunc()

	nsmRegistryClient, err := sr1.NsmRegistryClient()
	Expect(err).To(BeNil())

	n := 30
	errorCh := make(chan error)

	// concurrent modification of k8s registry
	go func() {
		for i := 0; i < n; i++ {
			url := fmt.Sprintf("1.1.1.%d:1", i)
			_, err := nsmRegistryClient.RegisterNSM(context.Background(), &registry.NetworkServiceManager{
				Url: url,
			})
			if err != nil {
				errorCh <- err
				break
			}
			logrus.Infof("RegisterNSM(%v)", url)
			<-time.After(100 * time.Millisecond)
		}
		close(errorCh)
	}()

	for {
		select {
		case <-time.Tick(10 * time.Millisecond):
			k8s.CleanupCRDs()
			logrus.Info("CleanupCRDs")
		case err := <-errorCh:
			Expect(err).To(BeNil())
			return
		}
	}
}
