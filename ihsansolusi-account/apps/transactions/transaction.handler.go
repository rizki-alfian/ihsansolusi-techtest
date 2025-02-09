package transactions

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	Service TransactionService
}

func NewTransactionHandler(service TransactionService) *TransactionHandler {
	return &TransactionHandler{Service: service}
}

func (th *TransactionHandler) Deposit(c echo.Context) error {
	var req struct {
		AccountNumber 	string  `json:"account_number"`
		Amount 			float64 `json:"amount"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}
	err := th.Service.ProcessTransaction(req.AccountNumber, req.Amount, "deposit")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Deposit successful"})
}

func (th *TransactionHandler) Withdraw(c echo.Context) error {
	var req struct {
		AccountNumber string    `json:"account_number"`
		Amount float64 `json:"amount"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}
	err := th.Service.ProcessTransaction(req.AccountNumber, req.Amount, "withdraw")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Withdraw successful"})
}

func (th *TransactionHandler) Transfer(c echo.Context) error {
	var req struct {
		SenderID   string    `json:"sender_id"`
		ReceiverID string    `json:"receiver_id"`
		Amount     float64 `json:"amount"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}
	err := th.Service.ProcessTransfer(req.SenderID, req.ReceiverID, req.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Transfer successful"})
}
