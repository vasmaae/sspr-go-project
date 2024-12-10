package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func MatrixHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var matrix [][]int
	err := json.NewDecoder(r.Body).Decode(&matrix)
	if err != nil {
		http.Error(w, "Ошибка декодирования json", http.StatusBadRequest)
		return
	}

	fmt.Println(time.Now())
	for _, row := range matrix {
		fmt.Println(row)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Матрица успешно получена")
	avg := getAverage(&matrix)
	fmt.Fprintf(w, "Среднее арифметическое матрицы: %.2f", avg)
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

func main() {
	http.HandleFunc("/", MatrixHandler)

	fmt.Println("Сервер запущен на http://127.0.0.1:1337")
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера: ", err)
	}
}
