package services

import (
	"testing"
)

func TestCounter(t *testing.T) {
	initialStatus := GetStatus()

	UpdateStatus()

	if GetStatus().Counter != 1 {
		t.Errorf("Expected: %v, got: %v", 1, status.Counter)
	}
	if GetStatus().LastChanged != initialStatus.LastChanged {
		t.Errorf("Unexpected lastChanged change")
	}

	UpdateStatus()

	if GetStatus().Counter != 2 {
		t.Errorf("Expected: %v, got: %v", 2, status.Counter)
	}
	if GetStatus().LastChanged != initialStatus.LastChanged {
		t.Errorf("Unexpected lastChanged change")
	}
}

func TestChangingStatus(t *testing.T) {
	initialStatus := GetStatus()
	
	Shutdown()

	if GetStatus().Status != false {
		t.Errorf("Expected: %v, got: %v", false, status.Status)
	}
	if GetStatus().Counter != initialStatus.Counter {
		t.Errorf("Unexpected counter change")
	}

	PowerOn()

	if GetStatus().Status != true {
		t.Errorf("Expected: %v, got: %v", true, status.Status)
	}
	if GetStatus().Counter != initialStatus.Counter {
		t.Errorf("Unexpected counter change")
	}
}
