package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

type Env struct {
	PgInterface interface {
		CreateClient(client Client) (Client, error)
		CreateTransaction(trans Transaction) (Transaction, error)
		GetBalance(client_id int) (int, error)
	}
}

const (
	defaultDSN = "postgres://postgres:test1234@localhost:5432/postgres?sslmode=disable"
)

func main() {
	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		dsn = defaultDSN
	}
	log.Println("Connecting to", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}
	log.Println("Pinged ", dsn)

	env := &Env{
		PgInterface: PgModel{DB: db},
	}

	router := mux.NewRouter()
	router.HandleFunc("/", env.HomeHandler)
	router.HandleFunc("/client/new/{deposit}", env.NewClientHandler).Methods("POST")
	router.HandleFunc("/transaction", env.NewTransactionHandler).Methods("POST")
	router.HandleFunc("/client/{id}/balance", env.BalanceHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func (env *Env) HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`{"status":"Ok"}`))
}

func (env *Env) NewTransactionHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t Transaction
	err := decoder.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	new_transaction, err := env.NewTransaction(t.From_client_id, t.To_client_id, t.Amount)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(writer).Encode(new_transaction)
}

func (env *Env) NewClientHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["deposit"]
	i, _ := strconv.Atoi(s)
	client, err := env.NewClient(i)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(writer).Encode(client)
}

func (env *Env) BalanceHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["id"]
	i, _ := strconv.Atoi(s)
	balance, err := env.CheckBalance(i)
	if err != nil {
		log.Fatal(err)
	}
	response := Balance{balance}

	json.NewEncoder(writer).Encode(response)
}
