package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите целое N (не меньше нуля):")
	if _, err := fmt.Scan(&n); err != nil || n < 0 {
		fmt.Println("Ошибка: N должно быть целым неотрицательным числом.")
		return
	}
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}
	fmt.Printf("Сумма подходящих чисел: %d\n", sum)
}
