package tests

import (
	"github.com/premsvmm/goapi/app"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWelcomeRouter(t *testing.T) {
	app.StartTestApplication()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	app.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "welcome Prem", w.Body.String())

}
