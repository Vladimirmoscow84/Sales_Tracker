package storage

import (
	"context"
	"errors"
	"time"

	"github.com/Vladimirmoscow84/Sales_Tracker/internal/model"
)

type transactionRepo interface {
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

type Storage struct {
	transactions transactionRepo
}

func New(trans transactionRepo) (*Storage, error) {
	if trans == nil {
		return nil, errors.New("[storage] repo is nil")
	}
	return &Storage{
		transactions: trans,
	}, nil
}

func (s *Storage) Create(ctx context.Context, t *model.Transaction) (int, error) {
	return s.transactions.Create(ctx, t)
}

func (s *Storage) GetByID(ctx context.Context, id int) (*model.Transaction, error) {
	return s.transactions.GetByID(ctx, id)
}

func (s *Storage) GetAll(ctx context.Context) ([]model.Transaction, error) {
	return s.transactions.GetAll(ctx)
}

func (s *Storage) Update(ctx context.Context, t *model.Transaction) error {
	return s.transactions.Update(ctx, t)
}

func (s *Storage) Delete(ctx context.Context, id int) error {
	return s.transactions.Delete(ctx, id)
}

func (s *Storage) GetSum(ctx context.Context, from, to time.Time) (int64, error) {
	return s.transactions.GetSum(ctx, from, to)
}

func (s *Storage) GetAvg(ctx context.Context, from, to time.Time) (float64, error) {
	return s.transactions.GetAvg(ctx, from, to)
}

func (s *Storage) GetCount(ctx context.Context, from, to time.Time) (int64, error) {
	return s.transactions.GetCount(ctx, from, to)
}

func (s *Storage) GetMedian(ctx context.Context, from, to time.Time) (float64, error) {
	return s.transactions.GetMedian(ctx, from, to)
}

func (s *Storage) GetPercentile90(ctx context.Context, from, to time.Time) (float64, error) {
	return s.transactions.GetPercentile90(ctx, from, to)
}
