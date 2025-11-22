package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Создаем файл для записи
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Создаем новый CSV писатель
	writer := csv.NewWriter(file)

	// Пример данных для записи
	data := [][]string{
		{"Name", "Age", "Occupation"},
		{"John Doe", "30", "Developer"},
		{"Jane Doe", "29", "Designer"},
	}

	// Записываем все данные в CSV
	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	// Записываем буфер в файл
	writer.Flush()

	fmt.Println("CSV file created successfully")
}

// https://purpleschool.ru/knowledge-base/article/csv/
