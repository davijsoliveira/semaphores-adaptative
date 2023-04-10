/*
************************************************************************************************************************************************
Author: Davi Oliveira
Description: This code implements a simple app for semaphores time control. The time of the semphores may change according to the traffic flow.
Date: 06/03/2023
************************************************************************************************************************************************
*/
package trafficApp

import (
	"semaphores-adaptative/constants"
)

type TrafficSignal struct {
	Id         int
	TimeGreen  int
	TimeYellow int
	TimeRed    int
	TrafficJam int
}

type TrafficSignalSystem struct {
	Semaphores []TrafficSignal
}

func NewTrafficSignal(id int) TrafficSignal {
	s := TrafficSignal{Id: id, TimeGreen: constants.DefaultGreen, TimeYellow: constants.DefaultYellow, TimeRed: constants.DefaultRed}

	return s
}

/*func NewSemaphoreSystem(num int) *[]TrafficSignal {
	system := make([]TrafficSignal, num)
	for i := 0; i < num; i++ {
		system[i] = NewTrafficSignal(i)
	}
	return &system
}*/

func NewSemaphoreSystem(num int) *TrafficSignalSystem {
	s := make([]TrafficSignal, num)
	system := TrafficSignalSystem{Semaphores: s}
	for i := 0; i < num; i++ {
		system.Semaphores[i] = NewTrafficSignal(i)
	}
	return &system
}

func (s *TrafficSignalSystem) Exec(changes map[int][]int) {
	//for {
	for k, v := range changes {
		adaptation := v
		for z := 0; z < len(adaptation); z++ {
			switch {
			case z == 0:
				s.Semaphores[k].TimeGreen = adaptation[z]
			case z == 1:
				s.Semaphores[k].TimeYellow = adaptation[z]
			case z == 2:
				s.Semaphores[k].TimeRed = adaptation[z]
			}
		}

		//fmt.Println("TrafficSignal ID:", s.Semaphores[i].Id, "Green:", s.Semaphores[i].TimeGreen, "Yellow:", s.Semaphores[i].TimeYellow, "Red:", s.Semaphores[i].TimeRed)
		//= TrafficSignal  {TimeGreen: timeGreen, TimeYellow: timeYellow, TimeRed: timeRed}
	}

	//}
}
