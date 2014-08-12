// Package ints provides some utilities for floats.
package floats

import "strconv"

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

// Parse is a simple wrapper around strconv.ParseFloat.
func Parse(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
