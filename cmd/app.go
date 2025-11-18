package main

import (
	"fmt"
	"net/http"
	"pavlyysh/calculator/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Calculator)

	fmt.Println("starting server on port :8080")
	http.ListenAndServe(":8080", nil)
}
