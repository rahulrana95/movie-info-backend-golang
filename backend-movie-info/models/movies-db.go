package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// get for one movie
func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `select id, title, description, year, release_date, runtime, mpaa_rating, created_at, updated_at, poster_url, url from movies where id = $1 `
	fmt.Println(id)
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	fmt.Println(row)

	err := row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.ReleaseDate, &movie.Runtime, &movie.MPAARating, &movie.CreatedAt, &movie.UpdatedAt, &movie.PosterURL, &movie.URL)

	if err != nil {
		fmt.Println("Failed at row scan")
		fmt.Println(err.Error())

		return nil, err
	}

	return &movie, nil
}

// returns all movies and errors
func (m *DBModel) All() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `select id, title, description, year, release_date, runtime, mpaa_rating, created_at, updated_at, poster_url, url from movies `
	rows, error := m.DB.QueryContext(ctx, query)

	if error != nil {
		fmt.Println("Failed at row scan")
		fmt.Println(error.Error())

		return nil, error
	}
	var movies []*Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.ReleaseDate, &movie.Runtime, &movie.MPAARating, &movie.CreatedAt, &movie.UpdatedAt, &movie.PosterURL, &movie.URL)
		if err != nil {
			fmt.Println("Failed at row scan")
			fmt.Println(err.Error())

			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}
