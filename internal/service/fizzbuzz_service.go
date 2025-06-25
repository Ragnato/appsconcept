package service

import (
	"context"
	"fmt"

	"appsconcept/internal/domain"
	"appsconcept/internal/repository"
)

type FizzBuzzService struct {
	repo repository.FizzBuzzRepository
}

func NewFizzBuzzService(r repository.FizzBuzzRepository) *FizzBuzzService {
	return &FizzBuzzService{repo: r}
}

func (s *FizzBuzzService) GenerateFizzBuzz(ctx context.Context, params domain.FizzBuzzParams) ([]string, error) {
	if params.Int1 <= 0 || params.Int2 <= 0 || params.Limit <= 0 {
		return nil, fmt.Errorf("int1, int2, and limit must be positive integers")
	}

	if params.Str1 == "" || params.Str2 == "" {
		return nil, fmt.Errorf("str1 and str2 must not be empty")
	}

	err := s.repo.SaveRequest(ctx, params)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, params.Limit)
	for i := 1; i <= params.Limit; i++ {
		switch {
		case i%params.Int1 == 0 && i%params.Int2 == 0:
			result = append(result, params.Str1+params.Str2)
		case i%params.Int1 == 0:
			result = append(result, params.Str1)
		case i%params.Int2 == 0:
			result = append(result, params.Str2)
		default:
			result = append(result, fmt.Sprintf("%d", i))
		}
	}

	return result, nil
}
func (s *FizzBuzzService) GetStats(ctx context.Context) (*domain.StatsResponse, error) {
	result, err := s.repo.GetTopRequest(ctx)
	return result, err
}
