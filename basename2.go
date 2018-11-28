package main

import (
	"fmt"
	"os"
	"strings"
)

func bansename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	for _, value := range os.Args[1:] {
		path := bansename(value)
		fmt.Printf("the pathname is %s\n", path)
	}
}
