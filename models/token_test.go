package models

import (
	"testing"
)

func TestNewToken(t *testing.T) {
	token := NewToken()
	if token == nil {
		t.Fail()
	}
}
