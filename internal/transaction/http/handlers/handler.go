package handlers

import (
	"bwg_test/internal/transaction/models"
	"context"
	"github.com/rs/zerolog"
)

type IService interface {
	Input(ctx context.Context, transaction *models.Transaction) error
	Output(ctx context.Context, transaction *models.Transaction) error
	GetTransactions(ctx context.Context, userID int) ([]*models.Transaction, error)
	GetBalance(ctx context.Context, userID int) (*models.Balance, error)
}

type Handler struct {
	logger  zerolog.Logger
	service IService
}

func New(logger zerolog.Logger, svc IService) *Handler {
	return &Handler{
		logger:  logger,
		service: svc,
	}
}
