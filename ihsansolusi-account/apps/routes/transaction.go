package routes

import (
	"ihsansolusi-account/apps/transactions"

	"github.com/labstack/echo/v4"
)

func RegisterTransactionRoutes(e *echo.Echo, transactionHandler *transactions.TransactionHandler) {
	e.POST("/transaction/deposit", transactionHandler.Deposit)
	e.POST("/transaction/withdraw", transactionHandler.Withdraw)
	e.POST("/transaction/transfer", transactionHandler.Transfer)
}