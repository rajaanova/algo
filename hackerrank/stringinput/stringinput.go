package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	personwithphonenum := make(map[string]int)
	i, j := 0, 0
	fmt.Scanln(&i)
	for j < i {
		v := ""
		n := 0
		fmt.Scanln(&v, &n)
		fmt.Println(v)
		personwithphonenum[v] = n
		j++
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		v := ""
		fmt.Scanln(&v)

		if val, ok := personwithphonenum[v]; ok {
			fmt.Printf("%s=%d\n", v, val)
		} else {
			fmt.Println("Not found")
		}
		i--
	}
	//Enter your code here. Read input from STDIN. Print output to STDOUT
}
