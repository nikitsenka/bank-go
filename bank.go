package main

import (
	_ "log"
	_ "net/http"

	_ "github.com/gorilla/mux"
)

func NewClient(balance int) (Client, error) {
	client := Client{0, "", "", ""}
	client, err := CreateClient(client)
	if err != nil {
		return client, err
	}
	transaction := Transaction{0, 0, client.Id, balance}
	CreateTransaction(transaction)
	return client, err
}

func NewTransaction(from_client_id int, to_client_id int, amount int) (Transaction, error) {
	transaction := Transaction{0, from_client_id, to_client_id, amount}
	transaction, err := CreateTransaction(transaction)
	return transaction, err
}

func CheckBalance(client_id int) (int, error) {
	return GetBalance(client_id)
}
