package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{3, 2, 0, 0, 2, 0, 1}
	min := getSmallest(a, 0)
	fmt.Println(min)
}

func getSmallest(a []int, i int) int {
	if i >= len(a)-1 {
		return 0
	}
	mint := math.MaxUint8
	if a[i] > 0 {
		j := i
		i = i + 1
		for ; i < len(a) && i <= a[j]+j; i++ {
			fmt.Println(a[j], j, i, a[j]+j)
			ret := getSmallest(a, i) + 1
			if ret < mint {
				mint = ret
			}
		}
	}
	return mint
}
