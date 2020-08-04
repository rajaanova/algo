package student

import "testing"

func BenchmarkPutStudent(t *testing.B) {
	t.ReportAllocs()
	for i := 0; i < 100000; i++ {
		PutStudent1()
	}
}
