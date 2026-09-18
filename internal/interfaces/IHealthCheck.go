package interfaces

import "github.com/gin-gonic/gin"

type IHealthCheckServices interface {
	HealthCheckServices() (string, error)
}

type IHealthCheckRepo interface {
	HealthCheckRepository() (string, error)
}

type IHealthCheckAPI interface {
	HealthCheckHandlerHTTP(c *gin.Context)
}