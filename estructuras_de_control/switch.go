package estructurasdecontrol

import "fmt"

func Switch() {
	condition := 1
	// no es necesario return ni break
	switch condition {
	case 1:
		fmt.Println("The condition is 1")
	case 2:
		fmt.Println("The condition is 2")
	case 3:
		fmt.Println("The condition is 3")
	default:
		fmt.Println("The condition is not 1, 2 or 3")
	}

	switch {
	case condition > 0:
		fmt.Println("The condition is greater than 0")
	case condition < 0:
		fmt.Println("The condition is less than 0")
	default:
		fmt.Println("The condition is 0")
	}

	expression1 := 1
	expression2 := 2
	expression3 := 3

	switch condition {
	case expression1, expression2, expression3:
		fmt.Println("The condition is 1, 2 or 3")
	default:
		fmt.Println("The condition is not 1, 2 or 3")
	}

	// === TYPE SWITCH ===
	// Permite verificar el tipo de una interface{}
	typeSwitch("Hola")
	typeSwitch(42)
	typeSwitch(3.14)
	typeSwitch(true)
	typeSwitch(Persona{Nombre: "Juan", Edad: 25})
	typeSwitch([]int{1, 2, 3})
}

// Struct para el ejemplo
type Persona struct {
	Nombre string
	Edad   int
}

func typeSwitch(h interface{}) {
	// v recibe el valor convertido al tipo del case
	switch v := h.(type) {
	case string:
		fmt.Printf("Es string: %q (len: %d)\n", v, len(v))
	case int:
		fmt.Printf("Es int: %d (el doble: %d)\n", v, v*2)
	case float64:
		fmt.Printf("Es float64: %.2f\n", v)
	case bool:
		fmt.Printf("Es bool: %t\n", v)
	case Persona:
		// Aquí v ya es tipo Persona, puedes acceder a sus campos
		fmt.Printf("Es Persona: %s tiene %d años\n", v.Nombre, v.Edad)
	case []int:
		fmt.Printf("Es slice de int: %v (len: %d)\n", v, len(v))
	default:
		fmt.Printf("Tipo desconocido: %T\n", v)
	}
}
