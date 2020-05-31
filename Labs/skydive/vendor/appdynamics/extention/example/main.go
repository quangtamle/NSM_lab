package main

import (
	appdynamics "appdynamics"
	appd "appdynamics/extention"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
	This exemple show you what could be represented with appdynamics.

	It is divided into 3 go routines (but only 1 Agent!)
    - each of them will have 1 app (central for main/ app1 and app2 for 2 other)
    - A common app will be added

    - each app will have two tiers/nodes (except common app, with only 1 tier/node)
    - central app will have a backend (contacted by go routines)



    These calls :
    - from central1 to central2 (central app)
    - from central2 to backend
    - from central1 to commonTier
    - from go routine tier2 to common
    - from go routine tier1 to backend
    - from go routine tier1 to tier2

    These reports :
    - both tier from go routine will report custom data
    - tier from central2app will report userdata
*/
func main() {

	//some configuration for the appdynamics agent
	//there is one agent per go routine

	agent := &appd.Agent
	agent.Config = &appdynamics.Config{}
	agent.Config.Controller.Host = "ffavelin.saas.appdynamics.com"
	agent.Config.Controller.Port = 443
	agent.Config.Controller.UseSSL = true
	agent.Config.Controller.Account = "ffavelin"
	agent.Config.Controller.AccessKey = "c4y8cmzzktg7"
	agent.Config.InitTimeoutMs = 1000

	//Central Application + Common
	central := appd.AddApp("centrale")
	common := appd.AddApp("common")
	go goroutine("1")
	go goroutine("2")
	//Tiers creation
	central1 := central.AddTier("centrale1", "TIER") //false because it is a tier and not just a backend
	central2 := central.AddTier("cenrale2", "TIER")  //false because it is a tier and not just a backend
	backend := central.AddTier("backend", "HTTP")
	commonTier := common.AddTier("commonTier", "TIER")

	dict := make(map[string]string)
	dict["nbrcall"] = "0"
	i := 0
	for {
		call, err := appd.StartCall(central1, central2)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		call2b, err := appd.StartCall(central2, backend)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		call1c, err := appd.StartCall(central1, commonTier)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		i += 3

		dict["nbrcall"] = strconv.Itoa(i)

		//central2.ReportUserData(dict)

		time.Sleep(10 * time.Millisecond)

		call.StopCall()
		call2b.StopCall()
		call1c.StopCall()

		time.Sleep(50 * time.Millisecond)
	}

}

func goroutine(name string) {

	app := appd.AddApp("app" + name)
	common := appd.AddApp("common")
	central := appd.AddApp("centrale")
	tier1 := app.AddTier("tier"+name+"/1", "TIER")
	tier2 := app.AddTier("tier"+name+"/2", appd.APPD_TIER_APP) // equivalent to "TIER"
	commonTier := common.AddTier("commonTier", "TIER")
	backend := central.AddTier("backend", appd.APPD_TIER_HTTP)

	dict := make(map[string]int, 0)

	dict["App"+name] = 0

	for {
		call1, err := appd.StartCall(tier1, tier2)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		call2b, err := appd.StartCall(tier2, backend)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		call1c, err := appd.StartCall(tier1, commonTier)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		time.Sleep(5 * time.Millisecond)
		dict["App"+name] += 1

		tier2.ReportMetric(dict)

		call1.StopCall()
		call2b.StopCall()
		call1c.StopCall()

		time.Sleep(50 * time.Millisecond)

	}
}
