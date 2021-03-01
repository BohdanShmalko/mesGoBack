package apiserver

import (
	"github.com/BohdanShmalko/mesGoBack/internal/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersFunc(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", nil)

	srv := newServer(teststore.New())
	srv.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
