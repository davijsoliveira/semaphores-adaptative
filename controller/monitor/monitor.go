package monitor

import (
	"semaphores-adaptative/commons"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller/knowledge"
)

// tipo monitor
type Monitor struct{}

// tipo sintoma
type Symptom struct {
	SemaphoreID    int
	CurrentRate    int
	CongestionRate string
	TimeGreen      int
	TimeYellow     int
	TimeRed        int
}

// tipo grupo de sintomas
type Symptoms struct {
	SymptomsGroup []Symptom
}

// instancia um sintoma
func NewSymptom() *Symptom {
	return &Symptom{}
}

// instancia um grupo de sintomas
func NewSymptoms(n int) *Symptoms {
	s := make([]Symptom, n)
	symptoms := Symptoms{SymptomsGroup: s}
	return &symptoms
}

// cria um novo monitor -
func NewMonitor() *Monitor {
	return &Monitor{}
}

// executa o monitor
func (Monitor) Exec(fromTrafficApp chan []commons.TrafficSignal, toAnalyser chan []Symptom, fromGoalConfiguration chan string) {
	for {
		// recebe as informações dos sinais enviadas pelo agente e recebidas pelo frontend
		ta := <-fromTrafficApp

		// recebe a meta gerada dinâmicamente
		g := <-fromGoalConfiguration
		commons.Goal = g

		// gera um sintoma para cada sinal
		Symptoms := NewSymptoms(constants.TrafficSignalNumber)
		for i := 0; i < constants.TrafficSignalNumber; i++ {
			switch {
			case ta[i].Congestion <= 10:
				Symptoms.SymptomsGroup[i].CongestionRate = constants.Low
			case ta[i].Congestion <= 20 && ta[i].Congestion > 10:
				Symptoms.SymptomsGroup[i].CongestionRate = constants.Medium
			case ta[i].Congestion > 20:
				Symptoms.SymptomsGroup[i].CongestionRate = constants.Intense
			}
			Symptoms.SymptomsGroup[i].SemaphoreID = i
			Symptoms.SymptomsGroup[i].CurrentRate = ta[i].Congestion
			Symptoms.SymptomsGroup[i].TimeGreen = ta[i].TimeGreen
			Symptoms.SymptomsGroup[i].TimeYellow = ta[i].TimeYellow
			Symptoms.SymptomsGroup[i].TimeRed = ta[i].TimeRed
			knowledge.KnowledgeDB.LastSignalSymptom[i] = Symptoms.SymptomsGroup[i].CongestionRate
		}

		// envia para o analisador o sintoma dos sinais
		toAnalyser <- Symptoms.SymptomsGroup
	}

}
