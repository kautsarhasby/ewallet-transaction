package constants

const (
	SuccessMessage = "success"
	ErrFailedBadRequest = "data not match"
	ErrServerError = "server went error"
)

const (
	TransactionStatusPending = "PENDING"
	TransactionStatusSuccess = "SUCCESS"
	TransactionStatusFailed = "FAILED"
	TransactionStatusReversed = "REVERSED"
)

const (
	TransactionTypeTopUp = "TOPUP"
	TransactionTypePurchase = "PURCHASE"
	TransactionTypeRefund = "REFUND"
)

var MapTransactionType = map[string] bool{
	TransactionTypeTopUp : true,
	TransactionTypeRefund : true,
	TransactionTypePurchase : true,
}