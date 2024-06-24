package main
import "fmt"


func main(){
     n:=10
       a,b:=0,1
       
       fmt.Printf("fibonacci till %d is \n",n)
   fmt.Printf("%d, %d ",a,b)
       
        for i:=2;i <=n ; i++{
        a,b = b, a+b
        fmt.Printf("%d\n",b)
    }
  fmt.Println()
   
}