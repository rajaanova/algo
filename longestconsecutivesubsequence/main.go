package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestConsecutiveNumber([]int{1, 9, 3, 10, 4, 20, 2}))
}

type mem struct {
	init int
	end  int
}

func longestConsecutiveNumber(a []int) (int, int, []int) {
	i := 0
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	var x mem
	max := 0
	for k := 1; k < len(a); k++ {
		if a[k-1]+1 != a[k] {
			if max < k-i {
				max = k - i
				x.end = k - 1
				x.init = i
			}
			i = k
		}
	}
	fmt.Println(a)
	return x.init, x.end, a[x.init : x.end+1]
}
