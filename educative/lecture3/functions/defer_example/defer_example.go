package defer_example

import (
	"fmt"
	"io"
)

func defer1() {
	fmt.Println("from defer1")
	defer f1()
	fmt.Println("end of defer1")
}

func f1() {
	fmt.Println("from function f1")
}

func defer2() {
	fmt.Println("Defer 2 example, printing reverse")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func defer3(s string) (n int, err error) {
	fmt.Println("Defer 3 example. logging")
	defer func() {
		fmt.Printf("func1(%q) = %d, %v\n\n", s, n, err)
	}()

	return 7, io.EOF
}

func Run() {
	defer1()
	defer2()
	defer3("GO")
}
