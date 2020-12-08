package functions

import "fmt"

func f1(a, b, c int) {
	fmt.Printf("a: %d, b: %d, c: %d \n\n", a, b, c)
}

func f2(a, b int) (int, int, int) {
	return a, b, a + b
}

func SumProductDiff(i, j int) (s int, p int, d int) { // named version
	s, p, d = i+j, i*j, i-j
	return
}

func sumOfNumbers(nums ...int) {
	var sum int = 0
	for i := range nums {
		sum += nums[i]
	}

	fmt.Printf("Sum of %v is: %d\n\n", nums, sum)
}

func Run() {
	f1(1, 2, 3)
	a, b, sum := f2(1, 2)
	fmt.Printf("a: %d, b: %d, sum: %d \n\n", a, b, sum)
	SumProductDiff(1, 2)
	sumOfNumbers(1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5)
}
