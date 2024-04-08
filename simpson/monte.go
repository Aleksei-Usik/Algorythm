package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func function(x float64) float64 {
	//return 1/(5-4*math.Sin(x)) + x
	//return (math.Sin(x) * math.Sin(x)) / (9 + 0.3*math.Cos(x))
	return x * x * x * x * math.Log(x+math.Sqrt(x*x-0.36))
	//return math.Sin(x)
}

func main() {
	var a, b float64
	var numSamples int
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("Точек = ")
	fmt.Scan(&numSamples)
	sum := 0.0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numSamples; i++ {
		x := a + rand.Float64()*(b-a)
		sum += function(x)
	}

	integral := sum / float64(numSamples) * (b - a)

	fmt.Println("Ответ:", integral)
}
