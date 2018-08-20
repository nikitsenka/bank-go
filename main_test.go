package main

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	Init()
	retCode := m.Run()
	os.Exit(retCode)
}
