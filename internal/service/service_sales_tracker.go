package service

import (
	"context"
	"errors"
	"time"

	"github.com/Vladimirmoscow84/Sales_Tracker/internal/model"
)

func (s *Service) Create(ctx context.Context, t *model.Transaction) (int, error) {
	if t.Amount <= 0 {
		return 0, errors.New("amount must be > 0")
	}
	if t.Category == "" {
		return 0, errors.New("category is required")
	}
	return s.storageCRUD.Create(ctx, t)
}

func (s *Service) GetByID(ctx context.Context, id int) (*model.Transaction, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	return s.storageCRUD.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]model.Transaction, error) {
	return s.storageCRUD.GetAll(ctx)
}

func (s *Service) Update(ctx context.Context, t *model.Transaction) error {
	if t.ID <= 0 {
		return errors.New("invalid id")
	}
	return s.storageCRUD.Update(ctx, t)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	return s.storageCRUD.Delete(ctx, id)
}

func (s *Service) GetSum(ctx context.Context, from, to time.Time) (int64, error) {
	return s.storageAnalytics.GetSum(ctx, from, to)
}

func (s *Service) GetAvg(ctx context.Context, from, to time.Time) (float64, error) {
	return s.storageAnalytics.GetAvg(ctx, from, to)
}

func (s *Service) GetCount(ctx context.Context, from, to time.Time) (int64, error) {
	return s.storageAnalytics.GetCount(ctx, from, to)
}

func (s *Service) GetMedian(ctx context.Context, from, to time.Time) (float64, error) {
	return s.storageAnalytics.GetMedian(ctx, from, to)
}

func (s *Service) GetPercentile90(ctx context.Context, from, to time.Time) (float64, error) {
	return s.storageAnalytics.GetPercentile90(ctx, from, to)
}
