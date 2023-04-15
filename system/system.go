/*
***********************************************************************************************************************************************************
Author: Davi Oliveira
Description: This code implements a simple MAPE-K for traffic signal timing control. The time of the signal traffic may change according to the traffic flow.
Date: 06/03/2023
***********************************************************************************************************************************************************
*/
package main

import (
	"semaphores-adaptative/commons"

	"semaphores-adaptative/controller"
	srvcontroller "semaphores-adaptative/controller/server"
	"semaphores-adaptative/goal"
	"sync"
)

func main() {
	// variável wait group para controlar as go routines
	var wg sync.WaitGroup

	// cria os canais
	appToMonitor := make(chan []commons.TrafficSignal)
	goalToController := make(chan string)
	executeToApp := make(chan []commons.TrafficSignal)

	// instancia o componente de configuração da meta, o controller e o frontend
	gl := goal.NewGoalConfiguration()
	ctl := controller.NewController()
	srv := srvcontroller.NewControllerSrv()

	// executa os componentes
	wg.Add(7)
	go gl.Exec(goalToController)
	go ctl.Exec(appToMonitor, executeToApp, goalToController)
	go srv.Run(appToMonitor, executeToApp)
	wg.Wait()
}
