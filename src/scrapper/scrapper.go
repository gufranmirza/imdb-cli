package scrapper

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gufranmirza/imdb-cli/src/models"
)

type scrap struct{}

// NewScrapper returns the scrapper implementation
func NewScrapper() Scrapper {
	return &scrap{}
}

// GetTopMoviesURL returns the list of top movies for a given url
func (s *scrap) GetTopMoviesURL(url string) ([]string, error) {
	movies := make([]string, 250)
	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return movies, err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return movies, err
	}

	// Save each .titleColumn as a list
	doc.Find(".titleColumn").Each(func(i int, s *goquery.Selection) {
		link, found := s.Find("a").Attr("href")
		if found {
			movies[i] = "https://www.imdb.com/" + link
		}
	})

	return movies, nil
}

// GetMovie returns the movie details from given url
func (s *scrap) GetMovie(url string) (*models.Movie, error) {
	movie := &models.Movie{}
	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return movie, err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return movie, err
	}

	doc.Find(".title_wrapper").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h1").Text()
		year := s.Find("h1 span a").Text()
		movie.Title = strings.TrimSpace(title)
		movie.MovieReleaseYear = strings.TrimSpace(year)
	})

	doc.Find(".subtext").Each(func(i int, s *goquery.Selection) {
		duration := s.Find("div time").Text()
		genre := s.Find("div a").First().Text()
		movie.Duration = strings.TrimSpace(duration)
		movie.Genre = strings.TrimSpace(genre)
	})

	doc.Find(".summary_text").Each(func(i int, s *goquery.Selection) {
		desc := s.Text()
		movie.Summary = strings.TrimSpace(desc)
	})

	doc.Find(".ratingValue").Each(func(i int, s *goquery.Selection) {
		rating := s.Find("strong span").First().Text()
		movie.IMDBRating = strings.TrimSpace(rating)
	})

	return movie, nil
}
