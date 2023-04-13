package monitor

import (
	"fmt"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller/knowledge"
	"semaphores-adaptative/traffic"
	"semaphores-adaptative/trafficApp"
	"time"
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
func (Monitor) Exec(fromTrafficApp chan []trafficApp.TrafficSignal, toAnalyser chan []Symptom, flow *traffic.TrafficFlow) {
	for {
		// interval para monitor coletar os dados de congestionamento
		time.Sleep(10 * time.Second)

		// coleta dos dados de congestionamento
		trafficFlowRate := flow.Sense()

		//TODO inserir no symptom o valor do verde a partir do monitoramento da aplicação

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
			Symptoms.SymptomsGroup[i].TimeGreen = knowledge.KnowledgeDB.LastSignalConfiguration[i].TimeGreen
			Symptoms.SymptomsGroup[i].TimeYellow = knowledge.KnowledgeDB.LastSignalConfiguration[i].TimeYellow
			Symptoms.SymptomsGroup[i].TimeRed = knowledge.KnowledgeDB.LastSignalConfiguration[i].TimeRed
		}
		fmt.Println("################### MONITOR #########################################################")
		fmt.Println("Semáforo", 0, " tem o seguinte sintoma: ", Symptoms.SymptomsGroup[0].CongestionRate)
		fmt.Println("Semáforo", 1, " tem o seguinte sintoma: ", Symptoms.SymptomsGroup[1].CongestionRate)
		fmt.Println("Semáforo", 2, " tem o seguinte sintoma: ", Symptoms.SymptomsGroup[2].CongestionRate)
		fmt.Println("#####################################################################################")
		toAnalyser <- Symptoms.SymptomsGroup
	}

}
