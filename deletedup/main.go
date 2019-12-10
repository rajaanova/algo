package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 4, 6, 3, 2, 9, 19, 2}
	deleteDup(a)
}
func deleteDup(a []int) {
	sort.Ints(a)
	j := len(a) - 1
	for i := 0; i < j; i++ {
		if a[i] == a[i+1] {
			a = append(a[0:i], a[i+1:]...)
			j--
		}

	}
	fmt.Println(a)
}
