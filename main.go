package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/v1/precipitation", func(c echo.Context) error {

		// API endpoint.
		endpoint := "https://www.yr.no/api/v0/locations/1-72837/forecast/now"

		// Get data from the API.
		response, err := http.Get(endpoint)

		// Guard for empty response.
		if err != nil {
			log.Fatal(err.Error())
		}

		// Get the JSON data out of the response.
		responseData, err := ioutil.ReadAll(response.Body)

		// Guard for something wrong with reading the content.
		if err != nil {
			log.Fatal(err.Error())
		}

		var yr Yr
		_ = json.Unmarshal(responseData, &yr)

		var doesItRain = yr.Points[1].Precipitation.Intensity > 0
		var lastUpdatedAt = yr.Points[1].Time

		return c.JSON(http.StatusOK, struct {
			DoesItRain   bool
			DataFromTime string
		}{DoesItRain: doesItRain, DataFromTime: lastUpdatedAt})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	log.Fatal(e.Start(":" + httpPort))
}
