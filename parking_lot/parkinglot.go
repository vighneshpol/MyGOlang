package main

import (
	"fmt"
	"sync"
)

// Vehicle represents a vehicle with a registration number and color.
type Vehicle struct {
	RegistrationNumber string
	Color              string
}

// Slot represents a parking slot.
type Slot struct {
	Number  int
	Vehicle *Vehicle
}

// ParkingLot represents a parking lot with slots.
type ParkingLot struct {
	Slots []Slot
	mu    sync.Mutex
}

// NewParkingLot creates a new parking lot with the given number of slots.
func NewParkingLot(capacity int) *ParkingLot {
	slots := make([]Slot, capacity)
	for i := range slots {
		slots[i] = Slot{Number: i + 1}
	}
	return &ParkingLot{Slots: slots}
}

// Park parks a vehicle in the first available slot.
func (pl *ParkingLot) Park(vehicle Vehicle) (int, error) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	for i := range pl.Slots {
		if pl.Slots[i].Vehicle == nil {
			pl.Slots[i].Vehicle = &vehicle
			return pl.Slots[i].Number, nil
		}
	}
	return 0, fmt.Errorf("parking lot is full")
}

// Leave removes the vehicle from the given slot number.
func (pl *ParkingLot) Leave(slotNumber int) error {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	if slotNumber <= 0 || slotNumber > len(pl.Slots) {
		return fmt.Errorf("invalid slot number")
	}

	if pl.Slots[slotNumber-1].Vehicle == nil {
		return fmt.Errorf("slot already empty")
	}

	pl.Slots[slotNumber-1].Vehicle = nil
	return nil
}

// Status prints the current status of the parking lot.
func (pl *ParkingLot) Status() {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	fmt.Println("Slot No.\tRegistration No.\tColor")
	for _, slot := range pl.Slots {
		if slot.Vehicle != nil {
			fmt.Printf("%d\t\t%s\t\t%s\n", slot.Number, slot.Vehicle.RegistrationNumber, slot.Vehicle.Color)
		}
	}
}
