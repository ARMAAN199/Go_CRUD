package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	// "math/rand"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Year   string  `json:"year"`
	Author *Author `json:"author"`
}

type Author struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {

	books = append(books, Book{ID: "1", Title: "Book One", Year: "2018", Author: &Author{ID: "1", Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Title: "Book Two", Year: "2019", Author: &Author{ID: "2", Firstname: "John2", Lastname: "Doe2"}})
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	// router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/addbook", createBook).Methods("POST")
	router.HandleFunc("/api/updatebook/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/deletebook/{id}", deleteBook).Methods("DELETE")

	log.Println("%s", "Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func checkserver(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var ind int
	for index, book := range books {
		if book.ID == params["id"] {
			ind = index
		}
	}
	books = append(books[:ind], books[ind+1:]...)
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// books = append(books, Book{ID: strconv.Itoa(len(books)) , Title: params["name"], Year: "2018", Author: &Author{ID: "1", Firstname: "John", Lastname: "Doe"}})
	// var book Book
	// dec := json.NewDecoder(r.Body)
	// err := dec.Decode(&book)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	r.ParseForm()
	if len(r.Form["title"]) == 0 || len(r.Form["year"]) == 0 || len(r.Form["authorf"]) == 0 || len(r.Form["authorl"]) == 0 {
		fmt.Fprintf(w, "Please enter all req feilds")
		return
	}
	books = append(books, Book{ID: strconv.Itoa(rand.Intn(10000)), Title: r.Form["title"][0], Year: r.Form["year"][0], Author: &Author{ID: "1", Firstname: r.Form["authorf"][0], Lastname: r.Form["authorl"][0]}})
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	r.ParseForm()
	var ind int
	for index, book := range books {
		if book.ID == params["id"] {
			ind = index
		}
	}
	log.Println(ind)
	log.Println(r.Form)
	if len(r.Form["title"]) != 0 {
		log.Println("here", books[ind].Title, r.Form["title"][0])
		books[ind].Title = r.Form["title"][0]
	}
	if len(r.Form["year"]) != 0 {
		books[ind].Year = r.Form["year"][0]
	}
	if len(r.Form["authorf"]) != 0 {
		books[ind].Author.Firstname = r.Form["authorf"][0]
	}
	if len(r.Form["authorl"]) != 0 {
		books[ind].Author.Lastname = r.Form["authorl"][0]
	}
	json.NewEncoder(w).Encode(books)
}
