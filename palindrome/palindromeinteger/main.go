package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func main() {
	log.Println("hello raj")
	isPalindromInt(101)
}
func isPalindromInt(i int) {
	reverse := 0
	j := i
	for i > 0 {
		remaine := i % 10
		i = i / 10
		reverse = reverse*10 + remaine
	}
	fmt.Println(j, reverse)
	if j == reverse {
		fmt.Println("its a palindrome")
	}
}
