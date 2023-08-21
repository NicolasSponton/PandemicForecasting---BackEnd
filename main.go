package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	data, err := scrapeAndProcessData()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Data: ", data)

	// Sample data: daily COVID-19 cases
	cases := []float64{100, 120, 150, 180, 200, 220, 250, 280, 300, 320}

	// Initialize ARIMA model
	model, err := arima.NewARIMA(1, 1, 1)
	if err != nil {
		fmt.Println("Error creating ARIMA model:", err)
		return
	}

	// Fit ARIMA model to data
	err = model.Fit(cases)
	if err != nil {
		fmt.Println("Error fitting ARIMA model:", err)
		return
	}

	// Forecast the next 7 days
	forecast, err := model.Forecast(7)
	if err != nil {
		fmt.Println("Error forecasting:", err)
		return
	}

	fmt.Println("Forecast for the next 7 days:")
	for i, val := range forecast {
		fmt.Printf("Day %d: %.2f\n", i+1, val)
	}

}

func scrapeAndProcessData() ([]int, error) {
	dataUrl := "https://www.worldometers.info/coronavirus/country/south-africa/"

	res, err := http.Get(dataUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	htmlSource, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile("covid_data.html", htmlSource, 0644)
	if err != nil {
		return nil, err
	}

	fmt.Println("HTML source saved to covid_data.html")

	document, err := goquery.NewDocumentFromReader(strings.NewReader(string(htmlSource)))
	if err != nil {
		return nil, err
	}

	var dailyCasesData []int

	// Find the script tag containing the chart data
	scriptSelector := "script[type='text/javascript']"
	document.Find(scriptSelector).Each(func(index int, scriptHtml *goquery.Selection) {
		scriptText := scriptHtml.Text()

		if strings.Contains(scriptText, "Highcharts.chart('graph-cases-daily'") {
			// Find the start of the 'Daily Cases' data array
			dailyCasesIndex := strings.Index(scriptText, "name: 'Daily Cases',")
			dataStartIndex := strings.Index(scriptText[dailyCasesIndex:], "data: [") + dailyCasesIndex

			// Extract the 'Daily Cases' data array
			dataEndIndex := strings.Index(scriptText[dataStartIndex:], "]")
			dailyCasesDataStr := scriptText[dataStartIndex+7 : dataStartIndex+dataEndIndex]

			// Split the comma-separated string and process each element
			for _, valueStr := range strings.Split(dailyCasesDataStr, ",") {
				if valueStr == "null" {
					dailyCasesData = append(dailyCasesData, 0)
				} else {
					value, err := strconv.Atoi(valueStr)
					if err != nil {
						return
					}
					dailyCasesData = append(dailyCasesData, value)
				}
			}
		}
	})

	return dailyCasesData, nil
}
