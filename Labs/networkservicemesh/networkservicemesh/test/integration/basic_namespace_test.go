// +build basic

package nsmd_integration_tests

import (
	"testing"

	"github.com/networkservicemesh/networkservicemesh/test/kubetest"
	"k8s.io/api/core/v1"

	. "github.com/onsi/gomega"
)

func Test_createNSMNamespace(t *testing.T) {
	RegisterTestingT(t)

	k8s, err := kubetest.NewK8s(true)
	defer k8s.Cleanup()

	namespaceName := k8s.GetK8sNamespace()
	namespace, err := k8s.GetNamespace(namespaceName)

	Expect(err).To(BeNil())
	Expect(namespace.Status.Phase).To(Equal(v1.NamespaceActive))
}
