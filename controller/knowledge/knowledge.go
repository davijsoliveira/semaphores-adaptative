package knowledge

import "semaphores-adaptative/constants"

type Knowledge struct {
	LastSemaphoreSymptom map[int]string
}

var KnowledgeDB = NewKnowledge()

func NewKnowledge() *Knowledge {
	k := make(map[int]string, constants.NumberSemaphores)
	knw := Knowledge{k}
	for i := 0; i < constants.NumberSemaphores; i++ {
		knw.LastSemaphoreSymptom[i] = "low"
	}
	return &knw
}
