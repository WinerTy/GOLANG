package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	// Разбиваем число на цифры и сохраняем их в массив
	digits := [3]int{number / 100, (number / 10) % 10, number % 10}

	// Проверяем, все ли цифры различны
	if digits[0] != digits[1] && digits[0] != digits[2] && digits[1] != digits[2] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
