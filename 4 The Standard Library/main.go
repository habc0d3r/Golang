package main

import ( // Go Standard Library : https://pkg.go.dev/std
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "Hello there"

	fmt.Println(strings.Contains(greetings, "Hello"))
	fmt.Println(strings.Contains(greetings, "world")) //Contains method: this return "true" or "false" i.e. boolean value

	fmt.Println(strings.ReplaceAll(greetings, "there", "World!"))
	fmt.Println("Original string value = ", greetings) // this returns a new string with new given substring but doesn't replace the old one in the original string

	fmt.Println(strings.ToUpper("make me uppercase")) // returns with uppercase, not replaces

	fmt.Println(strings.Index(greetings, "there"))

	splitted := strings.Split(greetings, "e") // Split slices s into all substrings separated by sep and
	// returns a slice of the substrings between those separators
	fmt.Println(splitted)
	fmt.Println(strings.Split(greetings, "")) // If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.
	// fmt.Printf("%T", splitted)
	ages := []int{22, 34, 53, 11, 51, 36, 57, 40}

	sort.Ints(ages) //this method sort & alter the original slice in increasing order
	fmt.Println(ages)

	index := sort.SearchInts(ages, 11)
	index2 := sort.SearchInts(ages, 52) // SearchInts searches for x in a sorted slice of ints and returns the index as specified by Search.
	//  The return value is the index to insert x if x is not present (it could be len(a)). The slice must be sorted in ascending order.
	fmt.Println(index, index2)

	names := []string{"rudy", "anwar", "bach", "lele", "inanna"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "adam"))
	fmt.Println(sort.SearchStrings(names, "inanna"))
}
