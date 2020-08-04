package main

import "fmt"

func main() {
	var U+1D70B = 22/7
	fmt.Println(@)
	sli := []int{1, 2, 3, 4, 5}
	num := 1
	getArrayofArray(sli, num)
}

func getArrayofArray(sli []int, num int) {
	xxx := make([][]int, 0)
	for i := 0; i <= len(sli)-num; i++ {
		xxx = append(xxx, sli[i:i+num])
	}
	fmt.Println(xxx)
}
