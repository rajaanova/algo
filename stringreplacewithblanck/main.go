package main

import "fmt"

func main() {
	st := "Mr John Smith    "
	str := []byte(st)
	ifFirstCheck := false
	i, j := len(str)-1, len(str)-1
	for i > -1 {
		if string(str[i]) != " " {
			ifFirstCheck = true
			str[j] = str[i]
			j--
		}
		if string(str[i]) == " " && ifFirstCheck {
			str[j] = '0'
			str[j-1] = '2'
			str[j-2] = '%'
			j = j - 3
		}
		i--
	}
	fmt.Println(string(str))
}
