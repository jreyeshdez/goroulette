package goroulette

import (
	"log"
	"strings"
	"testing"
)

func TestMoviesByTitle(t *testing.T) {
	movies, err := GetMoviesByTitle("Bad Boys")
	if err != nil {
		t.Errorf("Got unexpected error for requested title 'Bad Boys': %v", err)
		return
	}
	if (Movie{}) == movies {
		t.Errorf("Unexpected: could not get any movie with title 'Bad Boys'.")
	}
	if movies.GetTitle() != "Bad Boys" {
		t.Errorf("Unexpected title for movie 'Bad Boys': expected '%s', got '%s' instead.", "Bad Boys", movies.GetTitle())
	}
	log.Printf("Fetched: %v", movies)
}

func TestMoviesByTitleAndYear(t *testing.T) {
	movies, err := GetMoviesByTitleAndYear("Bad Boys", "1995")
	if err != nil {
		t.Errorf("Got unexpected error for requested title 'Bad Boys': %v", err)
		return
	}
	if (Movie{}) == movies {
		t.Errorf("Unexpected: could not get any movie with title 'Bad Boys'.")
	}
	if movies.GetTitle() != "Bad Boys" {
		t.Errorf("Unexpected title for movie 'Bad Boys': expected '%s', got '%s' instead.", "Bad Boys", movies.GetTitle())
	}
	if movies.GetReleaseYear() != "1995" {
		t.Errorf("Unexpected release year for movie 'Bad Boys': expected '%s', got '%s' instead.", "1995", movies.GetReleaseYear())
	}
	log.Printf("Fetched: %v", movies)
}

func TestMoviesByDirector(t *testing.T) {
	movies, err := GetMoviesByDirector("Quentin Tarantino")
	if err != nil {
		t.Errorf("Got unexpected error for requested director 'Quentin Tarantino': %v", err)
		return
	}
	if len(movies) < 1 {
		t.Errorf("Unexpected: could not get any movie with director 'Quentin Tarantino'.")
	}
	if (movies)[0].GetDirector() != "Quentin Tarantino" {
		t.Errorf("Unexpected director for director 'Quentin Tarantino': expected '%s', got '%s' instead.", "Quentin Tarantino", (movies)[0].GetDirector())
	}
	log.Printf("Fetched: %v", (movies)[0])
}

func TestMoviesByActor(t *testing.T) {
	movies, err := GetMoviesByActor("Martin Lawrence")
	if err != nil {
		t.Errorf("Got unexpected error for requested actor 'Martin Lawrence': %v", err)
		return
	}
	if len(movies) < 1 {
		t.Errorf("Unexpected: could not get any movie with actor 'Martin Lawrence'.")
	}
	actors := strings.Split((movies)[0].GetShowCast(), ",")
	for i := 0; i < len(actors); i++ {
		if actors[i] == "Martin Lawrence" {
			log.Printf("actor: %v", actors[i])
			break
		}
		t.Errorf("Unexpected show cast for actor 'Martin Lawrence'.")
		break
	}
	log.Printf("Fetched: %v", (movies)[0])
}
