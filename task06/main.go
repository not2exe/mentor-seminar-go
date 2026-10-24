package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите количество чисел N (не меньше нуля):")
	if _, err := fmt.Scan(&n); err != nil || n < 0 {
		fmt.Println("Ошибка: N должно быть целым неотрицательным числом.")
		return
	}
	fmt.Printf("Введите %d чисел:\n", n)
	positive, negative, zeros := 0, 0, 0
	for i := 0; i < n; i++ {
		var value float64
		if _, err := fmt.Scan(&value); err != nil {
			fmt.Printf("Ошибка: число №%d введено неверно.\n", i+1)
			return
		}
		if value > 0 {
			positive++
		} else if value < 0 {
			negative++
		} else {
			zeros++
		}
	}
	fmt.Printf("Положительных: %d\nОтрицательных: %d\nНулей: %d\n", positive, negative, zeros)
}
