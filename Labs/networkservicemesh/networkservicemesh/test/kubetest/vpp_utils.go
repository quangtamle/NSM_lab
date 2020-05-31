package kubetest

import (
	"fmt"
	"github.com/networkservicemesh/networkservicemesh/dataplane/pkg/common"
	"github.com/networkservicemesh/networkservicemesh/test/kubetest/pods"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"net"
	"strings"
	"time"
)

// DeployVppAgentICMP - Setup VPP Agent based ICMP responder NSE
func DeployVppAgentICMP(k8s *K8s, node *v1.Node, name string, timeout time.Duration) *v1.Pod {
	return deployICMP(k8s, node, name, timeout,
		pods.VppTestCommonPod("vppagent-icmp-responder-nse", name, "icmp-responder-nse", node,
			defaultICMPEnv(k8s.UseIPv6()),
		),
	)
}

// DeployVppAgentNSC - Setup Default VPP Based Client
func DeployVppAgentNSC(k8s *K8s, node *v1.Node, name string, timeout time.Duration) *v1.Pod {
	return deployNSC(k8s, node, name, "vppagent-nsc", timeout,
		pods.VppTestCommonPod("vppagent-nsc", name, "vppagent-nsc", node,
			defaultNSCEnv(),
		),
	)
}

// CheckVppAgentNSC - Perform check of VPP based agent operations.
func CheckVppAgentNSC(k8s *K8s, nscPodNode *v1.Pod) *NSCCheckInfo {
	if !k8s.UseIPv6() {
		return checkVppAgentNSCConfig(k8s, nscPodNode, "172.16.1.1")
	}
	return checkVppAgentNSCConfig(k8s, nscPodNode, "100::1")
}

func checkVppAgentNSCConfig(k8s *K8s, nscPodNode *v1.Pod, checkIP string) *NSCCheckInfo {
	info := &NSCCheckInfo{}
	response, errOut, _ := k8s.Exec(nscPodNode, nscPodNode.Spec.Containers[0].Name, "vppctl", "show int addr")
	if strings.Contains(response, checkIP) {
		info.ipResponse = response
		info.errOut = errOut
	}
	Expect(info.ipResponse).ShouldNot(Equal(""))
	Expect(info.errOut).Should(Equal(""))
	logrus.Printf("NSC IP status Ok")
	Expect(true, IsVppAgentNsePinged(k8s, nscPodNode))

	return info
}

// GetVppAgentNSEAddr - GetVppAgentNSEAddr - Return vpp agent NSE address
func GetVppAgentNSEAddr(k8s *K8s, nsc *v1.Pod) (net.IP, error) {
	return getNSEAddr(k8s, nsc, parseVppAgentAddr, "vppctl", "show int addr")
}

func parseVppAgentAddr(ipReponse string) (string, error) {
	spitedResponse := strings.Split(ipReponse, "L3 ")
	if len(spitedResponse) < 2 {
		return "", fmt.Errorf("bad ip response %v", ipReponse)
	}
	return spitedResponse[1], nil
}

// IsVppAgentNsePinged - Check if vpp agent NSE is pinged
func IsVppAgentNsePinged(k8s *K8s, from *v1.Pod) (result bool) {
	nseIP, err := GetVppAgentNSEAddr(k8s, from)
	Expect(err).Should(BeNil())
	logrus.Infof("%v trying vppctl ping to %v", from.Name, nseIP)
	response, _, _ := k8s.Exec(from, from.Spec.Containers[0].Name, "vppctl", "ping", nseIP.String())
	logrus.Infof("ping result: %s", response)
	if strings.TrimSpace(response) != "" && !strings.Contains(response, "100% packet loss") && !strings.Contains(response, "Fail") {
		result = true
		logrus.Info("Ping successful")
	}

	return result
}

// DefaultPlaneVariablesVPP - Default variables for VPP deployment
func DefaultPlaneVariablesVPP() map[string]string {
	return map[string]string{
		common.DataplaneMetricsEnabledKey: "false",
	}
}
