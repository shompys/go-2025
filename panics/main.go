package panics

import "fmt"

// el defer siemrpe se ejecuta al final de todo inclusive de un panic, nos sirve para prevenir que se rompa la aplicación
func Panics() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperado de un panic", r)
		}
	}()

	panic("soy un panic")
}
