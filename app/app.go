package app

import (
	"go-banking/app/handlers"
	"log"
	"net/http"
)

func App() {
	// define routes
	http.HandleFunc("/greet", handlers.Greet)

	http.HandleFunc("/customers", handlers.GetAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}
