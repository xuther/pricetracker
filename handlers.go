package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/labstack/echo"
)

func getAllPricesHandler(c echo.Context) error {
	prices, err := getAllPrices()
	if err != nil {
		return err
	}
	return sendJSONResponse(&prices, c)
}

func getAllLocationsHandler(c echo.Context) error {
	locs, err := getLocations()
	if err != nil {
		return err
	}
	return sendJSONResponse(&locs, c)
}

func getLocationByIDHandler(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("LocationID"), 10, 64)
	if err != nil {
		return err
	}
	loc, err := getLocationByID(id)
	if err != nil {
		return err
	}
	return sendJSONResponse(&loc, c)
}

func addLocationHandler(c echo.Context) error {
	bits, err := ioutil.ReadAll(c.Request().Body())
	if err != nil {
		return err
	}
	var loc location
	err = json.Unmarshal(bits, &loc)
	if err != nil {
		return err
	}

	err = insertLocation(&loc)
	if err != nil {
		return err
	}

	m := make(map[string]string)
	m["message"] = "success"
	return sendJSONResponse(&m, c)
}

func updatePricesHandler(c echo.Context) error {
	results, err := updatePrices()
	if err != nil {
		return err
	}
	checkSuccess(results, true)
	//TODO retry values that fail
	return sendJSONResponse(&results, c)
}

func getLocationPriceByIDHandler(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("LocationID"), 10, 64)
	if err != nil {
		return err
	}
	locationPrice, err := getLastPriceFromDBByID(id)
	return sendJSONResponse(&locationPrice, c)
}

func sendJSONResponse(toMarshal interface{}, c echo.Context) error {
	b, err := json.Marshal(toMarshal)
	if err != nil {
		return err
	}

	c.Response().Write(b)
	return nil
}
