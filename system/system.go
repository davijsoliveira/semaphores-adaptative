package main

import (
	//"fmt"
	"semaphores-adaptative/constants"
	"sync"

	/*"semaphores-adaptative/controller/analyser"
	"semaphores-adaptative/controller/executor"
	"semaphores-adaptative/controller/monitor"
	"semaphores-adaptative/controller/planner"*/
	"semaphores-adaptative/traffic"
	//"semaphores-adaptative/trafficApp"
)

func main() {

	// instantiate semaphore system
	/*trafSystem := trafficApp.NewTrafficSignalSystem(constants.TrafficSignalNumber)
	for _, v := range trafSystem.TrafficSignals {
		fmt.Println("########################### INITIAL VALUES OF SEMAPHORES ############################################")
		fmt.Println("TrafficSignal ID:", v.Id, "Green:", v.TimeGreen, "Yellow:", v.TimeYellow, "Red:", v.TimeRed)
		fmt.Println("#####################################################################################################")
	}*/
	var wg sync.WaitGroup
	// instantiate traffic flow
	trafFlow := traffic.NewTrafficFlow(constants.TrafficSignalNumber)
	wg.Add(1)
	trafFlow.Exec()
	wg.Wait()
	/*mon := monitor.NewMonitor()
	anl := analyser.NewAnalyser()
	pln := planner.NewPlanner()
	exc := executor.NewExecutor()
	m := mon.Exec(trafFlow)
	cr := anl.Exec(m)
	plan := pln.Exec(cr, trafSystem)
	changeSignals := exc.Exec(plan)
	trafSystem.Exec(changeSignals)*/

}
