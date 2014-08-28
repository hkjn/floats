// floats_test.go; Tests for package floats.
package floats

import (
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		s    string
		want float64
	}{
		{"0.0", 0.0},
		{"0.1", 0.1},
		{"-0.1", -0.1},
	}
	for i, tt := range cases {
		got, err := Parse(tt.s)
		if err != nil {
			t.Errorf("[%d] Parse(%q) got error %q, want %f\n", i, tt.s, err, tt.want)
		} else if got != tt.want {
			t.Errorf("[%d] Parse(%q) => %v, want %f\n", i, tt.s, got, tt.want)
		}
		// TODO: test failure cases
	}
}

func TestRound(t *testing.T) {
	cases := []struct {
		v    float64
		want int
	}{
		{0.0, 0},
		{1.51, 2},
		{0.51, 1},
		{-0.4999999999999, 0},
		{-0.5000000000001, -1},
	}
	for i, tt := range cases {
		got, err := Round(tt.v)
		if err != nil {
			t.Errorf("[%d] Round(%f) got error %q, want %d\n", i, tt.v, err, tt.want)
		} else if got != tt.want {
			t.Errorf("[%d] Round(%f) => %d, want %d\n", i, tt.v, got, tt.want)
		}
	}
}
