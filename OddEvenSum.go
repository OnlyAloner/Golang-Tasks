package main

import "fmt"

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

func main() {

	fmt.Println(OddEvenSum(5))
	fmt.Println(OddEvenSum(10))
	fmt.Println(OddEvenSum(30))
	fmt.Println(OddEvenSum(50))
	fmt.Println(OddEvenSum(99))
	fmt.Println()
	//Comparing with FastOddEvenSum function

	fmt.Println(FastOddEvenSum(5))
	fmt.Println(FastOddEvenSum(10))
	fmt.Println(FastOddEvenSum(30))
	fmt.Println(FastOddEvenSum(50))
	fmt.Println(FastOddEvenSum(99))
}
