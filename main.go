package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type movieResults struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIds         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

func main() {

	var executionLoop bool = true

	for executionLoop {
		var input string = ""
		fmt.Println("BEN MOVIE TRACKER")
		fmt.Println("type exit to exit\ntype search to enter search mode")
		fmt.Scanln(&input)
		if input == "exit" {
			executionLoop = false
		}
		if input == "search" {
			var searchMode bool = true
			for searchMode {
				var searchInput string = ""
				fmt.Println("type in the name of the movie you're looking for")
				fmt.Scanln(&searchInput)
				if searchInput == "exit" {
					searchMode = false
					break
				}
				url := "https://api.themoviedb.org/3/search/movie?query=" + searchInput + "&include_adult=false&language=en-US&page=1"
				req, _ := http.NewRequest("GET", url, nil)

				req.Header.Add("accept", "application/json")
				req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI0ODZhYWRjOTBiZTU4ZTBmZTFmYjkyYzVmYTIyOGUyZCIsIm5iZiI6MTc1ODA3MDcyMy4zODEsInN1YiI6IjY4Y2EwN2MzOWQ2YjZjMDEwOTEzZjQyMiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.0wTlnba8SavPvlv-Q8n3sOMFFLE-jfkgnI7RRPoCUTQ")

				res, _ := http.DefaultClient.Do(req)

				defer res.Body.Close()
				body, _ := io.ReadAll(res.Body)

				var results movieResults

				err := json.Unmarshal(body, &results)
				if err != nil {
					panic(err)
				}

				for index, element := range results.Results {
					releaseYear := ""
					if element.ReleaseDate == "" {
						releaseYear = "unknown"
					} else {
						releaseYear = element.ReleaseDate[0:4]
					}
					fmt.Printf("%d. "+element.Title+" ("+releaseYear+")\n", index+1)
				}
			}
		}

	}
}
