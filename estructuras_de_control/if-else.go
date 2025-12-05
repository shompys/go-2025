package estructurasdecontrol

func IfElse() string {

	condition := 1

	// if condition > 0 {
	// 	return "The condition is true"
	// }

	// if condition > 0 {
	// 	return "The condition is true"
	// } else {
	// 	return "The condition is false"
	// }

	if condition > 5 {
		return "The condition is true"
	} else if condition > 6 {
		return "The condition is true"
	} else {
		return "condition is not equal to 5 or 6"
	}
}
