package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maxMin(5, []int32{4504,
		1520,
		5857,
		4094,
		4157,
		3902,
		822,
		6643,
		2422,
		7288,
		8245,
		9948,
		2822,
		1784,
		7802,
		3142,
		9739,
		5629,
		5413,
		7232}))
}
func maxMin(k int32, arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	keep := arr
	fairness := math.MaxInt32
	for i := 0; i < len(keep)-(int(k)+1); i++ {
		arr = keep[i : int(k)+i]
		if int(arr[len(arr)-1]-arr[0]) < fairness {
			fairness = int(arr[len(arr)-1] - arr[0])
		}

	}
	return int32(fairness)
}
