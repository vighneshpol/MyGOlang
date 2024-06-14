package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var parkingLot *ParkingLot
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")

		switch args[0] {
		case "create_parking_lot":
			if len(args) != 2 {
				fmt.Println("Usage: create_parking_lot <capacity>")
				continue
			}
			capacity, err := strconv.Atoi(args[1])
			if err != nil || capacity <= 0 {
				fmt.Println("Invalid capacity")
				continue
			}
			parkingLot = NewParkingLot(capacity)
			fmt.Printf("Created a parking lot with %d slots\n", capacity)

		case "park":
			if len(args) != 3 {
				fmt.Println("Usage: park <registration_number> <color>")
				continue
			}
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			slotNumber, err := parkingLot.Park(Vehicle{RegistrationNumber: args[1], Color: args[2]})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Allocated slot number: %d\n", slotNumber)
			}

		case "leave":
			if len(args) != 2 {
				fmt.Println("Usage: leave <slot_number>")
				continue
			}
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			slotNumber, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid slot number")
				continue
			}
			err = parkingLot.Leave(slotNumber)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Slot number %d is free\n", slotNumber)
			}

		case "status":
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			parkingLot.Status()

		case "exit":
			return

		default:
			fmt.Println("Unknown command")
		}
	}
}
