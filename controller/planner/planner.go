package planner

import (
	"fmt"
	"semaphores-adaptative/commons"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller/analyser"
	"semaphores-adaptative/controller/knowledge"
	"semaphores-adaptative/trafficApp"
)

// tipo planejador
type Planner struct{}

// tipo plano de mudança
type Plan struct {
	Decision       string
	TrafficSignals []trafficApp.TrafficSignal
}

func NewPlan() *Plan {
	cp := []trafficApp.TrafficSignal{}
	changePlan := Plan{
		TrafficSignals: cp,
	}
	return &changePlan
}

// instância um novo planejador
func NewPlanner() *Planner {
	return &Planner{}
}

func (Planner) Exec(fromAnalyser chan analyser.ChangeRequest, toExecutor chan Plan) {
	for {
		c := <-fromAnalyser

		// instância um plano com a decisão de mudança
		changePlan := NewPlan()

		// caso seja necessário a mudança, os semáforos afetados são incluídos no plano de mudança
		// com uma nova configuração conforme o nível de congestionamento
		if c.Decision == constants.Change {
			changePlan.Decision = constants.Change
			for _, affect := range c.SemaphoresAffects {
				signalsNewConf := trafficApp.TrafficSignal{}
				// prepara o plano conforme a meta e a porcentagem de congestionamento
				switch commons.Goal {
				case constants.GoalLowCongestion:
					switch {
					case c.Congestion < constants.CongestionBasePercent:
						//signalsNewConf.Id = affect
						signalsNewConf.TimeGreen = 90
						signalsNewConf.TimeYellow = 5
						signalsNewConf.TimeRed = 25
						knowledge.KnowledgeDB.LastSignalPlan[affect] = constants.GoalLowCongestionP1
					case c.Congestion < constants.CongestionMaxPercent && c.Congestion >= constants.CongestionBasePercent:
						//signalsNewConf.Id = affect
						signalsNewConf.TimeGreen = 100
						signalsNewConf.TimeYellow = 5
						signalsNewConf.TimeRed = 20
						knowledge.KnowledgeDB.LastSignalPlan[affect] = constants.GoalLowCongestionP2
					case c.Congestion >= constants.CongestionMaxPercent:
						//signalsNewConf.Id = affect
						signalsNewConf.TimeGreen = 120
						signalsNewConf.TimeYellow = 5
						signalsNewConf.TimeRed = 15
						knowledge.KnowledgeDB.LastSignalPlan[affect] = constants.GoalLowCongestionP3
					}
					signalsNewConf.Id = affect
					changePlan.TrafficSignals = append(changePlan.TrafficSignals, signalsNewConf)
				case constants.GoalMediumCongestion:
					switch {
					case c.Congestion < constants.CongestionBasePercent:
						signalsNewConf.TimeGreen = 70
						signalsNewConf.TimeYellow = 15
						signalsNewConf.TimeRed = 50
					case c.Congestion < constants.CongestionMaxPercent && c.Congestion >= constants.CongestionBasePercent:
						signalsNewConf.TimeGreen = 80
						signalsNewConf.TimeYellow = 10
						signalsNewConf.TimeRed = 40
					case c.Congestion >= constants.CongestionMaxPercent:
						signalsNewConf.TimeGreen = 90
						signalsNewConf.TimeYellow = 5
						signalsNewConf.TimeRed = 30
					}
					signalsNewConf.Id = affect
					changePlan.TrafficSignals = append(changePlan.TrafficSignals, signalsNewConf)
				}
			}
		}

		fmt.Println("################### PLANNER ##########################################################")
		for _, signal := range changePlan.TrafficSignals {
			fmt.Println("O semáforo de ID:", signal.Id, "vai ser adaptado com o tempo de verde:", signal.TimeGreen)
			fmt.Println("O semáforo de ID:", signal.Id, "vai ser adaptado com o tempo de amarelo:", signal.TimeYellow)
			fmt.Println("O semáforo de ID:", signal.Id, "vai ser adaptado com o tempo de vermelho:", signal.TimeRed)
		}
		fmt.Println("######################################################################################")
		toExecutor <- *changePlan
	}
}
