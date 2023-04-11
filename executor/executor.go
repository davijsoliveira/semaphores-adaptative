package executor

import (
	"fmt"
	"semaphores-adaptative/planner"
	"semaphores-adaptative/trafficApp"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (Executor) Exec(p planner.Plan) []trafficApp.TrafficSignal {
	for _, signal := range p.TrafficSignals {
		fmt.Println("No executor o semáforo,", signal.Id, "vai ser adaptado com o tempo de verde:", signal.TimeGreen)
	}
	return p.TrafficSignals
}
