package calculate

import (
	"fmt"
)

func ForMainCalculate() {
	fmt.Println("Программа-калькулятор")

	var num1, num2 float64
	var op string

	_, err := fmt.Scan(&num1, &op, &num2)
	if err != nil {
		fmt.Println("Ошибка ввода", err)
		return
	}

	result, err := Calculate(num1, num2, op)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	} else {
		fmt.Printf("Результат: %v", result)
	}
}

func Calculate(num1, num2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("На ноль делить нельзя")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("Неверный оператор: %s", op)
	}

}
