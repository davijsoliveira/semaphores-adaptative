package main

import (
	//"fmt"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller"
	"semaphores-adaptative/goal"
	"semaphores-adaptative/signalControlApp"
	"sync"

	/*"semaphores-adaptative/controller/analyser"
	"semaphores-adaptative/controller/executor"
	"semaphores-adaptative/controller/monitor"
	"semaphores-adaptative/controller/planner"*/
	"semaphores-adaptative/traffic"
	//"semaphores-adaptative/signalControlApp"
)

func main() {
	// variável wait group para controlar as go routines
	var wg sync.WaitGroup

	// cria os canais
	appToMonitor := make(chan []signalControlApp.TrafficSignal)
	goalToMonitor := make(chan string)
	executeToApp := make(chan []signalControlApp.TrafficSignal)

	// instancia a app, o controller e a componente de configuração da meta
	trafFlow := traffic.NewTrafficFlow(constants.TrafficSignalNumber)
	trafSystem := signalControlApp.NewTrafficSignalSystem(constants.TrafficSignalNumber)
	gl := goal.NewGoalConfiguration()
	ctl := controller.NewController()

	// executa os componentes
	wg.Add(8)
	go trafFlow.Exec()
	go trafSystem.Exec(appToMonitor, executeToApp)
	go gl.Exec(goalToMonitor)
	go ctl.Exec(trafFlow, appToMonitor, executeToApp, goalToMonitor)
	wg.Wait()
}
