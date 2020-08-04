package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
)

//func TestMain(m *testing.M) {
//	b := &testing.B{}
//	BenchmarkTe(b)
//	BenchmarkTes(b)
//}
func BenchmarkTe(b *testing.B) {
	runtime.GOMAXPROCS(2)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/raj", strings.NewReader(`{"nam":"raj","roll":"4"}`))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(raj)
		handler.ServeHTTP(rr, req)
		fmt.Println(rr.Body.String())
	}
}

type str struct {
	a string
	b string
}

func TestCheck1(t *testing.T) {
	Hello()
}
func TestCheck(t *testing.T) {

	tests := []str{
		{a: "ths", b: "rs"},
		{a: "thisis", b: "helor"},
	}
	for i, k := range tests {
		fmt.Println(i, k)
	}
	for i := 1; i <= 10; i++ {
		ThisIsToCheck()
		if i%3 == 0 {
			t.Fatalf("some error again %d", i)
		}
	}
}

//
//func BenchmarkTes(b *testing.B) {
//	runtime.GOMAXPROCS(4)
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		getStruct()
//	}
//}
