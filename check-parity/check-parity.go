package checkparity

import "fmt"

func ForMainCheckParity() {
	fmt.Println("Программа, которая проверяет четное число или нет")
	var num int
	fmt.Scan(&num)
	fmt.Println(check(num))
}

func check(num int) bool {
	result := num % 2
	if result == 0 {
		return true
	} else {
		return false
	}
}
