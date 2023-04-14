package knowledge

import (
	"semaphores-adaptative/constants"
	"semaphores-adaptative/signalControlApp"
)

type Knowledge struct {
	LastDecision            string
	LastSignalSymptom       map[int]string
	LastSignalPlan          map[int]string
	LastSignalConfiguration []signalControlApp.TrafficSignal
}

var KnowledgeDB = NewKnowledge()

func NewKnowledge() *Knowledge {
	k := make(map[int]string, constants.TrafficSignalNumber)
	p := make(map[int]string, constants.TrafficSignalNumber)

	// inicializa os semáforos com os valores default
	signals := make([]signalControlApp.TrafficSignal, constants.TrafficSignalNumber)
	for i := 0; i < constants.TrafficSignalNumber; i++ {
		signals[i] = signalControlApp.NewTrafficSignal(i)
	}

	// inicializa o knowledge com valores default
	knw := Knowledge{constants.NoChange, k, p, signals}
	for i := 0; i < constants.TrafficSignalNumber; i++ {
		knw.LastSignalSymptom[i] = "low"
	}
	return &knw
}
