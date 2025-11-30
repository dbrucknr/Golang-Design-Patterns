package main

import (
	"os"
	"testing"

	"github.com/dbrucknr/go-design-patterns/models"
)

var testApp Application

func TestMain(m *testing.M) {

	testApp = Application{
		Models: *models.New(nil),
	}

	os.Exit(m.Run())
}
