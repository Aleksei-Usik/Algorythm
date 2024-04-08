package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	//return 1 / (math.Abs(x) + 0.01) // Пиковая функция
	//return math.Sin(x) //Периодическая
	//return 1/(5-4*math.Sin(x)) + x
	return (math.Sin(x) * math.Sin(x)) / (9 + 0.3*math.Cos(x))
	//return x * x * x * x * math.Log(x+math.Sqrt(x*x-0.36))
}

func simpson(a, b float64) float64 {
	return (b - a) / 6 * (f(a) + 4*f((a+b)/2) + f(b))
}

func adaptiveSimpson(a, b, e float64, i, zetta int) float64 {
	simp := simpson(a, (a+b)/2) + simpson((a+b)/2, b)

	delta := (simp - simpson(a, b)) / 15
	if math.Abs(delta) <= e || i > zetta {
		return simp + delta
	}

	i++

	left := adaptiveSimpson(a, (a+b)/2, e/2, i, zetta)
	right := adaptiveSimpson((a+b)/2, b, e/2, i, zetta)

	return left + right
}

func main() {
	var a, b, e float64
	var zetta int
	fmt.Println("Ведите границы функции, точность и максимальную глубину рекурсии")
	fmt.Scan(&a, &b, &e, &zetta)
	if b < a || e <= 0 || zetta < 1 {
		fmt.Println("Ошибка, введённые данные некорректны")
		return
	}
	fmt.Println("Интеграл от ", a, "до", b, " = ", adaptiveSimpson(a, b, e, 1, zetta))
}
