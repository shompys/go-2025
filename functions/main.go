package functions

import (
	"errors"
	"fmt"
)

func sum(a, b int) (val int) {
	val = a + b
	return val
}
func notaciónPuntos(name string, values ...float64) (float64, string) {
	var val float64
	for _, value := range values {
		val += value
	}

	return val, name
}
func withCallback(name string, callback func(name string)) {
	callback(name)
}

func division(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("denominator cannot be 0")
	}
	return numerator / denominator, nil
}

func retornoValoresNombrados(name string) (sameName string) {
	sameName = name
	return sameName
}

func retornoFunctions(saludo string) func(name string) string {
	return func(name string) string {
		return saludo + " " + name
	}
}

func Functions() (int, float64, string) {
	result := sum(10, 20)
	result2, _ := notaciónPuntos("juancito", 10.5, 20.5, 30.5)

	withCallback("juancito", func(name string) {
		fmt.Println("Hello,", name)
	})

	result3, err := division(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Result:", result3)
	name := retornoValoresNombrados("jonathan")

	saludo := retornoFunctions("Hello")
	fmt.Println(saludo("jonathan"))
	return result, result2, name
}
