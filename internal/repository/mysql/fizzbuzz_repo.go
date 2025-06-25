package mysql

import (
	"context"

	"appsconcept/internal/domain"
	"appsconcept/internal/repository"
	"database/sql"
)

type FizzBuzzRepo struct {
	db *sql.DB
}

func NewFizzBuzzRepo(db *sql.DB) repository.FizzBuzzRepository {
	//hack to create table instead of running migrations
	/**
	query := `
	CREATE TABLE IF NOT EXISTS fizzbuzz_requests (
		id INT AUTO_INCREMENT PRIMARY KEY,
		int1 INT NOT NULL,
		int2 INT NOT NULL,
		limit_val INT NOT NULL,
		str1 VARCHAR(50) NOT NULL,
		str2 VARCHAR(50) NOT NULL,
		count INT DEFAULT 1,
		UNIQUE KEY unique_request (int1, int2, limit_val, str1, str2)
	)`
	db.Exec(query)
	*/

	return &FizzBuzzRepo{db: db}
}

func (r *FizzBuzzRepo) SaveRequest(ctx context.Context, params domain.FizzBuzzParams) error {
	query := "INSERT INTO fizzbuzz_requests (`int1`, `int2`, `limit_val`, `str1`, `str2`) VALUES (?, ?, ?, ?, ?)"

	_, err := r.db.ExecContext(ctx, query, params.Int1, params.Int2, params.Limit, params.Str1, params.Str2)

	return err
}

func (r *FizzBuzzRepo) GetTopRequest(ctx context.Context) (*domain.StatsResponse, error) {
	query := "SELECT `int1`, `int2`, `limit_val`, `str1`, `str2`, COUNT(*) as `count_val` FROM `fizzbuzz_requests` GROUP BY `int1`, `int2`, `limit_val`, `str1`, `str2` ORDER BY `count_val` DESC LIMIT 1"

	row := r.db.QueryRowContext(ctx, query)

	var int1, int2, limit, count int
	var str1, str2 string

	err := row.Scan(&int1, &int2, &limit, &str1, &str2, &count)
	if err != nil {
		return nil, err
	}

	response := &domain.StatsResponse{
		Request: domain.FizzBuzzParams{
			Int1:  int1,
			Int2:  int2,
			Limit: limit,
			Str1:  str1,
			Str2:  str2,
		},
		Count: count,
	}

	return response, nil
}
