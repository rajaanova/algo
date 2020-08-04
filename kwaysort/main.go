package main

import (
	"container/heap"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type Sequence int

const (
	Increasing Sequence = iota
	Decreasing
)

type sortarr []int

func (a sortarr) Len() int {
	return len(a)
}
func (a sortarr) Swap(i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func (a sortarr) Push(x interface{}) {
	a = append(a, x.(int))
}
func (a sortarr) Pop() interface{} {
	return a[a.Len()-1]
}
func (a sortarr) Less(i, j int) bool {
	return a[i] < a[j]
}
func main() {
	stt := "ofoof"
	fmt.Println(strings.Trim(stt, "fo"))
	s := "012"
	s = strings.Map(func(r rune) rune {
		return r + 2
	}, s)
	fmt.Println(s)

	ss := url.PathEscape("a b c")
	fmt.Println(ss)
	fmt.Println(ss)
	//fmt.Println([]byte(s)[1:4])
	//fmt.Println([]byte("012")[1:4])
	//sss := []int{12, 99, 32, 78, 43, 119, 324, 39, 56, 89}
	som := sortarr([]int{89, 323, 789, 12, 32, 43})

	heap.Init(som)
	fmt.Println(heap.Pop(som))
	fmt.Println(heap.Pop(som))
	//fmt.Println()
	////	outplaceDistribution(sss)
	//for x := range sss {
	//	fmt.Println(x)
	//}
	//inplaceDistribution(sss)
}
func inplaceDistribution(sss []int) {
	sort.Ints(sss)
	//	fmt.Println(sss)
	//sort.Reverse()
	dist := len(sss) / 4
	for i := 1; i < 4; i++ {
		temp := sss[2*i : (2*i)+dist]
		copy(sss[2*i:], sss[2*(i+1):2*(i+1)+dist])
		copy(sss[2*(i+1):2*(i+1)+dist], temp)

	}
	fmt.Println(sss)
}
func outplaceDistribution(sss []int) {
	sort.Slice(sss, func(i, j int) bool {
		return sss[i] < sss[j]
	})

	totaldistribution := len(sss) / 4
	numofarrasy := make([][]int, 4)
	k := 0
	for i := 0; i < 3; i++ {
		numofarrasy[i] = sss[k : k+totaldistribution]
		k = k + totaldistribution
	}
	numofarrasy[len(numofarrasy)-1] = sss[k:]
	//fmt.Println(numofarrasy)
	newArr := make([]int, 0, len(sss))
	sequenceStatus := Increasing
	for j := 0; j < len(numofarrasy); j++ {
		if sequenceStatus == Increasing {
			newArr = append(newArr, numofarrasy[j]...)
			sequenceStatus = Decreasing
		} else {
			sort.Slice(numofarrasy[j], func(i, k int) bool {
				return numofarrasy[j][k] < numofarrasy[j][i]
			})
			newArr = append(newArr, numofarrasy[j]...)
			sequenceStatus = Increasing
		}
	}
	fmt.Println(newArr)
}
