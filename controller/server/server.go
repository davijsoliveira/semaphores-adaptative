package main

import (
	"encoding/json"
	"log"
	"net"
)

type Signal struct {
	ID         int `json:"id"`
	TimeGreen  int `json:"timegreen"`
	TimeYellow int `json:"timeyellow"`
	TimeRed    int `json:"timered"`
}

type Message struct {
	Text string `json:"text"`
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Lê a mensagem enviada pelo cliente
	decoder := json.NewDecoder(conn)
	//var msg Message
	var signals []Signal
	if err := decoder.Decode(&signals); err != nil {
		log.Println(err)
		return
	}

	// Imprime a mensagem recebida
	log.Printf("Recebido: %d", signals[0].ID)

	// Responde com uma mensagem de confirmação
	response := Message{Text: "Mensagem recebida com sucesso!"}
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(response); err != nil {
		log.Println(err)
		return
	}
}

func main() {

}
