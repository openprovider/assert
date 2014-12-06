package main

import (
	"encoding/json"
	"fmt"

	"github.com/takama/assert"
)

func main() {
	exampleJSON := `
	{
		"personA": {
			"name": "John",
			"age": 22,
			"height": 178,
			"weight": 82.59,
			"married": true,
			"hobby": ["fishing", "dancing", "cooking"]
		},
		"personB": {
			"name": "Jack",
			"age": 22,
			"height": 170.22,
			"weight": 65,
			"married": false,
			"hobby": ["photography", "singing", "cooking"]
		}
  	}`

	persons := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(exampleJSON), &persons); err == nil {
		pesronA := persons["personA"].(map[string]interface{})
		pesronB := persons["personB"].(map[string]interface{})
		fmt.Println("Comparing of a personal details:", pesronA["name"], "and", pesronB["name"])
		if assert.Equal(pesronA["age"], pesronB["age"]) {
			fmt.Println("Persons of the same age")
		}
		if assert.GreaterThan(pesronA["height"], pesronB["height"]) {
			fmt.Println(pesronA["name"], "higher than", pesronB["name"])
		}
		if assert.LessThan(pesronB["weight"], pesronA["weight"]) {
			fmt.Println(pesronB["name"], "weighs less than", pesronA["name"])
		}
		if assert.NotEqual(pesronA["married"], pesronB["married"]) {
			fmt.Println("Persons have a different marital status")
		}
		if assert.In("cooking", pesronA["hobby"]) && assert.In("cooking", pesronB["hobby"]) {
			fmt.Println("Both persons like to cook")
		}
		if assert.In("dancing", pesronA["hobby"]) && assert.NotIn("dancing", pesronB["hobby"]) {
			fmt.Println("Only", pesronA["name"], "likes to dance")
		}
	}
}
