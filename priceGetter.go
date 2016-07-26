package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getAllPrices() ([]price, error) {
	locs, err := getLocations()
	if err != nil {
		printErr(err, "Getting locations")
		return []price{}, nil
	}

	channel := make(chan locPrice, len(locs))
	for _, loc := range locs {
		go getPriceAsync(loc.Address, loc.Regex, loc.ID, channel)
	}

	return gatherPrices(channel, len(locs))
}

func gatherPrices(receivingChan <-chan locPrice, dispatched int) ([]price, error) {
	var prices []price

	for i := 0; i < dispatched; i++ {
		curPrice := <-receivingChan
		if curPrice.Success {
			fmt.Printf("Got price for %v, price: %v\n", curPrice.LocationID, curPrice.Price)
			prices = append(prices, price{LocationID: curPrice.LocationID, Price: curPrice.Price, AccessDate: time.Now()})
		}
	}

	return prices, nil
}

func updatePrices() ([]updateResults, error) {
	results := []updateResults{}

	locs, err := getLocations()
	if err != nil {
		printErr(err, "Getting locations")
		return results, err
	}

	for _, loc := range locs {
		status := true
		err = getAndStorePrice(loc)

		if err != nil {
			status = false
		} else {
			status = true
		}
		results = append(results, updateResults{status, loc.ID, err})
	}
	return results, nil
}

func getAndStorePrice(loc location) error {
	fmt.Printf("Getting and storing the price for %v:%v\n", loc.ID, loc.Name)
	p, err := getPrice(loc.Address, loc.Regex)
	if err != nil {
		return err
	}
	return insertPrice(&price{LocationID: loc.ID, Price: p, AccessDate: time.Now()})
}

func getPriceAsync(address string, regex string, locID int64, toSend chan<- locPrice) {
	vals, err := getPrice(address, regex)
	if err != nil {
		toSend <- locPrice{-1, -1, false}
	} else {
		toSend <- locPrice{vals, locID, true}
	}
}

//get the webpage from address, run the regex on it and return the first capture group as the price
func getPrice(address string, regex string) (int64, error) {
	resp, err := http.Get(address)
	if err != nil {
		printErr(err, "Getting the webpage")
		return -1, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		printErr(err, "decoding the response")
		return -1, err
	}

	re := regexp.MustCompile(regex)
	matches := re.FindStringSubmatch(string(bytes))

	if matches == nil {
		err = errors.New("No match found.")
		return -1, err
	}

	fmt.Printf("%v\n", matches)

	price, err := strconv.Atoi(strings.Replace(matches[1], ",", "", -1))
	if err != nil {
		printErr(err, "parsing the price from the regex results")
		return -1, err
	}
	return int64(price), nil
}

func printErr(err error, loc string) {
	fmt.Printf("ERRROR in %s: %s\n", loc, err.Error())
}
