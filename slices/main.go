package slices

func Slices() ([]string, []int, []string, []string) {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fruits := []string{"apple", "banana", "cherry"}

	var vegetables []string

	vegetables = append(vegetables, "tomato", "onion", "lettuce")

	longitud := 0
	capacidad := 3
	names := make([]string, longitud, capacidad)

	return fruits, numbers, vegetables, names[0:2]
}
