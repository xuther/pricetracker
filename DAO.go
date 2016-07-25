package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func getLocations() ([]location, error) {
	fmt.Printf("Getting all locations\n")
	locations := []location{}

	db, err := sql.Open("sqlite3", "file:"+config.DBAddress+"?cache=shared&mode=rwc")
	if err != nil {
		return locations, err
	}
	defer db.Close()

	stmt, err := db.Prepare("Select * FROM locations")
	if err != nil {
		fmt.Printf("Couldn't prepare query.\n")
		return locations, err
	}

	res, err := stmt.Query()
	if err != nil {
		fmt.Printf("Couldn't execute query.\n")
		return locations, err
	}

	for res.Next() {
		var id int64
		var name string
		var description string
		var address string
		var regex string

		err = res.Scan(&id, &name, &description, &address, &regex)
		if err != nil {
			fmt.Printf("Couldn't scan for data.")
			continue
		}

		locations = append(locations, location{ID: id,
			Name:        name,
			Description: description,
			Address:     address,
			Regex:       regex})
	}

	return locations, nil
}

func insertLocation(loc *location) error {
	fmt.Printf("Inserting location %s\n", loc.Name)

	db, err := sql.Open("sqlite3", "file:"+config.DBAddress+"?cache=shared&mode=rwc")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO locations (name, description, address) values(?,?,?)")
	if err != nil {
		fmt.Printf("Couldn't prepare query.\n")
		return err
	}

	res, err := stmt.Exec(loc.Name, loc.Description, loc.Address)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	loc.ID = id

	fmt.Printf("Done.")
	return nil
}
