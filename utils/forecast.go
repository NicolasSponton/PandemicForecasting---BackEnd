package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func ExponentialSmoothing(data []int, alpha float64, forecastDays int) []int {
	forecast := make([]int, len(data)+forecastDays)

	// Initialize the forecast with available data
	for i, value := range data {
		forecast[i] = value
	}

	for i := len(data); i < len(forecast); i++ {
		forecast[i] = int(alpha*float64(data[i-len(data)]) + (1-alpha)*float64(forecast[i-1]))
	}

	return forecast
}

func SaveCasesAsCSV(newCases []int) {
	tomorrow := time.Now().AddDate(0, 0, 1)
	filePath := "files/new_cases.csv"

	if err := os.MkdirAll("files", os.ModePerm); err != nil {
		fmt.Println("Error creating 'files' folder:", err)
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"date", "new_cases"}
	err = writer.Write(headers)
	if err != nil {
		fmt.Println("Error writing headers:", err)
		return
	}

	for _, cases := range newCases {
		dateStr := tomorrow.Format("2006-01-02")

		row := []string{dateStr, fmt.Sprintf("%d", cases)}
		err := writer.Write(row)
		if err != nil {
			fmt.Println("Error writing row:", err)
			return
		}

		tomorrow = tomorrow.AddDate(0, 0, 1)
	}
}
