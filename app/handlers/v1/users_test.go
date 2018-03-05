package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, UserCreate(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"create ok\"", rec.Body.String())
	}
}

func TestUserList(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, UserList(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"list ok\"", rec.Body.String())
	}
}

func TestUserUpdate(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, UserUpdate(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"update ok\"", rec.Body.String())
	}
}

func TestUserDelete(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, UserDelete(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"delete ok\"", rec.Body.String())
	}
}
