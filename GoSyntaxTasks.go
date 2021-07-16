package main

import "fmt"

func Nthfibonacci(a int) int { // Time Complexity O(n), Space Complexity O(1)

	var (
		fib1 = 1
		fib2 = 1
	)

	for i := 2; i < a; i++ {
		temp := fib2
		fib2 = fib1 + fib2
		fib1 = temp
	}

	return fib2
}

func FizzBuzz(a int) int {
	for i := 1; i <= a; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz") // Print Fizz if number divides by 3
		}
		if i%5 == 0 {
			fmt.Print("Buzz") // Print Buzz if number divides by 5
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Print(i) // Print number if he not divides 3 and 5
		}
		fmt.Println()
	}
	fmt.Println("End")
	fmt.Println()
	return 0
}

func isPalindrome(Text string) bool {
	var textSize = len(Text)
	for i := 0; i < textSize; i++ {
		if Text[i] != Text[textSize-1-i] {
			return false // Text is not palindrome
		}
	}
	return true // Text is palidnrome
}

func OddEvenSum(Range int) (int, int) { // Time Complexity O(Range).
	var (
		Odd  = 0
		Even = 0
	)
	for i := 1; i <= Range; i++ {
		if i%2 == 0 {
			Even += i
		} else {
			Odd += i
		}
	}
	return Odd, Even
}
func FastOddEvenSum(Range int) (int, int) { // Time Complexity O(1) with arifmethic progression

	var (
		Odd  = int((Range - 1) / 2)
		Even = int(Range / 2)
	)
	Even = 2 * Even * (Even + 1) / 2
	Odd = (Odd + 1) + 2*Odd*(Odd+1)/2
	return Odd, Even
}

func HasDuplicate(Array []int) bool { // Time Complexity O(N^2)
	for i := 0; i < len(Array)-1; i++ {
		for j := i + 1; j < len(Array); j++ {
			if Array[i] == Array[j] { // Check each pair of numbers and if we found pair where we have 2 same numbers we return true
				return true
			}
		}
	}
	return false
}

func main() {

	/* Tests for fibonacci function
	fmt.Println(Nthfibonacci(1))
	fmt.Println(Nthfibonacci(2))
	fmt.Println(Nthfibonacci(3))
	fmt.Println(Nthfibonacci(4))
	fmt.Println(Nthfibonacci(5))
	fmt.Println(Nthfibonacci(6))
	fmt.Println(Nthfibonacci(7))
	fmt.Println(Nthfibonacci(8))
	fmt.Println(Nthfibonacci(9))
	fmt.Println(Nthfibonacci(10))
	fmt.Println(Nthfibonacci(20))
	*/

	/* Tests for FizzBuzz function
	FizzBuzz(1)
	FizzBuzz(3)
	FizzBuzz(15)
	FizzBuzz(30)
	*/

	/* Tests for isPalindrome function
	fmt.Println("IsPalindrome")
	fmt.Println("a : ", isPalindrome("a"))
	fmt.Println("abcba : ", isPalindrome("abcba"))
	fmt.Println("abcdf : ", isPalindrome("abcdf"))
	fmt.Println("aaaaaaa : ", isPalindrome("aaaaaaa"))
	fmt.Println("Golanggnalog : ", isPalindrome("Golanggnalog"))
	fmt.Println("GolanggnaloG : ", isPalindrome("GolanggnaloG"))
	fmt.Println("Golang gnaloG : ", isPalindrome("Golang gnaloG"))
	*/

	/* Tests for OddEvenSum function
	fmt.Println(OddEvenSum(5))
	fmt.Println(OddEvenSum(10))
	fmt.Println(OddEvenSum(30))
	fmt.Println(OddEvenSum(50))
	fmt.Println(OddEvenSum(99))

	Comparing with FastOddEvenSum function

	fmt.Println(FastOddEvenSum(5))
	fmt.Println(FastOddEvenSum(10))
	fmt.Println(FastOddEvenSum(30))
	fmt.Println(FastOddEvenSum(50))
	fmt.Println(FastOddEvenSum(99))
	*/

	/* Tests for HasDuplicate function
	A := make([]int, 10, 10)
	A = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(HasDuplicate(A))
	A = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	fmt.Println(HasDuplicate(A))
	A = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(HasDuplicate(A))
	A = []int{2, 1, 2, 1, 2, 1, 2, 1, 2, 1}
	fmt.Println(HasDuplicate(A))
	*/
}
