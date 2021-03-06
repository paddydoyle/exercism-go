// Determine if a triangle is equilateral, isosceles, or scalene.
package triangle

import (
	"math"
	"sort"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// All sides zero length?
	if (a + b + c) == 0 {
		return NaT
	}

	// Use a list to find the largest value
	sidesList := []float64{a, b, c}
	sort.Float64s(sidesList)

	// Any side length a NaN? It's always put first in a sorted list.
	if math.IsNaN(sidesList[0]) {
		return NaT
	}

	// Any side length a Inf? First or last in a sorted list.
	if math.IsInf(sidesList[0], -1) || math.IsInf(sidesList[2], 1) {
		return NaT
	}

	// Triangle inequality failure?
	if sidesList[0]+sidesList[1] < sidesList[2] {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
