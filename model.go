package main

import (
	"database/sql"
	"fmt"
)

type product struct {
	ID    int     `json:"id"`
	Code  any     `json:"code"` //temporary set to any coz it will be generated later by API add product
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// fetch all products from db
func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, code, name, price FROM products")

	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	fmt.Println(err)

	defer rows.Close()

	// set initial value of products as empty slice
	products := []product{}

	// loop through data and append to products variable
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Code, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
