package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/gufranmirza/imdb-cli/src/scrapper"

	"github.com/gufranmirza/imdb-cli/src/models"
	"github.com/spf13/cobra"
)

var (
	url   string
	limit int
)

// topMovies represents the top movies command
var topMovies = &cobra.Command{
	Use:   "top-movies",
	Short: "Fetch top IMDB movies",
	Long:  `Command line interface to fetch top movies from given IMDB Url`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	RootCmd.AddCommand(topMovies)

	topMovies.Flags().StringVar(&url, "url", "", "IMDB url of top movies")
	topMovies.Flags().IntVar(&limit, "limit", 10, "Display limit")
}

func run() {
	var wg sync.WaitGroup
	movies := make([]*models.Movie, limit)
	scrap := scrapper.NewScrapper()
	if len(url) == 0 {
		fmt.Printf("[ERROR] URL %s can not be empty\n", url)
		os.Exit(1)
	}

	links, err := scrap.GetTopMoviesURL(url)
	if err != nil {
		fmt.Printf("[ERROR] Failed to fetch URL %s with error %v\n", url, err)
		os.Exit(1)
	}

	for index, item := range links {
		if index >= limit {
			break
		}

		wg.Add(1)
		go func(item string, index int, wg *sync.WaitGroup) {
			defer wg.Done()
			movie, err := scrap.GetMovie(item)
			if err != nil {
				fmt.Printf("[ERROR] Failed to fetch URL %s with error %v\n", url, err)
				return
			}
			movies[index] = movie
		}(item, index, &wg)
	}

	wg.Wait()

	data, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		fmt.Printf("[ERROR] Failed to unmarshal movies details with error %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
