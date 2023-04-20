package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}

var db *sql.DB

//Main function
func main() {
	var err error
	db, err = sql.Open("mysql", "<username>:<password>@tcp(<host>:<port>)/<database>")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

  //Register routes
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

//Create Book API function
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	stmt, err := db.Prepare("INSERT INTO books(id, title, author) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(book.ID, book.Title, book.Author)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(book)
}

//Get Book detail API by Id
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var book Book

	row := db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(book)
}

//List all the Books API
func getBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book

	rows, err := db.Query("SELECT id, title, author FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

//Update book detail by id API
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	stmt, err := db.Prepare("UPDATE books SET title=?, author=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(book.Title, book.Author, id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(book)
}

//Delete API Book by Id
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	stmt, err := db.Prepare("DELETE FROM books WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Fprintf(w, "Book with ID = %s has been deleted", id)
  
  response := map[string]string{"message": fmt.Sprintf("Book with ID = %s has been deleted", id)}
	json.NewEncoder(w).Encode(response)
}
