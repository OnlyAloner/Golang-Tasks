import package 
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
func main(){
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
}