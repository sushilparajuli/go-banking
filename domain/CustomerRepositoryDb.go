package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sushilparajuli/go-banking/errors"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d *CustomerRepositoryDb) FindAll() ([]Customer, *errors.AppError) {

	findAllSql := "select * from customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while query customer table")
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer table" + err.Error())
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d *CustomerRepositoryDb) ById(id string) (*Customer, *errors.AppError) {
	customerSql := "select * from customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while scanning customer table" + err.Error())
			return nil, errors.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB() *CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:root@tcp(db:3306)/banking")
	if err != nil {
		panic(err)
	}

	// Important settings
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return &CustomerRepositoryDb{client}
}
