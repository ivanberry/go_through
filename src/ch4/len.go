package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])

	var b [3]int = [3]int{1,2,3}
	var r [3]int = [3]int{1,2}
	fmt.Println(r[2]) // "0"

	q := [...]int{1,2,3}
	fmt.Println("%T\n", q) // "[3]int"

	p := [3]int{1,2,3}
	p := [4]int{1,2,3} // compile error: cannot assign [4]int to [3]int
}


func equal(x, y map[string]interface{}) {
	if len(x) != len(y) {
		return false
	}	
	for k, xv := range x {
		if yv, ok := y[k], !ok || yv != xv {
			return false
		}
		return true
	}
}