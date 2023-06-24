package srvcontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"semaphores-adaptative/commons"
	"semaphores-adaptative/constants"
	"time"
)

type ControllerSrv struct{}

type TrafficSignal struct {
	Id         int `json:"id"`
	TimeGreen  int `json:"timegreen"`
	TimeYellow int `json:"timeyellow"`
	TimeRed    int `json:"timered"`
	Congestion int `json:"congestion"`
}

func NewControllerSrv() *ControllerSrv {
	return &ControllerSrv{}
}

func (ControllerSrv) HandleConnection(toController, fromController chan []commons.TrafficSignal) {
	for {
		time.Sleep(5 * time.Second)

		signals := getTrafficSignalData()
		// envia os dados para o monitor
		toController <- signals

		// recebe do executor a nova configuração dos sinais
		ts := <-fromController

		var signalsUpdated TrafficSignal
		if len(ts) > 0 {
			for i, _ := range ts {
				signalsUpdated.TimeRed = ts[i].TimeRed
				signalsUpdated.TimeGreen = ts[i].TimeGreen
				signalsUpdated.TimeYellow = ts[i].TimeYellow
				signalsUpdated.Id = ts[i].Id
				signalsUpdated.Congestion = ts[i].Congestion
				sendTrafficSignalData(signalsUpdated)
			}
		}

	}

}

func sendTrafficSignalData(trafficSignal TrafficSignal) error {
	payload, err := json.Marshal(trafficSignal)
	if err != nil {
		return err
	}

	resp, err := http.Post(constants.UrlPost, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func getTrafficSignalData() []commons.TrafficSignal {
	// Cria uma requisição GET
	req, err := http.NewRequest("GET", constants.UrlGet, nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)

	}

	// Envia a requisição para o servidor
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar requisição:", err)

	}
	defer resp.Body.Close()

	// Verifica o status da resposta
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erro na requisição. Status:", resp.Status)

	}

	// Lê o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler corpo da resposta:", err)

	}

	// Decodifica o JSON para o struct TrafficSignalSystem
	var signals []commons.TrafficSignal
	err = json.Unmarshal(body, &signals)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)

	}
	return signals
}
