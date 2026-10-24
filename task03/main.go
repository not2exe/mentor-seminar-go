package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите целое N больше нуля:")
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		fmt.Println("Ошибка: N должно быть целым числом больше нуля.")
		return
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	average := float64(sum) / float64(n)
	fmt.Printf("Сумма: %d\nСреднее: %g\n", sum, average)
}
