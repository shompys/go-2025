package paquetes

import (
	"encoding/json"
	"fmt"

	"net/http"
)

func NetPackage() string {
	// getReqExample("https://pokeapi.co/api/v2/pokemon/ditto")
	get("https://pokeapi.co/api/v2/pokemon/ditto")
	return "NetPackage"

}

// func getReqExample(url string) ([]byte, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

//		defer resp.Body.Close()
//		return body, nil
//	}
type Pokemon struct {
	Name   string `json:"name"`
	Height uint32 `json:"height"`
}

func get(url string) (*Pokemon, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// respBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }
	var pokemon Pokemon
	json.NewDecoder(resp.Body).Decode(&pokemon)
	fmt.Println("-> ", pokemon)
	return &pokemon, nil
}
