package main

import (
	"fmt"
	"math"
)

func rem_empty(matrix [][]float64) [][]float64 { //Удаление пустых строк
	// Используем индекс для отслеживания новой позиции в матрице
	index := 0

	// Проходим по каждой строке в матрице
	for _, row := range matrix {
		// Проверяем, не является ли строка пустой
		notEmpty := false
		for _, j := range row {
			if j != 0 {
				notEmpty = true
				break
			}
		}

		// Если строка не пустая, перемещаем ее на новую позицию
		if notEmpty {
			matrix[index] = row
			index++
		}
	}

	matrix = matrix[:index]
	return matrix
}

func non_validation(matrix [][]float64) bool { //Проверка определенности и правильности матрицы
	var flag bool
	n := len(matrix)
	if n == 0 || len(matrix[0]) < 2 || n < len(matrix[0])-1 {
		return true
	}

	for i := 0; i < n-1; i++ {
		if len(matrix[i]) != len(matrix[i+1]) {
			return true
		}
	}

	for _, row := range matrix {
		flag = true
		for i := 0; i < len(row)-1; i++ {
			if row[i] != 0 {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		flag = true
		for i := 0; i < n; i++ {
			if matrix[i][j] != 0 {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

func x_Det(square, matrix []float64, j int) float64 { //Определение x
	sum := matrix[len(matrix)-1]
	for i := 0; i < len(square); i++ {
		if i != j {
			sum -= square[i] * matrix[i]
		}
	}
	return sum / matrix[j]
}

func d_max(vector, vector_l []float64) float64 {
	var max float64
	for i := 0; i < len(vector); i++ {
		tmp := math.Abs(vector[i] - vector_l[i])
		if max < tmp {
			max = tmp
		}
	}
	return max
}

func iter(matrix [][]float64, e float64, k int) []float64 {
	h := len(matrix[0]) - 1
	eq_l := make([]float64, h)
	eq := make([]float64, h)
	for j := 0; j < k; j++ {
		copy(eq_l, eq)
		for i := 0; i < h; i++ {
			eq[i] = x_Det(eq_l, matrix[i], i)
		}
		if d_max(eq, eq_l) <= e {
			return eq
		}
	}
	return eq
}

func main() {

	matrix := [][]float64{
		{3, 1, -1, -2},
		{1, 5, -1, 8},
		{20, 0, 3, 1},
	}

	//var e float64

	if non_validation(rem_empty(matrix)) {
		fmt.Println("Матрица заполнена неверно или она неопределенна")
		return
	}

	//fmt.Scan(&e)

	for i, x := range iter(matrix, 10e-4, 100) {
		fmt.Println("X_", i+1, " == ", x)
	}
}
