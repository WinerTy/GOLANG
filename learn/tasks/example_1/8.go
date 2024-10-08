// Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.

// Sample Input:

// 1 5

// Sample Output:

// 15

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	var sum int
	for i:= a; i <= b; i++{
		sum+=i
	}
	fmt.Println(sum)

}

// 1 2 3 4 5 1+2+3+4+5
