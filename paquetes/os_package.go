package paquetes

import (
	"fmt"
	"os"
	"path/filepath"
)

func OsPackage() string {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Sprintln("Error al obtener el directorio: ", err)
	}
	dir_asset := filepath.Join(dir, "paquetes", "test.txt") // compatibilidad de OS

	file, err := os.Open(dir_asset) // abre un archivo en el directorio actual

	if err != nil {
		file, err = os.Create(dir_asset)
		if err != nil {
			return fmt.Sprintln("Error al crear el archivo: ", err)
		}

		file.WriteString("hola 5")
		file.Seek(0, 0) // mover 0 bytes, desde 0=inicio 1= actual, 2=final
	}
	defer file.Close()

	data := make([]byte, 100)
	c, err := file.Read(data)
	if err != nil {
		return fmt.Sprintln("Error al leer el archivo: ", err)
		// os.Exit(1) si queremos terminar

	}
	fmt.Println(c, "bytes leidos", string(data[:c]))

	user := os.Getenv("USER")
	os.Setenv("VACIO", "")
	vacio := os.Getenv("VACIO")
	notExist, ok := os.LookupEnv("NOT_EXIST")

	stringFromEnv := os.ExpandEnv("user: ${USER} $USER")

	return fmt.Sprintln("OsPackage", user, dir, vacio, notExist, ok, os.Getppid(), stringFromEnv)
}
