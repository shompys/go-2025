package goroutines

import "fmt"

func enviar(c chan <- int){
c <-10
}
func recibir(c <- chan int){
	fmt.Println(<-c)
}
func Channel(){
	ca := make(chan int)
	go enviar(ca)
	recibir(ca)
	go func (){
		ca <- 42
	}()
	fmt.Println(<-ca)

}



func ChannelBuffer(){
	ca := make(chan int, 1)
	ca<-43
	fmt.Println(<-ca)
	fmt.Printf("%T\n", ca)
}
func SendOnly(){//send
	ca:= make(chan <- int, 1)
	ca<-2
}
func RecibirOnly(){//receive
	ca:= make(<-chan int, 1)
	fmt.Println(<-ca)
}

func ChannelRanger(){

	c:= make(chan int)

	go func(){
		for i := range 5 {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("fin")

}
