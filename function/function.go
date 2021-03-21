package main

import "fmt"

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func call_by_value(a int) {
	a = 10
}

func call_by_ref(a *int) {
	*a = 10
}

func main() {
	var n1 int
	n1 = 10

	var n2 = 20

	n3 := 30
	n4 := 40

	fmt.Println("sum : ", sum(n1, n2))

	result := sum(n3, n4)
	fmt.Println("sum : ", result)

	a := 0
	call_by_value(a)
	fmt.Println("a : ", a)

	call_by_ref(&a)
	fmt.Println("a : ", a)
}
