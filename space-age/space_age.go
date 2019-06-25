// Given an age in seconds, calculate how old someone would be on
// various planets.
package space

type Planet string

const EarthPeriodSeconds = 31557600

// The following are in Earth-Years
var orbitalPeriods = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Given an age in seconds, calculate how old someone would be on
// various planets.
func Age(seconds float64, planet Planet) float64 {
	return seconds / EarthPeriodSeconds / orbitalPeriods[planet]
}
