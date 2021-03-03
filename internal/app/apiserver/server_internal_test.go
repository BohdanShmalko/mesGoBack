package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/BohdanShmalko/mesGoBack/internal/app/models"
	"github.com/BohdanShmalko/mesGoBack/internal/store/teststore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersFunc(t *testing.T) {
	srv := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name:         "bad request",
			payload:      "some bad inf",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "not valid data",
			payload: map[string]string{
				"badkey": "badvalue",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "valid data",
			payload: map[string]string{
				"lastname": "testlastname",
				"email":    "test1@gmail.com",
				"password": "testpassword",
				"nickname": "testnick",
				"name":     "testname",
			},
			expectedCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/users", b)
			srv.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode , rec.Code)
		})
	}
}

func TestServer_HandleLoginFunc(t *testing.T) {
	u := models.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	srv := newServer(store, sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid data",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.RowPassword,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "bad request",
			payload:      "some bad inf",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "not valid data",
			payload: map[string]string{
				"email":    "bademail",
				"password": "badpassword",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/login", b)
			srv.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := models.TestUser(t)
	store.User().Create(u)

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id" : u.Id,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "not authenticated",
			cookieValue: nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	secureKey := []byte("secret")
	srv := newServer(store, sessions.NewCookieStore(secureKey))
	sc := securecookie.New(secureKey, nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _,tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)

			rowCookie := sessionName + "=" + cookieStr
			req.Header.Set("Cookie", rowCookie)
			srv.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
