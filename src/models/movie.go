package models

// Movie holds a IMDB movie
type Movie struct {
	Title            string `json:"title"`
	MovieReleaseYear string `json:"movie_release_year"`
	IMDBRating       string `json:"imdb_rating"`
	Summary          string `json:"summary"`
	Duration         string `json:"duration"`
	Genre            string `json:"genre"`
}
