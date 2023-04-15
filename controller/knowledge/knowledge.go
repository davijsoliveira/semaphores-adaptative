package knowledge

import (
	"semaphores-adaptative/commons"
	"semaphores-adaptative/constants"
)

type Knowledge struct {
	LastDecision            string
	LastSignalSymptom       map[int]string
	LastSignalPlan          map[int]string
	LastSignalConfiguration []commons.TrafficSignal
}

var KnowledgeDB = NewKnowledge()

func NewKnowledge() *Knowledge {
	k := make(map[int]string, constants.TrafficSignalNumber)
	p := make(map[int]string, constants.TrafficSignalNumber)

	// inicializa os semáforos com os valores default
	signals := make([]commons.TrafficSignal, constants.TrafficSignalNumber)
	for i := 0; i < constants.TrafficSignalNumber; i++ {
		signals[i] = commons.NewTrafficSignal(i)
	}

	// inicializa o knowledge com valores default
	knw := Knowledge{constants.NoChange, k, p, signals}
	for i := 0; i < constants.TrafficSignalNumber; i++ {
		knw.LastSignalSymptom[i] = "low"
	}
	return &knw
}
