// Copyright 2014 Igor Dolzhikov. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

/*
Package assert 0.1.0

This package helps to compare values of undefined types like interface{}.
Convenient using with JSON values (json.Number/float64, string, bool, nil).
The package is intended not only for testing purposes.
The package allows compare ambiguous values between each other.

Example 1:

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

A simple Assert package for interface/json values
*/
package assert

import (
	"encoding/json"
)

// Equal checks if values equal to each other
func Equal(left, right interface{}) bool {
	return isCompareTrue(equal, left, right)
}

// NotEqual checks if values not equal to each other
func NotEqual(left, right interface{}) bool {
	return isCompareTrue(notEqual, left, right)
}

// LessThan checks if value "left" less than value "right"
func LessThan(left, right interface{}) bool {
	return isCompareTrue(lessThan, left, right)
}

// GreaterThan checks if value "left" greater than value "right"
func GreaterThan(left, right interface{}) bool {
	return isCompareTrue(greaterThan, left, right)
}

// LessEqual checks if value "left" less than value "right" or equal to it
func LessEqual(left, right interface{}) bool {
	return isCompareTrue(lessEqual, left, right)
}

// GreaterEqual checks if value "left" greater than value "right" or equal to it
func GreaterEqual(left, right interface{}) bool {
	return isCompareTrue(greaterEqual, left, right)
}

// In checks if array value "right" contains value "left"
func In(left, right interface{}) bool {
	return isCompareTrue(in, left, right)
}

// In checks if array value "right" not contains value "left"
func NotIn(left, right interface{}) bool {
	return isCompareTrue(notIn, left, right)
}

const (
	equal = iota
	notEqual
	lessThan
	greaterThan
	lessEqual
	greaterEqual
	in
	notIn
)

