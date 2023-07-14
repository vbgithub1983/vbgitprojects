package main

import (
	"bytes"
	"encoding/json"
	"handlers"
	"model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestReturnAllArticles(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles", nil)
	req.Header.Add("x-access-token", "556671")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ReturnAllArticles)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	//expected := `[{"Title":"Go Lang API Testing","author":"Vijaykumar Bharoliya","link":"https://www.amazon.com/dp/B089KVK23P","Id":"1"}]`
	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "")
	}
}

func TestLoadArticleFromFile(t *testing.T) {

	req, err := http.NewRequest("GET", "/articlefromfile/5", nil)
	req.Header.Add("x-access-token", "556671")
	rr := httptest.NewRecorder()

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc("/articlefromfile/{id}", handlers.LoadArticleFromFile)
	router.ServeHTTP(rr, req)

	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var articleToReceive model.Article
	json.Unmarshal(rr.Body.Bytes(), &articleToReceive)

	// Check the response body is what we expect.
	if articleToReceive.ID != "5" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			articleToReceive.ID, "5")
	}
}

func TestUploadArticleToFile(t *testing.T) {

	var jsonStr = []byte(`{"Title":"Go Lang API Testing","author":"Vijaykumar Bharoliya","link":"https://www.amazon.com/dp/B089KVK23P","Id":"5"}`)

	req, err := http.NewRequest("POST", "/uploadarticle", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.UploadArticleToFile)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}
