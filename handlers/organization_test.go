package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/dao/mocks"
	"github.com/knuls/horus/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestOrganizationHandler(t *testing.T) {
	t.Parallel()

	// mocks
	logger, err := logger.New()
	if err != nil {
		t.Error(err)
	}
	defer logger.GetLogger().Sync()
	factory := &mocks.Factory{}
	errFactory := &mocks.ErrFactory{}
	id := primitive.NewObjectIDFromTimestamp(time.Now())

	// tests
	cases := []*struct {
		name               string
		factory            dao.Factory
		method             string
		path               string
		body               map[string]interface{}
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "getOrganization",
			factory:            factory,
			method:             http.MethodGet,
			path:               "/",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "getOrganizationErr",
			factory:            errFactory,
			method:             http.MethodGet,
			path:               "/",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "getOrganizationById",
			factory:            factory,
			method:             http.MethodGet,
			path:               fmt.Sprintf("/%s", id.Hex()),
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "getOrganizationByIdErr",
			factory:            errFactory,
			method:             http.MethodGet,
			path:               fmt.Sprintf("/%s", id.Hex()),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "postOrganization",
			factory:            factory,
			method:             http.MethodPost,
			path:               "/",
			body:               nil,
			expectedStatusCode: http.StatusCreated,
			expectedBody:       "",
		},
		{
			name:               "postOrganizationErr",
			factory:            errFactory,
			method:             http.MethodPost,
			path:               "/",
			body:               nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "",
		},
	}

	// execute
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			// target
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
		})
	}
}
