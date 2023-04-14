package srvcontroller

import (
	"encoding/json"
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

	// Lê a mensagem enviada pelo cliente
	decoder := json.NewDecoder(conn)
	//var msg Message
	var signals []signalControlApp.TrafficSignal
	if err := decoder.Decode(&signals); err != nil {
		log.Println(err)
		return
	}

	// repassa para o monitor
	toMonitor <- signals
	ts := <-fromExecutor

	// Responde com uma mensagem de confirmação
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(ts); err != nil {
		log.Println(err)
		return
	}

	// Imprime a mensagem recebida
	//for _, signal := range signals {
	//	fmt.Printf("ID: %d, Verde: %d, Amarelo: %d, Vermelho: %d\n", signal.Id, signal.TimeGreen, signal.TimeYellow, signal.TimeRed)
	//}
}

func (ControllerSrv) Run(toMonitor chan []signalControlApp.TrafficSignal, fromExecutor chan []signalControlApp.TrafficSignal) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// Aceita a conexão do cliente
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// Lida com a mensagem recebida
		go HandleConnection(conn, toMonitor, fromExecutor)
	}
}
