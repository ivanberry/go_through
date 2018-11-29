package main

import (
	"fmt"
)

// Currency
type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "^", GBP: "*", RMB: "#"}
	fmt.Println(RMB, symbol[RMB])

	a := [3]int{1, 2}
	fmt.Printf("%v", a)
}
