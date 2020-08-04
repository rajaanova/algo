package main

import "fmt"

func main() {
	x := -201
	fmt.Println(string(49))
	intToString(x)
}
func intToString(a int) {
	st := ""
	var isNegative bool
	if a < 0 {
		a = a * -1
		isNegative = true
	}
	for a > 0 {
		re := a % 10
		st = st + string('0'+re)
		a = a / 10
	}
	st = reverse(st)
	if isNegative {
		st = "-" + st
	}
	fmt.Println(st + "raj")
}

func reverse(st string) string {
	runesInString := []rune(st)
	lenofRune := len(runesInString) - 1
	for i, j := 0, lenofRune; i < j; i, j = i+1, j-1 {
		runesInString[i], runesInString[j] = runesInString[j], runesInString[i]
	}
	return string(runesInString)
}
