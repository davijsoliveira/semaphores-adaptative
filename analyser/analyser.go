package analyser

import (
	"fmt"
	"math"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller/monitor"
)

type Analyser struct{}

type ChangeRequest struct {
	decision         string
	semaphoresAfects []int
}

func NewAnalyser() *Analyser {
	return &Analyser{}
}

func (Analyser) Exec(s []monitor.Sympton) {
	// instancia a requisição de mudança
	change := ChangeRequest{
		decision:         "NoChange",
		semaphoresAfects: make([]int, constants.NumberSemaphores),
	}
	/*fmt.Println("Semáforo", 0, " tem o seguinte sintoma: ", s[0].CongestionRate)
	fmt.Println("Semáforo", 1, " tem o seguinte sintoma: ", s[1].CongestionRate)
	fmt.Println("Semáforo", 2, " tem o seguinte sintoma: ", s[2].CongestionRate)*/
	//  inicializa as variáveis para contabilizar o número de semáforos com congestionamento médio ou alto
	numLow := 0
	numMedium := 0
	numIntensive := 0

	// inicializa um slice para contabilizar os semáforos
	//idSemaphores := make([]int, constants.NumberSemaphores)

	// contabiliza os semáforos com baixo, médio e intenso congestionamento
	for i, sympton := range s {
		switch {
		case sympton.CongestionRate == "low":
			numLow++
		case sympton.CongestionRate == "medium":
			numMedium++
			change.semaphoresAfects[i] = sympton.SemaphoreID
		case sympton.CongestionRate == "intensive":
			numIntensive++
			change.semaphoresAfects[i] = sympton.SemaphoreID
		}
	}

	fmt.Println("O número de low:", numLow)
	fmt.Println("O número de medium:", numMedium)
	fmt.Println("O número de intensive:", numIntensive)

	// calcula a porcentagem de semáforos com congestionamento médio ou intenso
	totalCongestion := numMedium + numIntensive
	percentCongestion := math.Round((float64(totalCongestion) * float64(100)) / float64(constants.NumberSemaphores))

	switch constants.Goal {
	case "LowCongestion":
		switch {
		case percentCongestion <= 40:
			change.decision = "NoChange"
		case percentCongestion > 40:
			change.decision = "Change"
		}
	case "MediumCongestion":
		switch {
		case percentCongestion <= 60:
			change.decision = "NoChange"
		case percentCongestion > 60:
			change.decision = "Change"
		}
	case "IntensiveCongestion":
		switch {
		case percentCongestion <= 80:
			change.decision = "NoChange"
		case percentCongestion > 80:
			change.decision = "Change"
		}
	}

	fmt.Println("A decisão foi de:", change.decision)
}
