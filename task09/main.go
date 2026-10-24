package main

import "fmt"

func main() {
	var balance int
	fmt.Println("Введите начальный баланс в целых рублях:")
	if _, err := fmt.Scan(&balance); err != nil || balance < 0 {
		fmt.Println("Ошибка: баланс должен быть целым неотрицательным числом.")
		return
	}
	fmt.Println("1 — Посмотреть баланс\n2 — Пополнить баланс\n3 — Снять средства")
	var action int
	if _, err := fmt.Scan(&action); err != nil {
		fmt.Println("Ошибка: введите номер действия.")
		return
	}
	switch action {
	case 1:
		fmt.Printf("Баланс: %d руб.\n", balance)
	case 2, 3:
		var amount int
		fmt.Println("Введите сумму в целых рублях (больше нуля):")
		if _, err := fmt.Scan(&amount); err != nil || amount <= 0 {
			fmt.Println("Ошибка: сумма должна быть целым числом больше нуля.")
			return
		}
		if action == 2 {
			balance += amount
			fmt.Printf("Баланс после пополнения: %d руб.\n", balance)
		} else {
			if amount > balance {
				fmt.Printf("Недостаточно средств. Баланс: %d руб.\n", balance)
				return
			}
			balance -= amount
			fmt.Printf("Баланс после снятия: %d руб.\n", balance)
		}
	default:
		fmt.Println("Ошибка: неизвестное действие.")
	}
}
