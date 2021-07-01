package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type MoviePractice struct {
	Title  string
	Year   int    `json:"Released"`
	Color  string `json:"color,omitempty"`
	Actors []string
}

var moviePractice = []MoviePractice{
	{Title: "Victory", Year: 1948, Color: "Blue", Actors: []string{"Jack", "Jim"}},
	{Title: "Cut", Year: 1950, Color: "Yellow", Actors: []string{"Henry", "Cat"}},
}


func main() {
	// marshal
	data, err := json.Marshal(moviePractice)
	if err != nil {
		log.Fatalf("marshal error with err： %s", err)
	}
	fmt.Printf("marhsal data is: %s \n", data)

	data, err = json.MarshalIndent(moviePractice, "", "    ")
	if err != nil {
		log.Fatalf("marshal error with err: %s", err)
	}
	fmt.Printf("marshalIndent is: %s \n", data)

	// unmarshal
	var titles []struct{Title string}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("unmarshal error with err: %s", err)
	}
	fmt.Printf("unmarshal titles is: %s \n", titles)

}
