// https://pkg.go.dev/builtin#int
package main

import "fmt"

func main() {

	fmt.Println("Hello World! I hope every life is doing well.")
	// strings
	var name0 string = "yo"
	var name1 = "adam"
	var name2 string

	fmt.Println(name0, name1, name2)

	name0 = "waheed"
	name2 = "anwar"

	fmt.Println(name0, name1, name2)

	name4 := "lele" //short hand of variable declaration, ':' is only used 1st time to declare variable but not to update var. we can't use this way OUTSIDE A FUNTION
	fmt.Println(name4)

	// ints

	var ageZero int = 22
	var ageOne = 11
	ageTwo := 30
	age3, age4 := 33, 44

	fmt.Println(ageZero, ageOne, ageTwo, age3, ",", age4)

	// bits & memorys
	var numZero int8 = 25
	var numOne int8 = -128
	var numTwo uint16 = 256
	fmt.Println(numZero, numOne, numTwo)

	// float
	var scoreOne float32 = 25.97
	var scoreTwo float64 = 99357739349347
	scoreThree := 3.6

	fmt.Println(scoreOne, scoreThree, scoreTwo)
}
