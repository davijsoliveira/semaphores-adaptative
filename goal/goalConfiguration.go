package goal

import (
	"fmt"
	"semaphores-adaptative/constants"
	"strings"
)

type GoalConfiguration struct{}

func NewGoalConfiguration() *GoalConfiguration {
	return &GoalConfiguration{}
}

func (GoalConfiguration) Exec(toController chan string) {
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
		fmt.Println("")
		fmt.Println("************************ A Meta Atual é:", strings.ToUpper(goal), "****************************")
		fmt.Println("")
		toController <- goal
	}
}
