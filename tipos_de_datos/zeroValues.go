package tiposdedatos

import "fmt"

func ZeroValues() {
	// Booleano
	var b bool // false

	// Enteros
	var i int     // 0
	var i8 int8   // 0
	var i16 int16 // 0
	var i32 int32 // 0
	var i64 int64 // 0

	// Enteros sin signo
	var u uint     // 0
	var u8 uint8   // 0
	var u16 uint16 // 0
	var u32 uint32 // 0
	var u64 uint64 // 0

	// Flotantes
	var f32 float32 // 0.0
	var f64 float64 // 0.0

	// Complejos
	var c64 complex64   // (0+0i)
	var c128 complex128 // (0+0i)

	// String
	var s string // "" (string vacío)

	// Byte y Rune (alias)
	var by byte // 0 (alias de uint8)
	var r rune  // 0 (alias de int32)

	// Puntero
	var p *int // nil

	// Función
	var fn func() // nil

	// Interface
	var iface interface{} // nil

	// Slice
	var sl []int // nil

	// Map
	var m map[string]int // nil

	// Channel
	var ch chan int // nil

	// Struct
	var st struct {
		Name string
		Age  int
	} // {"", 0}

	// Array (cada elemento tiene su valor cero)
	var arr [3]int // [0, 0, 0]

	fmt.Println("=== ZERO VALUES EN GO ===")
	fmt.Println()

	fmt.Println("-- Booleano --")
	fmt.Printf("bool:        %v\n", b)
	fmt.Println()

	fmt.Println("-- Enteros con signo --")
	fmt.Printf("int:         %v\n", i)
	fmt.Printf("int8:        %v\n", i8)
	fmt.Printf("int16:       %v\n", i16)
	fmt.Printf("int32:       %v\n", i32)
	fmt.Printf("int64:       %v\n", i64)
	fmt.Println()

	fmt.Println("-- Enteros sin signo --")
	fmt.Printf("uint:        %v\n", u)
	fmt.Printf("uint8:       %v\n", u8)
	fmt.Printf("uint16:      %v\n", u16)
	fmt.Printf("uint32:      %v\n", u32)
	fmt.Printf("uint64:      %v\n", u64)
	fmt.Println()

	fmt.Println("-- Flotantes --")
	fmt.Printf("float32:     %v\n", f32)
	fmt.Printf("float64:     %v\n", f64)
	fmt.Println()

	fmt.Println("-- Complejos --")
	fmt.Printf("complex64:   %v\n", c64)
	fmt.Printf("complex128:  %v\n", c128)
	fmt.Println()

	fmt.Println("-- String --")
	fmt.Printf("string:      %q (vacío)\n", s)
	fmt.Println()

	fmt.Println("-- Byte y Rune --")
	fmt.Printf("byte:        %v (alias uint8)\n", by)
	fmt.Printf("rune:        %v (alias int32)\n", r)
	fmt.Println()

	fmt.Println("-- Tipos de referencia (nil) --")
	fmt.Printf("*int:        %v\n", p)
	fmt.Printf("func():      %v\n", fn)
	fmt.Printf("interface{}: %v\n", iface)
	fmt.Printf("[]int:       %v\n", sl)
	fmt.Printf("map:         %v\n", m)
	fmt.Printf("chan:        %v\n", ch)
	fmt.Println()

	fmt.Println("-- Struct --")
	fmt.Printf("struct:      %+v\n", st)
	fmt.Println()

	fmt.Println("-- Array --")
	fmt.Printf("[3]int:      %v\n", arr)
}

