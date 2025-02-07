package typeconversion

import (
	"fmt"
	"math"
)

func ForMainTypeConversion() {
	fmt.Println("Программа, которая преобразует число с плавающей точкой в целое число")
	var num float64
	fmt.Scan(&num)
	fmt.Println(typeconversion(num))
}

func typeconversion(num float64) float64 {
	result := math.Round(num)
	return result
}