func isCompareTrue(comparison uint8, valueLeft, valueRight interface{}) bool {
	if valueLeft == nil && valueRight == nil && comparison == equal {
		return true
	}
	switch vLeft := valueLeft.(type) {
	case json.Number:
		if left, err := vLeft.Float64(); err == nil {
			switch vRight := valueRight.(type) {
			case json.Number:
				if right, err := vRight.Float64(); err == nil {
					switch comparison {
					case equal:
						if left == right {
							return true
						}
					case notEqual:
						if left != right {
							return true
						}
					case lessThan:
						if left < right {
							return true
						}
					case greaterThan:
						if left > right {
							return true
						}
					case lessEqual:
						if left <= right {
							return true
						}
					case greaterEqual:
						if left >= right {
							return true
						}
					}
				}
			case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
				right := float(valueRight)
				switch comparison {
				case equal:
					if left == right {
						return true
					}
				case notEqual:
					if left != right {
						return true
					}
				case lessThan:
					if left < right {
						return true
					}
				case greaterThan:
					if left > right {
						return true
					}
				case lessEqual:
					if left <= right {
						return true
					}
				case greaterEqual:
					if left >= right {
						return true
					}
				}
			case []interface{}:
				switch comparison {
				case in:
					for _, v := range vRight {
						switch value := v.(type) {
						case json.Number:
							if item, err := value.Float64(); err == nil && item == left {
								return true
							}
						case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
							item := float(v)
							if item == left {
								return true
							}
						}
					}
				case notIn:
					for _, v := range vRight {
						switch value := v.(type) {
						case json.Number:
							if item, err := value.Float64(); err == nil && item == left {
								return false
							}
						case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
							item := float(v)
							if item == left {
								return false
							}
						}
					}
					return true
				}
			case []json.Number:
				switch comparison {
				case in:
					for _, value := range vRight {
						if item, err := value.Float64(); err == nil && item == left {
							return true
						}
					}
				case notIn:
					for _, value := range vRight {
						if item, err := value.Float64(); err == nil && item == left {
							return false
						}
					}
					return true
				}
			}
		}
	case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		left := float(valueLeft)
		switch vRight := valueRight.(type) {
		case json.Number:
			if right, err := vRight.Float64(); err == nil {
				switch comparison {
				case equal:
					if left == right {
						return true
					}
				case notEqual:
					if left != right {
						return true
					}
				case lessThan:
					if left < right {
						return true
					}
				case greaterThan:
					if left > right {
						return true
					}
				case lessEqual:
					if left <= right {
						return true
					}
				case greaterEqual:
					if left >= right {
						return true
					}
				}
			}
		case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			right := float(valueRight)
			switch comparison {
			case equal:
				if left == right {
					return true
				}
			case notEqual:
				if left != right {
					return true
				}
			case lessThan:
				if left < right {
					return true
				}
			case greaterThan:
				if left > right {
					return true
				}
			case lessEqual:
				if left <= right {
					return true
				}
			case greaterEqual:
				if left >= right {
					return true
				}
			}
		case []interface{}:
			switch comparison {
			case in:
				for _, v := range vRight {
					switch value := v.(type) {
					case json.Number:
						if item, err := value.Float64(); err == nil && item == left {
							return true
						}
					case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
						item := float(v)
						if item == left {
							return true
						}
					}
				}
			case notIn:
				for _, v := range vRight {
					switch value := v.(type) {
					case json.Number:
						if item, err := value.Float64(); err == nil && item == left {
							return false
						}
					case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
						item := float(v)
						if item == left {
							return false
						}
					}
				}
				return true
			}
		case []float32, []float64, []int, []uint, []int8, []uint8, []int16, []uint16, []int32, []uint32, []int64, []uint64:
			vr := floatArray(valueRight)
			switch comparison {
			case in:
				for _, v := range vr {
					item := float(v)
					if item == left {
						return true
					}
				}
			case notIn:
				for _, v := range vr {
					item := float(v)
					if item == left {
						return false
					}
				}
				return true
			}
		case []json.Number:
			switch comparison {
			case in:
				for _, value := range vRight {
					if item, err := value.Float64(); err == nil && item == left {
						return true
					}
				}
			case notIn:
				for _, value := range vRight {
					if item, err := value.Float64(); err == nil && item == left {
						return false
					}
				}
				return true
			}
		}
	case string:
		left := vLeft
		switch right := valueRight.(type) {
		case string:
			switch comparison {
			case equal:
				if left == right {
					return true
				}
			case notEqual:
				if left != right {
					return true
				}
			case lessThan:
				if left < right {
					return true
				}
			case greaterThan:
				if left > right {
					return true
				}
			case lessEqual:
				if left <= right {
					return true
				}
			case greaterEqual:
				if left >= right {
					return true
				}
			}
		case []interface{}:
			switch comparison {
			case in:
				for _, v := range right {
					if item, ok := v.(string); ok && item == left {
						return true
					}
				}
			case notIn:
				for _, v := range right {
					if item, ok := v.(string); ok && item == left {
						return false
					}
				}
				return true
			}
		case []string:
			switch comparison {
			case in:
				for _, v := range right {
					if v == left {
						return true
					}
				}
			case notIn:
				for _, v := range right {
					if v == left {
						return false
					}
				}
				return true
			}
		}
	case bool:
		left := vLeft
		switch right := valueRight.(type) {
		case bool:
			switch comparison {
			case equal:
				if left == right {
					return true
				}
			case notEqual:
				if left != right {
					return true
				}
			}
		}
	}

	return false
}

func float(value interface{}) float64 {

	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return float64(v)
	case int:
		return float64(v)
	case uint:
		return float64(v)
	case int8:
		return float64(v)
	case uint8:
		return float64(v)
	case int16:
		return float64(v)
	case uint16:
		return float64(v)
	case int32:
		return float64(v)
	case uint32:
		return float64(v)
	case int64:
		return float64(v)
	case uint64:
		return float64(v)
	}

	panic("never happen in that implementation")
}

func floatArray(value interface{}) []float64 {
	var result []float64
	switch v := value.(type) {
	case []float32:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []float64:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []int:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []uint:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []int8:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []uint8:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []int16:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []uint16:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []int32:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []uint32:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []int64:
		for _, item := range v {
			result = append(result, float64(item))
		}
	case []uint64:
		for _, item := range v {
			result = append(result, float64(item))
		}
	}

	return result
}
