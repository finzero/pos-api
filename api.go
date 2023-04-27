package main

import (
	"net/http"
	"strconv"
)

// set response for client to fetch
func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	// prevent error by user input, set max limit to 10
	if count > 10 || count < 1 {
		count = 10
	}

	// prevent error by user input, set min page to 0
	if start < 0 {
		start = 0
	}

	products, err := getProducts(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}
