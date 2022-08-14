package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/sushilparajuli/go-banking/errors"
	"github.com/sushilparajuli/go-banking/logger"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d *CustomerRepositoryDb) FindAll(status string) ([]Customer, *errors.AppError) {
	var err error
	var findAllSql string
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql = "select * from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql = "select * from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while query customer table")
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d *CustomerRepositoryDb) ById(id string) (*Customer, *errors.AppError) {
	customerSql := "select * from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSql, id)
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
	client, err := sqlx.Open("mysql", viper.GetString("MYSQL_URI"))
	if err != nil {
		panic(err)
	}

	// Important settings
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return &CustomerRepositoryDb{client}
}
