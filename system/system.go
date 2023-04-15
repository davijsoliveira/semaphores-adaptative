package main

import (
	//"fmt"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller"
	srvcontroller "semaphores-adaptative/controller/server"
	"semaphores-adaptative/goal"
	"semaphores-adaptative/signalControlApp"
	"semaphores-adaptative/traffic"
	"sync"
)

func main() {
	// variável wait group para controlar as go routines
	var wg sync.WaitGroup

	// cria os canais
	appToMonitor := make(chan []signalControlApp.TrafficSignal)
	goalToController := make(chan string)
	executeToApp := make(chan []signalControlApp.TrafficSignal)

	// instancia o ambiente, o componente de configuração da meta, o controller e o frontend
	trafFlow := traffic.NewTrafficFlow(constants.TrafficSignalNumber)
	gl := goal.NewGoalConfiguration()
	ctl := controller.NewController()
	srv := srvcontroller.NewControllerSrv()

	// executa os componentes
	wg.Add(8)
	go trafFlow.Exec()
	go gl.Exec(goalToController)
	go ctl.Exec(trafFlow, appToMonitor, executeToApp, goalToController)
	go srv.Run(appToMonitor, executeToApp)
	wg.Wait()
}
