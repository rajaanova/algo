package main

import "fmt"

func main() {
	fib(5)
	fmt.Println()
	recursefib(5)
}

func fib(a int) {
	x, y := 1, 0
	count := 1
	for count <= a {
		fmt.Print(x, " ")
		x, y = y+x, x
		count++
	}

}
func recursefib(a int) {
	for i := 0; i < a; i++ {
		fib := getFib(i + 1)
		fmt.Print(fib, " ")
	}
}

func getFib(i int) int {
	if i < 2 {
		return i
	}
	fib := getFib(i-1) + getFib(i-2)
	return fib
}
