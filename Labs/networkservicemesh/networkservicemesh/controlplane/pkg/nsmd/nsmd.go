package nsmd

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/crossconnect"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/nsm"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/nsmdapi"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/registry"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/remote/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/remote/networkservice"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/model"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/monitor"
	monitor_crossconnect "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/monitor/crossconnect"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/monitor/remote"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/nseregistry"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/remote/network_service_server"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/serviceregistry"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/services"
	"github.com/networkservicemesh/networkservicemesh/pkg/tools"
)

const (
	NsmdDeleteLocalRegistry = "NSMD_LOCAL_REGISTRY_DELETE"
	DataplaneTimeout        = 1 * time.Hour
	NSEAliveTimeout         = 1 * time.Second
)

type NSMServer interface {
	Stop()
	StartDataplaneRegistratorServer() error
	StartAPIServerAt(sock net.Listener)

	XconManager() *services.ClientConnectionManager
	Manager() nsm.NetworkServiceManager

	MonitorManager
	EndpointManager
}

type nsmServer struct {
	sync.Mutex
	workspaces       map[string]*Workspace
	model            model.Model
	serviceRegistry  serviceregistry.ServiceRegistry
	manager          nsm.NetworkServiceManager
	locationProvider serviceregistry.WorkspaceLocationProvider
	localRegistry    *nseregistry.NSERegistry
	registerServer   *grpc.Server
	registerSock     net.Listener
	regServer        *dataplaneRegistrarServer

	xconManager             *services.ClientConnectionManager
	crossConnectMonitor     monitor_crossconnect.MonitorServer
	remoteConnectionMonitor remote.MonitorServer
}

func (nsm *nsmServer) XconManager() *services.ClientConnectionManager {
	return nsm.xconManager
}

func (nsm *nsmServer) Manager() nsm.NetworkServiceManager {
	return nsm.manager
}

func (nsm *nsmServer) LocalConnectionMonitor(workspace string) monitor.Server {
	if ws := nsm.workspaces[workspace]; ws != nil {
		return ws.MonitorConnectionServer()
	}

	return nil
}

func (nsm *nsmServer) CrossConnectMonitor() monitor_crossconnect.MonitorServer {
	return nsm.crossConnectMonitor
}

func (nsm *nsmServer) RemoteConnectionMonitor() monitor.Server {
	return nsm.remoteConnectionMonitor
}

func RequestWorkspace(serviceRegistry serviceregistry.ServiceRegistry, id string) (*nsmdapi.ClientConnectionReply, error) {
	client, con, err := serviceRegistry.NSMDApiClient()
	if err != nil {
		logrus.Fatalf("Failed to connect to NSMD: %+v...", err)
	}
	defer con.Close()

	reply, err := client.RequestClientConnection(context.Background(), &nsmdapi.ClientConnectionRequest{Workspace: id})
	if err != nil {
		return nil, err
	}
	logrus.Infof("nsmd allocated workspace %+v for client operations...", reply)
	return reply, nil
}

func (nsm *nsmServer) RequestClientConnection(context context.Context, request *nsmdapi.ClientConnectionRequest) (*nsmdapi.ClientConnectionReply, error) {
	logrus.Infof("Requested client connection to nsmd : %+v", request)

	workspace, err := NewWorkSpace(nsm, request.Workspace, false)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	logrus.Infof("New workspace created: %+v", workspace)

	err = nsm.localRegistry.AppendClientRequest(workspace.Name())
	if err != nil {
		logrus.Errorf("Failed to store Client information into local registry: %v", err)
		return nil, err
	}
	nsm.Lock()
	nsm.workspaces[workspace.Name()] = workspace
	nsm.Unlock()
	reply := &nsmdapi.ClientConnectionReply{
		Workspace:       workspace.Name(),
		HostBasedir:     workspace.locationProvider.HostBaseDir(),
		ClientBaseDir:   workspace.locationProvider.ClientBaseDir(),
		NsmServerSocket: workspace.locationProvider.NsmServerSocket(),
		NsmClientSocket: workspace.locationProvider.NsmClientSocket(),
	}
	logrus.Infof("returning ClientConnectionReply: %+v", reply)
	return reply, nil
}

