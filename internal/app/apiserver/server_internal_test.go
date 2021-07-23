package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pyankovzhe/auth/internal/app/model"
	"github.com/pyankovzhe/auth/internal/app/producer/testproducer"
	"github.com/pyankovzhe/auth/internal/app/store/teststore"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestServer_Signup(t *testing.T) {
	s := newServer(teststore.New(), logrus.New(), "fakeaddr", &testproducer.Producer{})

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":    "login1",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid login",
			payload: map[string]string{
				"login":    "",
				"password": "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"login":    "login2",
				"password": "12",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/signup", b)
			s.Handler.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_Signin(t *testing.T) {
	a := model.TestAccount(t)
	store := teststore.New()
	store.Account().Create(a)

	s := newServer(store, logrus.New(), "fakeaddr", &testproducer.Producer{})
	testcases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":    a.Login,
				"password": a.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "incorrect login",
			payload: map[string]string{
				"login":    "incorrect",
				"password": a.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "incorrect password",
			payload: map[string]string{
				"login":    a.Login,
				"password": "incorrect",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/signin", b)
			s.Handler.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)

			// if tc.expectedCode == 200 {
			// 	resp := &signingResponse{}
			// 	json.NewDecoder(rec.Body).Decode(resp)
			// 	assert.NotNil(t, resp.Token)
			// }
		})
	}
}
