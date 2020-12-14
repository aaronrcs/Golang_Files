package main

import "fmt"

func main() {
	// & and * Operators

	// & --> Pointer or reference of the memeory location where value is stored in computer
	// * --> Derefrence or access the value in which its pointing to

	x := 7
	y := &x

	fmt.Println("X: ", x, " Y: ", y)

	// Using a '*' to derefrence from the pointer to modify the value of y
	*y = 8

	fmt.Println("X: ", x, " Y: ", y)

}
