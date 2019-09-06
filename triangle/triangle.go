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
	Equ = iota // equilateral
	Iso = iota // isosceles
	Sca = iota // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	//var k Kind

	// All sides zero length?
	if (a + b + c) == 0 {
		return NaT
	}

	// Use a list to find the largest value
	sidesList := [3]float64{a, b, c}
	sort.Float64s(sidesList[:])

	// Any side length a NaN? It's always put first in a sorted list.
	if sidesList[0] == math.NaN() {
		return NaT
	}

	// Any side length a Inf?
	if sidesList[0] == math.Inf(-1) || sidesList[2] == math.Inf(1) {
		return NaT
	}

	// Triangle inequality failure?
	if sidesList[0]+sidesList[1] < sidesList[2] {
		return NaT
	}

	// Use a map to find how many unique lengths
	sidesMap := make(map[float64]bool)

	sidesMap[a] = true
	sidesMap[b] = true
	sidesMap[c] = true

	if len(sidesMap) == 1 {
		return Equ
	}

	if len(sidesMap) == 2 {
		return Iso
	}

	return Sca
}
