//Program to print 1 to 10 using sync.waitgroup

package main

import (
	"fmt"
	"sync"
	"time"
)

func replitGoRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
	time.Sleep(1000* time.Millisecond)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go replitGoRoutine(&wg)
	wg.Wait()

}
