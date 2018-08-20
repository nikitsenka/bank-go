package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"home-work-1/bank"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/client/{deposit}", ClientHandler).Methods("POST")
	router.HandleFunc("/client/{id}/balance", ClientHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func ClientHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	s := params["id"]
	i, _ := strconv.Atoi(s)
	balance := bank.CheckBalance(i)
	response := BalanceJson{balance}
	json.NewEncoder(writer).Encode(response)
}

type BalanceJson struct {
	Balance int `json:"balance"`
}
