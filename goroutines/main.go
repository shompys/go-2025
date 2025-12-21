package goroutines

import (
	"fmt"
	"time"
)

func myProcess(p string) {
	i := 0
	for i < 15 {
		time.Sleep(time.Millisecond * 150)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
}
func myProcessWithchannel(p string, c chan string) {
	i := 0
	for i < 15 {
		time.Sleep(time.Millisecond * 150)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
	c <- "done"
}

func Goroutines() {
	go myProcess("A")
	go myProcess("B")
	go myProcess("C")

	myFirstChannel := make(chan string)
	go myProcessWithchannel("D", myFirstChannel)
	result := <-myFirstChannel
	fmt.Println("result: ", result)
	close(myFirstChannel)

}
