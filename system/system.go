package main

import (
	"fmt"
	"semaphores-adaptative/analyser"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller/monitor"
	"semaphores-adaptative/traffic"
	semaphore_app "semaphores-adaptative/trafficApp"
)

func main() {

	// instantiate semaphore system
	semSystem := semaphore_app.NewSemaphoreSystem(constants.NumberSemaphores)
	for _, v := range semSystem.Semaphores {
		fmt.Println("########################### INITIAL VALUES OF SEMAPHORES ############################################")
		fmt.Println("Semaphore ID:", v.Id, "Green:", v.TimeGreen, "Yellow:", v.TimeYellow, "Red:", v.TimeRed)
		fmt.Println("#####################################################################################################")
	}

	/*var c = make(map[int][]int)
	s0 := []int{120, 15, 15}
	s1 := []int{90, 20, 20}
	c[0] = s0
	c[2] = s1
	semSystem.Exec(c)

	for _, v := range semSystem.Semaphores {
		fmt.Println("########################### CHANGED VALUES OF SEMAPHORES ############################################")
		fmt.Println("Semaphore ID:", v.Id, "Green:", v.TimeGreen, "Yellow:", v.TimeYellow, "Red:", v.TimeRed)
		fmt.Println("#####################################################################################################")
	}*/

	// instantiate traffic flow
	trafFlow := traffic.NewTrafficFlow(constants.NumberSemaphores)
	mon := monitor.NewMonitor()
	anl := analyser.NewAnalyser()
	m := mon.Exec(trafFlow)
	anl.Exec(m)
	/*for i, v := range trafFlow.TrafficPerSemaphore {
		fmt.Println("########################### INITIAL VALUES BY TRAFFIC### ############################################")
		fmt.Println("Semaphore ID:", i, "Jam:", v)
		fmt.Println("#####################################################################################################")

	}
	trafFlow.Exec()*/
	/*t := trafFlow.Sense()
	fmt.Println("########################### VALUES BY TRAFFIC########################################################")
	for i, v := range t.TrafficPerSemaphore {
		fmt.Println("Semaphore ID:", i, "Jam:", v)
	}
	fmt.Println("#####################################################################################################")
	*/

}
