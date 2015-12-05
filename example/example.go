package main

import (
	"encoding/json"
	"fmt"

	"github.com/openprovider/assert"
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
		personA := persons["personA"].(map[string]interface{})
		personB := persons["personB"].(map[string]interface{})
		fmt.Println(
			"Comparing of a personal details:",
			personA["name"], "and", personB["name"],
		)
		if assert.Equal(personA["age"], personB["age"]) {
			fmt.Println("Persons of the same age")
		}
		if assert.GreaterThan(personA["height"], personB["height"]) {
			fmt.Println(personA["name"], "higher than", personB["name"])
		}
		if assert.LessThan(personB["weight"], personA["weight"]) {
			fmt.Println(personB["name"], "weighs less than", personA["name"])
		}
		if assert.NotEqual(personA["married"], personB["married"]) {
			fmt.Println("Persons have a different marital status")
		}
		if assert.In("cooking", personA["hobby"]) &&
			assert.In("cooking", personB["hobby"]) {
			fmt.Println("Both persons like to cook")
		}
		if assert.In("dancing", personA["hobby"]) &&
			assert.NotIn("dancing", personB["hobby"]) {
			fmt.Println("Only", personA["name"], "likes to dance")
		}
	}
}