func (nsm *nsmServer) DeleteClientConnection(context context.Context, request *nsmdapi.DeleteConnectionRequest) (*nsmdapi.DeleteConnectionReply, error) {
	socket := request.Workspace
	logrus.Infof("Delete connection for workspace %s", socket)

	workspace, ok := nsm.workspaces[socket]
	if !ok {
		err := fmt.Errorf("no connection exists for workspace %s", socket)
		return &nsmdapi.DeleteConnectionReply{}, err
	}

	err := nsm.localRegistry.DeleteClient(workspace.Name())
	if err != nil {
		logrus.Errorf("Failed to delete Client information into local registry: %v", err)
		return nil, err
	}

	workspace.Close()
	nsm.Lock()
	delete(nsm.workspaces, socket)
	nsm.Unlock()

	return &nsmdapi.DeleteConnectionReply{}, nil
}

func (nsm *nsmServer) EnumConnection(context context.Context, request *nsmdapi.EnumConnectionRequest) (*nsmdapi.EnumConnectionReply, error) {
	nsm.Lock()
	defer nsm.Unlock()
	workspaces := []string{}
	for w := range nsm.workspaces {
		if len(w) > 0 {
			workspaces = append(workspaces, w)
		}
	}
	return &nsmdapi.EnumConnectionReply{Workspace: workspaces}, nil
}

func (nsm *nsmServer) restoreClients(registeredEndpoints *registry.NetworkServiceEndpointList) {

	if os.Getenv(NsmdDeleteLocalRegistry) == "true" {
		logrus.Errorf("Delete of local nse/client registry... by ENV VAR: %s", NsmdDeleteLocalRegistry)
		nsm.localRegistry.Delete()
	}
	clients, nses, err := nsm.localRegistry.LoadRegistry()
	if err != nil {
		logrus.Errorf("NSMServer: Error Loading stored NSE registry: %v", err)
		return
	}

	updatedClients := []string{}
	updatedNSEs := map[string]nseregistry.NSEEntry{}
	if len(clients) > 0 {
		logrus.Infof("NSMServer: Creating workspaces for existing clients...")
		nsm.Lock()
		defer nsm.Unlock()
		for _, client := range clients {
			if client == "" {
				continue
			}
			workspace, err := NewWorkSpace(nsm, client, true)
			if err != nil {
				logrus.Errorf("NSMServer: Failed to create workspace %s %v. Ignoring...", client, err)
				continue
			}
			nsm.workspaces[workspace.Name()] = workspace
			updatedClients = append(updatedClients, client)
		}
	}

	existingEndpoints := map[string]*registry.NetworkServiceEndpoint{}
	for _, ep := range registeredEndpoints.NetworkServiceEndpoints {
		existingEndpoints[ep.EndpointName] = ep
	}

	if len(nses) > 0 {
		// Restore NSEs
		client, err := nsm.serviceRegistry.NseRegistryClient()
		if err != nil {
			err = fmt.Errorf("Failed to get RegistryClient: %s", err)
			return
		}

		for endpointId, nse := range nses {
			if ws, ok := nsm.workspaces[nse.Workspace]; ok {
				logrus.Infof("Checking NSE %s is alive at %v...", endpointId, ws.NsmClientSocket())
				if !ws.isConnectionAlive(NSEAliveTimeout) {
					logrus.Errorf("Unable to connect to local nse %v. Skipping", nse.NseReg)
					if err = nsm.deleteEndpointWithClient(endpointId, client); err != nil {
						logrus.Errorf("Remove NSE: NSE %v", err)
					}
					continue
				}

				logrus.Infof("NSE %s is alive at %v...", endpointId, ws.NsmClientSocket())

				if _, ok := existingEndpoints[endpointId]; !ok {
					newReg, err := ws.registryServer.RegisterNSEWithClient(context.Background(), nse.NseReg, client)
					if err != nil {
						endpointName := nse.NseReg.NetworkserviceEndpoint.EndpointName
						logrus.Warnf("Failed to register NSE with name %v: %v", endpointName, err)
						logrus.Infof("Try to register NSE with new name...")
						nse.NseReg.NetworkserviceEndpoint.EndpointName = ""
						newReg, err = ws.registryServer.RegisterNSEWithClient(context.Background(), nse.NseReg, client)
						if err == nil {
							nsm.manager.NotifyRenamedEndpoint(endpointName, newReg.NetworkserviceEndpoint.EndpointName)
						}
					}

					if err != nil {
						logrus.Errorf("Failed to register NSE: %v", err)
					} else {
						updatedNSEs[newReg.NetworkserviceEndpoint.EndpointName] = nseregistry.NSEEntry{
							Workspace: ws.Name(),
							NseReg:    newReg,
						}
					}
				} else {
					nse.NseReg.NetworkServiceManager = nsm.model.GetNsm()
					nse.NseReg.NetworkserviceEndpoint.NetworkServiceManagerName = nse.NseReg.NetworkServiceManager.Name
					nsm.model.AddEndpoint(&model.Endpoint{
						Endpoint:       nse.NseReg,
						Workspace:      nse.Workspace,
						SocketLocation: ws.NsmClientSocket(),
					})
					updatedNSEs[endpointId] = nse
				}
			}
		}
	} else {
		// We do not have NSE's, need to unregister all of them.
		// Restore NSEs
		client, err := nsm.serviceRegistry.NseRegistryClient()
		if err != nil {
			err = fmt.Errorf("Failed to get RegistryClient: %s", err)
			return
		}

		for _, nse := range existingEndpoints {
			if _, err := client.RemoveNSE(context.Background(), &registry.RemoveNSERequest{
				EndpointName: nse.EndpointName,
			}); err != nil {
				logrus.Errorf("Remove NSE: NSE %v", err)
			}
		}
	}
	if len(updatedClients) > 0 || len(updatedNSEs) > 0 {
		if err := nsm.localRegistry.Save(updatedClients, updatedNSEs); err != nil {
			logrus.Errorf("Store updated NSE local registry...")
		}
	}
	logrus.Infof("NSMD: Restore of NSE/Clients Complete...")
}

