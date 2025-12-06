package typeassertion

import "fmt"

func Typeassertion() {
	var i interface{} = "hello"
	s, ok := i.(float32)
	fmt.Println(s, ok)
}
