package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"handlers"
	"log"
	"model"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Articles ...
var Articles []model.Article

type application struct {
	username string
	password string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	webapp := new(application)

	webapp.username = "admin" //os.Getenv("AUTH_USERNAME")
	webapp.password = "admin" //os.Getenv("AUTH_PASSWORD")

	if webapp.username == "" {
		log.Fatal("Illegal username provided")
	}

	if webapp.password == "" {
		log.Fatal("Illegal password provided")
	}

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(CommonMiddleware)
	//myRouter.HandleFunc("/", homePage).Methods(http.MethodGet)
	myRouter.HandleFunc("/", handlers.Ping).Methods(http.MethodGet)
	//myRouter.HandleFunc("/articles", returnAllArticles).Methods(http.MethodGet)
	//myRouter.HandleFunc("/articles", handlers.ReturnAllArticles).Methods(http.MethodGet)
	myRouter.HandleFunc("/articles", basicAuth(handlers.ReturnAllArticles))

	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods(http.MethodGet)
	myRouter.HandleFunc("/articlefromfile/{id}", handlers.LoadArticleFromFile).Methods(http.MethodGet)

	myRouter.Use(JwtVerify)
	//myRouter.HandleFunc("/uploadarticle", handlers.UploadArticleToFile).Methods(http.MethodPost)
	myRouter.HandleFunc("/uploadarticle", handlers.UploadArticleToFile).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

/* Old methods without files */
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Articles)
	//xml.NewEncoder(w).Encode(Articles)
}

/* Old methods without files */
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Key -: ", key, ", Endpoint Hit: returnSingleArticle")
	for _, article := range Articles {
		if article.ID == key {
			w.Header().Add("Content-Type", "application/json")
			article.Print()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(article)
		}
	}
}

func main() {
	handleRequests()
}

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte("admin"))
			expectedPasswordHash := sha256.Sum256([]byte("admin"))
			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)
			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")
		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)
		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		} //else {
		//json.NewEncoder(w).Encode(fmt.Sprintf("Token found. Value %s", header))
		//}
		next.ServeHTTP(w, r)
	})
}

func loadArticleInFile(articleList []model.Article, fileName string) error {
	var err error

	for _, articleLocal := range articleList {
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(articleLocal)
		if _, err = f.Write(reqBodyBytes.Bytes()); err != nil {
			panic(err)
		}
	}
	return err
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
