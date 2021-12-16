package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var con *sql.DB
var start = time.Now()
var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s search_path=%s sslmode=disable",
	"postgres-net", "5432", "postgres", "root", "web_service", "bookstore")

type Book struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func init() {
	db, err := sql.Open("postgres", DataSourceName)
	if err != nil {
		panic("failed to open db connection")
	}
	con = db
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rows, err := con.Query("SELECT * FROM book")
		checkErr(err)

		books := []Book{}
		defer rows.Close()
		for rows.Next() {
			var r Book
			rows.Scan(&r.Id, &r.Name)
			books = append(books, r)
		}

		json.NewEncoder(rw).Encode(books)
	})

	log.Println("Server is running!")
	fmt.Printf("Startup took aprox %v\n", time.Since(start))
	http.ListenAndServe(":4000", router)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
