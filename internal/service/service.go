package service

import (
	"context"
	"errors"
	"time"

	"github.com/Vladimirmoscow84/Sales_Tracker/internal/model"
)

type salesTrackerRepo interface {
	Create(ctx context.Context, t *model.Transaction) (int, error)
	GetByID(ctx context.Context, id int) (*model.Transaction, error)
	GetAll(ctx context.Context) ([]model.Transaction, error)
	Update(ctx context.Context, t *model.Transaction) error
	Delete(ctx context.Context, id int) error

	GetSum(ctx context.Context, from, to time.Time) (int64, error)
	GetAvg(ctx context.Context, from, to time.Time) (float64, error)
	GetCount(ctx context.Context, from, to time.Time) (int64, error)
	GetMedian(ctx context.Context, from, to time.Time) (float64, error)
	GetPercentile90(ctx context.Context, from, to time.Time) (float64, error)
}

type Service struct {
	storage salesTrackerRepo
}

func New(storage salesTrackerRepo) (*Service, error) {
	if storage == nil {
		return nil, errors.New("[service] storage is nil")
	}
	return &Service{
		storage: storage,
	}, nil
}
