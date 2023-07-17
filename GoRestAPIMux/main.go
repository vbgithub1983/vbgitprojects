package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article ...
type NewsArticle struct {
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

// Articles ...
var NewsArticles []NewsArticle

func homePageNews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleNewsRequests() {
	http.HandleFunc("/", homePageNews)
	// add our articles route and map it to our
	// returnAllArticles function like so
	http.HandleFunc("/newsarticles", returnAllNewsArticles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnAllNewsArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(NewsArticles)
	//xml.NewEncoder(w).Encode(Articles)
}

func main1() {
	NewsArticles = []NewsArticle{
		{Title: "Aajtak Dangal",
			Author: "Diya",
			Link:   "https://www.amazon.com/dp/B089KVK23P"},
		{Title: "ABP News Hallabol",
			Author: "Pinal",
			Link:   "https://www.amazon.com/dp/B089WH12CR"},
		{Title: "Zee News Black and White",
			Author: "Maahi",
			Link:   "https://www.amazon.com/dp/B089S58WWG"},
	}
	fmt.Println("This is testing from nano editor")
	handleNewsRequests()
}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article ...
type NewsArticle struct {
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

// Articles ...
var NewsArticles []NewsArticle

func homePageNews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleNewsRequests() {
	http.HandleFunc("/", homePageNews)
	// add our articles route and map it to our
	// returnAllArticles function like so
	http.HandleFunc("/newsarticles", returnAllNewsArticles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnAllNewsArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(NewsArticles)
	//xml.NewEncoder(w).Encode(Articles)
}

func main1() {
	NewsArticles = []NewsArticle{
		{Title: "Aajtak Dangal",
			Author: "Diya",
			Link:   "https://www.amazon.com/dp/B089KVK23P"},
		{Title: "ABP News Hallabol",
			Author: "Pinal",
			Link:   "https://www.amazon.com/dp/B089WH12CR"},
		{Title: "Zee News Black and White",
			Author: "Maahi",
			Link:   "https://www.amazon.com/dp/B089S58WWG"},
	}
	fmt.Println("This is testing from nano editor")
	handleNewsRequests()
}
