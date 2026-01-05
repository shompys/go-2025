package tiposdedatos

import "fmt"

func Float() (float32, float64) {
	// | 1 bit | 8 bits    | 23 bits  |
	// | Signo | Exponente | Mantisa  |
	var number float32 = 10.5 // hasta 24 digitos
	// | 1 bit | 11 bits    | 52 bits  |
	// | Signo | Exponente | Mantisa  |
	var number64 float64 = 10.5 // hasta 53 digitos
	fmt.Printf("%.1f\n", number64)//forma de imprimir especificamente el numero de decimales
	fmt.Printf("%.1f\n", number)
	return number, number64
}
