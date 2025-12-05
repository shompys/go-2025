package tiposdedatos

func Integer() (int, int8, int16, int32, int64) {
	var number int = 20     // -9223372036854775808 al 9223372036854775807
	var number8 int8 = 20   // -128 al 127
	var number16 int16 = 20 // -32768 al 32767
	var number32 int32 = 20 // -2147483648 al 2147483647 rune es igual a int32
	var number64 int64 = 20 // -9223372036854775808 al 9223372036854775807

	return number, number8, number16, number32, number64
}
