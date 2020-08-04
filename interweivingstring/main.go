package main

import (
	"fmt"
)

func main() {
	one := "aabcc"
	two := "dbbca"
	three := "aadbbcbcac"
	i := 0
	j := 0
	fmt.Println(InterweavingStrings(one, two, three, i, j))
}
func InterweavingStrings(one, two, three string, i, j int) bool {
	k := i + j

	if k == len(one)+len(two) {
		return true
	}

	if three[k] == one[i] {
		if InterweavingStrings(one, two, three, i+1, j) {
			return true
		}
	}

	if three[k] == two[j] {
		return InterweavingStrings(one, two, three, i, j+1)
	}

	return false
}
