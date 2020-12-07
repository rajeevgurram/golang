package strings

import (
	"fmt"
	"strings"
)

var str string = "Welcome to golang strings"

func hasPrefix() {
	prefix := "Welcome"
	fmt.Printf("String %s has prefix %s ? %v", str, prefix, strings.HasPrefix(str, prefix))
}

// Run runs all above functions
func Run() {
	hasPrefix()
}
