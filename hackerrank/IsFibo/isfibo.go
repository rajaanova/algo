package main

import "fmt"

func main() {
	fmt.Println(solve(3))
}

func solve(n int64) string {
	a := int64(0)
	b := int64(1)
	temp := a
	for a <= n {
		if a == n {
			return "IsFibo"
		}
		temp = a
		a = a + b
		b = temp
	}
	return "IsNotFibo"
}
