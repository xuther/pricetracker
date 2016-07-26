package main

import "time"

type configuration struct {
	DBType             string
	DBConnectionString string
}

type location struct {
	ID          int64
	Name        string
	Description string
	Address     string
	Regex       string
}

type price struct {
	ID         int64
	LocationID int64
	Price      int64
	AccessDate time.Time
}

type updateResults struct {
	Success    bool
	LocationID int64
	err        error
}

type locPrice struct {
	Price      int64
	LocationID int64
	Success    bool
}
