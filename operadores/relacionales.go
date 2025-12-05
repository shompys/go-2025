package operadores

func Relacionales() (equal, notEqual, greater, less, greaterOrEqual, lessOrEqual bool) {
	var number1 int = 20
	var number2 int = 10

	equal = number1 == number2
	notEqual = number1 != number2
	greater = number1 > number2
	less = number1 < number2
	greaterOrEqual = number1 >= number2
	lessOrEqual = number1 <= number2
	return equal, notEqual, greater, less, greaterOrEqual, lessOrEqual
}
