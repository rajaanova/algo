package main

import "fmt"

func main() {
	fmt.Println(rotLeft([]int32{1, 2, 3, 4, 5}, 4))
}

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	a = append(a[d:], a[0:d]...)
	return a
}
