package commons

import "semaphores-adaptative/constants"

var Goal string

type TrafficSignal struct {
	Id         int `json:"id"`
	TimeGreen  int `json:"timegreen"`
	TimeYellow int `json:"timeyellow"`
	TimeRed    int `json:"timered"`
}

// instancia um semáforo
func NewTrafficSignal(id int) TrafficSignal {
	s := TrafficSignal{Id: id, TimeGreen: constants.DefaultGreen, TimeYellow: constants.DefaultYellow, TimeRed: constants.DefaultRed}

	return s
}
