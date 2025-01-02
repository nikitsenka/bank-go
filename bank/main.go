package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/alexbrainman/odbc"
	"github.com/gorilla/mux"
)

var db *sql.DB

const (
	defaultDSN = "postgres://postgres:test1234@localhost:5432/postgres?sslmode=disable"
)

func main() {
	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		dsn = defaultDSN
	}
	log.Println("Connecting to", dsn)
	db, err := sql.Open("odbc", "DSN=ODBC;Driver=PostgreSQL")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}
	log.Println("Pinged ", dsn)

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/client/new/{deposit}", NewClientHandler).Methods("POST")
	router.HandleFunc("/transaction", NewTransactionHandler).Methods("POST")
	router.HandleFunc("/client/{id}/balance", BalanceHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`{"status":"Ok"}`))
}

func NewTransactionHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t Transaction
	err := decoder.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	new_transaction, err := NewTransaction(db, t.From_client_id, t.To_client_id, t.Amount)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(writer).Encode(new_transaction)
}

func NewClientHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["deposit"]
	i, _ := strconv.Atoi(s)
	client, err := NewClient(db, i)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(writer).Encode(client)
}

func BalanceHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["id"]
	i, _ := strconv.Atoi(s)
	balance, err := CheckBalance(db, i)
	if err != nil {
		log.Fatal(err)
	}
	response := Balance{balance}

	json.NewEncoder(writer).Encode(response)
}
