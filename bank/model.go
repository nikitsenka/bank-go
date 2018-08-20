package bank

type Client struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Transaction struct {
	id             int
	from_client_id int
	to_client_id   int
	amount         int
}

type Balance struct {
	Balance int `json:"balance"`
}
