package main

import (
	"fmt"
)

const aConst int = 64

func main() {

	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt) // see that default int value is 0

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variable's type is %T\n", anotherInt)

	myString := "This also a string"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T\n", myString)

	// for const aConst
	fmt.Println(aConst)
	fmt.Printf("The variable's type is %T\n", aConst)

	//fmt.Println("Hello from Go!")
}
