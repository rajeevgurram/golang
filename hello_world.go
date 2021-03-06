package main

import (
	"fmt"

	"./educative/lecture1/calculator"
	"./educative/lecture1/pointers"
	"./educative/lecture1/strings"
	"./educative/lecture1/temperature/conversion"
	"./educative/lecture2/loops"
	"./educative/lecture2/season"
	"./educative/lecture3/functions"
	"./educative/lecture3/functions/defer_example"
	"./educative/lecture3/functions/recursion"
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

	newLecture(2, "Temparature Conversion")
	var celcius conversion.Celcius = 10
	fmt.Printf("%vc is %vf \n", celcius, conversion.ToFahrenheit(celcius))

	newLecture(3, "Strings")
	strings.Run()

	newLecture(4, "Pointers")
	pointers.Run()

	newLecture(5, "Switch")
	season.Run()

	newLecture(6, "For")
	loops.Run()

	newLecture(7, "Functions")
	functions.Run()

	newLecture(8, "Defer Functions")
	defer_example.Run()

	newLecture(9, "Recursion")
	recursion.Run()

	newLecture(10, "Factorial")
	fmt.Printf("Factorial for 10 is %d\n\n", recursion.Factorial(10))
}

func newLecture(lectureNumber int, lectureName string) {
	fmt.Printf("\n\nLecture %d, %s\n", lectureNumber, lectureName)
	fmt.Printf("_______________________________\n\n")
}
