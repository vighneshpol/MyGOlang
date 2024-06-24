package main

import (
	"testing"
)

func TestParkingLot(t *testing.T) {
	parkingLot := NewParkingLot(2)

	// Test parking vehicles
	slotNumber, err := parkingLot.Park(Vehicle{RegistrationNumber: "KA-01-HH-1234", Color: "White"})
	if err != nil || slotNumber != 1 {
		t.Errorf("expected slot number 1, got %d, err: %v", slotNumber, err)
	}

	slotNumber, err = parkingLot.Park(Vehicle{RegistrationNumber: "KA-01-HH-9999", Color: "Black"})
	if err != nil || slotNumber != 2 {
		t.Errorf("expected slot number 2, got %d, err: %v", slotNumber, err)
	}

	// Test parking when lot is full
	_, err = parkingLot.Park(Vehicle{RegistrationNumber: "KA-01-HH-0001", Color: "Blue"})
	if err == nil {
		t.Errorf("expected error for full parking lot, got nil")
	}

	// Test leaving a slot
	err = parkingLot.Leave(1)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Test leaving an already empty slot
	err = parkingLot.Leave(1)
	if err == nil {
		t.Errorf("expected error for empty slot, got nil")
	}
}
