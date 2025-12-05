package main

import (
	"fmt"

	"github.com/shompys/go-2025/estructuras"
	"github.com/shompys/go-2025/functions"
	"github.com/shompys/go-2025/maps"
	"github.com/shompys/go-2025/slices"

	"github.com/shompys/go-2025/constantes"
	"github.com/shompys/go-2025/variables"
)

func main() {
	fmt.Println(variables.Variables())
	fmt.Println(constantes.Constantes())
	fmt.Println(slices.Slices())
	fmt.Println(maps.Maps())
	fmt.Println(functions.Functions())
	c := estructuras.Circle{}
	c.SetRadius(1)
	fmt.Println(c.Area())
}
