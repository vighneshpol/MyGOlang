package main

import "fmt"

func main() {
	var eng, phy, chem, math, com float32
	var total, percentage float32
	var name string

	fmt.Print("\n")
	fmt.Print("\tPercentage of Student\n")
	// fmt.Scan("\n",name)
	fmt.Println("Name of student")
	fmt.Scanln(&name)
	fmt.Scan(&eng, &phy, &chem, &math, &com)

	total = eng + phy + chem + math + com
	percentage = (total / 500.0) * 100.0
	fmt.Print(name)

	fmt.Print("\tTotal Marks =")
	fmt.Printf("\t%0.2f", total)
	fmt.Print("\n")

	fmt.Print("\tPercentage =")
	fmt.Printf("\t", percentage)
	fmt.Print("%\n")

	// fmt.Println("Hello GoLang")
}
