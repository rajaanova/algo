package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(isStringPermutaion("string", "gnstsri"))
	fmt.Println(isStringPermutaionMapCount("string", "gstinsr"))
}
func isStringPermutaionMapCount(s, a string) bool {
	var mapOfCount map[int32]int = make(map[int32]int)
	for _, value := range s {
		mapOfCount[value] = mapOfCount[value] + 1
	}
	for _, value := range a {
		mapOfCount[value] = mapOfCount[value] - 1
	}
	for _, v := range mapOfCount {
		if v != 0 {
			return false
		}
	}
	return true
}
func isStringPermutaion(s, a string) bool {
	splited1 := strings.Split(s, "")
	splited2 := strings.Split(a, "")
	sort.Strings(splited2)
	sort.Strings(splited1)
	split1 := strings.Join(splited1, "")
	split2 := strings.Join(splited2, "")
	if split1 == split2 {
		return true
	}
	return false
}
