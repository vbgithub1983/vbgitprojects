<<<<<<< HEAD
package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"model"

	"github.com/gorilla/mux"
)

func LoadArticleFromFile(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./article/article.txt") // strings.NewReader("article.txt")
	fmt.Fprintln(os.Stderr, "Error Reading From File:", err)

	//var Articles []*Article
	var singleArticleFromFile model.Article
	var singleArticleFromFileResponse model.Article

	scanner := bufio.NewScanner(file)
	vars := mux.Vars(r)
	key := vars["id"]

	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		err := json.Unmarshal([]byte(scanner.Text()), &singleArticleFromFile)
		if err != nil {
			panic(err)
		}
		id := singleArticleFromFile.ID
		if id == key {
			singleArticleFromFile.Print()
			singleArticleFromFileResponse = singleArticleFromFile
		}

	}
	if len(singleArticleFromFileResponse.Title) != 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(singleArticleFromFileResponse)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No Book found with ID....")
	}

}

/* Will create one article in file and store */
func UploadArticleToFile(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(os.Stdout, "Received Body - ", string(body))
	var articleToAdd model.Article
	json.Unmarshal(body, &articleToAdd)

	f, err := os.OpenFile("./article/article.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(articleToAdd)
	if _, err = f.Write(reqBodyBytes.Bytes()); err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Article Created")

}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am alive....")
	json.NewEncoder(w).Encode("I am alive....")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ReturnAllArticles")
	w.Header().Add("Content-Type", "application/json")

	file, err := os.Open("./article/article.txt") // strings.NewReader("article.txt")
	fmt.Fprintln(os.Stderr, "Error Reading From File:", err)

	var articles []model.Article
	//articles := make([]model.Article, 100)
	var singleArticleFromFile model.Article

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}

		err := json.Unmarshal([]byte(scanner.Text()), &singleArticleFromFile)
		if err != nil {
			panic(err)
		}
		articles = append(articles, singleArticleFromFile)
		//articles[i] = singleArticleFromFile
		i++
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}
=======
package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"model"

	"github.com/gorilla/mux"
)

func LoadArticleFromFile(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./article/article.txt") // strings.NewReader("article.txt")
	fmt.Fprintln(os.Stderr, "Error Reading From File:", err)

	//var Articles []*Article
	var singleArticleFromFile model.Article
	var singleArticleFromFileResponse model.Article

	scanner := bufio.NewScanner(file)
	vars := mux.Vars(r)
	key := vars["id"]

	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		err := json.Unmarshal([]byte(scanner.Text()), &singleArticleFromFile)
		if err != nil {
			panic(err)
		}
		id := singleArticleFromFile.ID
		if id == key {
			singleArticleFromFile.Print()
			singleArticleFromFileResponse = singleArticleFromFile
		}

	}
	if len(singleArticleFromFileResponse.Title) != 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(singleArticleFromFileResponse)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No Book found with ID....")
	}

}

/* Will create one article in file and store */
func UploadArticleToFile(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(os.Stdout, "Received Body - ", string(body))
	var articleToAdd model.Article
	json.Unmarshal(body, &articleToAdd)

	f, err := os.OpenFile("./article/article.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(articleToAdd)
	if _, err = f.Write(reqBodyBytes.Bytes()); err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Article Created")

}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am alive....")
	json.NewEncoder(w).Encode("I am alive....")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ReturnAllArticles")
	w.Header().Add("Content-Type", "application/json")

	file, err := os.Open("./article/article.txt") // strings.NewReader("article.txt")
	fmt.Fprintln(os.Stderr, "Error Reading From File:", err)

	var articles []model.Article
	//articles := make([]model.Article, 100)
	var singleArticleFromFile model.Article

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}

		err := json.Unmarshal([]byte(scanner.Text()), &singleArticleFromFile)
		if err != nil {
			panic(err)
		}
		articles = append(articles, singleArticleFromFile)
		//articles[i] = singleArticleFromFile
		i++
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}
>>>>>>> 1924c79e50cabb512fc0038add33ee24bc86cbca
