package main

import (
	"database/sql"
	_ "log"
	_ "net/http"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/gorilla/mux"
)

func NewClient(p *sql.DB, balance int) (Client, error) {
	client := Client{0, "", "", ""}
	client, err := CreateClient(p, client)
	if err != nil {
		return client, err
	}
	transaction := Transaction{0, 0, client.Id, balance}
	CreateTransaction(p, transaction)
	return client, err
}

func NewTransaction(p *sql.DB, from_client_id int, to_client_id int, amount int) (Transaction, error) {
	transaction := Transaction{0, from_client_id, to_client_id, amount}
	transaction, err := CreateTransaction(p, transaction)
	return transaction, err
}

func CheckBalance(p *sql.DB, client_id int) (int, error) {
	return GetBalance(p, client_id)
}
