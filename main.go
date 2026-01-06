package main

import (
	"fmt"

	"github.com/shompys/go-2025/contexto"
	"github.com/shompys/go-2025/errores"
	"github.com/shompys/go-2025/estructuras"
	"github.com/shompys/go-2025/functions"
	"github.com/shompys/go-2025/goroutines"
	"github.com/shompys/go-2025/maps"
	"github.com/shompys/go-2025/panics"
	"github.com/shompys/go-2025/paquetes"
	"github.com/shompys/go-2025/punteros"
	"github.com/shompys/go-2025/slices"
	typeassertion "github.com/shompys/go-2025/type_assertion"

	"github.com/shompys/go-2025/constantes"
	"github.com/shompys/go-2025/variables"
)

func examples() {
	fmt.Println(variables.Variables())
	fmt.Println(constantes.Constantes())
	fmt.Println(slices.Slices())
	fmt.Println(maps.Maps())
	fmt.Println(functions.Functions())
	c := estructuras.Circle{}
	c.SetRadius(1)
	fmt.Println(c.Area())
	punteros.Punteros()
	typeassertion.Typeassertion()
	fmt.Println(errores.Errores())
	panics.Panics()
	fmt.Println(paquetes.TimePackage())
	fmt.Println(paquetes.OsPackage())
	paquetes.LogPackage()
	fmt.Println(paquetes.StrconvPackage())
	fmt.Println(paquetes.NetPackage())
	goroutines.Goroutines()
	contexto.Contexto()
	goroutines.Example()
	goroutines.Atomic()
	goroutines.ChannelRanger()
}

func main() {
	
}
