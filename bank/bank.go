package bank

import (
_	"log"
_	"net/http"
_	"github.com/gorilla/mux"
)

func NewClient(balance int) (Client){
	client := Client{0, "", "", ""}
	client = SaveClient(client)
	transaction := Transaction{0, 0, client.id, balance}
	CreateTransaction(transaction)
	return client
}

func NewTransaction(from_client_id int, to_client_id int, amount int) (Transaction){
	transaction := Transaction{0, from_client_id, to_client_id, amount}
	transaction = CreateTransaction(transaction)
	return transaction
}

func CheckBalance(client_id int) (int){
	return GetBalance(client_id)
}

