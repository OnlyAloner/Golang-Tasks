import package 
import "fmt"
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
func main(){
	FizzBuzz(1)
	FizzBuzz(3)
	FizzBuzz(15)
	FizzBuzz(30)
}