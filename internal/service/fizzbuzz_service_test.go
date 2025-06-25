package service

import (
	"appsconcept/internal/domain"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	saveErr error
}

func (m *mockRepo) SaveRequest(ctx context.Context, p domain.FizzBuzzParams) error {
	return m.saveErr
}
func (m *mockRepo) GetTopRequest(ctx context.Context) (*domain.StatsResponse, error) {
	return &domain.StatsResponse{}, nil
}

func TestGenerateFizzBuzz_ValidInput(t *testing.T) {
	ctx := context.Background()
	repo := &mockRepo{}
	svc := NewFizzBuzzService(repo)

	params := domain.FizzBuzzParams{
		Int1:  3,
		Int2:  5,
		Limit: 15,
		Str1:  "fizz",
		Str2:  "buzz",
	}

	expected := []string{
		"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz",
		"11", "fizz", "13", "14", "fizzbuzz",
	}

	result, err := svc.GenerateFizzBuzz(ctx, params)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGenerateFizzBuzz_InvalidInput(t *testing.T) {
	ctx := context.Background()
	repo := &mockRepo{}
	svc := NewFizzBuzzService(repo)

	invalidParams := []domain.FizzBuzzParams{
		{Int1: 0, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"},
		{Int1: 3, Int2: -1, Limit: 15, Str1: "fizz", Str2: "buzz"},
		{Int1: 3, Int2: 5, Limit: 0, Str1: "fizz", Str2: "buzz"},
		{Int1: 3, Int2: 5, Limit: 15, Str1: "", Str2: "buzz"},
	}

	for _, params := range invalidParams {
		result, err := svc.GenerateFizzBuzz(ctx, params)
		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

func TestGenerateFizzBuzz_SaveRequestFails(t *testing.T) {
	ctx := context.Background()
	repo := &mockRepo{saveErr: errors.New("db error")}
	svc := NewFizzBuzzService(repo)

	params := domain.FizzBuzzParams{
		Int1:  3,
		Int2:  5,
		Limit: 15,
		Str1:  "fizz",
		Str2:  "buzz",
	}

	result, err := svc.GenerateFizzBuzz(ctx, params)
	assert.Error(t, err)
	assert.Nil(t, result)
}
