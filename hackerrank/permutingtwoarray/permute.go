package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoArrays(5, []int32{2, 4}, []int32{3, 1}))
}

// Complete the twoArrays function below.
func twoArrays(k int32, A []int32, B []int32) string {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})
	//fmt.Println(A)
	//fmt.Println(B)
	l := len(A) - 1
	for i := 0; i <= l; i++ {
		fmt.Println(A[i], B[l-i])
		if (A[i] + B[l-i]) < k {
			return "NO"
		}
	}
	return "YES"
}
