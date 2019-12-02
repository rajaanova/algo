package main

import "fmt"

func main() {
	permute("", "abc")
}
func permute(prefix, st string) {
	if len(st) == 0 {
		fmt.Println(prefix)
		return
	}
	for i, _ := range st {
		permute(prefix+string(st[i]), st[:i]+st[i+1:])
	}
}
