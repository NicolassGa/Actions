package main

import (
	"fmt"

	"net/http"

	"strconv"
)

func Add(a, b int) int {

	return a + b

}

func additionHandler(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()

	aStr := values.Get("a")

	bStr := values.Get("b")

	a, err := strconv.Atoi(aStr)

	if err != nil {

		http.Error(w, "Parameter 'a' must be an integer", http.StatusBadRequest)

		return

	}

	b, err := strconv.Atoi(bStr)

