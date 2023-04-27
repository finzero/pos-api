// main.go

package main

import (
	"fmt"
)

const (
	HOST        = "localhost"
	PORT        = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "password"
	DB_NAME     = "point_of_sale"
	API_PORT    = 8000
)

func main() {
	a := App{}
	a.Initialize(
		DB_USER,
		DB_PASSWORD,
		DB_NAME)

	fmt.Printf("listening at %v\n", API_PORT)
	a.Run(":8000")
	// a.CloseDB()
}
