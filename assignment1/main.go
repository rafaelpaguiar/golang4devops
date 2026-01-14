package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Assignment1 struct {
	Page         string             `json:"page"`
	Words        []string           `json:"words"`
	Percenteges  map[string]float32 `json:"percentages"`
	Special      []string           `json:"special"`
	ExtraSpecial []any              `json:"extraSpecial"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/assignment1")
	if err != nil {
		fmt.Printf("Error: %s\n", err)

	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	var assignment1 Assignment1

	err = json.Unmarshal(body, &assignment1)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("%+v\n", assignment1)

}
