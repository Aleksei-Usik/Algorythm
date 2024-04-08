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

func sorted_abs(matrix [][]float64, ind int) [][]float64 { // Пузырёк, но по модулю
	var timed []float64
	var flag = true
	var n = len(matrix)
	for flag {
		flag = false
		for i := ind; i < n-1; i++ {
			if math.Abs(matrix[i][ind]) < math.Abs(matrix[i+1][ind]) {
				timed = matrix[i+1]
				matrix[i+1] = matrix[i]
				matrix[i] = timed
				flag = true
			}
		}
	}
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

func gauss_fw(matrix [][]float64) [][]float64 { //Прямой ход
	matrix = rem_empty(matrix)
	n := len(matrix)
	h := len(matrix[0])
	for i := 0; i < n-1; i++ {
		matrix = sorted_abs(matrix, i)
		if matrix[i][i] == 0 {
			continue
		}
		for j := i + 1; j < n; j++ {
			factor := matrix[j][i] / matrix[i][i]
			for k := 0; k < h; k++ {
				matrix[j][k] -= factor * matrix[i][k]
			}
		}
	}

	return rem_empty(matrix)
}

func gauss_bw(matrix [][]float64) []float64 { //Обратный ход
	n := len(matrix[0]) - 1
	var sum float64
	var vector = make([]float64, n)
	for i := len(matrix) - 1; i >= 0; i-- {
		sum = matrix[i][n]
		for k := n - 1; k > i; k-- {
			sum -= vector[k] * matrix[i][k]
		}
		vector[i] = sum / matrix[i][i]
	}
	return vector
}

func main() {

	matrix := [][]float64{
		{3, 1, -1, -2},
		{1, 5, -1, 8},
		{2, 0, 3, 1},
	}

	if non_validation(matrix) {
		fmt.Println("Матрица заполнена неверно или она неопределенна")
		return
	}
	matrix = gauss_fw(matrix)
	for _, x := range matrix {
		fmt.Println(x)
	}
	for i, x := range gauss_bw(gauss_fw(matrix)) {
		fmt.Println("X_", i+1, " == ", x)
	}
}
