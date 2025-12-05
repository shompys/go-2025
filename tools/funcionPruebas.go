package tools

import "errors"

var ErrIndexOutOfRange = errors.New("index out of range")

func ElementAtIndex(slice []int, index int) (result int, err error) {
	if index < 0 || index >= len(slice) {
		return 0, ErrIndexOutOfRange
	}

	return slice[index], nil
}

func Sum(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}
