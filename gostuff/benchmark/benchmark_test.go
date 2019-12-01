package benchmark

import "testing"

func BenchmarkFloatABS(t *testing.B) {
	for i := 0; i < t.N; i++ {
		floatABS(-i)
	}
}

func BenchmarkMultABS(t *testing.B) {
	for i := 0; i < t.N; i++ {
		multABS(-i)
	}
}

// run below command
//go test -bench=.
