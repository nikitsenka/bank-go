package bank

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateClient(p *sql.DB, client Client) Client {
	var id int
	err := p.QueryRow(
		"INSERT INTO client(name, email, phone) VALUES ($1, $2, $3) RETURNING id",
		client.Name, client.Email, client.Phone).Scan(&id)
	fmt.Println("Created client with id", id)
	checkErr(err)
	client.Id = id
	return client
}

func CreateTransaction(p *sql.DB, trans Transaction) Transaction {
	var id int
	err := p.QueryRow(
		"INSERT INTO transaction(from_client_id, to_client_id, amount) VALUES ($1, $2, $3) RETURNING id",
		trans.From_client_id, trans.To_client_id, trans.Amount).Scan(&id)
	fmt.Println("Created transaction with id", id)
	checkErr(err)
	trans.Id = id
	return trans
}

func GetBalance(p *sql.DB, client_id int) int {
	var balance int
	err := p.QueryRow(`
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
	checkErr(err)
	fmt.Println("Calculated balance with client id", client_id)
	return balance
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
