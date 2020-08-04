package main

import (
	"fmt"
	"math"
	"net/http"
)

var client = http.Client{}

func main() {

	//reques, _ := http.NewRequest("POST", "localhost:8080", nil)
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//http.DefaultTransport
	//
	//reques.WithContext(ctx)
	//client.Do()
	//cotx := req.cont
	fmt.Println(kangaroo(0, 3, 4, 2))
}

// Complete the kangaroo function below.
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if (x2 > x1 && (v1 < v2 || v1 == v2)) || (x1 > x2 && (v1 > v2 || v1 == v2)) {
		return "NO"
	}
	stpe := int32(math.Abs(float64(x1 - x2)))
	v2v1sub := int32(math.Abs(float64(v1 - v2)))
	if stpe%v2v1sub == 0 {
		return "YES"
	}
	return "NO"
}
