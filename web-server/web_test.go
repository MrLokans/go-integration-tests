package main

import (
	"fmt"
	"testing"
)

var testServerPort = 8991
var a = authService{Base: fmt.Sprintf("http://localhost:%d", testServerPort)}

func TestUserCannotLoginWithWrongPassword(t *testing.T) {
	if a.Login("user1", "wrongpass").Token != "" {
		t.Fail()
	}
}

func TestUserCannotLoginWithWrongUsername(t *testing.T) {
	if a.Login("non-existentuser", "pass1").Token != "" {
		t.Fail()
	}
}

func TestUserCanLoginWithCorrectUsernameAndPass(t *testing.T) {
	if a.Login("user1", "pass1").Token == "" {
		t.Fail()
	}
}
