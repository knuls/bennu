package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/models"
	"github.com/knuls/horus/logger"
)

type MockTokenDao struct {
}

func (m *MockTokenDao) Find(ctx context.Context, filter dao.Where) ([]*models.Token, error) {
	tokens := []*models.Token{
		models.NewToken(),
		models.NewToken(),
		models.NewToken(),
	}
	return tokens, nil
}
func (m *MockTokenDao) FindOne(ctx context.Context, filter dao.Where) (*models.Token, error) {
	return models.NewToken(), nil
}
func (m *MockTokenDao) Create(ctx context.Context, token *models.Token) (string, error) {
	return "", nil
}
func (m *MockTokenDao) Update(ctx context.Context, token *models.Token) (*models.Token, error) {
	return nil, nil
}

type MockErrTockenDao struct {
}

func (m *MockErrTockenDao) Find(ctx context.Context, filter dao.Where) ([]*models.Token, error) {
	return nil, errors.New("some error")
}
func (m *MockErrTockenDao) FindOne(ctx context.Context, filter dao.Where) (*models.Token, error) {
	return nil, errors.New("some error")
}
func (m *MockErrTockenDao) Create(ctx context.Context, org *models.Token) (string, error) {
	return "", errors.New("some error")
}
func (m *MockErrTockenDao) Update(ctx context.Context, org *models.Token) (*models.Token, error) {
	return nil, nil
}

func TestAuthHandler(t *testing.T) {
	logger, err := logger.New()
	if err != nil {
		t.Error(err)
	}
	// factory := &MockFactory{}
	// errFactory := &MockErrFactory{}

	// tests
	cases := []*testCase{
		// {
		// 	factory:            factory,
		// 	method:             http.MethodGet,
		// 	path:               "/csrf",
		// 	expectedStatusCode: http.StatusOK,
		// },
		// {
		// 	factory:            factory,
		// 	method:             http.MethodPost,
		// 	path:               "/login",
		// 	expectedStatusCode: http.StatusOK,
		// },
		// {
		// 	factory:            factory,
		// 	method:             http.MethodPost,
		// 	path:               "/register",
		// 	expectedStatusCode: http.StatusOK,
		// },
		// {
		// 	factory:            factory,
		// 	method:             http.MethodPost,
		// 	path:               "/reset-password",
		// 	expectedStatusCode: http.StatusOK,
		// },
		// {
		// 	factory:            factory,
		// 	method:             http.MethodPost,
		// 	path:               "/logout",
		// 	expectedStatusCode: http.StatusOK,
		// },
		// {
		// 	factory:            factory,
		// 	method:             http.MethodPost,
		// 	path:               "/verify/email",
		// 	expectedStatusCode: http.StatusOK,
		// },
		// {
		// 	factory:            factory,
		// 	method:             http.MethodPost,
		// 	path:               "/verify/reset-password",
		// 	expectedStatusCode: http.StatusOK,
		// },
		// {
		// 	factory:            factory,
		// 	method:             http.MethodPost,
		// 	path:               "/token/refresh",
		// 	expectedStatusCode: http.StatusOK,
		// },
	}

	for _, testCase := range cases {
		// execute
		handler := NewOrganizationHandler(logger, testCase.factory)
		body, err := json.Marshal(testCase.body)
		if err != nil {
			t.Error(err)
		}
		req := httptest.NewRequest(testCase.method, testCase.path, bytes.NewReader(body))
		rr := httptest.NewRecorder()

		// serve
		handler.Routes().ServeHTTP(rr, req)

		// assert
		res := rr.Result()
		if res.StatusCode != testCase.expectedStatusCode {
			t.Fatalf("result expected to be %d, got %d", testCase.expectedStatusCode, res.StatusCode)
		}
		var p map[string]interface{}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&p); err != nil {
			t.Error(err)
		}
	}
}
