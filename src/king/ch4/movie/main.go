package main

import (
	"encoding/json"
	"github.com/prometheus/common/log"
	"fmt"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Paul Newman"},
	},
	{
		Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
	},
}

func main() {
	{
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
	}

	{
		data, err := json.MarshalIndent(movies, "", "	")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)

		var titles []struct{Title string}
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(titles)
	}
}
