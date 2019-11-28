package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	reverseString("helloraj")
	reverseWithRuneCount("The quick brown 狐 jumped over the lazy 犬")
	reverseWords("The quick brown 狐 jumped over the lazy 犬")
}

func reverseWords(st string) {
	str := strings.Fields(st)
	//str := strings.Split(st, " ")
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	fmt.Println(strings.Join(str, " "))
}
func reverseWithRuneCount(st string) {
	o := make([]rune, utf8.RuneCountInString(st))
	i := len(o)
	for _, v := range st {
		i--
		o[i] = v
	}
	fmt.Println(string(o))
}

func reverseString(st string) {
	s := []rune(st)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(string(s))
}
