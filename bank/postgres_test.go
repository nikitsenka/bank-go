package bank

import (
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestSaveClient(t *testing.T) {
	client := Client{0, "name1", "email1", "phone1"}
	actual := SaveClient(client)
	expected := Client{1, "name1", "email1", "phone1"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestSaveClient() = %v, want %v", actual, expected)
	}
}

func TestUpdateClient(t *testing.T) {

	client := Client{0, "name1", "email1", "phone1"}
	client = SaveClient(client)
	client.name = "name2"
	actual := SaveClient(client)
	if !reflect.DeepEqual(actual, client) {
		t.Errorf("TestUpdateClient() = %v, want %v", actual, client)
	}
}

func TestCreateTransaction(t *testing.T) {
	trans := Transaction{0,0,0,10}
	actual := CreateTransaction(trans)
	expected := Transaction{1, 0, 0, 10}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestCreateTransaction() = %v, want %v", actual, expected)
	}
}
