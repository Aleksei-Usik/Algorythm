package main

import (
	"fmt"
	"math"
	"time"
)

func form(x float64) float64 {
	//return 1 / (math.Abs(x) + 0.01) //Пиковая
	return (math.Sin(x) * math.Sin(x)) / (9 + 0.3*math.Cos(x))
	//return x * x * x * x * math.Log(x+math.Sqrt(x*x-0.36))
	//return x * x //Возрастающая
	//return math.Sin(x) //Периодическая
	//return (math.Sin(x) * math.Sin(0.1*x)) / (math.Abs(x) + 0.01)
}

func main() {
	var x_n, result0, result1, delta, e, a, b float64
	var n, mnoj int
	result0 = 0
	result1 = 0
	n = 5
	fmt.Scan(&a, &b, &e)
	if b < a || n < 2 {
		fmt.Println("Ошмбка, введённые данные некорректны")
		return
	}
	start := time.Now()
	for i := 1; i <= 25; i++ {
		x_n = a
		result1 = form(a)
		mnoj = 4
		for j := 1; j < n; j++ {
			x_n += (b - a) / float64(n)
			result1 += form(x_n) * float64(mnoj)
			if mnoj == 4 {
				mnoj = 2
			} else {
				mnoj = 4
			}
		}
		result1 += form(b)
		result1 *= (b - a) / (3.0 * float64(n))
		delta = (result1 - result0) / 15
		if math.Abs(delta) <= e {
			fmt.Println(result1 + delta)
			elapsed := time.Since(start)
			fmt.Println(elapsed)
			return
		}
		result0 = result1
		n *= 2
	}
	fmt.Println("Заданная точность не достигнута", result1+delta)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
