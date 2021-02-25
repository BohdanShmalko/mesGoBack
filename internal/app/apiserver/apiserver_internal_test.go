package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer_HandleTest(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/someroute", nil)
	s.testRoute().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "<h1>test page</h1>")
}
