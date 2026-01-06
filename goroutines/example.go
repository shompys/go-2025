package goroutines

import (
	"fmt"
	"runtime"
	"sync"
)

func foo(){
	defer wg.Done()// le va restando 1 al contador de goroutines
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}
func bar(){
	defer wg.Done()// le va restando 1 al contador de goroutines
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
var wg sync.WaitGroup

func Example(){
	wg.Add(2) //contador de goroutines
	
	go foo()
	go bar()
	defer fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Number of CPUs:", runtime.NumCPU())
}