package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/dao/mocks"
	"github.com/knuls/horus/logger"
)

func TestAuthHandler(t *testing.T) {
	t.Parallel()

	// mocks
	logger, err := logger.New()
	if err != nil {
		t.Error(err)
	}
	defer logger.GetLogger().Sync()
	factory := &mocks.Factory{}
	// errFactory := &mocks.ErrFactory{}

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
			name:               "getCSRF",
			factory:            factory,
			method:             http.MethodGet,
			path:               "/csrf",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "postLogin",
			factory:            factory,
			method:             http.MethodPost,
			path:               "/login",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "postRegister",
			factory:            factory,
			method:             http.MethodPost,
			path:               "/register",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "postResetPassword",
			factory:            factory,
			method:             http.MethodPost,
			path:               "/reset-password",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "postLogout",
			factory:            factory,
			method:             http.MethodPost,
			path:               "/logout",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "postVerifyEmail",
			factory:            factory,
			method:             http.MethodPost,
			path:               "/verify/email",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "postVerifyResetPassword",
			factory:            factory,
			method:             http.MethodPost,
			path:               "/verify/reset-password",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "postTokenRefresh",
			factory:            factory,
			method:             http.MethodPost,
			path:               "/token/refresh",
			expectedStatusCode: http.StatusOK,
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
