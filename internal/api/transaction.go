package api

import (
	"kautsarhasby/ewallet-transaction/constants"
	"kautsarhasby/ewallet-transaction/helpers"
	"kautsarhasby/ewallet-transaction/internal/interfaces"
	"kautsarhasby/ewallet-transaction/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionService interfaces.ITransactionService
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.Transaction
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}
	
	if err := req.Validate(); err != nil {
		log.Error("failed to parse request")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token data")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	tokenData, ok := token.(models.TokenData)
	if !ok {
		log.Error("failed to parse token data")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	if !constants.MapTransactionType[req.TransactionType]{
		log.Error("invalid transaction type")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	req.UserID = int(tokenData.UserID)
	resp, err := h.TransactionService.CreateTransaction(c.Request.Context(), &req)
	if err != nil {
		log.Error("failed to create transaction: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}
