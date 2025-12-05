package tiposdedatos

func UnsignedInteger() (uint, uint8, uint16, uint32, uint64) {
	var number uint = 20     // 0 al 18446744073709551615
	var number8 uint8 = 20   // 0 al 255
	var number16 uint16 = 20 // 0 al 65535
	var number32 uint32 = 20 // 0 al 4294967295
	var number64 uint64 = 20 // 0 al 18446744073709551615

	return number, number8, number16, number32, number64
}
