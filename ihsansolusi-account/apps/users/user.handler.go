package users

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (uh *UserHandler) RegisterUser(c echo.Context) error {
	var req struct {
		Name	string `json:"name"`
		NIK   	string `json:"nik"`
		Phone 	string `json:"phone"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	user, err := uh.Service.GetUserByNIKOrPhone(req.NIK, req.Phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if user != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "User already exists"})
	}

	user, err = uh.Service.RegisterUser(req.Name, req.NIK, req.Phone)
	if err != nil {
		return c.JSON(http.StatusConflict, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"account_number": user.AccountNumber})
}

func (uh *UserHandler) CheckBalance(c echo.Context) error {
	account_number := c.Param("account_number")

	if account_number == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Account number is required"})
	}

	balance, err := uh.Service.CheckBalance(string(account_number))
	if err != nil {
		if err.Error() == "User not found" {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch balance"})
	}

	return c.JSON(http.StatusOK, echo.Map{"balance": balance})
}

func (uh *UserHandler) GetUserByNIKOrPhone(c echo.Context) error {
	nik := c.QueryParam("nik")
	phone := c.QueryParam("phone")

	if nik == "" && phone == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "NIK or Phone is required"})
	}

	user, err := uh.Service.GetUserByNIKOrPhone(nik, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}