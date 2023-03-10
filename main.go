package main

import "fmt"

// Can call whenever you are constraining something to int64 or float64
// replace the same union inside the function (see SumIntsOr... and SumNumbers below)
type Number interface {
	int64 | float64
}

func main() {
	//intialize the map for ints
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	//intialize the map for floats
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	//All return the same values, with progressively more generalized functions
	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type params inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// adds together the values of m
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds togehter the values of m
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of a map m. Supports both int64 and float64 as types for m
//
//	⬇️Means any type that can be compared with == and !=
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. Supports ints and floats
// Just replaces the Union with a variable that is a union (compared to above)
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
