package services

import (
	"testing"
)

func TestCounter(t *testing.T) {
	UpdateStatus()

	if GetStatus().Counter != 1 {
		t.Errorf("Expected: %v, got: %v", 1, status.Counter)
	}

	UpdateStatus()

	if GetStatus().Counter != 2 {
		t.Errorf("Expected: %v, got: %v", 2, status.Counter)
	}
}

func TestChangingStatus(t *testing.T) {
	Shutdown()

	if GetStatus().Status != false {
		t.Errorf("Expected: %v, got: %v", false, status.Status)
	}

	PowerOn()

	if GetStatus().Status != true {
		t.Errorf("Expected: %v, got: %v", true, status.Status)
	}
}
