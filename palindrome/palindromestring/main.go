package palindromestring

import "fmt"

func main() {
	isPalindrome("madma")
	isPalindrome("madam")
}
func isPalindrome(st string) {
	runes := []rune(st)
	n := len(runes) - 1
	i := 0
	for i < n {
		if runes[i] != runes[n] {
			fmt.Println("its not  a palindrome")
			return
		}
		i++
		n--
	}
	fmt.Println("its a palindrome")
}
