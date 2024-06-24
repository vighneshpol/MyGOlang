package main

import (
	"fmt"
	"time"
)

type CarType int

const (
	Hatchback CarType = iota
	SUV
)

type Car struct {
	licensePlate string
	carType      CarType
	entryTime    time.Time
}

type ParkingLot struct {
	hatchbackCapacity int
	suvCapacity       int
	suvCount          int
	hatchbackRate     float64
	suvRate           float64
	cars              map[string]Car
	hatchbackCount    int
}

func NewParkingLot(hatchbackCapacity, suvCapacity int, hatchbackRate, suvRate float64) *ParkingLot{
	return &ParkingLot{
		hatchbackCapacity: hatchbackCapacity, suvCapacity: suvCapacity, suvRate: suvRate, cars: make(map[string]Car),
	}
}

func (p *ParkingLot) Enter(licensePlate string, carType CarType) bool {
	if carType == Hatchback {
		if p.hatchbackCount < p.hatchbackCapacity {
			p.hatchbackCount++
		} else if p.suvCount < p.suvCapacity {
			p.suvCount++
		} else {
			fmt.Println("No Space available for hatchback type")
			return false
		}
	}
	p.cars[licensePlate] = Car{licensePlate, carType, time.Now()}
	fmt.Println("Car registered:",licensePlate)
	return true
}

func (p *ParkingLot) Exit(licensePlate string)float64{
	car,exists:= p.cars[licensePlate]
	if !exists{
		fmt.Println("Car not found",licensePlate)
		return 0
	}
	duration:= time.Since(car.entryTime).Hours()
	var rate float64
	if car.carType == Hatchback{
		rate = p.hatchbackRate
		//checcks if hatchback vehicle are fully occupied if hatchback cout is less than the capacity then we are are entering the hatchback car
		if p.hatchbackCount > p.hatchbackCapacity{
			p.hatchbackCount--
		} else if p.suvCount > p.suvCapacity{
			p.suvCount--
		} else {
			rate = p.suvRate
			p.suvCount--
		}
		
	}	
	delete(p.cars,licensePlate)
		totalCost:= duration * rate
		fmt.Printf("car exited at %s, Total cost is %.2f\n",licensePlate,totalCost)
		return totalCost
}

func(p *ParkingLot)AdminView(){
	fmt.Println("Currently parked cars in basement are:")
	for _,car:= range p.cars{
		fmt.Printf("License palte: %s, Car type: %v, entry time:%v\n",car.licensePlate,car.carType,car.entryTime)

	}
}
