// Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

// Формат входных данных

// На вход подается номер билета - одно шестизначное  число.

// Формат выходных данных
// Выведите "YES", если билет счастливый, в противном случае - "NO".

// Sample Input:

// 613244

// Sample Output:

// YES

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Пример входного номера билета
	var ticketNumber string

	fmt.Scan(&ticketNumber)

	// Проверка, что номер билета действительно шестизначный
	if len(ticketNumber) != 6 {
		fmt.Println("NO")
		return
	}

	// Разделяем номер на две части
	firstHalf := ticketNumber[:3]
	secondHalf := ticketNumber[3:]

	// Вычисляем сумму цифр для первой половины
	sumFirstHalf := sumDigits(firstHalf)

	// Вычисляем сумму цифр для второй половины
	sumSecondHalf := sumDigits(secondHalf)

	// Сравниваем суммы
	if sumFirstHalf == sumSecondHalf {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// Функция для вычисления суммы цифр в строке
func sumDigits(s string) int {
	sum := 0
	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		sum += digit
	}
	return sum
}
