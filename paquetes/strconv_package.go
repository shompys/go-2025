package paquetes

import (
	"fmt"
	"strconv"
)

func StrconvPackage() string {
	number := 10
	str := strconv.Itoa(number)
	//number := strconv.ParseInt("10", 10, 64)
	// strFloat := strconv.FormatFloat(3.14159, 'f', 2, 64)
	// bool := strconv.FormatBool(true)
	// hex := strconv.FormatInt(10, 16)
	// oct := strconv.FormatInt(10, 8)
	// bin := strconv.FormatInt(10, 2)
	// base64 := strconv.FormatInt(10, 64)
	// base64url := strconv.FormatInt(10, 64)
	// base64url := strconv.FormatInt(10, 64)

	return fmt.Sprintf("StrconvPackage: %s", str)
}
