package main

import "fmt"

// array is a fixed-length sequence of zero or more elements of a particular type. Because of their fixed lenth, arryas are rarely used directly in Go. Slice, which can grow and shrink, are much more versatile, but to understand slices we must understand array first.

func main() {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) //print the last element

	// Print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// initialize with array literal
	var q [3]int = [3]int{1, 2, 3}
	for _, v := range q {
		fmt.Printf("%d\n", v)
	}

	// simplify
	d := [...]int{1, 2, 3, 4}
	for _, v := range d {
		fmt.Println(v)
	}
}
