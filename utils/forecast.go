package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func SimpleMovingAverage(dailyCases []int, numDays int) []int {
	forecast := make([]int, numDays)

	lastNumDays := len(dailyCases)
	if lastNumDays >= numDays {
		for i := 0; i < numDays; i++ {
			sum := 0
			for j := lastNumDays - numDays + i; j < lastNumDays; j++ {
				sum += dailyCases[j]
			}
			forecast[i] = sum / numDays
		}
	}

	return forecast
}

func SaveCasesAsCSV(newCases []int) {
	tomorrow := time.Now().AddDate(0, 0, 1)
	filePath := "files/new_cases.csv"

	if err := os.MkdirAll("files", os.ModePerm); err != nil {
		fmt.Println(err.Error())
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"date", "new_cases"}
	err = writer.Write(headers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, cases := range newCases {
		dateStr := tomorrow.Format("2006-01-02")

		row := []string{dateStr, fmt.Sprintf("%d", cases)}
		err := writer.Write(row)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		tomorrow = tomorrow.AddDate(0, 0, 1)
	}
}