func (nsm *nsmServer) deleteEndpointWithClient(name string, client registry.NetworkServiceRegistryClient) error {
	if _, err := client.RemoveNSE(context.Background(), &registry.RemoveNSERequest{
		EndpointName: name,
	}); err != nil {
		return err
	}

	nsm.model.DeleteEndpoint(name)

	return nil
}

// DeleteEndpointWithBrokenConnection deletes endpoint if it has no active connections
func (nsm *nsmServer) DeleteEndpointWithBrokenConnection(endpoint *model.Endpoint) error {
	// If endpoint has active client connection, it should be handled by MonitorNetNsInodeServer
	for _, clientConnection := range nsm.model.GetAllClientConnections() {
		if endpoint.EndpointName() == clientConnection.Endpoint.NetworkserviceEndpoint.EndpointName {
			return nil
		}
	}

	client, err := nsm.serviceRegistry.NseRegistryClient()
	if err != nil {
		return err
	}

	return nsm.deleteEndpointWithClient(endpoint.EndpointName(), client)
}

func (nsm *nsmServer) Stop() {
	if nsm.registerServer != nil {
		nsm.registerServer.GracefulStop()
	}
	if nsm.registerSock != nil {
		_ = nsm.registerSock.Close()
	}
	if nsm.regServer != nil {
		nsm.regServer.Stop()
	}
}

