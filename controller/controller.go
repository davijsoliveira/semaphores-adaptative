package controller

import (
	"semaphores-adaptative/commons"
	"semaphores-adaptative/controller/analyser"
	"semaphores-adaptative/controller/executor"
	"semaphores-adaptative/controller/monitor"
	"semaphores-adaptative/controller/planner"
	//"semaphores-adaptative/signalControl"
	"semaphores-adaptative/traffic"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (Controller) Exec(t *traffic.TrafficFlow, appToController chan []commons.TrafficSignal, controllerToApp chan []commons.TrafficSignal, goalToController chan string) {
	// cria os canais
	monitorToAnalyser := make(chan []monitor.Symptom)
	analyserToPlanner := make(chan analyser.ChangeRequest)
	plannerToExecute := make(chan planner.Plan)

	// instancia os componentes
	mon := monitor.NewMonitor()
	anl := analyser.NewAnalyser()
	pln := planner.NewPlanner()
	exc := executor.NewExecutor()

	// executa os módulos
	go mon.Exec(appToController, monitorToAnalyser, goalToController, t)
	go anl.Exec(monitorToAnalyser, analyserToPlanner)
	go pln.Exec(analyserToPlanner, plannerToExecute)
	go exc.Exec(plannerToExecute, controllerToApp)
}
