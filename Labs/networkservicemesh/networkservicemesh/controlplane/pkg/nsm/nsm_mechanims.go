package nsm

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/nsm/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/nsm/networkservice"
	remote "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/remote/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/model"
)

func (srv *networkServiceManager) updateMechanism(requestID string, connection connection.Connection, request networkservice.Request, dp *model.Dataplane) error {
	// 5.x
	if request.IsRemote() {
		if m, err := srv.selectRemoteMechanism(requestID, request, dp); err == nil {
			connection.SetConnectionMechanism(m.Clone())
		} else {
			return err
		}
	} else {
		for _, m := range request.GetRequestMechanismPreferences() {
			if dpMechanism := findMechanism(dp.LocalMechanisms, m.GetMechanismType()); dpMechanism != nil {
				connection.SetConnectionMechanism(m.Clone())
				break
			}
		}
	}

	if connection.GetConnectionMechanism() == nil {
		return fmt.Errorf("required mechanism are not found... %v ", request.GetRequestMechanismPreferences())
	}

	if connection.GetConnectionMechanism().GetParameters() == nil {
		connection.GetConnectionMechanism().SetParameters(map[string]string{})
	}

	return nil
}

func (srv *networkServiceManager) selectRemoteMechanism(requestID string, request networkservice.Request, dp *model.Dataplane) (connection.Mechanism, error) {
	for _, mechanism := range request.GetRequestMechanismPreferences() {
		dpMechanism := findMechanism(dp.RemoteMechanisms, remote.MechanismType_VXLAN)
		if dpMechanism == nil {
			continue
		}

		// TODO: Add other mechanisms support

		if mechanism.GetMechanismType() == remote.MechanismType_VXLAN {
			parameters := mechanism.GetParameters()
			dpParameters := dpMechanism.GetParameters()

			parameters[remote.VXLANDstIP] = dpParameters[remote.VXLANSrcIP]

			vni := srv.serviceRegistry.VniAllocator().Vni(dpParameters[remote.VXLANSrcIP], parameters[remote.VXLANSrcIP])
			parameters[remote.VXLANVNI] = strconv.FormatUint(uint64(vni), 10)
		}

		logrus.Infof("NSM:(5.1-%v) Remote mechanism selected %v", requestID, mechanism)
		return mechanism, nil
	}

	return nil, fmt.Errorf("failed to select mechanism, no matched mechanisms found")
}

func findMechanism(mechanismPreferences []connection.Mechanism, mechanismType connection.MechanismType) connection.Mechanism {
	for _, m := range mechanismPreferences {
		if m.GetMechanismType() == mechanismType {
			return m
		}
	}
	return nil
}
