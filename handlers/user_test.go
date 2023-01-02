package handlers

import (
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

type MockUserDao struct {
}

func (m *MockUserDao) Find(ctx context.Context, filter dao.Where) ([]*models.User, error) {
	users := []*models.User{
		models.NewUser(),
		models.NewUser(),
		models.NewUser(),
	}
	return users, nil
}
func (m *MockUserDao) FindOne(ctx context.Context, filter dao.Where) (*models.User, error) {
	return models.NewUser(), nil
}
func (m *MockUserDao) Create(ctx context.Context, user *models.User) (string, error) {
	return "", nil
}
func (m *MockUserDao) Update(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, nil
}

type MockFactory struct {
}

func (f *MockFactory) GetUserDao() dao.Dao[models.User] {
	return &MockUserDao{}
}
func (f *MockFactory) GetOrganizationDao() dao.Dao[models.Organization] {
	return nil
}
func (f *MockFactory) GetTokenDao() dao.Dao[models.Token] {
	return nil
}

type MockErrUserDao struct {
}

func (m *MockErrUserDao) Find(ctx context.Context, filter dao.Where) ([]*models.User, error) {
	return nil, errors.New("some error")
}
func (m *MockErrUserDao) FindOne(ctx context.Context, filter dao.Where) (*models.User, error) {
	return nil, errors.New("some error")
}
func (m *MockErrUserDao) Create(ctx context.Context, user *models.User) (string, error) {
	return "", nil
}
func (m *MockErrUserDao) Update(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, nil
}

type MockErrFactory struct {
}

func (f *MockErrFactory) GetUserDao() dao.Dao[models.User] {
	return &MockErrUserDao{}
}
func (f *MockErrFactory) GetOrganizationDao() dao.Dao[models.Organization] {
	return nil
}
func (f *MockErrFactory) GetTokenDao() dao.Dao[models.Token] {
	return nil
}

type testCase struct {
	factory            dao.Factory
	method             string
	path               string
	expectedStatusCode int
}

func TestUserHandler(t *testing.T) {
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
	}

	for _, testCase := range cases {
		handler := NewUserHandler(logger, testCase.factory)
		req := httptest.NewRequest(testCase.method, testCase.path, nil)
		rr := httptest.NewRecorder()
		handler.Routes().ServeHTTP(rr, req)
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
