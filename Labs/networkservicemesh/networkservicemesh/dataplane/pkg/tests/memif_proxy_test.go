package tests

import (
	"net"
	"os"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/networkservicemesh/networkservicemesh/dataplane/pkg/memifproxy"
)

func TestClosingOpeningMemifProxy(t *testing.T) {
	RegisterTestingT(t)
	proxy := memifproxy.NewCustomProxy("source.sock", "target.sock", "unix")
	for i := 0; i < 10; i++ {
		startProxy(proxy)
		time.Sleep(time.Millisecond * 10)
		stopProxy(proxy)
	}
}

func TestTransferBetweenMemifProxies(t *testing.T) {
	RegisterTestingT(t)
	proxy1 := memifproxy.NewCustomProxy("source.sock", "target.sock", "unix")
	proxy2 := memifproxy.NewCustomProxy("target.sock", "source.sock", "unix")
	proxy1.Start()
	proxy2.Start()
	time.Sleep(time.Millisecond * 10)
	connectAndSendMsg("source.sock")
	connectAndSendMsg("target.sock")
	time.Sleep(time.Millisecond * 10)
	stopProxy(proxy1)
	stopProxy(proxy2)

}

func TestStartProxyIfSocketFileIsExist(t *testing.T) {
	RegisterTestingT(t)
	_, err := os.Create("source.sock")
	Expect(err).Should(BeNil())
	proxy := memifproxy.NewCustomProxy("source.sock", "target.sock", "unix")
	startProxy(proxy)
	time.Sleep(time.Millisecond * 10)
	stopProxy(proxy)
}

func connectAndSendMsg(sock string) {
	addr, err := net.ResolveUnixAddr("unix", sock)
	Expect(err).To(BeNil())
	conn, err := net.DialUnix("unix", nil, addr)
	defer conn.Close()
	Expect(err).To(BeNil())
	_, err = conn.Write([]byte("secret"))
	Expect(err).To(BeNil())

}

func startProxy(proxy *memifproxy.Proxy) {
	err := proxy.Start()
	Expect(err).To(BeNil())
}

func stopProxy(proxy *memifproxy.Proxy) {
	err := proxy.Stop()
	Expect(err).To(BeNil())
}
