package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	dataUrl := "https://www.worldometers.info/coronavirus/country/south-africa/"

	res, err := http.Get(dataUrl)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()

	htmlSource, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = ioutil.WriteFile("covid_data.html", htmlSource, 0644)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("HTML source saved to covid_data.html")

	document, err := goquery.NewDocumentFromReader(strings.NewReader(string(htmlSource)))
	if err != nil {
		log.Fatal(err)
	}

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
			dailyCasesData := scriptText[dataStartIndex+7 : dataStartIndex+dataEndIndex]

			fmt.Println(dailyCasesData)
		}
	})

}
