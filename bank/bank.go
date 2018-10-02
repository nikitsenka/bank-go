package bank

import (
	"database/sql"
	_ "log"
	_ "net/http"

	_ "github.com/gorilla/mux"
)

func NewClient(p *sql.DB, balance int) Client {
	client := Client{0, "", "", ""}
	client = CreateClient(p, client)
	transaction := Transaction{0, 0, client.Id, balance}
	CreateTransaction(p, transaction)
	return client
}

func NewTransaction(p *sql.DB, from_client_id int, to_client_id int, amount int) Transaction {
	transaction := Transaction{0, from_client_id, to_client_id, amount}
	transaction = CreateTransaction(p, transaction)
	return transaction
}

func CheckBalance(p *sql.DB, client_id int) int {
	return GetBalance(p, client_id)
}
