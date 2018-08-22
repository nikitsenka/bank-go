package bank

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	"reflect"
)

func TestMain(m *testing.M) {
	Init()
	retCode := m.Run()
	os.Exit(retCode)
}

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

	NewTransaction(client1.Id, client2.Id, 1000)

	balance1 := CheckBalance(client1.Id)
	if balance1 != 0{
		t.Errorf("Incorrect balance = %v, want %v", balance1, 0)
	}

	balance2 := CheckBalance(client2.Id)
	if balance2 != 3000{
		t.Errorf("Incorrect balance = %v, want %v", balance2, 3000)
	}

	t.Log("Bank test finished")
}

func TestSaveClient(t *testing.T) {
	client := Client{0, "name1", "email1", "phone1"}
	actual := CreateClient(client)
	expected := Client{3, "name1", "email1", "phone1"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestSaveClient() = %v, want %v", actual, expected)
	}
}

func TestCreateTransaction(t *testing.T) {
	trans := Transaction{0,0,0,10}
	actual := CreateTransaction(trans)
	expected := Transaction{4, 0, 0, 10}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestCreateTransaction() = %v, want %v", actual, expected)
	}
}
