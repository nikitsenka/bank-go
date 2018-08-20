package bank

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "test1234"
	DB_NAME     = "postgres"
)

func Init() {
	db, _ := newDb()
	var e error
	_, e = db.Query("DROP TABLE IF EXISTS client")
	_, e = db.Query("DROP TABLE IF EXISTS account")
	_, e = db.Query("DROP TABLE IF EXISTS transaction")
	checkErr(e)
	_, e = db.Query("CREATE TABLE client(id SERIAL PRIMARY KEY NOT NULL, name VARCHAR(20), email VARCHAR(20), phone VARCHAR(20));")
	_, e = db.Query("CREATE TABLE transaction(id SERIAL PRIMARY KEY NOT NULL, from_client_id INTEGER, to_client_id INTEGER, amount INTEGER);")
	checkErr(e)
	db.Close()
}

func SaveClient(client Client) (Client) {
	db, err := newDb()
	checkErr(err)
	var id int;
	if (client.Id == 0) {
		err = db.QueryRow(
			"INSERT INTO client(name, email, phone) VALUES ($1, $2, $3) RETURNING id",
			client.Name, client.Email, client.Phone).Scan(&id)
		fmt.Println("Created client with id", id)
	} else {
		err = db.QueryRow(
			"UPDATE client SET name = $2, email = $3, phone = $4 WHERE id = $1 RETURNING id",
			client.Id, client.Name, client.Email, client.Phone).Scan(&id)
		fmt.Println("Updated client with id", id)
	}
	checkErr(err)
	client.Id = id
	return client
}

func CreateTransaction(trans Transaction) (Transaction) {
	db, err := newDb()
	checkErr(err)
	var id int;
	if (trans.id == 0) {
		err = db.QueryRow(
			"INSERT INTO transaction(from_client_id, to_client_id, amount) VALUES ($1, $2, $3) RETURNING id",
			trans.from_client_id, trans.to_client_id, trans.amount).Scan(&id)
		fmt.Println("Created transaction with id", id)
	}
	checkErr(err)
	trans.id = id
	return trans
}

func GetBalance(client_id int) int {
	db, err := newDb()
	var balance int
	if (client_id != 0) {
		err = db.QueryRow(`
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
	}
	checkErr(err)
	fmt.Println("Calculated balance with client id", client_id)
	return balance;
}

func newDb() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	return db, err
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err);
		panic(err)
	}
}
