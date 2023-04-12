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

// executa o planejador
func (Planner) Exec(c analyser.ChangeRequest, t *trafficApp.TrafficSignalSystem) Plan {
	// instância um plano com a decisão de mudança
	changePlan := NewPlan()
	/*cp := []trafficApp.TrafficSignal{}
	changePlan := Plan{
		Decision:       c.Decision,
		TrafficSignals: cp,
	}*/

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
						case c.Congestion <= 50:
							s.TimeGreen = 90
							s.TimeYellow = 5
							s.TimeRed = 25
						case c.Congestion <= 70 && c.Congestion > 50:
							s.TimeGreen = 100
							s.TimeYellow = 5
							s.TimeRed = 20
						case c.Congestion > 70:
							s.TimeGreen = 120
							s.TimeYellow = 5
							s.TimeRed = 15
						}
						changePlan.TrafficSignals = append(changePlan.TrafficSignals, s)
					case constants.GoalMediumCongestion:
						switch {
						case c.Congestion <= 50:
							s.TimeGreen = 70
							s.TimeYellow = 15
							s.TimeRed = 50
						case c.Congestion <= 70 && c.Congestion > 50:
							s.TimeGreen = 80
							s.TimeYellow = 10
							s.TimeRed = 40
						case c.Congestion > 70:
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

	for _, signal := range changePlan.TrafficSignals {
		fmt.Println("O semáforo,", signal.Id, "vai ser adaptado com o tempo de verde:", signal.TimeGreen)
	}

	return *changePlan
}
