package composition

func Composition() {
	car := Car{
		Engine:        Engine{Horsepower: 250, Type: "V4"},
		Chassis:       Chassis{Material: "Steel"},
		Bodywork:      Bodywork{Color: "Red"},
		NumberOfDoors: 4,
	}
	motorcycle := Motorcycle{
		Engine:   Engine{Horsepower: 150, Type: "V2"},
		Chassis:  Chassis{Material: "Aluminum"},
		Bodywork: Bodywork{Color: "Black"},
	}
	car.Start()
	//esto solo sucede porque estoy dentro del paquete en sí, de lo contrario ruido es inaccesible
	motorcycle.ruido()
}
