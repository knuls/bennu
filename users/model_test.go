package users

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserRender(t *testing.T) {
	user := NewUser()
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := user.Render(rr, req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEmptyUserFromJSON(t *testing.T) {
	user := NewUser()
	var b []byte
	err := user.FromJSON(bytes.NewReader(b))
	if err != io.EOF {
		t.Error(err)
	}
}

func TestUserFromJSON(t *testing.T) {
	user := NewUser()
	b := []byte(`{
		"email": "m@m.m",
		"firstName": "m",
		"lastName": "m",
		"password": "m"
	}`)
	err := user.FromJSON(bytes.NewReader(b))
	if err != nil {
		t.Error(err)
	}
	if user.FirstName != "m" {
		t.Fatalf("first name not m, got %s", user.FirstName)
	}
}

func TestUserHashAndComparePassword(t *testing.T) {
	u := NewUser()
	u.Password = "super-secret"
	err := u.HashPassword()
	if err != nil {
		t.Error(err)
	}
	err = u.ComparePassword("super-secret")
	if err != nil {
		t.Error(err)
	}
}
