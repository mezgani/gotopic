package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Topic struct {
	Id          int      `json:"id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

var topics []Topic

func main() {

	topic1 := Topic{Id: 1, Title: "Java", Description: "Object-oriented programming language", Tags: []string{"Sun", "James Gosling"}}
	topic2 := Topic{Id: 2, Title: "Python", Description: "imperative programming language", Tags: []string{"Guido van Rossum"}}
	topic3 := Topic{Id: 3, Title: "Go", Description: "Concurrent programming language", Tags: []string{"Rob Pike", "Google", "Ken Thompson"}}

	topics = append(topics, topic1)
	topics = append(topics, topic2)
	topics = append(topics, topic3)

	router := mux.NewRouter()
	router.HandleFunc("/topics", GetTopics).Methods("GET")
	router.HandleFunc("/topics", CreateTopic).Methods("POST")
	router.HandleFunc("/topics/{id}", GetTopic).Methods("GET")
	router.HandleFunc("/topics/{id}", UpdateTopic).Methods("PUT")
	router.HandleFunc("/topics/{id}", DeleteTopic).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetTopics(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(topics)
}

func GetTopic(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range topics {
		i, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal("Conversion Error: %s", err)
		}
		if item.Id == i {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	var topic Topic
	_ = json.NewDecoder(r.Body).Decode(&topic)
	topics = append(topics, topic)
	json.NewEncoder(w).Encode(topics)
}

func UpdateTopic(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var topic Topic
	for index, item := range topics {
		i, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal("Conversion Error: %s", err)
		}
		if item.Id == i {
			_ = json.NewDecoder(r.Body).Decode(&topic)
			topics[index] = topic
			json.NewEncoder(w).Encode(topics)

		}
	}
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range topics {
		i, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal("Conversion Error: %s", err)
		}
		if item.Id == i {
			topics = append(topics[:index], topics[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(topics)
	}
}
