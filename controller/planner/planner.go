package planner

import (
	"fmt"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller/analyser"
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

func (Planner) Exec(fromAnalyser chan analyser.ChangeRequest, toExecutor chan Plan, t *trafficApp.TrafficSignalSystem) {
	for {
		c := <-fromAnalyser

		// instância um plano com a decisão de mudança
		changePlan := NewPlan()

		// caso seja necessário a mudança, os semáforos afetados são incluídos no plano de mudança
		if c.Decision == constants.Change {
			changePlan.Decision = constants.Change
			for _, affect := range c.SemaphoresAffects {
				for _, s := range t.TrafficSignals {
					if affect == s.Id {
						// prepara o plano de acordo com a meta e a porcentagem de congestionamento
						switch constants.Goal {
						case constants.GoalLowCongestion:
							switch {
							case c.Congestion <= constants.CongestionBasePercent:
								s.TimeGreen = 90
								s.TimeYellow = 5
								s.TimeRed = 25
							case c.Congestion <= constants.CongestionMaxPercent && c.Congestion > constants.CongestionBasePercent:
								s.TimeGreen = 100
								s.TimeYellow = 5
								s.TimeRed = 20
							case c.Congestion > constants.CongestionMaxPercent:
								s.TimeGreen = 120
								s.TimeYellow = 5
								s.TimeRed = 15
							}
							changePlan.TrafficSignals = append(changePlan.TrafficSignals, s)
						case constants.GoalMediumCongestion:
							switch {
							case c.Congestion <= constants.CongestionBasePercent:
								s.TimeGreen = 70
								s.TimeYellow = 15
								s.TimeRed = 50
							case c.Congestion <= constants.CongestionMaxPercent && c.Congestion > constants.CongestionBasePercent:
								s.TimeGreen = 80
								s.TimeYellow = 10
								s.TimeRed = 40
							case c.Congestion > constants.CongestionMaxPercent:
								s.TimeGreen = 90
								s.TimeYellow = 5
								s.TimeRed = 30
							}
							changePlan.TrafficSignals = append(changePlan.TrafficSignals, s)
						}
					}
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
