package monitor

import (
	"fmt"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller/knowledge"
	"semaphores-adaptative/traffic"
	"time"
)

type Monitor struct{}

type Sympton struct {
	SemaphoreID    int
	CurrentRate    int
	CongestionRate string
}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func (Monitor) Exec(flow *traffic.TrafficFlow) []Sympton {

	// monitor interval
	time.Sleep(5 * time.Second)

	// data collect
	trafficFlowRate := flow.Sense()

	Symptoms := make([]Sympton, constants.NumberSemaphores)

	for i := 0; i < constants.NumberSemaphores; i++ {
		//fmt.Println("Trafego do semáforo", i, ": ", trafficFlowRate.TrafficPerSemaphore[i])
		//fmt.Println("O último sintoma do semáforo", i, " é: ", knowledge.KnowledgeDB.LastSemaphoreSymptom[i])
		switch {
		case trafficFlowRate.TrafficPerSemaphore[i] <= 10:
			Symptoms[i].SemaphoreID = i
			Symptoms[i].CurrentRate = trafficFlowRate.TrafficPerSemaphore[i]
			Symptoms[i].CongestionRate = "low"
			knowledge.KnowledgeDB.LastSemaphoreSymptom[i] = "low"
		case trafficFlowRate.TrafficPerSemaphore[i] <= 20 && trafficFlowRate.TrafficPerSemaphore[i] > 10:
			Symptoms[i].SemaphoreID = i
			Symptoms[i].CurrentRate = trafficFlowRate.TrafficPerSemaphore[i]
			Symptoms[i].CongestionRate = "medium"
			knowledge.KnowledgeDB.LastSemaphoreSymptom[i] = "medium"
		case trafficFlowRate.TrafficPerSemaphore[i] <= 30 && trafficFlowRate.TrafficPerSemaphore[i] > 20:
			Symptoms[i].SemaphoreID = i
			Symptoms[i].CurrentRate = trafficFlowRate.TrafficPerSemaphore[i]
			Symptoms[i].CongestionRate = "intensive"
			knowledge.KnowledgeDB.LastSemaphoreSymptom[i] = "intensive"
		}

		//fmt.Println("Semáforo", i, " tem o seguinte sintoma atual: ", knowledge.KnowledgeDB.LastSemaphoreSymptom[i])

	}
	fmt.Println("Semáforo", 0, " tem o seguinte sintoma: ", Symptoms[0].CongestionRate)
	fmt.Println("Semáforo", 1, " tem o seguinte sintoma: ", Symptoms[1].CongestionRate)
	fmt.Println("Semáforo", 2, " tem o seguinte sintoma: ", Symptoms[2].CongestionRate)
	return Symptoms
}
