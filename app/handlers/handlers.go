package handlers

import (
	"encoding/json"
	"fmt"
	"go-banking/app/types"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []types.Customer{
		{Name: "Sushil", City: "Kathmandu", Zipcode: "111111"},
		{Name: "Puja", City: "Kathmandu", Zipcode: "00000"},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
