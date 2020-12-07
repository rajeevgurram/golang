package pointers

import "fmt"

func addressOfVariable() {
	var i int = 10
	var address *int = &i
	fmt.Printf("value of i is %d\n address of i is %v\n\n", i, address)
}

// Run the above functions
func Run() {
	addressOfVariable()
}
