package main

import (
	handlers "group/handler"
	"net/http"
)

func main() {
	// Register the /dates handler
	http.HandleFunc("/api/dates", handlers.DatesHandler)
	// Add other routes here if needed

	// Start the server
	http.ListenAndServe(":8080", nil)
}
