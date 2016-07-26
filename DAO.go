package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

func getLocationByID(id int64) (location, error) {
	fmt.Printf("Getting location %v\n", id)

	db, err := sql.Open(config.DBType, config.DBConnectionString)
	if err != nil {
		return location{}, err
	}
	defer db.Close()

	stmt, err := db.Prepare("Select * FROM locations WHERE LocationID = ?")
	if err != nil {
		fmt.Printf("Couldn't prepare query.\n")
		return location{}, err
	}

	res, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("Couldn't execute query.\n")
		return location{}, err
	}

	res.Next()
	var locID int64
	var name string
	var description string
	var address string
	var regex string

	err = res.Scan(&id, &name, &description, &address, &regex)
	if err != nil {
		fmt.Printf("Couldn't scan for data. %s\n", err.Error())
		return location{}, err
	}

	return location{ID: locID,
		Name:        name,
		Description: description,
		Address:     address,
		Regex:       regex}, nil
}

func insertPrice(toInsert *price) error {
	fmt.Printf("Inserting price for %v into DB\n", toInsert.LocationID)

	db, err := sql.Open(config.DBType, config.DBConnectionString)
	if err != nil {
		fmt.Printf("Error %v\n", err.Error())
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT into prices (locationId, price, accessDate) values (?,?,?)`)
	if err != nil {
		fmt.Printf("Couldn't prepare query.\n")
		return err
	}

	res, err := stmt.Exec(toInsert.LocationID, toInsert.Price, toInsert.AccessDate)
	if err != nil {
		fmt.Printf("Error %v\n", err.Error())
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Error %v\n", err.Error())
		return err
	}

	toInsert.ID = id

	fmt.Printf("Done.\n")
	return nil
}

func getLastPriceFromDBByID(id int64) (price, error) {
	fmt.Printf("Getting last price from DB for %v\n", id)

	db, err := sql.Open(config.DBType, config.DBConnectionString)
	if err != nil {
		return price{}, err
	}
	defer db.Close()

	stmt, err := db.Prepare(`Select * FROM prices where accessDate in (
    select max(accessDate) as accessDate
    from prices group by locationID
    ) AND locationID = ?`)
	if err != nil {
		fmt.Printf("Couldn't prepare query.\n")
		return price{}, err
	}

	res, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("Couldn't execute query: %v.\n", err.Error())
		return price{}, err
	}

	if res.Next() {
		var Priceid int64
		var locationID int64
		var value int64
		var accessDate time.Time

		err = res.Scan(&Priceid, &locationID, &value, &accessDate)
		if err != nil {
			fmt.Printf("Couldn't scan for data.")
			return price{}, err
		}

		fmt.Printf("Done.\n")
		return price{
			ID:         Priceid,
			LocationID: locationID,
			Price:      value,
			AccessDate: accessDate}, nil
	}

	return price{}, errors.New("No price results found for the id" + string(id))
}

func getLocations() ([]location, error) {
	fmt.Printf("Getting all locations\n")
	locations := []location{}

	db, err := sql.Open(config.DBType, config.DBConnectionString)
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

	db, err := sql.Open(config.DBType, config.DBConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO locations (name, description, address, regex) values(?,?,?,?)")
	if err != nil {
		fmt.Printf("Couldn't prepare query.\n")
		return err
	}

	res, err := stmt.Exec(loc.Name, loc.Description, loc.Address, loc.Regex)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	loc.ID = id

	fmt.Printf("Done.\n")
	return nil
}
