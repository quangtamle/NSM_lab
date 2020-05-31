# Network Service Mesh Dataplane

## Overview

The dataplane in Network Service Mesh is responsible for handling the connectivity between the client and the network service endpoint.
The following describes what needs to be done from a dataplane-provider point-of-view, so that support for other dataplane "drivers" can be developed in future.
It should be considered as a baseline that can be further extended when needed.

## Configuration

* Have the following package imported -

```go
"github.com/networkservicemesh/networkservicemesh/dataplane/pkg/common"
```

* This structure keeps the main dataplane configuration -

```go
type DataplaneConfigBase struct {
    Name                string
    NSMBaseDir          string
    RegistrarSocket     string
    RegistrarSocketType string
    DataplaneSocket     string
    DataplaneSocketType string
}
```

* The dataplane instance should implement the `NSMDataplane` interface expected by `CreateDataplane` having the following methods - `Init()`, `Close()`, `Request()` and `MonitorMechanisms()`.

* The dataplane is responsible for populating the following base configuration fields - `Name`, `DataplaneSocket` and `DataplaneSocketType` in its `Init()` handler. They are mandatory in order to proceed with the dataplane setup.

## Dataplane example

The following is an example using VPP as a dataplane.

* The configuration will look like -

```go
&DataplaneConfigBase{
    Name:                "vppagent"
    NSMBaseDir:          "/var/lib/networkservicemesh/"
    RegistrarSocket:     "/var/lib/networkservicemesh/nsm.dataplane-registrar.io.sock"
    RegistrarSocketType: "unix"
    DataplaneSocket:     "/var/lib/networkservicemesh/nsm-vppagent.dataplane.sock"
    DataplaneSocketType: "unix"
}
```

* This is the `main` -

```go
func main() {
    // Capture signals to cleanup before exiting
    c := make(chan os.Signal, 1)
    signal.Notify(c,
        syscall.SIGHUP,
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGQUIT)

    go common.BeginHealthCheck()

    vppagent := vppagent.CreateVPPAgent()

    registration := common.CreateDataplane(vppagent)

    select {
    case <-c:
        logrus.Info("Closing Dataplane Registration")
        registration.Close()
    }
}
```
