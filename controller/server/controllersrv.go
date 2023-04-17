package srvcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"semaphores-adaptative/commons"
)

type ControllerSrv struct{}

func NewControllerSrv() *ControllerSrv {
	return &ControllerSrv{}
}

func HandleConnection(conn net.Conn, toController, fromController chan []commons.TrafficSignal) {
	defer conn.Close()

	// recebe os dados dos sinais de trânsito
	decoder := json.NewDecoder(conn)
	//var msg Message
	var signals []commons.TrafficSignal
	if err := decoder.Decode(&signals); err != nil {
		log.Println(err)
		return
	}

	// envia os dados para o monitor
	toController <- signals

	// recebe do executor a nova configuração dos sinais
	ts := <-fromController

	// envia para o agente a nova configuração dos sinais
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(ts); err != nil {
		log.Println(err)
		return
	}
}

func (ControllerSrv) Run(toController chan []commons.TrafficSignal, fromController chan []commons.TrafficSignal) {
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
		go HandleConnection(conn, toController, fromController)
	}
}
