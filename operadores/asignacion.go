package operadores

func Asignacion() (assignment int) {
	var number1 int = 20
	var number2 int = 10

	number1 = number2
	assignment = number1
	number1 += number2

	number1 -= number2
	number1 *= number2
	number1 /= number2
	number1 %= number2

	return assignment
}
