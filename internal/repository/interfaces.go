package repository

import (
	"context"

	"appsconcept/internal/domain"
)

type FizzBuzzRepository interface {
	SaveRequest(c context.Context, params domain.FizzBuzzParams) error
	GetTopRequest() (domain.StatsResponse, error)
}
