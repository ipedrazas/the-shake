package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func loadBooks() BookList {
	file, e := ioutil.ReadFile("./books.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))
	var bookList BookList

	json.Unmarshal(file, &bookList)
	return bookList

}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love Shakespeare!")
}

var books BookList

func main() {
	books = loadBooks()
	http.HandleFunc("/", handler)
	http.HandleFunc("/books", booksHandler)
	http.ListenAndServe(":8080", nil)
	os.Exit(0)
}
