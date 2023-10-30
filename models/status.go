package models

import "time"

type Status struct {
	Counter     int       `json:"counter"`
	Status      bool      `json:"status"`
	LastChanged time.Time `json:"lastChanged"`
}
