package analyser

import (
	"fmt"
	"math"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller/knowledge"
	"semaphores-adaptative/controller/monitor"
)

// tipo analisador
type Analyser struct{}

// tipo requisição de mudança: contêm a decisão e os semáforos que devem sofrer adaptação
type ChangeRequest struct {
	Decision          string
	Congestion        float64
	SemaphoresAffects []int
}

func NewChangeRequest() *ChangeRequest {
	change := ChangeRequest{
		Decision:          "NoChange",
		SemaphoresAffects: []int{},
	}
	return &change
}

// instancia um analisador
func NewAnalyser() *Analyser {
	return &Analyser{}
}

// executa o analisador, recebendo os sintomas e decidindo se é necessário realizar a adaptação
// func (Analyser) Exec(fromMonitor chan []monitor.Symptom) ChangeRequest {
func (Analyser) Exec(fromMonitor chan []monitor.Symptom, toPlanner chan ChangeRequest) {
	for {
		s := <-fromMonitor

		// instancia a requisição de mudança
		change := NewChangeRequest()

		// inicializa os contadores de congestionamento
		numLow := 0
		numMedium := 0
		numIntensive := 0
		totalCongestion := 0
		percentCongestion := 0.0

		// contabiliza os semáforos com baixo, médio e intenso congestionamento
		for _, sympton := range s {
			switch {
			case sympton.CongestionRate == constants.Low:
				numLow++
			// verifica o sintoma e se o último sintoma estava em um nível mais alto de congestionamento
			//case sympton.CongestionRate == constants.Medium && knowledge.KnowledgeDB.LastSemaphoreSymptom[sympton.SemaphoreID] == constants.Low:
			case sympton.CongestionRate == constants.Medium:
				numMedium++
				change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
			//case sympton.CongestionRate == constants.Intense && knowledge.KnowledgeDB.LastSemaphoreSymptom[sympton.SemaphoreID] != constants.Intense:
			case sympton.CongestionRate == constants.Intense:
				numIntensive++
				change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
				// TODO knowledge.KnowledgeDB.LastSemaphoreSymptom[sympton.SemaphoreID] = sympton.CongestionRate
			}
		}

		// calcula a porcentagem de semáforos com congestionamento médio ou intenso
		totalCongestion = numMedium + numIntensive
		percentCongestion = math.Round((float64(totalCongestion) * float64(100)) / float64(constants.TrafficSignalNumber))

		// atribui a porcentagem de congestionamento para a requisição de mudança
		change.Congestion = percentCongestion

		//TODO if knowledge.KnowledgeDB.LastDecision == constants.Change

		// verifica se o congestionamento atual está de acordo com a meta e solicita ou não a mudança
		switch constants.Goal {
		case constants.GoalLowCongestion:
			switch {
			case percentCongestion <= 40:
				change.Decision = constants.NoChange
				knowledge.KnowledgeDB.LastDecision = constants.NoChange
			case percentCongestion > 40:
				change.Decision = constants.Change
				knowledge.KnowledgeDB.LastDecision = constants.Change
			}
		case constants.GoalMediumCongestion:
			switch {
			case percentCongestion <= 60:
				change.Decision = constants.NoChange
				knowledge.KnowledgeDB.LastDecision = constants.NoChange
			case percentCongestion > 60:
				change.Decision = constants.Change
				knowledge.KnowledgeDB.LastDecision = constants.Change
			}
		case constants.GoalIntensiveCongestion:
			change.Decision = constants.NoChange
			knowledge.KnowledgeDB.LastDecision = constants.NoChange
		}
		fmt.Println("################### ANALYSER #########################################################")
		fmt.Println("O número de semáforos com pouco congestionamento:", numLow)
		fmt.Println("O número de semáforos com congestionamento médio:", numMedium)
		fmt.Println("O número de semáforos com congestionamento intenso:", numIntensive)
		fmt.Println("Os semáforos afetados foram:", change.SemaphoresAffects)
		fmt.Println("A decisão foi de:", change.Decision)
		fmt.Println("######################################################################################")
		toPlanner <- *change

	}
}
