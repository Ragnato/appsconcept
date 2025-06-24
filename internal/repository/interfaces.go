package repository

import "appsconcept/internal/domain"

type FizzBuzzRepository interface {
	SaveRequest(p domain.FizzBuzzParams) error
	GetTopRequest() (domain.StatsResponse, error)
}
