package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Article struct
type Article struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

//Status struct to give response as success or failure of query
type Status struct {
	Status   string `json:"status"`
	Comments string `json:"comments"`
}

//Articles = In-memory database to store articles
var Articles []Article

func createDummyArticles() {
	Articles = append(Articles, Article{1, "First Article...", "Thomas"})
	Articles = append(Articles, Article{2, "Second Article...", "Alva"})
	Articles = append(Articles, Article{3, "Third Article...", "Edison"})
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>API to dive into Articles</h1><p>Usage<br>/article/create => Creates New Article <br>/article/{article_id} => Fetch an article <br>/article/delete/{article_id} => delete an article <br></p>")
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
	contentType := r.Header.Get("Content-type")
	myarticle := Article{}
	id := rand.Intn(9999999)
	if contentType == "application/x-www-form-urlencoded" {
		r.ParseForm()
		myarticle = Article{id, r.FormValue("content"), r.FormValue("author")}
	}
	if contentType == "application/json" {
		reqBody, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(reqBody, &myarticle)
		if err != nil {
			json.NewEncoder(w).Encode(Status{"fail", "Invalid Request Body"})
			return
		}
		myarticle.ID = id
	}
	Articles = append(Articles, myarticle)
	json.NewEncoder(w).Encode(myarticle)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
	params := mux.Vars(r)
	for _, article := range Articles {
		if strconv.Itoa(article.ID) == params["article_id"] {
			err := json.NewEncoder(w).Encode(article)
			if err != nil {
				json.NewEncoder(w).Encode(Status{"fail", "Internal Server Error"})
			}
			return
		}
	}
	json.NewEncoder(w).Encode(Status{"fail", "Invalid Article ID"})

}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
	json.NewEncoder(w).Encode(Articles)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
	params := mux.Vars(r)
	for index, article := range Articles {
		if strconv.Itoa(article.ID) == params["article_id"] {
			Articles = append(Articles[:index], Articles[index+1:]...)
			json.NewEncoder(w).Encode(Status{Status: "success"})
			return
		}
	}
	json.NewEncoder(w).Encode(Status{"fail", "Invalid Article ID"})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	createDummyArticles()
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", welcomePage)
	muxRouter.HandleFunc("/article/create", createArticle).Methods("POST")
	muxRouter.HandleFunc("/article/allArticles", getAllArticles).Methods("GET")
	muxRouter.HandleFunc("/article/{article_id}", getArticle).Methods("GET")
	muxRouter.HandleFunc("/article/delete/{article_id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", muxRouter))

}
