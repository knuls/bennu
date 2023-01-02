package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/models"
	"github.com/knuls/horus/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockOrganizationDao struct {
}

func (m *MockOrganizationDao) Find(ctx context.Context, filter dao.Where) ([]*models.Organization, error) {
	orgs := []*models.Organization{
		models.NewOrganization(),
		models.NewOrganization(),
		models.NewOrganization(),
	}
	return orgs, nil
}
func (m *MockOrganizationDao) FindOne(ctx context.Context, filter dao.Where) (*models.Organization, error) {
	return models.NewOrganization(), nil
}
func (m *MockOrganizationDao) Create(ctx context.Context, org *models.Organization) (string, error) {
	return "", nil
}
func (m *MockOrganizationDao) Update(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	return nil, nil
}

type MockErrOrganizationDao struct {
}

func (m *MockErrOrganizationDao) Find(ctx context.Context, filter dao.Where) ([]*models.Organization, error) {
	return nil, errors.New("some error")
}
func (m *MockErrOrganizationDao) FindOne(ctx context.Context, filter dao.Where) (*models.Organization, error) {
	return nil, errors.New("some error")
}
func (m *MockErrOrganizationDao) Create(ctx context.Context, org *models.Organization) (string, error) {
	return "", errors.New("some error")
}
func (m *MockErrOrganizationDao) Update(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	return nil, nil
}

func TestOrganizationHandler(t *testing.T) {
	logger, err := logger.New()
	if err != nil {
		t.Error(err)
	}
	factory := &MockFactory{}
	errFactory := &MockErrFactory{}
	id := primitive.NewObjectIDFromTimestamp(time.Now())
	cases := []*testCase{
		{
			factory:            factory,
			method:             http.MethodGet,
			path:               "/",
			expectedStatusCode: http.StatusOK,
		},
		{
			factory:            errFactory,
			method:             http.MethodGet,
			path:               "/",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			factory:            factory,
			method:             http.MethodGet,
			path:               fmt.Sprintf("/%s", id.Hex()),
			expectedStatusCode: http.StatusOK,
		},
		{
			factory:            errFactory,
			method:             http.MethodGet,
			path:               fmt.Sprintf("/%s", id.Hex()),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			factory:            factory,
			method:             http.MethodPost,
			path:               "/",
			body:               nil,
			expectedStatusCode: http.StatusCreated,
			expectedBody:       "",
		},
		{
			factory:            errFactory,
			method:             http.MethodPost,
			path:               "/",
			body:               nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "",
		},
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
