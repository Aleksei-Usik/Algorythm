package main

import (
	"fmt"
	"math"
	"math/rand"
)

func f(x float64) float64 {
	//return 1 / (math.Abs(x) + 0.01) // Пиковая
	//return math.Sin(x) // Периодическая
	//return 1/(5-4*math.Sin(x)) + x
	return (math.Sin(x) * math.Sin(x)) / (9 + 0.3*math.Cos(x))
	//return x * x * x * x * math.Log(x+math.Sqrt(x*x-0.36))
}

func max_min(a, b float64, hag int) (float64, float64) {
	h := b - a
	n := int(b-a) * hag
	rand_num := f(a + rand.Float64()*h)
	max := rand_num
	min := max
	for i := 0; i < n; i++ {
		rand_num = f(a + rand.Float64()*h)
		if rand_num > max {
			max = rand_num
		}
		if rand_num < min {
			min = rand_num
		}
	}
	return max, min
}

func monte(a, b, min, max float64, h int) float64 {
	n_p, n_m, count_p, count_m, S_p, S_m := 1, 1, 0, 0, 0.0, 0.0
	if max >= 0 {
		S_p = math.Abs((b - a) * max)
		n_p = int(S_p * float64(h))
	}
	if min < 0 {
		S_m = math.Abs((b - a) * min)
		n_m = int(S_m * float64(h))
	}
	for i := 0; i < n_p+n_m; i++ {
		rand_x := f(a + rand.Float64()*(b-a))
		rand_y := min + rand.Float64()*(max-min)
		if rand_x >= 0 {
			if rand_y < rand_x && rand_y >= 0 {
				count_p++
			}
		} else {
			if rand_y > rand_x && rand_y <= 0 {
				count_m++
			}
		}
	}
	return S_p*(float64(count_p)/float64(n_p)) - S_m*(float64(count_m)/float64(n_m))
}

func main() {
	var a, b, result, e, tmp float64
	var h, n int
	fmt.Println("Ведите границы функции, плотность точек и колличество итераций")
	fmt.Scan(&a, &b, &h, &n)
	if b < a || h <= 0 {
		fmt.Println("Ошибка, введённые данные некорректны")
		return
	}
	for i := 1; i <= n; i++ {
		max, min := max_min(a, b, h)
		tmp = monte(a, b, min, max, h)
		result += tmp
		e += tmp * tmp
	}
	e /= float64(n)
	result /= float64(n)
	e = math.Sqrt(math.Abs(e - result*result))
	fmt.Println("Интеграл от", a, " до ", b, " == ", result, " +-", e)
}
