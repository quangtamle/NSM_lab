package tests

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/registry"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/model"
)

func newModel() model.Model {
	return model.NewModel()
}

func TestModelAddRemove(t *testing.T) {
	RegisterTestingT(t)

	mdl := newModel()

	mdl.AddDataplane(&model.Dataplane{
		RegisteredName: "test_name",
		SocketLocation: "location",
	})

	Expect(mdl.GetDataplane("test_name").RegisteredName).To(Equal("test_name"))

	mdl.DeleteDataplane("test_name")

	Expect(mdl.GetDataplane("test_name")).To(BeNil())
}

func TestModelSelectDataplane(t *testing.T) {
	RegisterTestingT(t)

	mdl := newModel()

	mdl.AddDataplane(&model.Dataplane{
		RegisteredName: "test_name",
		SocketLocation: "location",
	})
	dp, err := mdl.SelectDataplane(nil)
	Expect(dp.RegisteredName).To(Equal("test_name"))
	Expect(err).To(BeNil())
}
func TestModelSelectDataplaneNone(t *testing.T) {
	RegisterTestingT(t)

	mdl := newModel()

	dp, err := mdl.SelectDataplane(nil)
	Expect(dp).To(BeNil())
	Expect(err.Error()).To(Equal("no appropriate dataplanes found"))
}

func TestModelAddEndpoint(t *testing.T) {
	RegisterTestingT(t)

	mdl := newModel()

	ep1 := createNSERegistration("golden-network", "ep1")
	mdl.AddEndpoint(ep1)
	Expect(mdl.GetEndpoint("ep1")).To(Equal(ep1))

	Expect(mdl.GetEndpointsByNetworkService("golden-network")[0]).To(Equal(ep1))
}

func createNSERegistration(networkServiceName string, endpointName string) *model.Endpoint {
	return &model.Endpoint{
		SocketLocation: "none",
		Workspace:      "nsm-1",
		Endpoint: &registry.NSERegistration{
			NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
				NetworkServiceName: networkServiceName,
				EndpointName:       endpointName,
			}, NetworkService: &registry.NetworkService{
				Name: networkServiceName,
			},
		},
	}
}

func TestModelTwoEndpoint(t *testing.T) {
	RegisterTestingT(t)

	model := newModel()

	ep1 := createNSERegistration("golden-network", "ep1")
	ep2 := createNSERegistration("golden-network", "ep2")
	model.AddEndpoint(ep1)
	model.AddEndpoint(ep2)
	Expect(model.GetEndpoint("ep1")).To(Equal(ep1))
	Expect(model.GetEndpoint("ep2")).To(Equal(ep2))

	Expect(len(model.GetEndpointsByNetworkService("golden-network"))).To(Equal(2))
}

func TestModelAddDeleteEndpoint(t *testing.T) {
	RegisterTestingT(t)

	model := newModel()

	ep1 := createNSERegistration("golden-network", "ep1")
	ep2 := createNSERegistration("golden-network", "ep2")
	model.AddEndpoint(ep1)
	model.AddEndpoint(ep2)
	model.DeleteEndpoint("ep1")
	Expect(model.GetEndpoint("ep1")).To(BeNil())
	Expect(model.GetEndpoint("ep2")).To(Equal(ep2))

	Expect(len(model.GetEndpointsByNetworkService("golden-network"))).To(Equal(1))
}

func TestModelRestoreIds(t *testing.T) {
	RegisterTestingT(t)

	mdl := newModel()
	Expect(mdl.ConnectionID()).To(Equal("1"))
	Expect(mdl.ConnectionID()).To(Equal("2"))
	mdl2 := newModel()
	mdl2.CorrectIDGenerator(mdl.ConnectionID())
	Expect(mdl2.ConnectionID()).To(Equal("4"))

}
