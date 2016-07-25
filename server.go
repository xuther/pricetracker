package main

import "fmt"

var config = configuration{}

func main() {

	val, err := getPrice("https://braddockandlogan-apts.securecafe.com/onlineleasing/rock-creek-ridge-townhomes/oleapplication.aspx?stepname=Floorplan&myOlePropertyid=45723",
		`<td *data-selenium-id *= *"Rent_1">\$([0-9,]+)-.+?<\/td>`)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	} else {
		fmt.Printf("Price: %s\n", val)
	}

	c, err := importConfiguration("./config.json")
	config = c

	test := location{Name: "North Bend Townhomes No Den",
		Description: "The lowest price townhomes in north bend. No den.",
		Address:     "https://braddockandlogan-apts.securecafe.com/onlineleasing/rock-creek-ridge-townhomes/oleapplication.aspx?stepname=Floorplan&myOlePropertyid=45723",
		Regex:       `<td *data-selenium-id *= *"Rent_1">\$([0-9,]+)-.+?<\/td>`}
	err = insertLocation(&test)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	/*
		if err != nil {
			fmt.Printf("Could not import the configuration file, check that it exists: %s\n", err.Error())
		}
		config = c

		e := echo.New()
		e.Pre(middleware.RemoveTrailingSlash())

		e.Run(standard.New(":8888"))

		//eventually we need to accept command line parameters.
	*/
}
