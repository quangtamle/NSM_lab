package main

import (
	appdynamics "appdynamics"
	appd "appdynamics/extention"
	"fmt"
	"os"
	"time"
)

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
	agent.Config.Logging.BaseDir = "/home/vtramier/go/src/appdynamics/extention/log/"
	agent.Config.Logging.MaxNumFiles = 100
	agent.Config.Logging.MaxFileSizeBytes = 1000000
	agent.Config.Logging.MinimumLevel = appdynamics.APPD_LOG_LEVEL_TRACE

	//Central Application + Common
	central := appd.AddApp("centrale")
	fmt.Println("central1")
	central1 := central.AddTier("centrale1", "TIER") //false because it is a tier and not just a backend
	fmt.Println("central2")
	central2 := central.AddTier("cenrale2", "TIER") //false because it is a tier and not just a backend

	for {
		call, err := appd.StartCall(central1, central2)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		time.Sleep(10 * time.Millisecond)

		call.StopCall()

		time.Sleep(50 * time.Millisecond)
	}

}
