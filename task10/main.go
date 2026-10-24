package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите количество целых чисел N больше нуля:")
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		fmt.Println("Ошибка: N должно быть целым числом больше нуля.")
		return
	}
	fmt.Printf("Введите %d целых чисел:\n", n)
	sum, minimum, maximum := 0, 0, 0
	even, odd := 0, 0
	for i := 0; i < n; i++ {
		var value int
		if _, err := fmt.Scan(&value); err != nil {
			fmt.Printf("Ошибка: число №%d должно быть целым.\n", i+1)
			return
		}
		sum += value
		if i == 0 {
			minimum, maximum = value, value
		}
		if value < minimum {
			minimum = value
		}
		if value > maximum {
			maximum = value
		}
		if value%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	average := float64(sum) / float64(n)
	fmt.Printf("Сумма: %d\nСреднее: %g\nМаксимум: %d\nМинимум: %d\nЧётных: %d\nНечётных: %d\n", sum, average, maximum, minimum, even, odd)
}
