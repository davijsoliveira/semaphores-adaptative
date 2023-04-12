package knowledge

import "semaphores-adaptative/constants"

type Knowledge struct {
	LastDecision         string
	LastSemaphoreSymptom map[int]string
}

var KnowledgeDB = NewKnowledge()

func NewKnowledge() *Knowledge {
	k := make(map[int]string, constants.TrafficSignalNumber)
	knw := Knowledge{constants.NoChange, k}
	for i := 0; i < constants.TrafficSignalNumber; i++ {
		knw.LastSemaphoreSymptom[i] = "low"
	}
	return &knw
}
