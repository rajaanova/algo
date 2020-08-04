package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][]int32{
		{8, 1},
		{4, 2},
		{5, 6},
		{3, 1},
		{4, 3},
	}
	fmt.Println(jimOrders(a))
}

type ss struct {
	in int
}

func jimOrders(orders [][]int32) []int32 {
	a := make(map[int32]int32)
	for i, v := range orders {
		totoal := v[0] + v[1]
		a[totoal] = int32(i) + 1
	}
	var sl []int32
	for k, _ := range a {
		sl = append(sl, k)
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
	listofss := make([]ss, 0)
	listofss = append(listofss, ss{in: 8})
	listofss = append(listofss, ss{in: 10})
	listofss = append(listofss, ss{in: 4})
	sort.Slice(listofss, func(i, j int) bool {
		return listofss[i].in < listofss[j].in

	})
	lsitOfCustomer := make([]int32, 0)
	for _, val := range sl {
		va := a[val]
		lsitOfCustomer = append(lsitOfCustomer, va)
	}
	return lsitOfCustomer
}
