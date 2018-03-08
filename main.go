package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Note struct {
	Id          int      `json:"id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Topic       *Topic   `json:"topic,omiempty"`
}

type Topic struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

var notes []Note

func main() {

	topic1 := Topic{Id: 1, Title: "Java", Description: "Object-oriented programming language"}
	topic2 := Topic{Id: 2, Title: "Python", Description: "Imperative programming language"}
	//topic3 := Topic{Id: 3, Title: "Go", Description: "Concurrent programming language"}

	note1 := Note{Id: 1, Title: "JDK vs JRE", Description: "The JDK is a superset of the JRE, and contains everything that is in the JRE, plus tools such as the compilers and debuggers necessary for developing applets and applications", Tags: []string{"Oracle", "Sun", "James Gosling"}, Topic: &topic1}
	note2 := Note{Id: 2, Title: "Flask Framework", Description: "", Tags: []string{"Python", "Microframework", "Pocoo"}, Topic: &topic2}
	note3 := Note{Id: 3, Title: "Rails Framework", Description: "Ruby on Rails is open source software, so not only is it free to use, you can also help make it better. More than 4,500 people already have contributed code to Rails", Tags: []string{"", ""}, Topic: &Topic{Id: 4, Title: "Ruby", Description: "Imperative programming language"}}

	notes = append(notes, note1)
	notes = append(notes, note2)
	notes = append(notes, note3)

	router := mux.NewRouter()
	router.HandleFunc("/notes", GetNotes).Methods("GET")
	router.HandleFunc("/notes", CreateNote).Methods("POST")
	router.HandleFunc("/notes/{id}", GetNote).Methods("GET")
	router.HandleFunc("/notes/{id}", UpdateNote).Methods("PUT")
	router.HandleFunc("/notes/{id}", DeleteNote).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(notes)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range notes {
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

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	notes = append(notes, note)
	json.NewEncoder(w).Encode(notes)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var note Note
	for index, item := range notes {
		i, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal("Conversion Error: %s", err)
		}
		if item.Id == i {
			_ = json.NewDecoder(r.Body).Decode(&note)
			notes[index] = note
			json.NewEncoder(w).Encode(notes)

		}
	}
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range notes {
		i, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal("Conversion Error: %s", err)
		}
		if item.Id == i {
			notes = append(notes[:index], notes[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(notes)
	}
}
