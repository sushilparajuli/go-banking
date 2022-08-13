package app

import (
	"github.com/gorilla/mux"
	"github.com/sushilparajuli/go-banking/domain"
	"github.com/sushilparajuli/go-banking/service"
	"log"
	"net/http"
	"time"
)

func App() {

	r := mux.NewRouter()

	// Wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	r.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	// starting server
	srv := &http.Server{
		Handler: r,
		Addr:    ":9000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
