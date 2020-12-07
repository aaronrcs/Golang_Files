package main

import "fmt"

func main() {
	fmt.Println("Learning Go, whooohooooo!")
	fmt.Println("The most complex Golang program ever created!")

	//Declaring Data Types

	var name string = "Skxter"
	var number uint16 = 260

	number = number + 10

	fmt.Println("Sum", number)
	fmt.Println("My name is", name)

	// Another way of declaring a variable

	age := 25 //Go guessing the type of value being assigned to variable

	fmt.Println("I am", age, "years old.")
}
