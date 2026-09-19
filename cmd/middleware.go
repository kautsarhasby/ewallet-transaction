package cmd

import (
	"kautsarhasby/ewallet-transaction/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func (d *Dependency) MiddlewareValidateToken(c *gin.Context) {
	var log =helpers.Logger
	authHeader := c.Request.Header.Get("Authorization")
    if authHeader == "" {
        helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
        c.Abort()
        return
    }

    token := strings.TrimPrefix(authHeader, "Bearer ")
    token = strings.TrimSpace(token)
	if token == "" {
		helpers.SendResponseHTTP(c,http.StatusUnauthorized, "unauthorized", nil)
		c.Abort()
		return
	}

	tokenData, err := d.External.ValidateToken(c.Request.Context(), token)
	if err != nil {
		log.Error(err)
		helpers.SendResponseHTTP(c,http.StatusUnauthorized, "unauthorized", nil)
		c.Abort()
		return
	}

	c.Set("token", tokenData)
	c.Next()
}