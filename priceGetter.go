package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//get the webpage from address, run the regex on it and return the first capture group as the price
func getPrice(address string, regex string) (string, error) {
	resp, err := http.Get(address)
	if err != nil {
		printErr(err, "Getting the webpage")
		return "", err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		printErr(err, "decoding the response")
		return "", err
	}

	re := regexp.MustCompile(regex)
	matches := re.FindStringSubmatch(string(bytes))

	if matches == nil {
		err := errors.New("No match found.")
		return "", err
	}

	fmt.Printf("%v\n", matches)

	return matches[1], nil
}

func printErr(err error, loc string) {
	fmt.Printf("ERRROR in %s: %s\n", loc, err.Error())
}
