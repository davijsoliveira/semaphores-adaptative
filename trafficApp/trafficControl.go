/*
************************************************************************************************************************************************
Author: Davi Oliveira
Description: This code implements a simple app for traffic signal timing control. The time of the signal traffic may change according to the traffic flow.
Date: 06/03/2023
************************************************************************************************************************************************
*/
package trafficApp

import (
	"fmt"
	"semaphores-adaptative/constants"
)

// tipo semáforo
type TrafficSignal struct {
	Id         int
	TimeGreen  int
	TimeYellow int
	TimeRed    int
	TrafficJam int
}

// tipo sistema de semáforos
type TrafficSignalSystem struct {
	TrafficSignals []TrafficSignal
}

// instancia um semáforo
func NewTrafficSignal(id int) TrafficSignal {
	s := TrafficSignal{Id: id, TimeGreen: constants.DefaultGreen, TimeYellow: constants.DefaultYellow, TimeRed: constants.DefaultRed}

	return s
}

/*func NewTrafficSignalSystem(num int) *[]TrafficSignal {
	system := make([]TrafficSignal, num)
	for i := 0; i < num; i++ {
		system[i] = NewTrafficSignal(i)
	}
	return &system
}*/

// instancia um sistema de semáforos
func NewTrafficSignalSystem(num int) *TrafficSignalSystem {
	s := make([]TrafficSignal, num)
	system := TrafficSignalSystem{TrafficSignals: s}
	for i := 0; i < num; i++ {
		system.TrafficSignals[i] = NewTrafficSignal(i)
	}
	return &system
}

// executa o sistema de semáforos
// func (s *TrafficSignalSystem) Exec(changes map[int][]int) {
func (s *TrafficSignalSystem) Exec(changes []TrafficSignal) {
	// itera sobre os semáforos alterados e os pertencentes ao sistema para aplicar as alterações
	for _, signalsChange := range changes {
		for j, signals := range s.TrafficSignals {
			if signalsChange.Id == signals.Id {
				s.TrafficSignals[j].TimeGreen = signalsChange.TimeGreen
				s.TrafficSignals[j].TimeYellow = signalsChange.TimeYellow
				s.TrafficSignals[j].TimeRed = signalsChange.TimeRed
			}

		}
	}
	for i := range s.TrafficSignals {
		fmt.Println("TrafficSignal ID:", s.TrafficSignals[i].Id, "Green:", s.TrafficSignals[i].TimeGreen, "Yellow:", s.TrafficSignals[i].TimeYellow, "Red:", s.TrafficSignals[i].TimeRed)
	}

	/*for {
	for k, v := range changes {
	adaptation := v
	for z := 0; z < len(adaptation); z++ {
		switch {
		case z == 0:
			s.TrafficSignals[k].TimeGreen = adaptation[z]
		case z == 1:
			s.TrafficSignals[k].TimeYellow = adaptation[z]
		case z == 2:
			s.TrafficSignals[k].TimeRed = adaptation[z]
		}
	}
	}

	}*/
}
