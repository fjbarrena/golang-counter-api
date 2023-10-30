package services

import (
	"counter-api/models"
	"time"
)

var status = &models.Status{
	Counter:     0,
	Status:      true,
	LastChanged: time.Now(),
}

func UpdateStatus() models.Status {
	status.Counter++
	
	// Returning a copy of status to don't allow modifications outside the service
	return *status
}

func Shutdown() {
	status.Status = false
	status.LastChanged = time.Now()
}

func PowerOn() {
	status.Status = true
	status.LastChanged = time.Now()
}

func GetStatus() models.Status {
	return *status
}
