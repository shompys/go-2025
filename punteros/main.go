package punteros

import "fmt"

func Punteros() {
	//referencia de nombre v que apunta a la direccion de memoria que contiene el valor 4
	var v int = 4
	// var p *int

	p := &v
	*p = 2
	fmt.Printf("El valor de v es %d\n", v)
	fmt.Printf("La posición de memoria de v es %v\n", &v)

	fmt.Printf("El valor de p es %d\n", *p)
	fmt.Printf("La posición de memoria de p es %v\n", p)

	//var pepe = new(int) //forma de declarar una reserva de memoria con valor 0 pero retornando la dirección en pepe es decir para acceder *pepe
	//hacer new(int) es lo mismo que
	//var x int  // vale 0
	//pepe := &x // puntero hacia x
}

// ejemplo claro:
func ModificaOriginal(value *string) string {
	aux := *value + "epetacular"
	return aux
}

var value string = "inicial"

func init() {
	fmt.Println("soy init")
	val := ModificaOriginal(&value)
	fmt.Println(val)
}
