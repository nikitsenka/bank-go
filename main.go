package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vincentserpoul/bank-go/bank"
	"github.com/vincentserpoul/bank-go/bank/utils"
)

// our main function
func main() {

	conf := Config{}
	conf.Host = utils.GetEnv("POSTGRES_HOST", "localhost")
	conf.User = utils.GetEnv("POSTGRES_USER", "postgres")
	conf.Password = utils.GetEnv("POSTGRES_PASSWORD", "test1234")
	conf.DbName = utils.GetEnv("POSTGRES_NAME", "postgres")

	pool, err := NewConnPool(conf)
	if err != nil {
		log.Fatalf("%v", err)
	}

	migrate(pool)

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/client/new/{deposit}", NewClientHandler(pool)).Methods("POST")
	router.HandleFunc("/transaction", NewTransactionHandler(pool)).Methods("POST")
	router.HandleFunc("/client/{id}/balance", BalanceHandler(pool)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Config is a conf for the postgres database
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

// NewConnPool connects to db and return a connection pool
func NewConnPool(c Config) (*sql.DB, error) {

	pool, err := sql.Open("postgres", getDSN(c))
	if err != nil {
		return nil, fmt.Errorf("NewConnPool: sqlx.Open %v", err)
	}

	errP := pool.Ping()
	if errP != nil {
		return nil, fmt.Errorf("NewConnPool: pool.Ping %v", errP)
	}

	return pool, nil
}

func getDSN(c Config) string {
	dsn := "postgres://" +
		c.User + ":" +
		c.Password + "@" +
		c.Host + ":" +
		c.Port + "/" +
		c.DbName

	// Secure connection?
	ssl := "?sslmode=disable"

	return dsn + ssl
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`{"status":"Ok"}`))
}

func NewTransactionHandler(p *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(writer htt.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		var t bank.Transaction
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		new_transaction := bank.NewTransaction(p, t.From_client_id, t.To_client_id, t.Amount)
		json.NewEncoder(writer).Encode(new_transaction)
	}
}

func NewClientHandler(p *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(writer htt.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		s := params["deposit"]
		i, _ := strconv.Atoi(s)
		client := bank.NewClient(p, i)
		json.NewEncoder(writer).Encode(client)
	}
}

func BalanceHandler(p *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(writer htt.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		s := params["id"]
		i, _ := strconv.Atoi(s)
		balance := bank.CheckBalance(p, i)
		response := bank.Balance{balance}

		json.NewEncoder(writer).Encode(response)
	}
}

func migrate(p *sql.DB) {
	var e error
	_, e = p.Query("DROP TABLE IF EXISTS client")
	_, e = p.Query("DROP TABLE IF EXISTS account")
	_, e = p.Query("DROP TABLE IF EXISTS transaction")
	checkErr(e)
	_, e = p.Query("CREATE TABLE client(id SERIAL PRIMARY KEY NOT NULL, name VARCHAR(20), email VARCHAR(20), phone VARCHAR(20));")
	_, e = p.Query("CREATE TABLE transaction(id SERIAL PRIMARY KEY NOT NULL, from_client_id INTEGER, to_client_id INTEGER, amount INTEGER);")
	checkErr(e)
}
