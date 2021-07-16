package main

import "fmt"

func isPalindrome(Text string) bool {
	var textSize = len(Text)
	for i := 0; i < textSize; i++ {
		if Text[i] != Text[textSize-1-i] {
			return false // Text is not palindrome
		}
	}
	return true // Text is palidnrome
}
func main() {
	fmt.Println("IsPalindrome")
	fmt.Println("a : ", isPalindrome("a"))
	fmt.Println("abcba : ", isPalindrome("abcba"))
	fmt.Println("abcdf : ", isPalindrome("abcdf"))
	fmt.Println("aaaaaaa : ", isPalindrome("aaaaaaa"))
	fmt.Println("Golanggnalog : ", isPalindrome("Golanggnalog"))
	fmt.Println("GolanggnaloG : ", isPalindrome("GolanggnaloG"))
	fmt.Println("Golang gnaloG : ", isPalindrome("Golang gnaloG"))
}
