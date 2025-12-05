package estructurasdecontrol

import "fmt"

func ForLoop() (int, int, int) {
	//standard for
	standarFor := 0
	for i := 0; i < 10; i++ {
		standarFor += i
	}
	//while
	whileFor := 0
	for whileFor <= 10 {
		whileFor++
	}
	//infinite for
	// infiniteFor := 0
	// for {
	// 	infiniteFor++
	// }
	// for true {
	// 	infiniteFor++
	// }

	//range for
	fruits := []string{"apple", "banana", "cherry", "orange", "pineapple", "strawberry", "watermelon"}
	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}

	fruitsMap := map[string]string{"apple": "red", "banana": "yellow", "cherry": "red"}
	for fruit, color := range fruitsMap {
		fmt.Println(fruit, color)
	}
	word := "word"
	for index, letter := range word {
		//imprime su unicode
		fmt.Println(index, letter)
		//formas de imprimir el valor de la letra
		fmt.Printf("%d %c\n", index, letter)
		fmt.Println(string(letter))
	}

	// continue salta a la siguiente iteración

	for index, fruit := range fruits {
		if index%2 == 0 {
			continue
		}
		fmt.Println("continue impares: ", index, fruit)
	}
	// break rompe con el ciclo
	breakFor := 0
	for {
		fmt.Println("break: ", breakFor)
		breakFor++
		if breakFor == 3 {
			break
		}

	}

	return standarFor, whileFor, breakFor

}
