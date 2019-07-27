package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Rating struct {
	Max       int
	Average   string
	NumRaters int
	min       int
}

type Author struct {
	Name string `json:"name"`
}

type Attrs struct {
	PubDate       []string `json:"pubdate"`
	Language      []string `json:"language"`
	Ttile         []string `json:"title"`
	Country       []string `json:"country"`
	Writer        []string `json:"writer"`
	Director      []string `json:"director"`
	Cast          []string `json:"cast"`
	MovieDuration []string `json:"movie_duration"`
	Year          []string `json:"year"`
	MovieType     []string `json:"movie_type"`
}

type Tag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Movie struct {
	Rating   Rating   `json:"rating"`
	Author   []Author `json:"author"`
	Attrs    Attrs    `json:"attrs"`
	Title    string   `json:"title"`
	AltTitle string   `json:"alt_title"`
	Image    string   `json:"image"`
	Alt      string   `json:"alt"`
	ID       string   `json:"id"`
	Tags     []Tag    `json:"tags"`
}

func loadHTTP() {
	key, _ := os.LookupEnv("DOUBAN_KEY")
	url := "https://api.douban.com/v2/movie/imdb/tt0111161?apikey=" + key
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("err %v", err)
		return
	}
	defer response.Body.Close()
	var m Movie
	// fmt.Println(response.Body.Read())
	err = json.NewDecoder(response.Body).Decode(&m)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Printf("%v", m)
}

func loadFile() {
	file, err := os.Open("movie.json")
	if err != nil {
		fmt.Printf("err %v", err)
		return
	}

	// var m Movie
	var m map[string]interface{}
	err = json.NewDecoder(file).Decode(&m)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Printf("%v", m)
}

func encode() {
	m := Movie{
		Title:    "The Shawshank Redemption",
		AltTitle: "肖申克的救赎 / 月黑高飞(港)",
		Image:    "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p480747492.webp",
		ID:       "xxxx",
		Alt:      "zzzz",
	}

	// data, err := json.Marshal(m)
	data, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	loadFile()
}
