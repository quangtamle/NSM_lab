package remote

import (
	"context"

	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/remote/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/monitor"
	"google.golang.org/grpc"
)

type eventStream struct {
	connection.MonitorConnection_MonitorConnectionsClient
}

func (s *eventStream) Recv() (interface{}, error) {
	return s.MonitorConnection_MonitorConnectionsClient.Recv()
}

func newEventStream(ctx context.Context, cc *grpc.ClientConn, selector *connection.MonitorScopeSelector) (monitor.EventStream, error) {
	stream, err := connection.NewMonitorConnectionClient(cc).MonitorConnections(ctx, selector)

	return &eventStream{
		MonitorConnection_MonitorConnectionsClient: stream,
	}, err
}

// NewMonitorClient creates a new monitor.Client for remote/connection GRPC API
func NewMonitorClient(cc *grpc.ClientConn, selector *connection.MonitorScopeSelector) (monitor.Client, error) {
	streamConstructor := func(ctx context.Context, grpcCC *grpc.ClientConn) (stream monitor.EventStream, e error) {
		return newEventStream(ctx, grpcCC, selector)
	}

	return monitor.NewClient(cc, &eventFactory{}, streamConstructor)
}
