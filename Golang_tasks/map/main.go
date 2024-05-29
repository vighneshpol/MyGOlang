package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)

	studentGrades["Vighnesh"] = 100
	studentGrades["Poppy"] = 90
	studentGrades["Dilawar"] = 76
	studentGrades["Gaurav"] = 87

	fmt.Println("Marks of vighnesh is:", studentGrades["Vighnesh"])
	studentGrades["Dilawar"] = 100
	fmt.Println("Marks of Dilawar is:", studentGrades["Dilawar"])
	// delete(studentGrades, "Dilawar")
	fmt.Println("Dilawar Marks", studentGrades["Dilawar"])

	//check key if exists
	grades, eists := studentGrades["Dilawar"]
	fmt.Println("Grades of Dilawar is:", grades)
	fmt.Println("Dilawar is exists:", eists)
	// studentGrades["Dilawar"] = 76
	

	for index,values:= range studentGrades{
		fmt.Printf("Key is %s and marks is %d\n",index,values)
	}

	employeeGrades:=map[string]int{
		"Alice": 100,
		"Bob": 90,
		"Charlie": 80,
		"Dave": 70,

	}
	for index,value:= range employeeGrades{
		fmt.Printf("Key ==========is %s and marks is %d\n",index,value)
	}
}
