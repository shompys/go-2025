package fmt_example

import (
	"fmt"
)

func Fmt_example() string {
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
	fmt.Printf("2 decimal: %.2f\n", pi)   // 3.14 el 2 indica que queremos mostrar 2 number despues del .
	fmt.Printf("Bool:      %t\n", active) // true
	fmt.Printf("Tipo:      %T\n", name)   // string retorna el tipo de  dato
	fmt.Printf("Cualquier: %v\n", age)

	// === FLAGS DE FORMATO ===
	type Persona struct {
		Nombre string
		Edad   int
	}
	p := Persona{"Juan", 25}
	nums := []int{1, 2, 3}

	// %v vs %+v vs %#v
	fmt.Println("\n-- Variantes de %v --")
	fmt.Printf("%%v  struct:  %v\n", p)   // {Juan 25}
	fmt.Printf("%%+v struct:  %+v\n", p)  // {Nombre:Juan Edad:25} ← muestra nombres de campos
	fmt.Printf("%%#v struct:  %#v\n", p)  // fmt_example.Persona{Nombre:"Juan", Edad:25} ← sintaxis Go
	fmt.Printf("%%#v slice:   %#v\n", nums) // []int{1, 2, 3}
	fmt.Printf("%%#v string:  %#v\n", name) // "Jonathan" ← con comillas

	// # con números (añade prefijos)
	fmt.Println("\n-- Flag # con números --")
	n := 255
	fmt.Printf("%%x  hex:     %x\n", n)   // ff
	fmt.Printf("%%#x hex:     %#x\n", n)  // 0xff ← con prefijo 0x
	fmt.Printf("%%o  octal:   %o\n", n)   // 377
	fmt.Printf("%%#o octal:   %#o\n", n)  // 0377 ← con prefijo 0
	fmt.Printf("%%b  binario: %b\n", n)   // 11111111
	fmt.Printf("%%#b binario: %#b\n", n)  // 0b11111111 ← con prefijo 0b

	// + con números (muestra signo)
	fmt.Println("\n-- Flag + con números --")
	pos := 42
	neg := -42
	fmt.Printf("%%d:  %d, %d\n", pos, neg)   // 42, -42
	fmt.Printf("%%+d: %+d, %+d\n", pos, neg) // +42, -42 ← siempre muestra signo

	// Ancho y alineación
	fmt.Println("\n-- Ancho y alineación --")
	fmt.Printf("|%%10s|:  |%10s|\n", "hola")  // |      hola| ← derecha (10 chars)
	fmt.Printf("|%%-10s|: |%-10s|\n", "hola") // |hola      | ← izquierda con -
	fmt.Printf("|%%010d|: |%010d|\n", 42)     // |0000000042| ← rellena con 0

	fmt.Print("\nIngresa tu nombre y edad: ")
	fmt.Scan(&name, &age) // Lee separado por espacio o enter

	fmt.Printf("Hola %s, tienes %d años\n", name, age)

	var day, month, year int

	fmt.Print("Ingresa fecha (dd/mm/yyyy): ")
	fmt.Scanf("%d/%d/%d", &day, &month, &year)

	fmt.Printf("Día: %d, Mes: %d, Año: %d\n", day, month, year)

	//para retornar strings
	str := fmt.Sprintf("strings %d", 5)
	return str

}
