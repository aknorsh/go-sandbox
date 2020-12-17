package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie data
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {

	var movies = []Movie{
		{Title: "Hoge", Year: 1999, Color: false,
			Actors: []string{"Fuu", "Bar", "Baz"}},
		{Title: "Fuga", Year: 2003, Color: true,
			Actors: []string{"Fuu", "Baz", "Piy"}},
		{Title: "Piyo", Year: 2007, Color: true,
			Actors: []string{"Baz", "Kru", "Tog"}},
	}

	// Marshal it
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("movie: %v\n", err)
	} else {
		fmt.Printf("%s\n", data)
	}

	// Marshal for human
	if data, err := json.MarshalIndent(movies, "", "- "); err != nil {
		log.Fatalf("movie: %v\n", err)
	} else {
		fmt.Printf("%s\n", data)
	}

	// Unmarshal
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("movie: %v\n", err)
	} else {
		fmt.Println(titles)
	}
}
