package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(toys([]int32{1, 2, 3, 21, 7, 12, 14, 21}))
}

func toys(w []int32) int32 {

	sort.Slice(w, func(i, j int) bool {
		return w[i] < w[j]
	})
	lessThan := int32(0)
	reset := true
	count := int32(0)
	for _, va := range w {
		if reset {
			lessThan = va + 4
			reset = false
			count++
		}
		if va <= lessThan {
			continue
		}
		reset = true
	}
	return count
}
