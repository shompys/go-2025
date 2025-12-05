package operadores

func Logicos() (and, or, not bool) {
	var number1 int = 20
	var number2 int = 10

	and = number1 > 0 && number2 > 0
	or = number1 > 0 || number2 > 0
	not = !(number1 > 0)
	return and, or, not
}
