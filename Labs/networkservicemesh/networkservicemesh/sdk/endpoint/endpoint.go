// Copyright 2018, 2019 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package endpoint

import (
	"context"
	"io"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/local/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/local/networkservice"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/registry"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/monitor"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/monitor/local"
	"github.com/networkservicemesh/networkservicemesh/pkg/tools"
	"github.com/networkservicemesh/networkservicemesh/sdk/common"
)

type nsmEndpoint struct {
	*common.NsmConnection
	service        *CompositeEndpoint
	grpcServer     *grpc.Server
	registryClient registry.NetworkServiceRegistryClient
	endpointName   string
	tracerCloser   io.Closer

	monitorServer local.MonitorServer
}

func (nsme *nsmEndpoint) MonitorServer() monitor.Server {
	return nsme.monitorServer
}

func (nsme *nsmEndpoint) setupNSEServerConnection() (net.Listener, error) {
	c := nsme.Configuration
	if err := tools.SocketCleanup(c.NsmClientSocket); err != nil {
		logrus.Errorf("nse: failure to cleanup stale socket %s with error: %v", c.NsmClientSocket, err)
		return nil, err
	}

	logrus.Infof("nse: listening socket %s", c.NsmClientSocket)
	connectionServer, err := net.Listen("unix", c.NsmClientSocket)
	if err != nil {
		logrus.Errorf("nse: failure to listen on a socket %s with error: %v", c.NsmClientSocket, err)
		return nil, err
	}
	return connectionServer, nil
}

func (nsme *nsmEndpoint) serve(listener net.Listener) {
	go func() {
		if err := nsme.grpcServer.Serve(listener); err != nil {
			logrus.Fatalf("nse: failed to start grpc server on socket %v with error: %v ", nsme.Configuration.NsmClientSocket, err)
		}
	}()
}

func (nsme *nsmEndpoint) Start() error {

	var grpcOptions []grpc.ServerOption
	if nsme.Configuration.TracerEnabled {
		tracer, closer := tools.InitJaeger(nsme.Configuration.AdvertiseNseName)
		grpcOptions = append(grpcOptions,
			grpc.UnaryInterceptor(
				otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())))
		grpcOptions = append(grpcOptions,
			grpc.StreamInterceptor(
				otgrpc.OpenTracingStreamServerInterceptor(tracer)))
		opentracing.SetGlobalTracer(tracer)
		nsme.tracerCloser = closer
	}

	nsme.grpcServer = grpc.NewServer(grpcOptions...)
	networkservice.RegisterNetworkServiceServer(nsme.grpcServer, nsme)
	connection.RegisterMonitorConnectionServer(nsme.grpcServer, nsme.monitorServer)

	listener, err := nsme.setupNSEServerConnection()

	if err != nil {
		logrus.Errorf("Unable to setup NSE")
		return err
	}

	// spawn the listnening thread
	nsme.serve(listener)

	// Registering NSE API, it will listen for Connection requests from NSM and return information
	// needed for NSE's dataplane programming.
	nse := &registry.NetworkServiceEndpoint{
		NetworkServiceName: nsme.Configuration.AdvertiseNseName,
		Payload:            "IP",
		Labels:             tools.ParseKVStringToMap(nsme.Configuration.AdvertiseNseLabels, ",", "="),
	}
	registration := &registry.NSERegistration{
		NetworkService: &registry.NetworkService{
			Name:    nsme.Configuration.AdvertiseNseName,
			Payload: "IP",
		},
		NetworkserviceEndpoint: nse,
	}

	nsme.registryClient = registry.NewNetworkServiceRegistryClient(nsme.GrpcClient)
	registeredNSE, err := nsme.registryClient.RegisterNSE(nsme.Context, registration)
	if err != nil {
		logrus.Fatalln("unable to register endpoint", err)
	}
	nsme.endpointName = registeredNSE.GetNetworkserviceEndpoint().GetEndpointName()
	logrus.Infof("NSE registered: %v", registeredNSE)
	logrus.Infof("NSE: channel has been successfully advertised, waiting for connection from NSM...")

	return nil
}

func (nsme *nsmEndpoint) Delete() error {
	if nsme.Configuration.TracerEnabled {
		_ = nsme.tracerCloser.Close()
	}
	// prepare and defer removing of the advertised endpoint
	removeNSE := &registry.RemoveNSERequest{
		EndpointName: nsme.endpointName,
	}
	_, err := nsme.registryClient.RemoveNSE(context.Background(), removeNSE)
	if err != nil {
		logrus.Errorf("Failed removing NSE: %v, with %v", removeNSE, err)
	}
	nsme.grpcServer.Stop()

	return err
}

func (nsme *nsmEndpoint) Request(ctx context.Context, request *networkservice.NetworkServiceRequest) (*connection.Connection, error) {
	logrus.Infof("Request for Network Service received %v", request)

	incomingConnection, err := nsme.service.Request(ctx, request)
	if err != nil {
		logrus.Errorf("The composite returned an error: %v", err)
		return nil, err
	}

	nsme.monitorServer.Update(incomingConnection)

	logrus.Infof("Responding to NetworkService.Request(%v): %v", request, incomingConnection)
	return incomingConnection, nil
}

func (nsme *nsmEndpoint) Close(ctx context.Context, incomingConnection *connection.Connection) (*empty.Empty, error) {
	_, _ = nsme.service.Close(ctx, incomingConnection)
	_, _ = nsme.NsClient.Close(ctx, incomingConnection)

	nsme.monitorServer.Delete(incomingConnection)

	return &empty.Empty{}, nil
}

// NewNSMEndpoint creates a new NSM endpoint
func NewNSMEndpoint(ctx context.Context, configuration *common.NSConfiguration, service *CompositeEndpoint) (*nsmEndpoint, error) {
	if configuration == nil {
		configuration = &common.NSConfiguration{}
	}
	configuration.CompleteNSConfiguration()

	if service == nil {
		service = &CompositeEndpoint{}
	}

	nsmConnection, err := common.NewNSMConnection(ctx, configuration)
	if err != nil {
		logrus.Errorf("Error: %v", err)
		return nil, err
	}

	endpoint := &nsmEndpoint{
		NsmConnection: nsmConnection,
		service:       service,

		monitorServer: local.NewMonitorServer(),
	}

	return endpoint, nil
}
