package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sushilparajuli/go-banking/service"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}

}
