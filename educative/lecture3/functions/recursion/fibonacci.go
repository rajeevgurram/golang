package recursion

import "fmt"

func fibonacci(number int) int {
	if number <= 1 {
		return 1
	}

	return fibonacci(number-1) + fibonacci(number-2)
}

func Run() {
	var number int = 11
	val := fibonacci(number)
	fmt.Printf("Fibonacci for %d is %d\n\n", number, val)
}
