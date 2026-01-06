package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)


func Atomic() {
	
	var wg sync.WaitGroup
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	
	// ch := make(chan int)
	var contador int64
	const gs = 1000
	wg.Add(gs)
	
	for range gs {
		go func(){
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Número de Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(contador)
}