package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"github.com/nikitsenka/bank-go/bank"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/client/new/{deposit}", NewClientHandler).Methods("POST")
	router.HandleFunc("/transaction", NewTransactionHandler).Methods("POST")
	router.HandleFunc("/client/{id}/balance", BalanceHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`{"status":"Ok"}`))
}

func NewTransactionHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t bank.Transaction
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	new_transaction := bank.NewTransaction(t.From_client_id, t.To_client_id, t.Amount)
	json.NewEncoder(writer).Encode(new_transaction)
}

func NewClientHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["deposit"]
	i, _ := strconv.Atoi(s)
	client := bank.NewClient(i)
	json.NewEncoder(writer).Encode(client)
}

func BalanceHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["id"]
	i, _ := strconv.Atoi(s)
	balance := bank.CheckBalance(i)
	response := bank.Balance{balance}
	json.NewEncoder(writer).Encode(response)
}