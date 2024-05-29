package main

import "fmt"

//write a go program to find longest string from a slice
func largestString(slice []string) string {
	if len(slice) == 0 {
		return ""
	}

	longestStr := slice[0]

	for _, str := range slice[1:] {
		if len(str) > len(longestStr) {
			longestStr = str
		}
	}
	return longestStr
}

func main() {
	slice := []string{"abcd", "def", "ghi", "Vighnesh ", "skdksddksjk","Gaurav Prakash Bornare"}
	fmt.Println(largestString(slice))
}
