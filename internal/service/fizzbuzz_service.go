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

func (s *FizzBuzzService) GenerateFizzBuzz(c context.Context, p domain.FizzBuzzParams) ([]string, error) {
	if p.Int1 <= 0 || p.Int2 <= 0 || p.Limit <= 0 {
		return nil, fmt.Errorf("int1, int2, and limit must be positive integers")
	}

	if p.Str1 == "" || p.Str2 == "" {
		return nil, fmt.Errorf("str1 and str2 must not be empty")
	}

	err := s.repo.SaveRequest(c, p)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, p.Limit)
	for i := 1; i <= p.Limit; i++ {
		switch {
		case i%p.Int1 == 0 && i%p.Int2 == 0:
			result = append(result, p.Str1+p.Str2)
		case i%p.Int1 == 0:
			result = append(result, p.Str1)
		case i%p.Int2 == 0:
			result = append(result, p.Str2)
		default:
			result = append(result, fmt.Sprintf("%d", i))
		}
	}

	return result, nil
}
func (s *FizzBuzzService) GetStats() (domain.StatsResponse, error) {
	return s.repo.GetTopRequest()
}
