package main

import (
	_ "log"
	_ "net/http"

	_ "github.com/gorilla/mux"
)

func (env *Env) NewClient(balance int) (Client, error) {
	client := Client{0, "", "", ""}
	client, err := env.PgInterface.CreateClient(client)
	if err != nil {
		return client, err
	}
	transaction := Transaction{0, 0, client.Id, balance}
	env.PgInterface.CreateTransaction(transaction)
	return client, err
}

func (env *Env) NewTransaction(from_client_id int, to_client_id int, amount int) (Transaction, error) {
	transaction := Transaction{0, from_client_id, to_client_id, amount}
	transaction, err := env.PgInterface.CreateTransaction(transaction)
	return transaction, err
}

func (env *Env) CheckBalance(client_id int) (int, error) {
	return env.PgInterface.GetBalance(client_id)
}
