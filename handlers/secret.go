package handlers

import (
	"net/http"

	"github.com/caiyeon/goldfish/vault"
	"github.com/gorilla/csrf"
	"github.com/labstack/echo"
)

func GetSecrets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var auth = &vault.AuthInfo{}
		defer auth.Clear()

		// fetch auth from cookie
		getSession(c, auth)

		path := c.QueryParam("path")
		if path == "" {
			conf := vault.GetConfig()
			path = conf.DefaultSecretPath
		}

		c.Response().Writer.Header().Set("X-CSRF-Token", csrf.Token(c.Request()))

		if path == "" || path[len(path)-1:] == "/" {
			// listing a directory
			if result, err := auth.ListSecret(path); err != nil {
				return logError(c, err.Error(), "Internal error")
			} else {
				return c.JSON(http.StatusOK, H{
					"result": result,
					"path":   path,
				})
			}
		} else {
			// reading a specific secret's key value pairs
			if result, err := auth.ReadSecret(path); err != nil {
				return logError(c, err.Error(), "Internal error")
			} else {
				return c.JSON(http.StatusOK, H{
					"result": result,
					"path":   path,
				})
			}
		}
	}
}

func PostSecrets() echo.HandlerFunc {
	return func(c echo.Context) error {
		return logError(c, "", "Secret writing not allowed in demo mode")
	}
}
