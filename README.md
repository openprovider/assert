Go Assert
=========

A simple Assert package for interface/json values have use in Go (golang)

[![Build Status](https://travis-ci.org/takama/assert.png?branch=master)](https://travis-ci.org/takama/assert)
[![GoDoc](https://godoc.org/github.com/takama/assert?status.svg)](https://godoc.org/github.com/takama/assert)

This package helps to compare values of undefined types like interface{}.
Convenient using with JSON values (json.Number/float64, string, bool, nil).
The package is intended not only for testing purposes.
The package allows compare ambiguous values between each other.

### Types and supported methods

| Type        | Equal/NotEqual | Less/Greater | In/NotIn |           JSON <-> Go          |
| ----------- |:--------------:|:------------:|:--------:|:------------------------------:|
| json.Number | X              | X            | X        | number <-> json.Number/float64 |
| float32     | X              | X            | X        | number <-> json.Number/float64 |
| float64     | X              | X            | X        | number <-> json.Number/float64 |
| int         | X              | X            | X        | number <-> json.Number/float64 |
| uint        | X              | X            | X        | number <-> json.Number/float64 |
| int8        | X              | X            | X        | number <-> json.Number/float64 |
| uint8       | X              | X            | X        | number <-> json.Number/float64 |
| int16       | X              | X            | X        | number <-> json.Number/float64 |
| uint16      | X              | X            | X        | number <-> json.Number/float64 |
| int32       | X              | X            | X        | number <-> json.Number/float64 |
| uint32      | X              | X            | X        | number <-> json.Number/float64 |
| int64       | X              | X            | X        | number <-> json.Number/float64 |
| uint64      | X              | X            | X        | number <-> json.Number/float64 |
| string      | X              | X            | X        | string <-> string              |
| bool        | X              |              |          | boolean <-> bool               |
| nil         | X              |              |          | null <-> nil                   |

### Example

```go
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
```

## Author

[Igor Dolzhikov](https://github.com/takama)

## Contributors

All the contributors are welcome. If you would like to be the contributor please accept some rules.
- The pull requests will be accepted only in "develop" branch
- All modifications or additions should be tested
- Sorry, I'll not accept code with any dependency, only standard library

Thank you for your understanding!

## License

[MIT Public License](https://github.com/takama/assert/blob/master/LICENSE)
