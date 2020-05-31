// +build basic

package nsmd_integration_tests

import (
	"testing"

	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/nsmd"
	"github.com/networkservicemesh/networkservicemesh/test/kubetest"
	"github.com/networkservicemesh/networkservicemesh/test/kubetest/pods"
	. "github.com/onsi/gomega"
)

func TestExcludePrefixCheck(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kubetest.NewK8s(true)
	defer k8s.Cleanup()
	Expect(err).To(BeNil())

	nodesCount := 1

	variables := map[string]string{
		nsmd.ExcludedPrefixesEnv:     "172.16.1.0/24",
		nsmd.NsmdDeleteLocalRegistry: "true",
	}

	if k8s.UseIPv6() {
		variables = map[string]string{
			nsmd.ExcludedPrefixesEnv:     "100::/64",
			nsmd.NsmdDeleteLocalRegistry: "true",
		}
	}

	nodes, err := kubetest.SetupNodesConfig(k8s, nodesCount, defaultTimeout, []*pods.NSMgrPodConfig{
		{
			Variables:          variables,
			DataplaneVariables: kubetest.DefaultDataplaneVariables(k8s.GetForwardingPlane()),
			Namespace:          k8s.GetK8sNamespace(),
		},
	}, k8s.GetK8sNamespace())
	Expect(err).To(BeNil())

	defer kubetest.FailLogger(k8s, nodes, t)

	icmp := kubetest.DeployICMP(k8s, nodes[0].Node, "icmp-responder-nse-1", defaultTimeout)

	clientset, err := k8s.GetClientSet()
	Expect(err).To(BeNil())

	nsc, err := clientset.CoreV1().Pods(k8s.GetK8sNamespace()).Create(pods.NSCPod("nsc", nodes[0].Node,
		map[string]string{
			"OUTGOING_NSC_LABELS": "app=icmp",
			"OUTGOING_NSC_NAME":   "icmp-responder",
		},
	))

	defer k8s.DeletePods(nsc)

	Expect(err).To(BeNil())
	k8s.WaitLogsContains(icmp, "", "IPAM: The available address pool is empty, probably intersected by excludedPrefix", defaultTimeout)

}
