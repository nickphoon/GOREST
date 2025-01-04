package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	dbDriver string
	dbUser   string
	dbPass   string
	dbName   string
)

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	dbDriver = os.Getenv("dbDriver")
	dbUser = os.Getenv("dbUser")
	dbPass = os.Getenv("dbPass")
	dbName = os.Getenv("dbName")

	// Ensure required variables are set
	if dbDriver == "" || dbUser == "" || dbPass == "" || dbName == "" {
		log.Fatal("Missing required environment variables")
	}
}

type Article struct {
	ID      int    `json:id`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hits")
}

func TitleExists(title string, id int) bool {
	for _, article := range Articles {
		if article.Title == title && article.ID != id {
			return true
		}
	}
	return false
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", getAllArticlesHandler).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", getArticleHandler).Methods("GET")
	myRouter.HandleFunc("/articles", createArticleHandler).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", updateArticleHandler).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", deleteArticleHandler).Methods("DELETE")
	log.Println("Server listening on http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", myRouter))
}

func main() {

	handleRequests()

}
