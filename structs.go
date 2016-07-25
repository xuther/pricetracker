package main

type configuration struct {
	DBAddress string
}

type location struct {
	ID          int64
	Name        string
	Description string
	Address     string
	Regex       string
}
