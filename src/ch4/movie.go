package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey", "Ingrid Bergman"}},
	{Title: "aasablanca", Year: 1942, Color: false, Actors: []string{"Humphrey", "Ingrid Bergman"}},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
