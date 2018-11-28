package main

import (
	"fmt"
	"os"
)

func basename(s string) string {
	// Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			// s = s[:i] //Discard current byte
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	s := os.Args[1]
	filename := basename(s)
	fmt.Printf("the basename is %s ", filename)
}
