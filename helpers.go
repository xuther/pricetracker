package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func importConfiguration(configPath string) (configuration, error) {
	fmt.Printf("Importing the configuration information from %v\n", configPath)

	f, err := ioutil.ReadFile(configPath)
	var c configuration

	if err != nil {
		fmt.Printf("Done. Error %s.\n", err.Error())
		return c, err
	}

	json.Unmarshal(f, &c)

	fmt.Printf("Done.\n")

	return c, err
}

//Returns true if all succeeded, else false.
func checkSuccess(res []updateResults, print bool) bool {
	var status bool
	for _, val := range res {
		if !val.Success {
			status = false
			if !print {
				return status
			}
		}
		if print {
			fmt.Printf("%v  -  %v\n", val.LocationID, val.Success)
		}
	}

	return status
}
