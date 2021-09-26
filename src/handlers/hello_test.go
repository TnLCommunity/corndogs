package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, Hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"msg":"Hello World!"}`+"\n", rec.Body.String())
	}
}
