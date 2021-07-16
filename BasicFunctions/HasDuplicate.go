package main

import "fmt"

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
	A := make([]int, 10, 10)
	A = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(HasDuplicate(A))
	A = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	fmt.Println(HasDuplicate(A))
	A = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(HasDuplicate(A))
	A = []int{2, 1, 2, 1, 2, 1, 2, 1, 2, 1}
	fmt.Println(HasDuplicate(A))
}
