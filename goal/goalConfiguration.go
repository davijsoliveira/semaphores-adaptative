package goal

import (
	"fmt"
	"semaphores-adaptative/constants"
)

type GoalConfiguration struct{}

func NewGoalConfiguration() *GoalConfiguration {
	return &GoalConfiguration{}
}

func (GoalConfiguration) Exec(toMonitor chan string) {
	iterations := 0
	goal := constants.GoalLowCongestion
	for {
		if iterations < 5 {
			iterations++
		} else {
			iterations = 0
			switch goal {
			case constants.GoalLowCongestion:
				goal = constants.GoalMediumCongestion
			case constants.GoalMediumCongestion:
				goal = constants.GoalIntensiveCongestion
			case constants.GoalIntensiveCongestion:
				goal = constants.GoalLowCongestion
			}
		}
		fmt.Println("************************ A META ATUAL É:", goal)
		toMonitor <- goal
	}
}
