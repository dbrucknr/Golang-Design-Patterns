package main

import (
	"os"
	"testing"

	"github.com/dbrucknr/go-design-patterns/config"
)

var testApp Application

func TestMain(m *testing.M) {

	testApp = Application{
		App: config.New(nil),
	}

	os.Exit(m.Run())
}
