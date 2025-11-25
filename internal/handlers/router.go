package handlers

import (
	"context"
	"time"

	"github.com/Vladimirmoscow84/Sales_Tracker/internal/model"
	"github.com/wb-go/wbf/ginext"
)

type transactionCRUDer interface {
	Create(ctx context.Context, t *model.Transaction) (int, error)
	GetByID(ctx context.Context, id int) (*model.Transaction, error)
	GetAll(ctx context.Context) ([]model.Transaction, error)
	Update(ctx context.Context, t *model.Transaction) error
	Delete(ctx context.Context, id int) error
}

type transactionAnalyticsGetter interface {
	GetSum(ctx context.Context, from, to time.Time) (int64, error)
	GetAvg(ctx context.Context, from, to time.Time) (float64, error)
	GetCount(ctx context.Context, from, to time.Time) (int64, error)
	GetMedian(ctx context.Context, from, to time.Time) (float64, error)
	GetPercentile90(ctx context.Context, from, to time.Time) (float64, error)
}

type Router struct {
	Router           *ginext.Engine
	tCRUDer          transactionCRUDer
	tAnalyticsGetter transactionAnalyticsGetter
}

func New(router *ginext.Engine, cruder transactionCRUDer, analyticGetter transactionAnalyticsGetter) *Router {
	return &Router{
		Router:           router,
		tCRUDer:          cruder,
		tAnalyticsGetter: analyticGetter,
	}
}

func (r *Router) Routes() {

}
