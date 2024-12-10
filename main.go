package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix, err := readMatrix()
	if err != nil {
		fmt.Println("Ошибка при чтении матрицы:", err)
		return
	}

	avg := getAverage(&matrix)
	fmt.Println("Среднее матрицы:", avg)
}

func getAverage(matrix *[][]int) float32 {
	elementsCount := len(*matrix) * len((*matrix)[0])
	sum := 0

	for _, row := range *matrix {
		for _, elem := range row {
			sum += elem
		}
	}

	return float32(sum) / float32(elementsCount)
}

func readMatrix() ([][]int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите количество строк: ")
	rowsInput, _ := reader.ReadString('\n')
	rowsInput = strings.TrimSpace(rowsInput)
	rows, err := strconv.Atoi(rowsInput)
	if err != nil {
		fmt.Printf("Ошибка преобразования количества строк: %v\n", err)
		return nil, err
	}

	fmt.Println("Введите количество столбцов: ")
	colsInput, _ := reader.ReadString('\n')
	colsInput = strings.TrimSpace(colsInput)
	cols, err := strconv.Atoi(colsInput)
	if err != nil {
		fmt.Printf("Ошибка преобразования количества столбцов: %v\n", err)
		return nil, err
	}

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		fmt.Printf("Введите %d элемент(-ов/-а) для строки %d (через пробел): ", cols, i+1)
		rowInput, _ := reader.ReadString('\n')
		rowInput = strings.TrimSpace(rowInput)
		elements := strings.Split(rowInput, " ")

		if len(elements) != cols {
			fmt.Printf("Ошибка ввода элементов. Ожидалось: %d, получено: %d.\n", cols, len(elements))
			return nil, err
		}

		for j := 0; j < cols; j++ {
			value, err := strconv.Atoi(elements[j])
			if err != nil {
				fmt.Printf("Ошибка преобразования элемента %s в число: %v\n", elements[j], err)
				return nil, err
			}
			matrix[i][j] = value
		}
	}

	return matrix, nil
}