func StartNSMServer(model model.Model, manager nsm.NetworkServiceManager, serviceRegistry serviceregistry.ServiceRegistry, apiRegistry serviceregistry.ApiRegistry) (NSMServer, error) {
	var err error
	if err = tools.SocketCleanup(ServerSock); err != nil {
		return nil, err
	}

	tracer := opentracing.GlobalTracer()
	locationProvider := serviceRegistry.NewWorkspaceProvider()

	nsm := &nsmServer{
		workspaces:       make(map[string]*Workspace),
		model:            model,
		serviceRegistry:  serviceRegistry,
		manager:          manager,
		locationProvider: locationProvider,
		localRegistry:    nseregistry.NewNSERegistry(locationProvider.NsmNSERegistryFile()),
	}

	nsm.registerServer = grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)))

	nsmdapi.RegisterNSMDServer(nsm.registerServer, nsm)

	nsm.registerSock, err = apiRegistry.NewNSMServerListener()
	if err != nil {
		logrus.Errorf("failed to start device plugin grpc server %+v", err)
		nsm.Stop()
		return nil, err
	}
	go func() {
		if err := nsm.registerServer.Serve(nsm.registerSock); err != nil {
			logrus.Error("failed to start device plugin grpc server")
		}
	}()
	endpoints, err := setLocalNSM(model, serviceRegistry)
	if err != nil {
		logrus.Errorf("failed to set local NSM %+v", err)
		return nil, err
	}

	// Check if the socket of NSM server is operation
	_, conn, err := serviceRegistry.NSMDApiClient()
	if err != nil {
		nsm.Stop()
		return nil, err
	}
	_ = conn.Close()
	logrus.Infof("NSM gRPC socket: %s is operational", nsm.registerSock.Addr().String())

	// Restore existing clients in case of NSMd restart.
	nsm.restoreClients(endpoints)

	nsm.initMonitorServers()
	return nsm, nil
}

func (nsm *nsmServer) initMonitorServers() {
	nsm.xconManager = services.NewClientConnectionManager(nsm.model, nsm.manager, nsm.serviceRegistry)
	// Start CrossConnect monitor server
	nsm.crossConnectMonitor = monitor_crossconnect.NewMonitorServer()
	// Start Connection monitor server
	nsm.remoteConnectionMonitor = remote.NewMonitorServer(nsm.xconManager)
}

func (nsm *nsmServer) StartDataplaneRegistratorServer() error {
	var err error
	nsm.regServer, err = StartDataplaneRegistrarServer(nsm.model)
	return err
}

func setLocalNSM(model model.Model, serviceRegistry serviceregistry.ServiceRegistry) (*registry.NetworkServiceEndpointList, error) {
	client, err := serviceRegistry.NsmRegistryClient()
	if err != nil {
		err = fmt.Errorf("Failed to get RegistryClient: %s", err)
		return nil, err
	}

	nsm, err := client.RegisterNSM(context.Background(), &registry.NetworkServiceManager{
		Url: serviceRegistry.GetPublicAPI(),
	})
	if err != nil {
		err = fmt.Errorf("Failed to get my own NetworkServiceManager: %s", err)
		return nil, err
	}

	endpoints, err := client.GetEndpoints(context.Background(), &empty.Empty{})
	if err != nil {
		err = fmt.Errorf("Failed to get list of own Endpoints: %s", err)
		return nil, err
	}

	logrus.Infof("Setting local NSM %v", nsm)
	model.SetNsm(nsm)

	return endpoints, nil
}

// StartAPIServerAt starts GRPC API server at sock
func (nsm *nsmServer) StartAPIServerAt(sock net.Listener) {
	tracer := opentracing.GlobalTracer()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)))

	crossconnect.RegisterMonitorCrossConnectServer(grpcServer, nsm.crossConnectMonitor)
	connection.RegisterMonitorConnectionServer(grpcServer, nsm.remoteConnectionMonitor)

	// Register Remote NetworkServiceManager
	remoteServer := network_service_server.NewRemoteNetworkServiceServer(nsm.model, nsm.manager, nsm.serviceRegistry, nsm.remoteConnectionMonitor)
	networkservice.RegisterNetworkServiceServer(grpcServer, remoteServer)

	// TODO: Add more public API services here.

	go func() {
		if err := grpcServer.Serve(sock); err != nil {
			logrus.Errorf("failed to start gRPC NSMD API server %+v", err)
		}
	}()
	logrus.Infof("NSM gRPC API Server: %s is operational", sock.Addr().String())
}
