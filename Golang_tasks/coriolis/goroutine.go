package main
import "fmt"

func printAbc(){
    fmt.Print("hello")
}
func main(){
	// your code goes here
 fmt.Println("In main")
 go printAbc()
}
