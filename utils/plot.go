package utils

import (
	"os"

	"github.com/wcharczuk/go-chart"
)

func CreatePlotAndSave(dailyValues []int, futureValues []int) error {
	futureValues = append([]int{dailyValues[len(dailyValues)-1]}, futureValues...)

	err := os.MkdirAll("files", os.ModePerm)
	if err != nil {
		return err
	}

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

	file, err := os.Create("files/plot.png")
	if err != nil {
		return err
	}
	defer file.Close()

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
		futureXValues[i] = float64((startX - 1) + i)
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
