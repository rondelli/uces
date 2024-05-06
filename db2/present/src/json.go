package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Sólo se marshalean los fields públicos
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

// Cristina OMIT
func main() {
	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Lilo & Stitch", Year: 2002, Color: true,
			Actors: []string{"Chris Sanders"}},
	}
	fmt.Printf("%v\n", movies)

	data, err := json.MarshalIndent(movies, "", "    ") //data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var películas []Movie
	err = json.Unmarshal(data, &películas)
	if err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}
	fmt.Printf("%v\n", películas)
}

//END OMIT
