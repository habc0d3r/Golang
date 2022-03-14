package main

import (
	"fmt"
	"sort"
)

func main() {

	x := 0
	for x < 5 { //there's no "while" keyword in golang, for can be used as while
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i < 10; i++ { // traditional for loop
		fmt.Println("Value :", i)
	}

	names := []string{"rudy", "anwar", "bach", "lele", "inanna"}
	names[2] = "adam"
	sort.Strings(names)
	for i := 0; i < len(names); i++ { // traditional for loop
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	for _, value := range names {
		// _(underscore) in Golang is known as the Blank Identifier. Identifiers are the user-defined name of the program components used for the identification purpose. Golang has a special feature to define and use the unused variable using Blank Identifier.
		fmt.Printf("the value is %v \n", value)
	}

	fmt.Println(names)

}
