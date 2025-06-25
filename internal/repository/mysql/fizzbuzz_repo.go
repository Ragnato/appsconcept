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

	if err != nil {
		return err
	}

	return nil
}

func (r *FizzBuzzRepo) GetTopRequest() (domain.StatsResponse, error) {
	var res domain.StatsResponse

	row := r.db.QueryRow(`
		SELECT int1, int2, limit_val, str1, str2, count
		FROM fizzbuzz_requests
		ORDER BY count DESC
		LIMIT 1;
	`)

	err := row.Scan(&res.Request.Int1, &res.Request.Int2, &res.Request.Limit,
		&res.Request.Str1, &res.Request.Str2, &res.Count)
	return res, err
}
