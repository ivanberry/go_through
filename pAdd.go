package main

import "fmt"

func main() {
	b := 12
	bPoint := &b

	fmt.Printf("the address of  b is %p\n", bPoint)

	bValue := *bPoint
	fmt.Printf("the value of b is %d", bValue)
}
