package estructuras

import "encoding/json"

type Human struct {
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Age        uint8  `json:"algo"`
	NumberZero uint8  `json:"noVoy,omitempty"` // omite si no hay valor
	Ignorado   string `json:"-"`               // omite siempre el campo
}

func Etiquetas() (Human, []byte, string) {
	human := Human{"jonathan", "molina", 34, 0, "ignorado"}
	humanAsJSON, _ := json.Marshal(human) // marshal transforma mi objeto GO en un buffer de bytes, si tiene una regla json:"algo" va a reemplazar la key por algo :D
	jsonString := string(humanAsJSON)

	return human, humanAsJSON, jsonString
}
