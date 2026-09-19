package repository

import (
	"context"
	"kautsarhasby/ewallet-transaction/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, trx *models.Transaction) error {
	return r.DB.WithContext(ctx).Create(trx).Error
}