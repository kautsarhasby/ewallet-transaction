package cmd

import (
	"kautsarhasby/ewallet-transaction/external"
	"kautsarhasby/ewallet-transaction/helpers"
	"kautsarhasby/ewallet-transaction/internal/api"
	"kautsarhasby/ewallet-transaction/internal/interfaces"
	"kautsarhasby/ewallet-transaction/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()
	r := gin.Default()


	r.GET("/health", d.HealthCheckAPI.HealthCheckHandlerHTTP)

	// walletV1 := r.Group("/wallet/v1")

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}


type Dependency struct {
	HealthCheckAPI interfaces.IHealthCheckAPI
	External interfaces.IExternal
}

func dependencyInject() Dependency {
	healthCheckSvc := &services.HealthCheck{}
	healthCheckAPI := &api.HealthCheck{
		HealthCheckServices: healthCheckSvc,
	}

	external := &external.External{}

	return Dependency{
		HealthCheckAPI: healthCheckAPI,
		External: external,
	}
}