package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your name: ")
	// var name string;
	// fmt.Scanln(&name)//This will scan until the space starts if you want to read the whole line use bufio.scan
	// fmt.Printf("Hello, %s", name)
	// fmt.Print("\n")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello Mr.", name)
}
