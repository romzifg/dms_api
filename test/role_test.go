package test

import (
	"dms_api/database"
	"dms_api/middlewares"
	"dms_api/modules/roles"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetRolesUnauthorized(t *testing.T) {
	database.DbConnect()

	app := fiber.New()
	app.Use(middlewares.ApiTokenMiddleware)
   	app.Get("/api/v1/role", roles.GetAll)

	req := httptest.NewRequest("GET", "/api/v1/role", nil)
	req.Header.Add("api_token", "123")
	resp, _ := app.Test(req)

	assert.Equal(t, 401, resp.StatusCode)
	assert.NotEqual(t, 200, resp.StatusCode)
}

func TestRolesAuthorized(t *testing.T) {
	database.DbConnect()

	app := fiber.New()
	app.Use(middlewares.ApiTokenMiddleware)
   	app.Get("/api/v1/role", roles.GetAll)

	req := httptest.NewRequest("GET", "/api/v1/role", nil)
	req.Header.Add("api_token", "P@ssw0rd123!!@@")
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEqual(t, 401, resp.StatusCode)
}