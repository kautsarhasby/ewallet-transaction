package services

import (
	"context"
	"kautsarhasby/ewallet-transaction/constants"
	"kautsarhasby/ewallet-transaction/helpers"
	"kautsarhasby/ewallet-transaction/internal/interfaces"
	"kautsarhasby/ewallet-transaction/internal/models"

	"github.com/pkg/errors"
)

type TransactionService struct {
	TransactionRepository interfaces.ITransactionRepository
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req *models.Transaction) (models.TransactionResponse, error) {
	var resp models.TransactionResponse

	req.TransactionStatus = constants.TransactionStatusPending
	req.Reference = helpers.GenerateReference()
	err := s.TransactionRepository.CreateTransaction(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "failed to create transaction")
	}

	resp.Reference = req.Reference
	resp.TransactionStatus = req.TransactionStatus
	return resp, nil
}