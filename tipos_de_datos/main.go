package tiposdedatos

func TiposDeDatos() (string, int, float64, bool) {
	// Nativos integer (con y sin signo), float, boolean, string
	var name string = "Jonathan"
	var number int = 20
	var price float64 = 10.5
	var stocked bool = true

	return name, number, price, stocked
}
