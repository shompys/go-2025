package constantes

// Deben asignarse al ser declaradas
func Constantes() (string, string, int, float64) {
	const status = "ok"

	const (
		product  = "Course"
		quantity = 20
		price    = 40.50
	)
	return status, product, quantity, price
}
