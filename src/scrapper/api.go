package scrapper

import "github.com/gufranmirza/imdb-cli/src/models"

// Scrapper implents the scrapper interface
type Scrapper interface {
	GetMovie(url string) (*models.Movie, error)
	GetTopMoviesURL(url string) ([]string, error)
}
