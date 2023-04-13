package analyser

import (
	"fmt"
	"math"
	"semaphores-adaptative/commons"
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
			case sympton.CongestionRate == constants.Low && knowledge.KnowledgeDB.LastSignalSymptom[sympton.SemaphoreID] != constants.Low:
				numLow++
			case sympton.CongestionRate == constants.Medium:
				numMedium++
				change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
			case sympton.CongestionRate == constants.Intense:
				numIntensive++
				change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
			}
			knowledge.KnowledgeDB.LastSignalSymptom[sympton.SemaphoreID] = sympton.CongestionRate
		}

		// calcula a porcentagem de semáforos com congestionamento médio ou intenso
		totalCongestion = numMedium + (numIntensive * 2)
		percentCongestion = math.Round((float64(totalCongestion) * float64(100)) / float64(constants.TrafficSignalNumber*2))

		// atribui a porcentagem de congestionamento para a requisição de mudança
		change.Congestion = percentCongestion

		plan := constants.GoalLowCongestionP1

		// obtém o plano a ser implementado
		switch commons.Goal {
		case constants.GoalLowCongestion:
			switch {
			case change.Congestion < constants.CongestionBasePercent:
				plan = constants.GoalLowCongestionP1
			case change.Congestion < constants.CongestionMaxPercent && change.Congestion >= constants.CongestionBasePercent:
				plan = constants.GoalLowCongestionP2
			case change.Congestion >= constants.CongestionMaxPercent:
				plan = constants.GoalLowCongestionP3
			}
		case constants.GoalMediumCongestion:
			switch {
			case change.Congestion < constants.CongestionBasePercent:
				plan = constants.GoalMediumCongestionP1
			case change.Congestion < constants.CongestionMaxPercent && change.Congestion >= constants.CongestionBasePercent:
				plan = constants.GoalMediumCongestionP2
			case change.Congestion >= constants.CongestionMaxPercent:
				plan = constants.GoalMediumCongestionP3
			}
		}

		// verifica a quantidade de semáforos cujo plano a ser aplicado é o atual
		equalsConfSignal := []int{}
		for i, lastPlan := range knowledge.KnowledgeDB.LastSignalPlan {
			for _, v := range change.SemaphoresAffects {
				if i == v {
					if lastPlan == plan {
						equalsConfSignal = append(equalsConfSignal, i)
					}
				}
			}
		}

		// verifica se o número de semáforos que tiveram seu sintoma alterado é igual ao número de semáforos com configuração já adaptada
		// para aquele cenário de congestionamento
		// se não existem semáforos para adaptação, a decisão é por checar o nível de congestionamento, pois deve ter havido uma redução no congestionamento
		// o que requer uma redução nos tempos dos semáforos
		if len(change.SemaphoresAffects) == len(equalsConfSignal) && len(change.SemaphoresAffects) != 0 {
			change.Decision = constants.NoChange
		} else {
			// verifica se o congestionamento atual está de acordo com a meta e solicita ou não a mudança
			switch commons.Goal {
			case constants.GoalLowCongestion:
				switch {
				case percentCongestion <= constants.PercentLowCongestion:
					// caso tenha ocorrido uma mudança anteriormente, a decisão é adaptar para retornar
					//a configuração dos semáforos para um tempo mais adequado ao fluxo
					if knowledge.KnowledgeDB.LastDecision == constants.NoChange && numLow == 1 {
						change.Decision = constants.NoChange
					} else {
						change.Decision = constants.Change
						for _, sympton := range s {
							if sympton.CongestionRate == constants.Low {
								change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
							}
						}
					}
				case percentCongestion > constants.PercentLowCongestion:
					change.Decision = constants.Change
				}
			case constants.GoalMediumCongestion:
				switch {
				case percentCongestion <= constants.PercentMediumCongestion:
					change.Decision = constants.NoChange
					//if knowledge.KnowledgeDB.LastDecision == constants.NoChange && numLow == 1 {
					//	change.Decision = constants.NoChange
					//} else {
					//	change.Decision = constants.Change
					//	for _, sympton := range s {
					//		if sympton.CongestionRate == constants.Low {
					//			change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
					//		}
					//	}
					//}
				case percentCongestion > constants.PercentMediumCongestion:
					change.Decision = constants.Change
				}
			case constants.GoalIntensiveCongestion:
				change.Decision = constants.NoChange
			}

		}

		//  atualiza o knowledge com a última decisão de adaptação do analyser
		knowledge.KnowledgeDB.LastDecision = change.Decision

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
