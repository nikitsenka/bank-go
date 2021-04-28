package main

type Client struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Transaction struct {
	Id             int `json:"id"`
	From_client_id int `json:"from_client_id"`
	To_client_id   int `json:"to_client_id"`
	Amount         int `json:"amount"`
}

type Balance struct {
	Balance int `json:"balance"`
}
