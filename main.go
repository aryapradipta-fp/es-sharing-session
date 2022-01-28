package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return
	}

	res, err := es.Info()
	if err != nil {
		return
	}
	defer res.Body.Close()

	fmt.Println(res)

	for _, m := range movies {
		data, _ := json.Marshal(m)
		req := esapi.IndexRequest{
			Index: "movies",
			Body:  bytes.NewReader(data),
		}
		esres, _ := req.Do(context.Background(), es)

		esres.Body.Close()
	}

}
