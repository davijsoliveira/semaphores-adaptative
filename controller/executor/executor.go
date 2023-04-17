package executor

import (
	"semaphores-adaptative/commons"
	"semaphores-adaptative/controller/knowledge"
	"semaphores-adaptative/controller/planner"
)

// tipo executor
type Executor struct{}

// instância um novo executor
func NewExecutor() *Executor {
	return &Executor{}
}

// repassa para a aplicação os semáforos que devem ter seu tempo alterado
func (Executor) Exec(fromPlanner chan planner.Plan, toTrafficApp chan []commons.TrafficSignal) {
	for {
		p := <-fromPlanner
		for _, signal := range p.TrafficSignals {
			for i, trafficSignal := range knowledge.KnowledgeDB.LastSignalConfiguration {
				if signal.Id == trafficSignal.Id {
					knowledge.KnowledgeDB.LastSignalConfiguration[i].TimeGreen = signal.TimeGreen
					knowledge.KnowledgeDB.LastSignalConfiguration[i].TimeYellow = signal.TimeYellow
					knowledge.KnowledgeDB.LastSignalConfiguration[i].TimeRed = signal.TimeRed
				}
			}
		}
		toTrafficApp <- p.TrafficSignals
	}
}
