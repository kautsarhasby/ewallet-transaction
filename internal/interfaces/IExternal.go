package interfaces

import (
	"context"
	"kautsarhasby/ewallet-transaction/internal/models"
)


type IExternal interface {
	ValidateToken(ctx context.Context, token string) (models.TokenData, error)
}