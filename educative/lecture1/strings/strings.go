package strings

import (
	"fmt"
	"strconv"
	"strings"
)

var str string = "Welcome to golang strings"

func hasPrefix() {
	prefix := "Welcome"
	fmt.Printf("String:\n %s \nhas prefix %s ? = %v\n\n", str, prefix, strings.HasPrefix(str, prefix))
}

func hasPostfix() {
	postfix := "test"
	fmt.Printf("String:\n %s \nhas postfix %s ? = %v\n\n", str, postfix, strings.HasSuffix(str, postfix))
}

func contains() {
	contains := "go"
	fmt.Printf("String:\n %s \ncontains %s ? = %v\n\n", str, contains, strings.Contains(str, contains))
}

func getIndex() {
	indexFor := "go"
	fmt.Printf("String:\n %s \nindex for %s is = %v\n\n", str, indexFor, strings.Index(str, indexFor))
}

func replace() {
	replace := "l"
	with := "a"
	fmt.Printf("String:\n %s \nreplace %s with %s is \n %s\n\n", str, replace, with, strings.Replace(str, replace, with, -1))
}

func toLower() {
	fmt.Printf("To Lower \n %s is \n %s \n\n", str, strings.ToLower(str))
}

func toUpper() {
	fmt.Printf("To Upper \n %s is \n %s \n\n", str, strings.ToUpper(str))
}

func splitString() {
	fmt.Printf("Splitting %s with space\n", str)
	split := strings.Fields(str)
	fmt.Println(split)
}

func splitString1() {
	fmt.Printf("Splitting %s with any value\n", str)
	split := strings.Split(str, "l")
	fmt.Println(split)
}

func joinStringArray() {
	fmt.Printf("\nSplitting %s with spaces\n", str)
	split := strings.Fields(str)
	fmt.Printf("\nJoining string with comma\n")
	fmt.Println(strings.Join(split, ","))
}

func convertToString() {
	value := "666"
	ivalue, _err := strconv.Atoi(value)
	if _err == nil {
		fmt.Printf("\n converted value %s \n", strconv.Itoa(ivalue+4))
	}
}

// Run runs all above functions
func Run() {
	hasPrefix()
	hasPostfix()
	contains()
	getIndex()
	replace()
	toLower()
	toUpper()
	splitString()
	splitString1()
	joinStringArray()
	convertToString()
}
