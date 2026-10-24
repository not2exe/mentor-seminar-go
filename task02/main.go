package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите целое число N:")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: требуется целое число.")
		return
	}
	if n > 0 {
		fmt.Println("Число положительное.")
	} else if n < 0 {
		fmt.Println("Число отрицательное.")
	} else {
		fmt.Println("Число равно нулю.")
	}
	if n%2 == 0 {
		fmt.Println("Число чётное.")
	} else {
		fmt.Println("Число нечётное.")
	}
	if n%3 == 0 {
		fmt.Println("Делится на 3: да.")
	} else {
		fmt.Println("Делится на 3: нет.")
	}
	if n%5 == 0 {
		fmt.Println("Делится на 5: да.")
	} else {
		fmt.Println("Делится на 5: нет.")
	}
}
