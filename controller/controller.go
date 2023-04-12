package controller

import (
	"semaphores-adaptative/traffic"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (Controller) Exec(t *traffic.TrafficFlow) {
	//mon := monitor.NewMonitor()
	//anl := analyser.NewAnalyser()
	//pln := planner.NewPlanner()
	//exc := executor.NewExecutor()

	//go mon.Exec(t)
	//go anl.Exec()
	//go pln.Exec()
	//go exc.Exec()
}
