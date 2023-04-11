package traffic

import (
	"math/rand"
	"semaphores-adaptative/constants"
	"time"
)

// tipo fluxo de trânsito representando o ambiente
type TrafficFlow struct {
	TrafficPerSemaphore []int
}

// cria um novo fluxo de trânsito
func NewTrafficFlow(num int) *TrafficFlow {
	t := make([]int, num)
	tf := TrafficFlow{TrafficPerSemaphore: t}
	for i := range tf.TrafficPerSemaphore {
		tf.TrafficPerSemaphore[i] = constants.DefaultTraffic
	}
	tf.TrafficPerSemaphore[0] = 30
	tf.TrafficPerSemaphore[2] = 35
	return &tf
}

// executa o fluxo de trânsito, gerando congestionamentos aleatórios
func (t *TrafficFlow) Exec() {
	for {
		for i := 0; i < constants.TrafficSignalNumber; i++ {
			time.Sleep(10 * time.Second)
			rand.Seed(time.Now().UnixNano())
			jam := rand.Intn(constants.MaxTraffic)
			t.TrafficPerSemaphore[i] = jam
		}
	}
}

// capta o congestionamento dos semáforos
func (t TrafficFlow) Sense() TrafficFlow {
	return t
}
