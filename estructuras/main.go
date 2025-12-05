package estructuras

type Person struct {
	Name       string
	Gender     string
	Age        uint8
	Profession string
}

func Estructuras() (Person, Person, Person, uint8) {
	person := Person{"jonathan", "male", 25, "cadete"}
	person.Age = 30
	person2 := Person{
		Name:       "Jonathan",
		Gender:     "Male",
		Age:        34,
		Profession: "Software Engineer",
	}
	var person3 Person
	person3.Name = "Juan"
	person3.Age = 20
	person3.Gender = "Male"
	person3.Profession = "Delivery"

	return person, person2, person3, person.Age
}
