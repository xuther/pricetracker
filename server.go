package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

var config = configuration{}

func main() {
	c, err := importConfiguration("./config.json")
	if err != nil {
		fmt.Printf("Could not import the configuration file, check that it exists: %s\n", err.Error())
	}
	config = c

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/api/locations/prices", getAllPricesHandler)
	e.GET("/api/locations", getAllLocationsHandler)
	e.GET("/api/locations/:LocationID", getLocationByIDHandler)
	e.GET("/api/locations/:LocationID/prices", getLocationPriceByIDHandler)
	e.POST("/api/locations", addLocationHandler)
	//this might not be needed, do internally?
	e.POST("/api/locations/prices/update", updatePricesHandler)
	e.Run(standard.New(":8888"))

	//eventually we need to accept command line parameters.

}
