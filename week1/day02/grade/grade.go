package main

import "fmt"

func main() {
	// здесь твои переменные
	score := 120
	if score >= 120 {
		fmt.Println("Отлично")
	} else if score >= 100 {
		fmt.Println("Хорошо")
	} else if score >= 75 {
		fmt.Println("Проходной балл")
	} else {
		fmt.Println("нужно пересдавать")
	}
}
