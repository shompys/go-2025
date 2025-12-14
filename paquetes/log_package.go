package paquetes

import (
	"log"
	"os"
)

func LogPackage() {
	log.Println("log package")
	// log.Fatal("log fatal package")
	// log.Panic("log panic package")
	file, err := os.Create("paquetes/myLogs.log")
	if err != nil {
		log.Println("Error al crear el archivo: ", err)
	}
	defer file.Close()
	// ejecutamos log pero genera el file
	l := log.New(file, "logger: ", log.LstdFlags|log.Llongfile)
	// l := log.New(os.Stdout, "logger: ", log.LstdFlags|log.Llongfile) // para imprimir en consola
	l.Println("log package con log.New")
}
