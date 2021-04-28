package main

import (
	"context"
	"log"
)

func CreateClient(client Client) (Client, error) {
	var id int
	err := db.QueryRow(context.Background(),
		"INSERT INTO client(name, email, phone) VALUES ($1, $2, $3) RETURNING id",
		client.Name, client.Email, client.Phone).Scan(&id)
	log.Println("Created client with id", id)
	client.Id = id
	return client, err
}

func CreateTransaction(trans Transaction) (Transaction, error) {
	var id int
	err := db.QueryRow(context.Background(),
		"INSERT INTO transaction(from_client_id, to_client_id, amount) VALUES ($1, $2, $3) RETURNING id",
		trans.From_client_id, trans.To_client_id, trans.Amount).Scan(&id)
	log.Println("Created transaction with id", id)
	trans.Id = id
	return trans, err
}

func GetBalance(client_id int) (int, error) {
	var balance int
	err := db.QueryRow(context.Background(), `
				SELECT debit - credit
				FROM
				  (
					SELECT COALESCE(sum(amount), 0) AS debit
					FROM transaction
					WHERE to_client_id = $1
				  ) a,
				  (
					SELECT COALESCE(sum(amount), 0) AS credit
					FROM transaction
					WHERE from_client_id = $1
				  ) b;
		`, client_id).Scan(&balance)
	log.Println("Calculated balance with client id", client_id)
	return balance, err
}
