package main

import (
	"fmt"
)

func main()  {
	fmt.Print(gcd(10, 12))
	fmt.Print(fib(12))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int  {
	x, y := 0, 1
	for i :=0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}