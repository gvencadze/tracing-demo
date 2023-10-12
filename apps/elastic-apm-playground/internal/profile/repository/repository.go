package repository

import (
	"context"
	"errors"
	"github.com/gvencadze/apm-playground/internal/profile/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrUserNotFound = errors.New("user by id not found")
)

type Impl struct {
	pool *pgxpool.Pool
}

func (r *Impl) GetProfile(ctx context.Context, id int) (*model.Profile, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var profile model.Profile
	err = conn.QueryRow(ctx, `SELECT first_name, last_name FROM users WHERE id = $1`, id).
		Scan(&profile.FirstName, &profile.LastName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &profile, nil
}

func NewRepository(p *pgxpool.Pool) *Impl {
	return &Impl{
		pool: p,
	}
}
