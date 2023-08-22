package main

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart"
)

type CaseData struct {
	Daily    []int     `json:"daily"`
	Forecast []float64 `json:"forecast"`
}

func main() {

	// data, err := utils.ScrapeAndProcessData()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	// dailyCases := []int{100, 120, 130, 110, 105, 125, 130, 115, 135, 140}

	// forecastCases := utils.ExponentialSmoothing(dailyCases, 0.2, 7)
	// utils.SaveCasesAsCSV(forecastCases)

	// Example input data
	dailyValues := []int{10, 15, 12, 18, 20, 22, 25}
	futureValues := []int{30, 35, 40, 45}

	// Call the function to create the plot and save it
	err := createPlotAndSave(dailyValues, futureValues)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Plot created and saved successfully.")
	}

	// server.StartServer()

}

func createPlotAndSave(dailyValues []int, futureValues []int) error {
	// Create the "files" directory if it doesn't exist
	err := os.MkdirAll("files", os.ModePerm)
	if err != nil {
		return err
	}

	// Create a new chart
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Days",
			NameStyle: chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Values",
			NameStyle: chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    "Daily Values",
				XValues: generateXValues(len(dailyValues)),
				YValues: generateFloatYValues(dailyValues),
			},
			chart.ContinuousSeries{
				Name:    "Future Values",
				XValues: generateFutureXValues(len(dailyValues), len(futureValues)),
				YValues: generateFloatYValues(futureValues),
			},
		},
	}

	// Create a PNG file
	file, err := os.Create("files/plot.png")
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the chart to the PNG file
	err = graph.Render(chart.PNG, file)
	if err != nil {
		return err
	}

	return nil
}

func generateXValues(numDays int) []float64 {
	xValues := make([]float64, numDays)
	for i := 0; i < numDays; i++ {
		xValues[i] = float64(i)
	}
	return xValues
}

func generateFutureXValues(startX, numFutureDays int) []float64 {
	futureXValues := make([]float64, numFutureDays)
	for i := 0; i < numFutureDays; i++ {
		futureXValues[i] = float64(startX + i)
	}
	return futureXValues
}

func generateFloatYValues(intValues []int) []float64 {
	floatValues := make([]float64, len(intValues))
	for i, v := range intValues {
		floatValues[i] = float64(v)
	}
	return floatValues
}
