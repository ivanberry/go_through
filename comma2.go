package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// PASS
	for _, value := range os.Args[1:] {
		v := depart(value)
		fmt.Printf("the comma string is %s\n", v)
	}
}

func depart(s string) string {
	// divide the string into two parts
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		first, last := s[:dot], s[dot+1:]
		return coma(first) + "." + coma(last)
	}
	return coma(s)
}

func coma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return coma(s[:n-3]) + "," + s[n-3:]
}
