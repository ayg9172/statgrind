package main

import "math"

const tolerance = 0.005

// TODO: Implement this floatEquals everywhere we compare floats with ==
func FloatEquals(a, b float64) bool {
	return math.Abs(a-b) < tolerance
}

func CreateIdFactory() func() int {
	idCounter := 0

	return func() int {
		idCounter++
		return idCounter
	}
}

func LastUnit(arr []Unit) Unit {
	if len(arr) == 0 {
		panic("LastUnit() called on empty slice!")
	}
	return arr[len(arr)-1]
}
