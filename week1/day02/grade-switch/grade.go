package main

import "fmt"

func main() {
	score := 73
	switch {
	case score >= 120:
		fmt.Println("Отлично")
	case score >= 100:
		fmt.Println("Хорошо")
	case score >= 75:
		fmt.Println("Проходной балл")
	default:
		fmt.Println("нужно пересдавать")
	}
}
