package executor

import (
	"semaphores-adaptative/planner"
	"semaphores-adaptative/trafficApp"
)

// tipo executor
type Executor struct{}

// instância um novo executor
func NewExecutor() *Executor {
	return &Executor{}
}

// repassa para a aplicação os semáforos que devem ter seu tempo alterado
func (Executor) Exec(p planner.Plan) []trafficApp.TrafficSignal {
	return p.TrafficSignals
}
