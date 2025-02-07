package stringcomparison

import (
	"fmt"
	"strings"
)

func ForMainStringComparsion() {
	fmt.Println("Программа, которая сравнивает две строки, введенные пользователем, и выводит, равны ли они или нет, игнорируя регистр")
	fmt.Println(stringcomparison())
}

func stringcomparison() bool {
	var str1, str2 string
	fmt.Scanln(&str1, &str2)
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	result := str1 == str2
	return result
}
