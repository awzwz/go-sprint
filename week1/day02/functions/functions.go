package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func isEven(n int) bool {
	return n%2 == 0
}

func greet(name string) {
	fmt.Printf("Привет, %s!\n", name)
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {

	fmt.Println(add(3, 5))
	fmt.Println(isEven(10))
	greet("Aziz")
	q, r := divide(17, 5)
	fmt.Printf("%d остаток %d\n", q, r)
}
