package monitor

import (
	"semaphores-adaptative/commons"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller/knowledge"
	"semaphores-adaptative/traffic"
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
func (Monitor) Exec(fromTrafficApp chan []commons.TrafficSignal, toAnalyser chan []Symptom, fromGoalConfiguration chan string, flow *traffic.TrafficFlow) {
	for {
		// intervalo para monitor coletar os dados de congestionamento
		//time.Sleep(10 * time.Second)

		ta := <-fromTrafficApp

		g := <-fromGoalConfiguration

		commons.Goal = g

		// coleta dos dados de congestionamento
		trafficFlowRate := flow.Sense()

		// Gera o sintoma de cada semáforo, verificando os semáforos que tem tráfego baixo/médio/intenso
		Symptoms := NewSymptoms(constants.TrafficSignalNumber)
		for i := 0; i < constants.TrafficSignalNumber; i++ {
			switch {
			case trafficFlowRate.TrafficPerSemaphore[i] <= 10:
				Symptoms.SymptomsGroup[i].CongestionRate = constants.Low
			case trafficFlowRate.TrafficPerSemaphore[i] <= 20 && trafficFlowRate.TrafficPerSemaphore[i] > 10:
				Symptoms.SymptomsGroup[i].CongestionRate = constants.Medium
			case trafficFlowRate.TrafficPerSemaphore[i] > 20:
				Symptoms.SymptomsGroup[i].CongestionRate = constants.Intense
			}
			Symptoms.SymptomsGroup[i].SemaphoreID = i
			Symptoms.SymptomsGroup[i].CurrentRate = trafficFlowRate.TrafficPerSemaphore[i]
			Symptoms.SymptomsGroup[i].TimeGreen = ta[i].TimeGreen
			Symptoms.SymptomsGroup[i].TimeYellow = ta[i].TimeYellow
			Symptoms.SymptomsGroup[i].TimeRed = ta[i].TimeRed
			knowledge.KnowledgeDB.LastSignalSymptom[i] = Symptoms.SymptomsGroup[i].CongestionRate
		}
		toAnalyser <- Symptoms.SymptomsGroup
	}

}
