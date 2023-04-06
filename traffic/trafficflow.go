package traffic

import (
	"math/rand"
	"semaphores-adaptative/constants"
	"time"
)

type TrafficFlow struct {
	TrafficPerSemaphore []int
}

func NewTrafficFlow(num int) *TrafficFlow {
	t := make([]int, num)
	tf := TrafficFlow{TrafficPerSemaphore: t}
	for i, _ := range tf.TrafficPerSemaphore {
		tf.TrafficPerSemaphore[i] = constants.DefaultTraffic
	}
	return &tf
}

func (t *TrafficFlow) Exec() {
	for {
		for i := 0; i < constants.NumberSemaphores; i++ {
			time.Sleep(10 * time.Second)
			rand.Seed(time.Now().UnixNano())
			jam := rand.Intn(constants.MaxTraffic)
			t.TrafficPerSemaphore[i] = jam
		}
	}
}

func (t TrafficFlow) Sense() TrafficFlow {
	return t
}
