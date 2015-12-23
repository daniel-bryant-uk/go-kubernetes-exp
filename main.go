package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Tweet struct {
	Id string `json:"title"`
	UserName string `json:"UserName"`
	Content  string `json:"content"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello/{name}", index).Methods("GET")
	router.HandleFunc("/tweets", handleTweets).Methods("GET")

	log.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Responding to /hello request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello:", name)
}

func handleTweets(w http.ResponseWriter, r *http.Request) {
	log.Println("Responding to /tweets request")
	log.Println(r.UserAgent())

	r.Header.Set("Content-Type", "application/json")

	var tweets = map[string]*Tweet {
		"1": &Tweet{Id: "1", UserName: "danielbryantuk", Content:"Hello World"},
		"2": &Tweet{Id: "2", UserName: "danielbryantuk", Content:"Just setting up my twitter"},
	}

	outgoingJSON, error := json.Marshal(tweets)

	if error != nil {
		log.Println(error.Error())
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, string(outgoingJSON))
}