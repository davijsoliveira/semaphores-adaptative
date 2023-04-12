package executor

import (
	"fmt"
	"semaphores-adaptative/controller/knowledge"
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
		fmt.Println("################### EXECUTOR ##########################################################")
		for _, signal := range p.TrafficSignals {
			for i, trafficSignal := range knowledge.KnowledgeDB.LastSignalConfiguration {
				if signal.Id == trafficSignal.Id {
					knowledge.KnowledgeDB.LastSignalConfiguration[i].TimeGreen = signal.TimeGreen
					knowledge.KnowledgeDB.LastSignalConfiguration[i].TimeYellow = signal.TimeYellow
					knowledge.KnowledgeDB.LastSignalConfiguration[i].TimeRed = signal.TimeRed
				}
			}
		}

		for _, v := range knowledge.KnowledgeDB.LastSignalConfiguration {
			fmt.Println("O semáforo de ID:", v.Id, "tem como último valor para o verde:", v.TimeGreen)
			fmt.Println("O semáforo de ID:", v.Id, "tem como último valor para o amarelo:", v.TimeYellow)
			fmt.Println("O semáforo de ID:", v.Id, "tem como último valor para o vermelho:", v.TimeRed)
		}
		fmt.Println("######################################################################################")
		toTrafficApp <- p.TrafficSignals

	}
}
