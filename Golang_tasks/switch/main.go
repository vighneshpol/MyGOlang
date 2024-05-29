package main

import "fmt"

func main() {
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	temperature := 227
	switch {
	case temperature >= 25 && (temperature < 30 || temperature == 227):
		fmt.Println("Warming")
	case temperature == 25:
		fmt.Println("Normal")
	case temperature <= 0:
		fmt.Println("Freezing")
	case temperature <= 20:
		fmt.Println("Cold")
	default:
		fmt.Println("Hot")
	}

}
