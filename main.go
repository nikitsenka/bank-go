package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

const (
	defaultDSN = "postgres://postgres:test1234@localhost:5432/postgres?sslmode=disable"
)

func main() {
	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		dsn = defaultDSN
	}
	log.Println("Connecting to", dsn)
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	db, err = pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(ctx); err != nil {
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
	new_transaction, err := NewTransaction(t.From_client_id, t.To_client_id, t.Amount)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(writer).Encode(new_transaction)
}

func NewClientHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["deposit"]
	i, _ := strconv.Atoi(s)
	client, err := NewClient(i)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(writer).Encode(client)
}

func BalanceHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["id"]
	i, _ := strconv.Atoi(s)
	balance, err := CheckBalance(i)
	if err != nil {
		log.Fatal(err)
	}
	response := Balance{balance}

	json.NewEncoder(writer).Encode(response)
}
