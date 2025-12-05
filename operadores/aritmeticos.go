package operadores

func Aritmeticos() (sum, subtraction, multiplication, division, modulo, increment, decrement int) {
	var number1 int = 20
	var number2 int = 10

	sum = number1 + number2
	subtraction = number1 - number2
	multiplication = number1 * number2
	division = number1 / number2
	modulo = number1 % number2

	// no retornan el valor, solo modifican el valor de la variable
	number1++
	increment = number1
	number1--
	decrement = number1

	return sum, subtraction, multiplication, division, modulo, increment, decrement

}
