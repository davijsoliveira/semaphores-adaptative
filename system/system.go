package main

import (
	//"fmt"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller"
	"semaphores-adaptative/goal"
	"semaphores-adaptative/trafficApp"
	"sync"

	/*"semaphores-adaptative/controller/analyser"
	"semaphores-adaptative/controller/executor"
	"semaphores-adaptative/controller/monitor"
	"semaphores-adaptative/controller/planner"*/
	"semaphores-adaptative/traffic"
	//"semaphores-adaptative/trafficApp"
)

func main() {
	// variável wait group para controlar as go routines
	var wg sync.WaitGroup

	// cria os canais
	appToMonitor := make(chan []trafficApp.TrafficSignal)
	goalToMonitor := make(chan string)
	//monitorToAnalyser := make(chan []monitor.Symptom)
	//analyserToPlanner := make(chan analyser.ChangeRequest)
	//plannerToExecute := make(chan planner.Plan)
	executeToApp := make(chan []trafficApp.TrafficSignal)

	// instancia a app e o componentes do mape-k
	trafFlow := traffic.NewTrafficFlow(constants.TrafficSignalNumber)
	trafSystem := trafficApp.NewTrafficSignalSystem(constants.TrafficSignalNumber)
	gl := goal.NewGoalConfiguration()
	ctl := controller.NewController()
	//mon := monitor.NewMonitor()
	//anl := analyser.NewAnalyser()
	//pln := planner.NewPlanner()
	//exc := executor.NewExecutor()

	// executa os componentes
	wg.Add(8)
	go trafFlow.Exec()
	go trafSystem.Exec(appToMonitor, executeToApp)
	go gl.Exec(goalToMonitor)
	go ctl.Exec(trafFlow, appToMonitor, executeToApp, goalToMonitor)
	//go mon.Exec(appToMonitor, monitorToAnalyser, goalToMonitor, trafFlow)
	//go anl.Exec(monitorToAnalyser, analyserToPlanner)
	//go pln.Exec(analyserToPlanner, plannerToExecute)
	//go exc.Exec(plannerToExecute, executeToApp)
	wg.Wait()
}
