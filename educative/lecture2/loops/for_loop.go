package loops

import "fmt"

func forLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d \n", i)
	}
}

func forRange() {
	fmt.Printf("For Range \n\n")
	var str string = "This is from go lang"

	for idx, char := range str {
		fmt.Printf("Index: %d, Char: %c \n", idx, char)
	}
}

func Run() {
	forLoop()
	forRange()
}
