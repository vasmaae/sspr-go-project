package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Передайте программе 2 аргумента: IP-адрес и порт сервера")
		return
	}

	matrix, err := readMatrix()
	if err != nil {
		fmt.Println("Ошибка при чтении матрицы:", err)
		return
	}

	for _, row := range matrix {
		fmt.Println(row)
	}

	ip := os.Args[1]
	port := os.Args[2]
	resp, err := sendMatrix(ip, port, &matrix)
	if err != nil {
		fmt.Printf("Ошибка при отправке: %v\n", err)
		return
	}

	responseMessage := getResponse(resp)
	fmt.Println("Ответ от сервера:", responseMessage)
}

func getResponse(resp *http.Response) string {
	var responseMessage string
	if resp.StatusCode == http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		responseMessage = buf.String()
	} else {
		responseMessage = fmt.Sprintf("Ошибка: %s", resp.Status)
	}

	return responseMessage
}

func sendMatrix(ip string, port string, matrix *[][]int) (*http.Response, error) {
	jsonMatrix, err := json.Marshal(*matrix)
	if err != nil {
		fmt.Printf("Ошибка сериализации матрицы: %v\n", err)
		return nil, err
	}

	address := "http://" + ip + ":" + port
	resp, err := http.Post(address, "application/json", bytes.NewBuffer(jsonMatrix))
	if err != nil {
		fmt.Printf("Ошибка отправки запроса: %v\n", err)
		return nil, err
	}

	return resp, nil
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
