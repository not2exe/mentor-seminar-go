package main

import "fmt"

func main() {
	var a, b float64
	var operation string
	fmt.Println("Введите два числа, затем операцию (+, -, *, /):")
	if _, err := fmt.Scan(&a, &b, &operation); err != nil {
		fmt.Println("Ошибка: введите два числа и операцию.")
		return
	}

	var result float64
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неизвестная операция.")
		return
	}
	fmt.Printf("Результат: %g\n", result)
}
