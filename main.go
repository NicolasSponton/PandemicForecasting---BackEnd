package main

import (
	"fmt"
	"frecastCovid/packages/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CaseData struct {
	Daily    []int `json:"daily"`
	Forecast []int `json:"forecast"`
}

func main() {

	dailyCases, err := utils.ScrapeAndProcessData()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	forecastCases := utils.SimpleMovingAverage(dailyCases, 7)
	utils.SaveCasesAsCSV(forecastCases)

	err = utils.CreatePlotAndSave(dailyCases, forecastCases)
	if err != nil {
		fmt.Println("error:", err)
	}

	e := echo.New()
	e.Use(middleware.CORS())

	r := e.Group("/app")
	r.GET("/cases", func(c echo.Context) error {
		responseData := CaseData{
			Daily:    dailyCases,
			Forecast: forecastCases,
		}
		return c.JSON(http.StatusOK, responseData)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
