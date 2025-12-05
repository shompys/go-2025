package maps

import "fmt"

func Maps() map[string]int {
	person := map[string]int{
		"John": 20,
		"Jane": 21,
		"Jim":  22,
	}
	myMap := make(map[string]int)
	//var myMapFail map[string]int // no crear asi
	fmt.Println("myMap: ", len(myMap))
	person["jonathan"] = 34
	delete(person, "jonathan")
	return person
}
