package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Call the GetUser function to fetch the user data from the database
	article, err := getArticles(db)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}
func getArticleHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	articleID, _ := strconv.Atoi(idStr)

	// Call the GetUser function to fetch the user data from the database
	article, err := getArticle(db, articleID)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}
func getArticle(db *sql.DB, id int) (*Article, error) {
	query := "SELECT * FROM articles WHERE id = ?"
	row := db.QueryRow(query, id)

	article := &Article{}
	err := row.Scan(&article.ID, &article.Title, &article.Desc, &article.Content)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func getArticles(db *sql.DB) ([]Article, error) {
	query := "SELECT * FROM articles "
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed after usage
	// Slice to hold the articles
	var articles []Article

	// Iterate over the rows
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Desc, &article.Content)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
