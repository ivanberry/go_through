package main

import "fmt"

// type tree struct {
// 	value       int
// 	left, right *tree
// }

type Employee struct {
	Salary        int
	Name, Address string
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

func main() {
	someone := Employee{Name: "test", Salary: 100}
	fmt.Println(someone)

	some1 := AwardAnnualRaise(*someone)
	fmt.Println(some1)
	fmt.Println(someone)
}

// func sort(values []int) {
// 	var root *tree
// 	for _, v := range values {
// 		root = add(root, v)
// 	}
// 	appendValues(values[:0], root)
// }

// func appendValues(values []int, t *tree) []int {
// 	if t != nil {
// 		values = appendValues(values, t.left)
// 		values = append(values, t.values)
// 		values = appendValues(values, t.right)
// 	}
// 	return values
// }

// func add(t *tree, value int) *tree {
// 	if t == nil {
// 		t = new(tree)
// 		t.value = value
// 		return t
// 	}
// 	if value < t.value {
// 		t.left = add(t.left, value)
// 	} else {
// 		t.right = add(t.right, value)
// 	}
// 	return t
// }
