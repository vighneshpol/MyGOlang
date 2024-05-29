package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Person_c_details Contact
	Person_a_details Address
}

type Contact struct {
	email      string
	phoneNumber string
}
type Address struct{
	house int
	area string
	state string
}

func main() {
	// Create a new instance of the Person struct

	person1 := Person{
		FirstName: "GaurAV",
		LastName:  "pAL",
		Age:       27,
		Person_c_details: Contact{
			email: "contact@gmail.com",
			phoneNumber: "+91 9876543210",
		},
		Person_a_details: Address{
			house: 123,
            area: "Vighnesh Nagar",
            state: "Maharashtra",
		},
	}
	fmt.Println("Vighnesh Person", person1)
}
