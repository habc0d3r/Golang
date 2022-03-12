package main

import (
	"fmt"
)

func main() {
	age := 35
	name := "Hengul Akash"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n") //Print method doesn't add a new line automatically

	// Println
	fmt.Println("Hi guys!!")
	fmt.Println("bye bye guys") //But Println function/method does add a new line automatically
	fmt.Println("My name is", name, "and my age is", age)

	// Printf (formatted strings) %_ = format specifier ( _ can be many letters) https://pkg.go.dev/fmt
	fmt.Printf("My name is %v and my age is %v \n", name, age) //default format specifier for variable: %v
	fmt.Printf("My name is %q and my age is %q \n", name, age)
	fmt.Printf("age is of type %T \n", age) // %T represents type of the value
	fmt.Printf("You scored %f points! \n", 255.55)
	fmt.Printf("You scored %0.1f points! \n", 255.55) // to round a float value

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My name is %v and my age is %v \n", name, age) // Sprintf returns a formatted string to us and we can save that in a var
	fmt.Println("The saved string is:", str)

}
