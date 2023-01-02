package models

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrganizationRender(t *testing.T) {
	org := NewOrganization()
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := org.Render(rr, req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEmptyOrganizationFromJSON(t *testing.T) {
	org := NewOrganization()
	var b []byte
	err := org.FromJSON(bytes.NewReader(b))
	if err != io.EOF {
		t.Error(err)
	}
}

func TestOrganizationFromJSON(t *testing.T) {
	org := NewOrganization()
	b := []byte(`{
		"name": "some-name"
	}`)
	err := org.FromJSON(bytes.NewReader(b))
	if err != nil {
		t.Error(err)
	}
	if org.Name != "some-name" {
		t.Fatalf("name not some-name, got %s", org.Name)
	}
}
