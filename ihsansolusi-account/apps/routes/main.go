package routes

import (
	"ihsansolusi-account/apps/databases/repositories"
	"ihsansolusi-account/apps/transactions"
	"ihsansolusi-account/apps/users"
	"ihsansolusi-account/config"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo               *echo.Echo
	UserHandler        *users.UserHandler
	TransactionHandler *transactions.TransactionHandler
}

func NewServer() *Server {
	config.InitDB()

	// Inisialisasi repository
	userRepo := repositories.NewUserRepository(config.DB)
	transactionRepo := repositories.NewTransactionRepository(config.DB)

	// Inisialisasi service
	userService := users.NewUserService(userRepo)
	transactionService := transactions.NewTransactionService(userRepo, transactionRepo)

	// Inisialisasi handler
	userHandler := users.NewUserHandler(userService)
	transactionHandler := transactions.NewTransactionHandler(transactionService)

	// Buat instance server
	server := &Server{
		Echo:               echo.New(),
		UserHandler:        userHandler,
		TransactionHandler: transactionHandler,
	}

	server.SetupRoutes()

	return server
}

func (s *Server) SetupRoutes() {
	RegisterUserRoutes(s.Echo, s.UserHandler)
	RegisterTransactionRoutes(s.Echo, s.TransactionHandler)
}
