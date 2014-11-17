// Package ints provides some utilities for floats.
package floats // import "hkjn.me/floats"

import (
	"log"
	"strconv"
)

// ParseStrings parses multiple strings as float64.
func ParseStrings(in ...string) ([]float64, error) {
	result := make([]float64, len(in))
	for i, s := range in {
		r, err := Parse(s)
		if err != nil {
			return []float64{}, err
		}
		result[i] = r
	}
	return result, nil
}

// Parse returns the string s as a float.
func Parse(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// MustParse returns the string s parsed as a float, or panics.
func MustParse(s string) float64 {
	v, err := Parse(s)
	if err != nil {
		log.Fatalf("bad string %q: %v", s, err)
	}
	return v
}

// Round returns the nearest integer value to specified value v.
func Round(v float64) int {
	x := strconv.FormatFloat(v, 'f', 0, 64)
	r, err := strconv.Atoi(x)
	if err != nil {
		// Shouldn't happen.
		log.Fatalf("failed to parse float %f with 0 decimal values as int: %err", x, err)
	}
	return r
}
