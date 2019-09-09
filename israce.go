package israce

import "testing"

// MustRace fails test if test runs with no -race flag.
func MustRace(test *testing.T) {
	if !Race {
		test.Fatalf("test must be run with -test flag")
	}
}
