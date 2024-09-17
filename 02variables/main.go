package main

import "fmt"

func main() {
	var userName string = "rahul"
	fmt.Printf("Variable is of type: %T with value: %q\n", userName, userName)

	var isLoggedIn bool = false
	fmt.Printf("Variable is of type: %T with value: %v\n", isLoggedIn, isLoggedIn)

	// unsigned int
	var smallInt uint16 = 256
	fmt.Printf("Variable is of type: %T with value: %v\n", smallInt, smallInt)

	// precision upto 4 decimal places
	var smallFloat float32 = 256.123456789
	fmt.Printf("Variable is of type: %T with value: %v\n", smallFloat, smallFloat)

	// precision upto 15 decimal places
	var bigFloat float64 = 256.1234567891234567
	fmt.Printf("Variable is of type: %T with value: %v\n", bigFloat, bigFloat)

	// default int puts 0
	var defaultInt int
	fmt.Printf("Variable is of type: %T with value: %v\n", defaultInt, defaultInt)

	// implicit type definition
	var implicitStrType = "auto assign type to str"
	fmt.Printf("Variable is of type: %T with value: %q\n", implicitStrType, implicitStrType)

	// no var declaration
	noVarVariable := 10
	fmt.Printf("Variable is of type: %T with value: %v\n", noVarVariable, noVarVariable)

}
