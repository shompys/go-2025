package errores

import (
	"errors"
	"fmt"
)

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func Errores() (error, error, error) {
	//son 3 tipos de formas de errores todos retornan el type error
	err := errors.New("error message package")
	var ErrCustom error = &CustomError{"Error Message"}
	errFmt := fmt.Errorf("Error message: %s, %w", "error", ErrCustom) //añadir un mensaje de error anterior %w
	fmt.Println("->", errors.Is(errFmt, ErrCustom))                   //revisa si errCustom se encuentra en la cadena de errores de errFmt return bool

	// -- si el error es del tipo los retorna en mi erro, entiendo que si tuviera varios diferentes los puedo bifurcar
	var erro *CustomError

	if errors.As(errFmt, &erro) {
		fmt.Println("->->", erro)
	}
	return ErrCustom, err, errFmt
}
