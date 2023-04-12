package executor

import (
	"semaphores-adaptative/controller/planner"
	"semaphores-adaptative/trafficApp"
)

// tipo executor
type Executor struct{}

// instância um novo executor
func NewExecutor() *Executor {
	return &Executor{}
}

// repassa para a aplicação os semáforos que devem ter seu tempo alterado
func (Executor) Exec(fromPlanner chan planner.Plan, toTrafficApp chan []trafficApp.TrafficSignal) {
	for {
		p := <-fromPlanner
		toTrafficApp <- p.TrafficSignals

	}
}
