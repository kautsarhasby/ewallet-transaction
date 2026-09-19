package cmd

import (
	"kautsarhasby/ewallet-transaction/external"
	"kautsarhasby/ewallet-transaction/helpers"
	"kautsarhasby/ewallet-transaction/internal/api"
	"kautsarhasby/ewallet-transaction/internal/interfaces"
	"kautsarhasby/ewallet-transaction/internal/repository"
	"kautsarhasby/ewallet-transaction/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()
	r := gin.Default()


	r.GET("/health", d.HealthCheckAPI.HealthCheckHandlerHTTP)

	transactionV1 := r.Group("/transaction/v1")
	transactionV1.POST("/create",d.MiddlewareValidateToken, d.Transaction.CreateTransaction)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}


type Dependency struct {
	HealthCheckAPI interfaces.IHealthCheckAPI
	External interfaces.IExternal
	Transaction interfaces.ITransactionAPI
}

func dependencyInject() Dependency {
	healthCheckSvc := &services.HealthCheck{}
	healthCheckAPI := &api.HealthCheck{
		HealthCheckServices: healthCheckSvc,
	}

	external := &external.External{}
	trxRepository := &repository.TransactionRepository{
		DB: helpers.DB,
	}
	trxSvc := &services.TransactionService{TransactionRepository: trxRepository,}
	trxAPI := &api.TransactionHandler{TransactionService: trxSvc}
	return Dependency{
		HealthCheckAPI: healthCheckAPI,
		External: external,
		Transaction: trxAPI,
	}
}