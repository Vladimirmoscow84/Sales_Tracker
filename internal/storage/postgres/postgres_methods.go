package postgres

import (
	"context"
	"time"

	"github.com/Vladimirmoscow84/Sales_Tracker/internal/model"
)

func (p *Postgres) Create(ctx context.Context, t *model.Transaction) (int, error) {
	query := `
        INSERT INTO transactions (name, description, amount, type, category, event_date)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id;
    `
	var id int
	err := p.DB.QueryRowContext(
		ctx, query,
		t.Name, t.Description, t.Amount, t.Type, t.Category, t.EventDate,
	).Scan(&id)
	return id, err
}

func (p *Postgres) GetByID(ctx context.Context, id int) (*model.Transaction, error) {
	query := `
        SELECT id, name, description, amount, type, category, event_date
        FROM transactions
        WHERE id = $1;
    `
	var t model.Transaction
	err := p.DB.GetContext(ctx, &t, query, id)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (p *Postgres) GetAll(ctx context.Context) ([]model.Transaction, error) {
	query := `
        SELECT id, name, description, amount, type, category, event_date
        FROM transactions
        ORDER BY event_date DESC;
    `
	var list []model.Transaction
	err := p.DB.SelectContext(ctx, &list, query)
	return list, err
}

func (p *Postgres) Update(ctx context.Context, t *model.Transaction) error {
	query := `
        UPDATE transactions
        SET name = $1,
            description = $2,
            amount = $3,
            type = $4,
            category = $5,
            event_date = $6
        WHERE id = $7;
    `
	_, err := p.DB.ExecContext(
		ctx, query,
		t.Name, t.Description, t.Amount, t.Type, t.Category, t.EventDate,
		t.ID,
	)
	return err
}

func (p *Postgres) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM transactions WHERE id = $1;`
	_, err := p.DB.ExecContext(ctx, query, id)
	return err
}

func (p *Postgres) GetSum(ctx context.Context, from, to time.Time) (int64, error) {
	query := `
        SELECT COALESCE(SUM(amount), 0)
        FROM transactions
        WHERE event_date BETWEEN $1 AND $2;
    `
	var sum int64
	err := p.DB.GetContext(ctx, &sum, query, from, to)
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func (p *Postgres) GetAvg(ctx context.Context, from, to time.Time) (float64, error) {
	query := `
        SELECT COALESCE(AVG(amount), 0)
        FROM transactions
        WHERE event_date BETWEEN $1 AND $2;
    `
	var avg float64
	err := p.DB.GetContext(ctx, &avg, query, from, to)
	if err != nil {
		return 0, err
	}
	return avg, nil
}

func (p *Postgres) GetCount(ctx context.Context, from, to time.Time) (int64, error) {
	query := `
        SELECT COUNT(*)
        FROM transactions
        WHERE event_date BETWEEN $1 AND $2;
    `
	var count int64
	err := p.DB.GetContext(ctx, &count, query, from, to)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *Postgres) GetMedian(ctx context.Context, from, to time.Time) (float64, error) {
	query := `
        SELECT COALESCE(
            percentile_cont(0.5) WITHIN GROUP (ORDER BY amount),
            0
        )
        FROM transactions
        WHERE event_date BETWEEN $1 AND $2;
    `
	var median float64
	err := p.DB.GetContext(ctx, &median, query, from, to)
	if err != nil {
		return 0, err
	}
	return median, nil
}

func (p *Postgres) GetPercentile90(ctx context.Context, from, to time.Time) (float64, error) {
	query := `
        SELECT COALESCE(
            percentile_cont(0.9) WITHIN GROUP (ORDER BY amount),
            0
        )
        FROM transactions
        WHERE event_date BETWEEN $1 AND $2;
    `
	var result float64
	err := p.DB.GetContext(ctx, &result, query, from, to)
	if err != nil {
		return 0, err
	}
	return result, nil
}
