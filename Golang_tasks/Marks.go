/*Write a program to generate student score card/exam report. Please take input from
user, it contains 4 subjects (Maths, Science, S.S, English) mark, name of the student
and standard. You will need to generate score report based on below cases.
i. Any subject mark less than 18 then student will be failed.
ii. Calculate percentage for 4 subjects. Make sure that user will give input mark out
of 50.
iii. Exam report will include class (failed, second, first, distinction) of the student
which student will get based on percentage. If percentage is


A. Less than 35, class should be failed (No need to calculate percentage)
B. Less than 60 but greater than or equals 35, class should be second.
C. Less than 70 but greater than or equals 60, class should be first.
D. An equals or above 70, class should be distinction.*/

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
	fmt.Print("\t", name)

	fmt.Print("\tTotal Marks =")
	fmt.Printf("\t%0.2f", total)
	fmt.Print("\n")

	fmt.Print("\tPercentage =")
	fmt.Printf("\t%0.2f", percentage)
	fmt.Print("%\n")

	if percentage >= 90 {
		print("Grade: A")
		fmt.Print("\n")
	} else if percentage >= 80 && percentage < 90 {
		print("Grade: B")
		fmt.Print("\n")
	} else if percentage >= 70 && percentage < 80 {
		print("Grade: C")
		fmt.Print("\n")
	} else if percentage >= 60 && percentage < 70 {
		print("Grade: D")
		fmt.Print("\n")
	} else {
		print("Grade: F")
		fmt.Print("\n")
	}
	// fmt.Println("Hello GoLang")
}
