package knowledge

import (
	"semaphores-adaptative/constants"
	"semaphores-adaptative/trafficApp"
)

type Knowledge struct {
	LastDecision            string
	LastSignalSymptom       map[int]string
	LastSignalConfiguration []trafficApp.TrafficSignal
}

var KnowledgeDB = NewKnowledge()

func NewKnowledge() *Knowledge {
	k := make(map[int]string, constants.TrafficSignalNumber)

	// inicializa os semáforos com os valores default
	signals := make([]trafficApp.TrafficSignal, constants.TrafficSignalNumber)
	for i := 0; i < constants.TrafficSignalNumber; i++ {
		signals[i] = trafficApp.NewTrafficSignal(i)
	}

	// inicializa o knowledge com valores default
	knw := Knowledge{constants.NoChange, k, signals}
	for i := 0; i < constants.TrafficSignalNumber; i++ {
		knw.LastSignalSymptom[i] = "low"
	}
	return &knw
}
