package main

import (
	"fmt"
	"semaphores-adaptative/analyser"
	"semaphores-adaptative/constants"
	"semaphores-adaptative/controller/monitor"
	"semaphores-adaptative/executor"
	"semaphores-adaptative/planner"
	"semaphores-adaptative/traffic"
	//semaphore_app "semaphores-adaptative/trafficApp"
	"semaphores-adaptative/trafficApp"
)

func main() {

	// instantiate semaphore system
	trafSystem := trafficApp.NewTrafficSignalSystem(constants.TrafficSignalNumber)
	for _, v := range trafSystem.TrafficSignals {
		fmt.Println("########################### INITIAL VALUES OF SEMAPHORES ############################################")
		fmt.Println("TrafficSignal ID:", v.Id, "Green:", v.TimeGreen, "Yellow:", v.TimeYellow, "Red:", v.TimeRed)
		fmt.Println("#####################################################################################################")
	}

	/*var c = make(map[int][]int)
	s0 := []int{120, 15, 15}
	s1 := []int{90, 20, 20}
	c[0] = s0
	c[2] = s1
	trafSystem.Exec(c)

	for _, v := range trafSystem.TrafficSignals {
		fmt.Println("########################### CHANGED VALUES OF SEMAPHORES ############################################")
		fmt.Println("TrafficSignal ID:", v.Id, "Green:", v.TimeGreen, "Yellow:", v.TimeYellow, "Red:", v.TimeRed)
		fmt.Println("#####################################################################################################")
	}*/

	// instantiate traffic flow
	trafFlow := traffic.NewTrafficFlow(constants.TrafficSignalNumber)
	mon := monitor.NewMonitor()
	anl := analyser.NewAnalyser()
	pln := planner.NewPlanner()
	exc := executor.NewExecutor()
	m := mon.Exec(trafFlow)
	cr := anl.Exec(m)
	plan := pln.Exec(cr, trafSystem)
	changeSignals := exc.Exec(plan)
	trafSystem.Exec(changeSignals)
	/*for i, v := range trafFlow.TrafficPerSemaphore {
		fmt.Println("########################### INITIAL VALUES BY TRAFFIC### ############################################")
		fmt.Println("TrafficSignal ID:", i, "Jam:", v)
		fmt.Println("#####################################################################################################")

	}
	trafFlow.Exec()*/
	/*t := trafFlow.Sense()
	fmt.Println("########################### VALUES BY TRAFFIC########################################################")
	for i, v := range t.TrafficPerSemaphore {
		fmt.Println("TrafficSignal ID:", i, "Jam:", v)
	}
	fmt.Println("#####################################################################################################")
	*/

}
