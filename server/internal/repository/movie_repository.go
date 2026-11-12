package repository

import (
	"context"
	"errors"
	"fmt"
	"movie-api/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrMovieNotFound = errors.New("movie not found")

type MovieRepository struct {
	pool *pgxpool.Pool
}

func NewMovieRepository(pool *pgxpool.Pool) *MovieRepository {
	return &MovieRepository{pool: pool}
}

func (r *MovieRepository) GetAll(ctx context.Context) ([]model.Movie, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, isbn, title, director_first_name, director_last_name
		FROM movies ORDER BY created_at
	`)
	if err != nil {
		return nil, fmt.Errorf("query movies: %w", err)
	}
	defer rows.Close()

	movies := make([]model.Movie, 0)
	for rows.Next() {
		var movie model.Movie
		movie.Director = &model.Director{}
		if err := rows.Scan(&movie.ID, &movie.Isbn, &movie.Title, &movie.Director.FirstName, &movie.Director.LastName); err != nil {
			return nil, fmt.Errorf("scan movie: %w", err)
		}
		movies = append(movies, movie)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate movies: %w", err)
	}
	return movies, nil
}

func (r *MovieRepository) GetByID(ctx context.Context, id string) (model.Movie, error) {
	var movie model.Movie
	movie.Director = &model.Director{}
	err := r.pool.QueryRow(ctx, `
		SELECT id, isbn, title, director_first_name, director_last_name
		FROM movies WHERE id = $1
	`, id).Scan(&movie.ID, &movie.Isbn, &movie.Title, &movie.Director.FirstName, &movie.Director.LastName)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Movie{}, ErrMovieNotFound
	}
	if err != nil {
		return model.Movie{}, fmt.Errorf("get movie: %w", err)
	}
	return movie, nil
}

func (r *MovieRepository) Create(ctx context.Context, movie model.Movie) (model.Movie, error) {
	err := r.pool.QueryRow(ctx, `
		INSERT INTO movies (isbn, title, director_first_name, director_last_name)
		VALUES ($1, $2, $3, $4) RETURNING id
	`, movie.Isbn, movie.Title, movie.Director.FirstName, movie.Director.LastName).Scan(&movie.ID)
	if err != nil {
		return model.Movie{}, fmt.Errorf("create movie: %w", err)
	}
	return movie, nil
}

func (r *MovieRepository) Update(ctx context.Context, id string, movie model.Movie) (model.Movie, error) {
	result, err := r.pool.Exec(ctx, `
		UPDATE movies SET isbn = $1, title = $2,
			director_first_name = $3, director_last_name = $4, updated_at = NOW()
		WHERE id = $5
	`, movie.Isbn, movie.Title, movie.Director.FirstName, movie.Director.LastName, id)
	if err != nil {
		return model.Movie{}, fmt.Errorf("update movie: %w", err)
	}
	if result.RowsAffected() == 0 {
		return model.Movie{}, ErrMovieNotFound
	}
	movie.ID = id
	return movie, nil
}

func (r *MovieRepository) Delete(ctx context.Context, id string) error {
	result, err := r.pool.Exec(ctx, `DELETE FROM movies WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete movie: %w", err)
	}
	if result.RowsAffected() == 0 {
		return ErrMovieNotFound
	}
	return nil
}
