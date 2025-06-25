package repository

import (
	"context"

	"appsconcept/internal/domain"
)

type FizzBuzzRepository interface {
	SaveRequest(ctx context.Context, params domain.FizzBuzzParams) error
	GetTopRequest(ctx context.Context) (*domain.StatsResponse, error)
}
