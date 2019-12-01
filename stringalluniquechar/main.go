package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	st := "ab"
	fmt.Println(isAllUnique(st))
}
func isAllUnique(st string) bool {
	sp := strings.Split(st, "")
	sort.Strings(sp)
	for i, value := range sp {
		if i < len(sp)-1 {
			if value == sp[i+1] {
				return false
			}
		}
	}
	return true
}
