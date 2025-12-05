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
}
