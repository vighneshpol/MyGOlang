package main

import (
	"fmt"
	"time"
)

func main() {
	parkingLot := NewParkingLot(2, 2, 10.0, 20.0)

	parkingLot.Enter("MH19", Hatchback)
	parkingLot.Enter("Mh1224", SUV)
	parkingLot.Enter("MH191", Hatchback)
	parkingLot.Enter("Mh122", SUV)
	parkingLot.Enter("MH1921", Hatchback)

	//for admin view
	fmt.Println("Checking admin view-------->")
	parkingLot.AdminView()

	time.Sleep(2 * time.Minute)
	parkingLot.Exit("MH19")
	parkingLot.Exit("MH1921")

	fmt.Println("Admin view------->")
	parkingLot.AdminView()

}