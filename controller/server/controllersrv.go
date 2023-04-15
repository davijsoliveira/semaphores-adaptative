package srvcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"semaphores-adaptative/signalControlApp"
)

type ControllerSrv struct{}

func NewControllerSrv() *ControllerSrv {
	return &ControllerSrv{}
}

func HandleConnection(conn net.Conn, toMonitor chan []signalControlApp.TrafficSignal, fromExecutor chan []signalControlApp.TrafficSignal) {
	defer conn.Close()

	// recebe os dados dos sinais de trânsito
	decoder := json.NewDecoder(conn)
	//var msg Message
	var signals []signalControlApp.TrafficSignal
	if err := decoder.Decode(&signals); err != nil {
		log.Println(err)
		return
	}

	// envia os dados para o monitor
	toMonitor <- signals

	// recebe do executor a nova configuração dos sinais
	ts := <-fromExecutor

	// envia para o agente a nova configuração dos sinais
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(ts); err != nil {
		log.Println(err)
		return
	}
}

func (ControllerSrv) Run(toMonitor chan []signalControlApp.TrafficSignal, fromExecutor chan []signalControlApp.TrafficSignal) {
	listener, err := net.Listen("tcp", ":8080")
	fmt.Println("Running Frontend on port 8080............")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// recebe as conexões do agente
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// encaminha os dados da conexão para processamento no MAPE-K
		go HandleConnection(conn, toMonitor, fromExecutor)
	}
}
