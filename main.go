package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	// router.HandleFunc("/api/updatebook/{id}", updateBook).Methods("PUT")
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
	log.Println(r.Form["hello"])
}

// func main(){
// 	fileserver := http.FileServer(http.Dir("./static"))
// 	http.Handle("/", fileserver)
// 	http.HandleFunc("/form", formhandle)
// 	http.HandleFunc("/hello", hellohandle)

// 	log.Println("Listening... on port `8080`")
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

// func hellohandle(w http.ResponseWriter, r *http.Request) {
// 	if(r.URL.Path != "/hello"){
// 		http.NotFound(w, r)
// 		return
// 	}
// 	fmt.Fprintf(w, "Hello, %s!", r.URL.Path)
// }

// func formhandle(w http.ResponseWriter, r *http.Request) {
// 	if(r.URL.Path != "/form"){
// 		http.NotFound(w, r)
// 		return
// 	}
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Fprintf(w, "Hello, %s!", r.FormValue("name"))
// }
