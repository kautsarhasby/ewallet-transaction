package interfaces

import (
	"context"
	"kautsarhasby/ewallet-transaction/internal/models"

	"github.com/gin-gonic/gin"
)

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, trx *models.Transaction) error
}

type ITransactionService interface {
	CreateTransaction(ctx context.Context, req *models.Transaction) (models.TransactionResponse, error) 
}

type ITransactionAPI interface {
	 CreateTransaction(c *gin.Context)
}