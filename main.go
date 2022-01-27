package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Cast   []string `json:"cast"`
	Genres []string `json:"genres"`
}

func main() {
	path := "movies.json"
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	var movies []Movie
	err = json.Unmarshal(jsonData, &movies)
	if err != nil {
		return
	}

	fmt.Printf("%+v\n", movies[0])
}
