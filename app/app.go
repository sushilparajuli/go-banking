package app

import (
	"github.com/sushilparajuli/go-banking/app/handlers"
	"log"
	"net/http"
)

func App() {

	mux := http.NewServeMux()
	mux.HandleFunc("/greet", handlers.Greet)

	mux.HandleFunc("/customers", handlers.GetAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:9000", mux))
}
