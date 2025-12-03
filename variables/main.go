package variables

func Variables() (string, int, int, string, float64, int, float64, bool) {

	var hours int
	hours = 20

	var number int = 20

	name := "jonathan"

	product, price := "Jean", 10.5

	var (
		quantity = 25
		prices   = 40.5
		stocked  = true
	)

	return name, hours, number, product, price, quantity, prices, stocked
}
