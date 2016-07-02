package goroulette

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseURL = "http://netflixroulette.net/api/api.php"

type Movie struct {
	Unit        int    `json:unit`
	ShowID      int    `json:show_id`
	ShowTitle   string `json:"show_title"`
	ReleaseYear string `json:"release_year"`
	Rating      string `json:"rating"`
	Category    string `json:"category"`
	ShowCast    string `json:"show_cast"`
	Director    string `json:"director"`
	Summary     string `json:"summary"`
	Poster      string `json:"poster"`
	MediaType   int    `json:mediatype`
	Runtime     string `json:"runtime"`
}

func GetURL(params map[string]string) (string, error) {
	var URL *url.URL
	URL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	parameters := url.Values{}
	for param, value := range params {
		parameters.Add(param, value)
		URL.RawQuery = parameters.Encode()
	}
	return URL.String(), nil
}

func CreateRouletteCall(params map[string]string) ([]byte, error) {
	url, err := GetURL(params)
	if err != nil {
		return []byte{}, err
	}
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		e := errors.New(fmt.Sprintf("Unexpected API status code %s", res.Status))
		return []byte{}, e
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func GetMoviesByTitle(title string) (Movie, error) {
	m := map[string]string{"title": title}
	data, err := CreateRouletteCall(m)
	if err != nil {
		return Movie{}, err
	}
	var movie Movie
	err = json.Unmarshal(data, &movie)
	if err != nil {
		return Movie{}, err
	}
	return movie, nil
}

func GetMoviesByTitleAndYear(title string, year string) (Movie, error) {
	m := map[string]string{"title": title, "year": year}
	data, err := CreateRouletteCall(m)
	var movie Movie
	if err != nil {
		return movie, err
	}
	err = json.Unmarshal(data, &movie)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func GetMoviesByDirector(director string) ([]Movie, error) {
	m := map[string]string{"director": director}
	data, err := CreateRouletteCall(m)
	if err != nil {
		return nil, err
	}
	var movie []Movie
	err = json.Unmarshal(data, &movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func GetMoviesByActor(actor string) ([]Movie, error) {
	m := map[string]string{"actor": actor}
	data, err := CreateRouletteCall(m)
	if err != nil {
		return nil, err
	}
	var movie []Movie
	err = json.Unmarshal(data, &movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (movie Movie) GetNetflixID() int {
	return movie.ShowID
}

func (movie Movie) GetTitle() string {
	return movie.ShowTitle
}

func (movie Movie) GetReleaseYear() string {
	return movie.ReleaseYear
}

func (movie Movie) GetRating() string {
	return movie.Rating
}

func (movie Movie) GetCategory() string {
	return movie.Category
}

func (movie Movie) GetShowCast() string {
	return movie.ShowCast
}

func (movie Movie) GetDirector() string {
	return movie.Director
}

func (movie Movie) GetSummary() string {
	return movie.Summary
}

func (movie Movie) GetPoster() string {
	return movie.Poster
}

func (movie Movie) GetMediaType() int {
	return movie.MediaType
}

func (movie Movie) GetRuntime() string {
	return movie.Runtime
}
