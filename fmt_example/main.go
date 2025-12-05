package fmt_example

import "fmt"

func Fmt_example() {
	fmt.Println("Hello, World!")

	name := "Jonathan"
	age := 25
	pi := 3.14159
	active := true

	fmt.Printf("String:    %s\n", name)   // Jonathan
	fmt.Printf("Entero:    %d\n", age)    // 25
	fmt.Printf("Binario:   %b\n", age)    // 11001
	fmt.Printf("Hex:       %x\n", age)    // 19
	fmt.Printf("Float:     %f\n", pi)     // 3.141590
	fmt.Printf("2 decimal: %.2f\n", pi)   // 3.14
	fmt.Printf("Bool:      %t\n", active) // true
	fmt.Printf("Tipo:      %T\n", name)   // string
	fmt.Printf("Cualquier: %v\n", age)

	fmt.Print("Ingresa tu nombre y edad: ")
	fmt.Scan(&name, &age) // Lee separado por espacio o enter

	fmt.Printf("Hola %s, tienes %d años\n", name, age)

	var day, month, year int

	fmt.Print("Ingresa fecha (dd/mm/yyyy): ")
	fmt.Scanf("%d/%d/%d", &day, &month, &year)

	fmt.Printf("Día: %d, Mes: %d, Año: %d\n", day, month, year)

}
