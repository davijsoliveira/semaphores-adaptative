/*
***********************************************************************************************************************************************************
Author: Davi Oliveira
Description: This code implements a simple app for traffic signal timing control. The time of the signal traffic may change according to the traffic flow.
Date: 06/03/2023
***********************************************************************************************************************************************************
*/
package signalControlApp

import (
	"fmt"
	"semaphores-adaptative/constants"
)

// tipo semáforo
//type TrafficSignal struct {
//	Id         int
//	TimeGreen  int
//	TimeYellow int
//	TimeRed    int
//}

type TrafficSignal struct {
	Id         int `json:"id"`
	TimeGreen  int `json:"timegreen"`
	TimeYellow int `json:"timeyellow"`
	TimeRed    int `json:"timered"`
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
func (s *TrafficSignalSystem) Exec(toMonitor chan []TrafficSignal, fromExecutor chan []TrafficSignal) {
	for {
		toMonitor <- s.TrafficSignals

		//conn, err := net.Dial("tcp", "localhost:8080")
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer conn.Close()
		//encoder := json.NewEncoder(conn)
		//if err := encoder.Encode(s.TrafficSignals); err != nil {
		//	log.Fatal(err)
		//}
		ts := <-fromExecutor

		// itera sobre os semáforos alterados e os pertencentes ao sistema para aplicar as alterações
		for _, signalsChange := range ts {
			for j, signals := range s.TrafficSignals {
				if signalsChange.Id == signals.Id {
					s.TrafficSignals[j].TimeGreen = signalsChange.TimeGreen
					s.TrafficSignals[j].TimeYellow = signalsChange.TimeYellow
					s.TrafficSignals[j].TimeRed = signalsChange.TimeRed
				}

			}
		}
		fmt.Println("################### APP TRAFFIC SIGNAL CONTROL #######################################")
		for i := range s.TrafficSignals {
			fmt.Println("O semáforo de ID:", s.TrafficSignals[i].Id, "tem agora os seguintes tempos, Verde:", s.TrafficSignals[i].TimeGreen, "Amarelo:", s.TrafficSignals[i].TimeYellow, "Vermelho:", s.TrafficSignals[i].TimeRed)
		}
		fmt.Println("######################################################################################")
	}

}
