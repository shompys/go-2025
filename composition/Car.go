package composition

import "fmt"

type Engine struct {
	Horsepower int
	Type       string
}

func (e Engine) ruido() {
	fmt.Println("brr brr brr!!")
}

type Chassis struct {
	Material string
}

type Bodywork struct {
	Color string
}

type Vehicle interface {
	Start()
	DisplayInfo()
}

//entidades

type Car struct {
	Engine
	Chassis
	Bodywork
	NumberOfDoors int
}

func (c Car) Start() {
	c.ruido()
	fmt.Println("The car is starting with a roar!")
}
func (c Car) DisplayInfo() {
	fmt.Printf("This car has %d doors, a %s body, a %s chassis, and a %d horsepower %s engine. \n", c.NumberOfDoors, c.Color, c.Material, c.Horsepower, c.Type)
}

type Motorcycle struct {
	Engine
	Chassis
	Bodywork
}

func (m Motorcycle) Start() {
	m.ruido()
	fmt.Println("The motorcycle is starting with a vroom!")
}
func (m Motorcycle) DisplayInfo() {
	fmt.Printf("This motorcycle has a %s body, a %s chassis, and a %d horsepower %s engine. \n", m.Bodywork.Color, m.Chassis.Material, m.Engine.Horsepower, m.Engine.Type)
}

// es como un overwrite sobre el metodo String que se ejecutal al consolear la clase
func (m Motorcycle) String() string {
	return fmt.Sprintln("MOTORCYCLE")
}
