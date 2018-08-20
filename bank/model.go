package bank

type Client struct {
	id    int
	name  string
	email string
	phone string
}

type Transaction struct {
	id             int
	from_client_id int
	to_client_id   int
	amount         int
}

