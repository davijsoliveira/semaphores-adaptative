/*
************************************************************************************************************************************************
Author: Davi Oliveira
Description: This code implements a simple app for semaphores time control. The time of the semphores may change according to the traffic flow.
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

// cria um conjunto de semáforos
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
//var TrafficSystem = NewTrafficSignalSystem(constants.TrafficSignalNumber)
// cria um sistema de semáforos
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
	for _, signalsChange := range changes {
		for _, signals := range s.TrafficSignals {
			if signalsChange.Id == signals.Id {
				fmt.Println("Entrou aqui!!!!!!!!!!!")
				signals.TimeGreen = signalsChange.TimeGreen
				signals.TimeYellow = signalsChange.TimeYellow
				signals.TimeRed = signalsChange.TimeRed
			}

		}
	}
	for i := range s.TrafficSignals {
		fmt.Println("TrafficSignal ID:", s.TrafficSignals[i].Id, "Green:", s.TrafficSignals[i].TimeGreen, "Yellow:", s.TrafficSignals[i].TimeYellow, "Red:", s.TrafficSignals[i].TimeRed)
	}

	//for {
	/*for k, v := range changes {
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
	}*/

	//fmt.Println("TrafficSignal ID:", s.TrafficSignals[i].Id, "Green:", s.TrafficSignals[i].TimeGreen, "Yellow:", s.TrafficSignals[i].TimeYellow, "Red:", s.TrafficSignals[i].TimeRed)
	//= TrafficSignal  {TimeGreen: timeGreen, TimeYellow: timeYellow, TimeRed: timeRed}
	//}

	//}
}
