package main
import (
    "fmt"
    )

func isPalindrome(x int) bool{
   if x<0{
       return false
   }
   
   original := x
   reversed:= 0
   
   for x!=0{
       reversed =reversed*10+ x%10 
       x/=10
   }
   return original ==reversed
}  

func main(){
	// your code goes here
	numbers:= []int {121,-121,10,12321}
	
	for _,num :=range numbers{
	    if isPalindrome(num){
	        fmt.Printf("%d is a Palindrome",num)
	    }else{
	        fmt.Printf("%d is not a palindrome",num)
	    }
	} 
}
