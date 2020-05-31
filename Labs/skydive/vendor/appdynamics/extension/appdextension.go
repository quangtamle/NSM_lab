package appdextension

import (
	appd "appdynamics"
	"errors"
	"fmt"
	"sync"
)

var Agent AppdAgent

type AppdAgent struct {
	Config *appd.Config
	Tiers  map[string]map[string]*Tier
	Calls  []*Call
	mux    sync.Mutex
}

type Tier struct {
	Name    string
	Backend string
	Added   bool
	Datas   map[string]int
	AppName string
}

// Backend possibilities. If different, it will be consider by appd agent as a custom backend (not a tier)
// If a tier is needed, put "TIER"
const (
	APPD_TIER_HTTP       = "HTTP"
	APPD_TIER_DB         = "DB"
	APPD_TIER_CACHE      = "CACHE"
	APPD_TIER_RABBITMQ   = "RABBITMQ"
	APPD_TIER_WEBSERVICE = "WEBSERVICE"
	APPD_TIER_JMS        = "JMS"
	APPD_TIER_APP        = "TIER"
)

type Call struct {
	Src      *Tier
	Dst      *Tier
	BTsrc    appd.BtHandle
	ExitCall appd.ExitcallHandle
	BTdst    appd.BtHandle //not exist if dst.Backend != "TIER"
	Agent    AppdAgent
}

func StartCall(src *Tier, dst *Tier) (*Call, error) {
	Agent.mux.Lock()
	srcName := src.Name
	dstName := dst.Name
	appName := src.AppName
	dstAppName := dst.AppName
	call := Call{
		Src: src,
		Dst: dst,
	}
	if Agent.Tiers[appName][srcName].Backend != "TIER" {
		return nil, errors.New("Could not start call from Backend")
	}
	call.BTsrc = appd.StartBTWithAppContext(Agent.Tiers[appName][srcName].Name, appName+"/"+Agent.Tiers[appName][srcName].Name+"/"+Agent.Tiers[dstAppName][dstName].Name, "")
	call.ExitCall = appd.StartExitcall(call.BTsrc, Agent.Tiers[dstAppName][dstName].Name)

	if Agent.Tiers[dstAppName][dstName].Backend == "TIER" {
		hdr := appd.GetExitcallCorrelationHeader(call.ExitCall)
		call.BTdst = appd.StartBTWithAppContext(Agent.Tiers[dstAppName][dstName].Name, appName+"/"+Agent.Tiers[appName][srcName].Name+"/"+Agent.Tiers[dstAppName][dstName].Name, hdr)
	}
	Agent.Calls = append(Agent.Calls, &call)
	Agent.mux.Unlock()
	return &call, nil
}

func (c *Call) StopCall() {
	index := 0
	find := false

	Agent.mux.Lock()

	for i, call := range c.Agent.Calls {
		if call.Src.Name == c.Src.Name && call.Dst.Name == c.Dst.Name && call.Src.AppName == c.Src.AppName && call.Dst.AppName == c.Dst.AppName {
			index = i
			find = true
			break
		}
	}

	if find == true {
		c.Agent.Calls = append(c.Agent.Calls[:index], c.Agent.Calls[index+1:]...)
	}

	if c.Dst.Backend == "TIER" {
		appd.EndBT(c.BTdst)
	}
	appd.EndExitcall(c.ExitCall)
	appd.EndBT(c.BTsrc)

	Agent.mux.Unlock()

}

func Update() error {

	for _, call := range Agent.Calls {
		call.StopCall()
	}
	appd.TerminateSDK()
	cfg := Agent.Config
	for _, app := range Agent.Tiers {
		for _, tier := range app {
			contextConfig := &appd.ContextConfig{
				AppName:  tier.AppName,
				TierName: tier.Name,
				NodeName: tier.Name,
			}
			if err := appd.AddAppContextToConfig(cfg, tier.Name, contextConfig); err != nil {
				return err
			}
		}
	}
	if err := appd.InitSDK(cfg); err != nil {
		return errors.New("Error re-initializing the Appdynamics SDK\n")
	} else {
		fmt.Printf("Re-initialized Appdynamics SDK successfully\n")
	}

	for _, app := range Agent.Tiers {
		for _, tier := range app {
			if tier.Backend == "TIER" {
				appd.AddBackend(tier.Name, tier.Backend, map[string]string{"id": tier.Name, "app": tier.AppName}, true)
			} else {
				appd.AddBackend(tier.Name, tier.Backend, map[string]string{"id": tier.Name, "app": tier.AppName}, false)
			}
		}
	}
	return nil
}

func AddTier(name string, backend string, app string) *Tier {

	Agent.mux.Lock()
	if _, ok := Agent.Tiers[app]; !ok {
		AddApp(app)
	}

	newtier := &Tier{
		Name:    name,
		Backend: backend,
		Datas:   make(map[string]int, 0),
		AppName: app,
	}

	if tier, ok := Agent.Tiers[app][name]; ok {
		Agent.mux.Unlock()
		return tier
	}

	Agent.Tiers[app][name] = newtier

	if backend == "TIER" {
		if err := Update(); err != nil {
			fmt.Println(err)
			for err != nil {
				err = Update()
				fmt.Println(err)
			}
		}
	} else {
		appd.AddBackend(name, backend, map[string]string{"id": name, "app": app}, false)
	}
	fmt.Printf(app + "/" + name + "\n")
	Agent.mux.Unlock()

	return newtier
}

func AddApp(name string) {
	if len(Agent.Tiers) == 0 {
		Agent.Config.AppName = name
		Agent.Config.TierName = "Agent"
		Agent.Config.NodeName = "Agent"
		Agent.Tiers = make(map[string]map[string]*Tier, 0)
	}
	if _, ok := Agent.Tiers[name]; !ok {
		Agent.Tiers[name] = make(map[string]*Tier, 0)
	}
}

func (t *Tier) AddMetric(path string) {
	appd.AddCustomMetric(t.Name, path, appd.APPD_TIMEROLLUP_TYPE_AVERAGE, appd.APPD_CLUSTERROLLUP_TYPE_INDIVIDUAL,
		appd.APPD_HOLEHANDLING_TYPE_RATE_COUNTER)

}

func (t *Tier) ReportMetric(dict map[string]int) {
	for key := range dict {
		if _, ok := t.Datas[key]; !ok {
			t.AddMetric("Custom Metrics|" + key)
		}
		appd.ReportCustomMetric(t.Name, "Custom Metrics|"+key, int64(dict[key]))
		t.Datas[key] = dict[key]
	}
}

func (t *Tier) ReportUserData(dict map[string]string) {
	BT := appd.StartBTWithAppContext(t.Name, t.Name+":UserData", "")
	appd.SetBTURL(BT, t.Name+":UserData")
	if appd.IsBTSnapshotting(BT) {
		for key := range dict {
			appd.AddUserDataToBT(BT, key, dict[key])
		}
	}
	appd.EndBT(BT)
}
