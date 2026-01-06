package goroutines

import "fmt"

func Select(){
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	//enviar
	go enviar2(par, impar, salir)

	//recibir
	recibir2(par, impar, salir)

}
//guardar valores en canales: chan <-
func enviar2(par, impar, salir chan <- int){
	for j := 0; j < 100; j++{
		if j % 2 == 0{
			par <- j
		} else{
			impar <- j
		}
	}
	close(par)
	close(impar)
	salir <- 0
}
//sacar valores de canales: <- chan
func recibir2(par, impar, salir <- chan int){

	for {
		select {
		case v, ok := <- par:
			if !ok {
				par = nil
				continue
			}
			fmt.Println("Desde el canal par:", v)
		case v, ok := <- impar:
			if !ok {
				impar = nil
				continue
			}
			fmt.Println("Desde el canal impar:", v)
		case v := <- salir:
			
			fmt.Println("Desde el canal salir:", v)
			return
		}
	}
}