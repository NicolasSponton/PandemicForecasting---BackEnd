package utils

import (
	"os"
	"time"

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
			Name:      "Date",
			NameStyle: chart.StyleShow(),
			ValueFormatter: func(v interface{}) string {
				if vf, isFloat := v.(float64); isFloat {
					today := time.Date(2023, time.August, 25, 0, 0, 0, 0, time.UTC)
					t := today.AddDate(0, 0, -len(dailyValues)+int(vf))
					return t.Format("2006-01-02")
				}
				return ""
			},
			TickPosition: chart.TickPositionBetweenTicks,
			TickStyle:    chart.StyleShow(),
			Style:        chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Cases",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
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
