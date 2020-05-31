package registryserver

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"

	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/registry"
)

type nsmRegistryService struct {
	nsmName string
	cache   RegistryCache
}

func newNsmRegistryService(nsmName string, cache RegistryCache) *nsmRegistryService {
	return &nsmRegistryService{
		nsmName: nsmName,
		cache:   cache,
	}
}

func (n *nsmRegistryService) RegisterNSM(ctx context.Context, nsm *registry.NetworkServiceManager) (*registry.NetworkServiceManager, error) {
	logrus.Infof("Received RegisterNSM(%v)", nsm)
	nsmCr := mapNsmToCustomResource(nsm)
	nsmCr.SetName(n.nsmName)

	registeredNsm, err := n.cache.CreateOrUpdateNetworkServiceManager(nsmCr)
	if err != nil {
		logrus.Errorf("Failed to create or update nsm: %s", err)
		return nil, err
	}

	nsm = mapNsmFromCustomResource(registeredNsm)
	logrus.Infof("RegisterNSM return: %v", nsm)
	return nsm, nil
}

func (n *nsmRegistryService) GetEndpoints(context.Context, *empty.Empty) (*registry.NetworkServiceEndpointList, error) {
	logrus.Info("Received GetEndpoints")

	endpoints := n.cache.GetEndpointsByNsm(n.nsmName)
	var response []*registry.NetworkServiceEndpoint
	for _, endpoint := range endpoints {
		ns, err := n.cache.GetNetworkService(endpoint.Spec.NetworkServiceName)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		response = append(response, mapNseFromCustomResource(endpoint, ns.Spec.Payload))
	}

	logrus.Infof("GetEndpoints return: %v", response)
	return &registry.NetworkServiceEndpointList{
		NetworkServiceEndpoints: response,
	}, nil
}
