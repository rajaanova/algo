package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumDistances([]int32{1, 1}))
	fmt.Println(minimumDistancesMap([]int32{1, 1}))
}

func minimumDistancesMap(a []int32) int32 {
	distance := math.MaxInt32
	mapOfDistance := make(map[int32]int, len(a))
	for i := 0; i < len(a); i++ {
		if v, ok := mapOfDistance[a[i]]; ok {
			if i-v < distance {
				distance = i - v
			}
		} else {
			mapOfDistance[a[i]] = i
		}
	}
	if distance == math.MaxInt32 {
		return -1
	}
	return int32(distance)
}

func minimumDistances(a []int32) int32 {
	distance := math.MaxInt32
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j <= len(a)-1; j++ {
			if a[i] == a[j] {
				if j-i < distance {
					distance = j - i
				}
			}
		}
	}
	if distance == math.MaxInt32 {
		return -1
	}
	return int32(distance)
}
