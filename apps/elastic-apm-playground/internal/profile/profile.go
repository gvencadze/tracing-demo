package profile

import (
	"context"
	"fmt"
	"github.com/gvencadze/apm-playground/internal/profile/model"
	"github.com/gvencadze/apm-playground/internal/profile/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IProfile interface {
	Get(ctx context.Context, id int) (*model.Profile, error)
}

func New(pool *pgxpool.Pool) IProfile {
	return &Service{
		db: repository.NewRepository(pool),
	}
}

type Service struct {
	db *repository.Impl
}

func (s *Service) Get(ctx context.Context, id int) (*model.Profile, error) {
	res, err := s.db.GetProfile(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get profile: %w", err)
	}

	return res, nil
}
