package appdextension

import (
	appd "appdynamics"
	"errors"
	"fmt"
	"sync"
	"time"
)

var Agent AppdAgent

type AppdAgent struct {
	Config *appd.Config
	Apps   map[string]*App
	Calls  []*Call
	mux    sync.Mutex
}

type App struct {
	Name  string
	agent *AppdAgent
	Tiers map[string]*Tier
}

type Tier struct {
	Name    string
	Backend string
	Added   bool
	Datas   map[string]int
	App     *App
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
	app := src.App
	call := Call{
		Src: src,
		Dst: dst,
	}

	if src.Backend != "TIER" {
		return nil, errors.New("Could not start call from Backend")
	}

	// call.BTsrc = appd.StartBTWithAppContext(src.Name, app.Name+"/"+src.Name+"/"+dst.Name, "")

	// call.ExitCall = appd.StartExitcall(call.BTsrc, dst.Name)

	if dst.Backend == "TIER" {
		hdr := appd.GetExitcallCorrelationHeader(call.ExitCall)
		call.BTdst = appd.StartBTWithAppContext(dst.Name, app.Name+"/"+src.Name+"/"+dst.Name, hdr)
	}

	Agent.Calls = append(Agent.Calls, &call)
	return &call, nil
}

func (c *Call) StopCall() {
	index := 0
	find := false
	for i, call := range c.Agent.Calls {
		if call.Src.Name == c.Src.Name && call.Dst.Name == c.Dst.Name && call.Src.App.Name == c.Src.App.Name && call.Dst.App.Name == c.Dst.App.Name {
			index = i
			find = true
			break
		}
	}

	if find == true {
		c.Agent.Calls = append(c.Agent.Calls[:index], c.Agent.Calls[index+1:]...)
	}
	// else {
	//     fmt.Println("Call not found")
	// }

	if c.Dst.Backend == "TIER" {
		appd.EndBT(c.BTdst)
	}
	appd.EndExitcall(c.ExitCall)
	appd.EndBT(c.BTsrc)
}

func Update() error {

	for _, call := range Agent.Calls {
		call.StopCall()
	}
	appd.TerminateSDK()
	cfg := Agent.Config
	for _, app := range Agent.Apps {
		for _, tier := range app.Tiers {
			contextConfig := &appd.ContextConfig{
				AppName:  app.Name,
				TierName: tier.Name,
				NodeName: tier.Name,
			}
			if err := appd.AddAppContextToConfig(cfg, tier.Name, contextConfig); err != nil {
				return err
			}
		}
	}
	if err := appd.InitSDK(cfg); err != nil {
		//fmt.Println(err)
		return errors.New("Error re-initializing the Appdynamics SDK\n")
	} else {
		fmt.Printf("Re-initialized Appdynamics SDK successfully\n")
	}

	for _, app := range Agent.Apps {
		for _, tier := range app.Tiers {
			if tier.Backend == "TIER" {
				appd.AddBackend(tier.Name, tier.Backend, map[string]string{"id": tier.Name, "app": tier.App.Name}, true)
			} else {
				appd.AddBackend(tier.Name, tier.Backend, map[string]string{"id": tier.Name, "app": tier.App.Name}, false)
			}
		}
	}
	return nil
}

func (app *App) AddTier(name string, backend string) *Tier {

	Agent.mux.Lock()
	if tier, ok := app.Tiers[name]; ok {
		return tier
	}
	newtier := &Tier{
		Name:    name,
		Backend: backend,
		Datas:   make(map[string]int, 0),
		App:     app,
	}

	app.Tiers[name] = newtier
	if backend == "TIER" {
		if err := Update(); err != nil {
			fmt.Println(err)
			for err != nil {
				err = Update()
				fmt.Println(err)
				time.Sleep(1 * time.Millisecond)
			}
		}
		fmt.Printf(app.Name + "/" + name + "\n")
	} else {
		appd.AddBackend(name, backend, map[string]string{"id": name, "app": app.Name}, false)
	}
	Agent.mux.Unlock()

	return newtier
}

func AddApp(name string) *App {

	if _, ok := Agent.Apps[name]; ok {
		return Agent.Apps[name]
	}
	newApp := &App{
		Name:  name,
		Tiers: make(map[string]*Tier, 0),
	}
	if len(Agent.Apps) == 0 {
		Agent.Config.AppName = name
		Agent.Config.TierName = "Agent"
		Agent.Config.NodeName = "Agent"
		Agent.Apps = make(map[string]*App, 0)
	}
	Agent.Apps[name] = newApp
	return newApp
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
