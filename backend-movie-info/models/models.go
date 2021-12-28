package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate string       `json:"releaseDate"`
	Runtime     int          `json:"runtime"`
	Rating      int          `json:"rating"`
	MPAARating  int          `json:"mpaa_rating"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	URL         string       `json:"url"`
	source      string       `json:"source"`
	PosterURL   string       `json:"poster_url"`
	MovieGenre  []MovieGenre `json:"movie_genre"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movie_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
