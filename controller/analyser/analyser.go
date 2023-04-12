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
				knowledge.KnowledgeDB.LastSignalSymptom[sympton.SemaphoreID] = sympton.CongestionRate
			case sympton.CongestionRate == constants.Medium:
				numMedium++
				change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
				knowledge.KnowledgeDB.LastSignalSymptom[sympton.SemaphoreID] = sympton.CongestionRate
			case sympton.CongestionRate == constants.Intense:
				numIntensive++
				change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
				knowledge.KnowledgeDB.LastSignalSymptom[sympton.SemaphoreID] = sympton.CongestionRate
			}
		}

		// calcula a porcentagem de semáforos com congestionamento médio ou intenso
		totalCongestion = numMedium + numIntensive
		percentCongestion = math.Round((float64(totalCongestion) * float64(100)) / float64(constants.TrafficSignalNumber))

		// atribui a porcentagem de congestionamento para a requisição de mudança
		change.Congestion = percentCongestion

		// verifica se o congestionamento atual está de acordo com a meta e solicita ou não a mudança
		//TODO se os valores dos sinais forem os mesmos a serem adaptados, não realizar a adaptação
		switch constants.Goal {
		case constants.GoalLowCongestion:
			switch {
			case percentCongestion <= constants.PercentLowCongestion:
				// caso tenha ocorrido uma mudança anteriormente, a decisão é adaptar para retornar
				//a configuração dos semáforos para um tempo mais adequado ao fluxo
				if knowledge.KnowledgeDB.LastDecision == constants.NoChange {
					change.Decision = constants.NoChange
					knowledge.KnowledgeDB.LastDecision = constants.NoChange
				} else {
					change.Decision = constants.Change
					knowledge.KnowledgeDB.LastDecision = constants.Change
					for _, sympton := range s {
						if sympton.CongestionRate == constants.Low {
							change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
						}
					}
				}
			case percentCongestion > constants.PercentLowCongestion:
				change.Decision = constants.Change
				knowledge.KnowledgeDB.LastDecision = constants.Change
			}
		case constants.GoalMediumCongestion:
			switch {
			case percentCongestion <= constants.PercentMediumCongestion:
				if knowledge.KnowledgeDB.LastDecision == constants.NoChange {
					change.Decision = constants.NoChange
					knowledge.KnowledgeDB.LastDecision = constants.NoChange
				} else {
					change.Decision = constants.Change
					knowledge.KnowledgeDB.LastDecision = constants.Change
					for _, sympton := range s {
						if sympton.CongestionRate == constants.Low {
							change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
						}
					}
				}
			case percentCongestion > constants.PercentMediumCongestion:
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
