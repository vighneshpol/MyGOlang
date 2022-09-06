/*
Write a program that asks the user for a long string containing multiple words. Print back to
the user the same string, except with the words in backwards order with first character of
each word in capital. For example :
Input: your name is xyz,
Output: Xyz Is Name Your
*/
package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	tobeReversed := []string{
		"My Name Is Xyz",
	}

	for i := 0; i < len(tobeReversed); i++ {
		fmt.Printf("Input ;%s\n%s\n\n",
			tobeReversed[i],
			reverseWords(tobeReversed[i]))
	}
}
