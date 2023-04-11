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

// tipo requisição de mudança: contêm a decisão e os semáforos que devem sofrer adaptação no seu tempo
type ChangeRequest struct {
	Decision          string
	Congestion        float64
	SemaphoresAffects []int
}

func NewAnalyser() *Analyser {
	return &Analyser{}
}

func (Analyser) Exec(s []monitor.Sympton) ChangeRequest {
	// instancia a requisição de mudança
	change := ChangeRequest{
		Decision:          "NoChange",
		SemaphoresAffects: []int{},
	}
	/*fmt.Println("Semáforo", 0, " tem o seguinte sintoma: ", s[0].CongestionRate)
	fmt.Println("Semáforo", 1, " tem o seguinte sintoma: ", s[1].CongestionRate)
	fmt.Println("Semáforo", 2, " tem o seguinte sintoma: ", s[2].CongestionRate)*/
	//  inicializa as variáveis para contabilizar o número de semáforos com congestionamento médio ou alto
	numLow := 0
	numMedium := 0
	numIntensive := 0

	// inicializa um slice para contabilizar os semáforos
	//idSemaphores := make([]int, constants.TrafficSignalNumber)

	// contabiliza os semáforos com baixo, médio e intenso congestionamento
	for _, sympton := range s {
		switch {
		case sympton.CongestionRate == constants.Low:
			numLow++
		// verifica o sintoma e se o último sintoma estava em um nível mais alto de congestionamento
		case sympton.CongestionRate == constants.Medium && knowledge.KnowledgeDB.LastSemaphoreSymptom[sympton.SemaphoreID] == constants.Low:
			numMedium++
			change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
		case sympton.CongestionRate == constants.Intense && knowledge.KnowledgeDB.LastSemaphoreSymptom[sympton.SemaphoreID] != constants.Intense:
			numIntensive++
			change.SemaphoresAffects = append(change.SemaphoresAffects, sympton.SemaphoreID)
		}
	}

	fmt.Println("O número de low:", numLow)
	fmt.Println("O número de medium:", numMedium)
	fmt.Println("O número de intensive:", numIntensive)
	fmt.Println("Os semáforos afetados foram:", change.SemaphoresAffects)

	// calcula a porcentagem de semáforos com congestionamento médio ou intenso
	totalCongestion := numMedium + numIntensive
	percentCongestion := math.Round((float64(totalCongestion) * float64(100)) / float64(constants.TrafficSignalNumber))

	// atribui a porcentagem de congestionamento para a requisição de mudança
	change.Congestion = percentCongestion

	// verifica se o congestionamento atual está de acordo com a meta e solicita ou não a mudança
	switch constants.Goal {
	case constants.GoalLowCongestion:
		switch {
		case percentCongestion <= 40:
			change.Decision = constants.NoChange
		case percentCongestion > 40:
			change.Decision = constants.Change
		}
	case constants.GoalMediumCongestion:
		switch {
		case percentCongestion <= 60:
			change.Decision = constants.NoChange
		case percentCongestion > 60:
			change.Decision = constants.Change
		}
	case constants.GoalIntensiveCongestion:
		//switch {
		//case percentCongestion <= 80:
		change.Decision = constants.NoChange
		//case percentCongestion > 80:
		//	change.Decision = constants.Change
		//}
	}

	fmt.Println("A decisão foi de:", change.Decision)
	return change
}
