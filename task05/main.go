package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите количество чисел N больше нуля:")
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		fmt.Println("Ошибка: N должно быть целым числом больше нуля.")
		return
	}
	fmt.Printf("Введите %d чисел:\n", n)
	var minimum, maximum float64
	for i := 0; i < n; i++ {
		var value float64
		if _, err := fmt.Scan(&value); err != nil {
			fmt.Printf("Ошибка: число №%d введено неверно.\n", i+1)
			return
		}
		if i == 0 {
			minimum, maximum = value, value
		}
		if value < minimum {
			minimum = value
		}
		if value > maximum {
			maximum = value
		}
	}
	fmt.Printf("Максимум: %g\nМинимум: %g\n", maximum, minimum)
}
