package crossconnect

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/crossconnect"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/monitor"
	"google.golang.org/grpc"
)

type eventStream struct {
	crossconnect.MonitorCrossConnect_MonitorCrossConnectsClient
}

func (s *eventStream) Recv() (interface{}, error) {
	return s.MonitorCrossConnect_MonitorCrossConnectsClient.Recv()
}

func newEventStream(ctx context.Context, cc *grpc.ClientConn) (monitor.EventStream, error) {
	stream, err := crossconnect.NewMonitorCrossConnectClient(cc).MonitorCrossConnects(ctx, &empty.Empty{})

	return &eventStream{
		MonitorCrossConnect_MonitorCrossConnectsClient: stream,
	}, err
}

// NewMonitorClient creates a new monitor.Client for crossconnect GRPC API
func NewMonitorClient(cc *grpc.ClientConn) (monitor.Client, error) {
	return monitor.NewClient(cc, &eventFactory{}, newEventStream)
}
