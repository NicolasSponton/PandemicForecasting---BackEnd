package utils

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeAndProcessData() ([]int, error) {
	dataUrl := "https://www.worldometers.info/coronavirus/country/south-africa/"

	res, err := http.Get(dataUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	htmlSource, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

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
