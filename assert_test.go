package assert

import (
	"testing"
)

func test(t *testing.T, expected bool, message string) {
	if !expected {
		t.Error(message)
	}
}

func testValue(t *testing.T, a, b, c, array interface{}, desc string) {
	test(t, Equal(a, b), desc+" values are not equal, must be equal")
	test(t, LessEqual(a, b), desc+" value A is not less or equal than value B, must be less or equal")
	test(t, GreaterEqual(a, b), desc+" value A is not greater or eqail than value B, must be greater or equal")
	test(t, !Equal(a, c), desc+" values are equal, must be not equal")
	test(t, NotEqual(a, c), desc+" values are equal, must be not equal")
	test(t, LessThan(a, c), desc+" value A is not less than value B, must be less")
	test(t, LessEqual(a, c), desc+" value A is not less or equal than value B, must be less or equal")
	test(t, GreaterThan(c, a), desc+" value B is not greater than value A, must be greater")
	test(t, GreaterEqual(c, a), desc+" value B is not greater or eqail than value A, must be greater or equal")
	test(t, In(a, array), desc+" value A is not in array, must be in array")
	test(t, NotIn(c, array), desc+" value B is in array, must be not in array")
	test(t, !NotIn(a, array), desc+" value A is not in array, must be in array")
}

func TestAsserts(t *testing.T) {
	test(t, Equal(nil, nil), "Nil values are not equal, must be equal")
	float64A := 2.24
	float64B := 2.24
	float64C := 8.0
	testValue(t, float64A, float64B, float64C, []interface{}{2.22, 2.23, 2.24}, "Float64")
	testValue(t, float64A, float64B, float64C, []float64{2.22, 2.23, 2.24}, "Float64")
	var float32A float32 = 1.24
	var float32B float32 = 1.24
	var float32C float32 = 4
	testValue(t, float32A, float32B, float32C, []float32{1.22, 1.23, 1.24}, "Float32")
	intA := 3
	intB := 3
	intC := 4
	testValue(t, intA, intB, intC, []interface{}{1, 2, 3}, "Int")
	testValue(t, intA, intB, intC, []int{1, 2, 3}, "Int")
	var int8A int8 = 3
	var int8B int8 = 3
	var int8C int8 = 4
	testValue(t, int8A, int8B, int8C, []int8{1, 2, 3}, "Int8")
	var int16A int16 = 3
	var int16B int16 = 3
	var int16C int16 = 4
	testValue(t, int16A, int16B, int16C, []int16{1, 2, 3}, "Int16")
	var int32A int32 = 3
	var int32B int32 = 3
	var int32C int32 = 4
	testValue(t, int32A, int32B, int32C, []int32{1, 2, 3}, "Int32")
	var int64A int64 = 3
	var int64B int64 = 3
	var int64C int64 = 4
	testValue(t, int64A, int64B, int64C, []int64{1, 2, 3}, "Int64")
	var uintA uint = 3
	var uintB uint = 3
	var uintC uint = 4
	testValue(t, uintA, uintB, uintC, []uint{1, 2, 3}, "Uint")
	var uint8A uint8 = 3
	var uint8B uint8 = 3
	var uint8C uint8 = 4
	testValue(t, uint8A, uint8B, uint8C, []uint8{1, 2, 3}, "Uint8")
	var uint16A uint16 = 3
	var uint16B uint16 = 3
	var uint16C uint16 = 4
	testValue(t, uint16A, uint16B, uint16C, []uint16{1, 2, 3}, "Uint16")
	var uint32A uint32 = 3
	var uint32B uint32 = 3
	var uint32C uint32 = 4
	testValue(t, uint32A, uint32B, uint32C, []uint32{1, 2, 3}, "Uint32")
	var uint64A uint64 = 3
	var uint64B uint64 = 3
	var uint64C uint64 = 4
	testValue(t, uint64A, uint64B, uint64C, []uint64{1, 2, 3}, "Uint64")
	var mixFloat32 float32 = 3.0
	var mixInt8 int8 = 3
	var mixUint64 uint64 = 4
	testValue(t, mixFloat32, mixInt8, mixUint64, []int32{1, 2, 3}, "Uint64")
	stringA := "string"
	stringB := "string"
	stringC := "string_b"
	testValue(t, stringA, stringB, stringC, []interface{}{"str", "string", "string_a"}, "String")
	testValue(t, stringA, stringB, stringC, []string{"str", "string", "string_a"}, "String")
	boolA := true
	boolB := true
	test(t, Equal(boolA, boolB), "Bool values are not equal, must be equal")
	boolB = false
	test(t, NotEqual(boolA, boolB), "Bool values are equal, must be not equal")
}
