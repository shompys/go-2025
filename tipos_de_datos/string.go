package tiposdedatos

func String() string {
	// codificación UTF-8
	// cada caracter ocupa 1 byte
	// existen caracteres especiales que ocupan mas bytes
	var name string = "Jonathan\"\n\t\r"
	return name
}
