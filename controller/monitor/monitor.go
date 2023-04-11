package monitor

import (
	"fmt"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/traffic"
	"time"
)

// tipo monitor
type Monitor struct{}

// tipo sintoma
type Sympton struct {
	SemaphoreID    int
	CurrentRate    int
	CongestionRate string
}

// cria um novo monitor -
func NewMonitor() *Monitor {
	return &Monitor{}
}

// executa o monitor
func (Monitor) Exec(flow *traffic.TrafficFlow) []Sympton {

	// interval para monitor coletar os dados de congestionamento
	time.Sleep(5 * time.Second)

	// coleta dos dados de congestionamento
	trafficFlowRate := flow.Sense()

	// Gera o sintoma de cada semáforo, verificando os semáforos que tem tráfego baixo/médio/intenso
	Symptoms := make([]Sympton, constants.TrafficSignalNumber)
	for i := 0; i < constants.TrafficSignalNumber; i++ {
		//fmt.Println("Trafego do semáforo", i, ": ", trafficFlowRate.TrafficPerSemaphore[i])
		//fmt.Println("O último sintoma do semáforo", i, " é: ", knowledge.KnowledgeDB.LastSemaphoreSymptom[i])
		switch {
		case trafficFlowRate.TrafficPerSemaphore[i] <= 10:
			Symptoms[i].SemaphoreID = i
			Symptoms[i].CurrentRate = trafficFlowRate.TrafficPerSemaphore[i]
			Symptoms[i].CongestionRate = constants.Low
		case trafficFlowRate.TrafficPerSemaphore[i] <= 20 && trafficFlowRate.TrafficPerSemaphore[i] > 10:
			Symptoms[i].SemaphoreID = i
			Symptoms[i].CurrentRate = trafficFlowRate.TrafficPerSemaphore[i]
			Symptoms[i].CongestionRate = constants.Medium
		case trafficFlowRate.TrafficPerSemaphore[i] > 20:
			Symptoms[i].SemaphoreID = i
			Symptoms[i].CurrentRate = trafficFlowRate.TrafficPerSemaphore[i]
			Symptoms[i].CongestionRate = constants.Intense
		}

		//fmt.Println("Semáforo", i, " tem o seguinte sintoma atual: ", knowledge.KnowledgeDB.LastSemaphoreSymptom[i])

	}
	fmt.Println("Semáforo", 0, " tem o seguinte sintoma: ", Symptoms[0].CongestionRate)
	fmt.Println("Semáforo", 1, " tem o seguinte sintoma: ", Symptoms[1].CongestionRate)
	fmt.Println("Semáforo", 2, " tem o seguinte sintoma: ", Symptoms[2].CongestionRate)
	return Symptoms
}
