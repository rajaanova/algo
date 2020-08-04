package main

import (
	"fmt"
)

func main() {
	fmt.Println(stringcount("aabbccc"))
}
func stringcount(s string) string {
	prv := s[0]
	str := ""
	count := 1
	for i := 1; i < len(s); i++ {
		if string(s[i]) != string(prv) {
			//fmt.Println(string(s[i]), string(prv), count)
			str = fmt.Sprintf("%s%d", str+string(prv), count)
			prv = s[i]
			count = 0
		}
		count++
		if i == len(s)-1 {
			str = fmt.Sprintf("%s%d", str+string(s[i]), count)
		}
	}
	return str
}
