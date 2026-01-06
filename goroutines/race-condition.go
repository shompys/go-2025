package goroutines

import (
	"fmt"
	"runtime"
	"sync"
)


func RaceCondition() {
	var mux sync.Mutex
	var wg sync.WaitGroup
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	
	// ch := make(chan int)
	contador := 0
	const gs = 1000000
	wg.Add(gs)
	
	for range gs {
		go func(){
			mux.Lock()
			contador++
			runtime.Gosched()//yield thread
			mux.Unlock()
			wg.Done()
		}()
		fmt.Println("Número de Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(contador)
}