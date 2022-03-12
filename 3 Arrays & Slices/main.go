package main

import "fmt"

func main() {
	// arrays
	// var years [2]int
	// var ages [3]int = [3]int{20, 23, 25}
	var ages = [3]int{20, 23, 25} // short hand

	names := [4]string{"adam", "anwar", "lele", "hannah"} // short hand assignment
	names[3] = "rudy"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25

	scores = append(scores, 85) // automatically it doesn't append it. it just returns a new slice and the new slice gets assigned to the variable. kinda overwrite the variable with a new slice

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3] // exclusive of the 2nd number but 1st
	rangeTwo := names[2:]
	rangeThree := names[:2]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "bach")
	fmt.Println(rangeOne)

}
