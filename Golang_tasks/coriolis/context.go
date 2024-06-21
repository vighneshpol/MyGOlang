package main
import (
    "fmt"
    "context"
    "time"
    )


func performsomeTask(ctx context.Context, taskID int){
    select{
	case <- time.After(2 * time.Second):
        fmt.Println("Task complete  %d \n",taskID)
	case <-ctx.Done():
		fmt.Println("Canceled:", ctx.Err())
        
    }
    // case <- ctx.Done():
    // fmt.Println("task cancelled %v\n",taskID,ctx.Err())
// default:
// 	// Handle the case where both channels are ready
// 	fmt.Println("Both channels are ready, unexpected behavior!")

}
func main(){
ctx:= context.Background()
ctx,cancel:=context.WithTimeout(1*time.Second)
defer cancel()
	// your code goes here
	go performsomeTask(ctx,1)
	time.Sleep(3*time.Second)
}
