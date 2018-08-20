package bank

import (
	"reflect"
	"testing"
)

func TestBank(t *testing.T) {

	client1 := NewClient(1000)
	expectedCleint1 := Client{1, "", "", ""}
	if !reflect.DeepEqual(client1, expectedCleint1) {
		t.Errorf("TestNewClient() = %v, want %v", client1, expectedCleint1)
	}

	client2 := NewClient(2000)
	expectedCleint2 := Client{2, "", "", ""}
	if !reflect.DeepEqual(client2, expectedCleint2) {
		t.Errorf("TestNewClient() = %v, want %v", client2, expectedCleint2)
	}

	NewTransaction(client1.id, client2.id, 1000)

	balance1 := CheckBalance(client1.id)
	if balance1 != 0{
		t.Errorf("Incorrect balance = %v, want %v", balance1, 0)
	}

	balance2 := CheckBalance(client2.id)
	if balance2 != 3000{
		t.Errorf("Incorrect balance = %v, want %v", balance2, 3000)
	}

	t.Log("Bank test finished")
}


