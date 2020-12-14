package main

import "fmt"

func changeValue(str *string) {
	*str = "String 1"
}

func changeValue2(str string) {
	str = "String 2"
}

func main() {
	// & and * Operators

	// & --> Pointer or reference of the memeory location where value is stored in computer
	// * --> Derefrence or access the value in which its pointing to

	// -- Simple Example using & and * Operators
	x := 7
	y := &x

	// fmt.Println("X: ", x, " Y: ", y)

	// Using a '*' to derefrence from the pointer to modify the value of x
	*y = 8

	// fmt.Println("X: ", x, " Y: ", y)

	//--Using functions, Example:

	toChange := "I was changed!"

	// Give an '&' or pointer to toChange in order to change value of string inside the function
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

	fmt.Println("----------------------------")

	// Using the changeValue2 function, the value of string inside function does not change because pointer was not passed

	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)

}
