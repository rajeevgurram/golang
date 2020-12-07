package main

import (
	"fmt"

	"./educative/lecture1/calculator"
)

func main() {
	fmt.Println("Hello World")

	// Calculator
	newLecture(1, "Calculator")
	var num1, num2 int = 10, 20
	fmt.Printf("Sum of %d + %d = %d \n", num1, num2, calculator.Sum(num1, num2))
	fmt.Printf("Sub of %d - %d = %d \n", num1, num2, calculator.Sub(num1, num2))
	fmt.Printf("Mul of %d X %d = %d \n", num1, num2, calculator.Mul(num1, num2))
	fmt.Printf("Div of %d / %d = %d \n", num1, num2, calculator.Div(num2, num1))
}

func newLecture(lectureNumber int, lectureName string) {
	fmt.Printf("\n\nLecture %d, %s\n", lectureNumber, lectureName)
	fmt.Printf("_______________________________\n\n")
}
