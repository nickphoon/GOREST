package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func createArticleHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Parse JSON data from the request body
	var article Article
	json.NewDecoder(r.Body).Decode(&article)

	err = createArticle(db, article.Title, article.Desc, article.Content)
	if err != nil {
		http.Error(w, "Failed to create Article", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Article created successfully")
}

func createArticle(db *sql.DB, title string, descriptions string, content string) error {
	// var article Article
	// err := json.NewDecoder(r.Body).Decode(&article)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// if titleExists(article.Title, article.ID) {
	// 	http.Error(w, "Article Title Exists", http.StatusConflict)
	// 	return
	// }

	// article.ID = len(Articles) + 1
	// Articles = append(Articles, article)
	// json.NewEncoder(w).Encode(article)

	query := "INSERT INTO articles (title, descriptions, content) VALUES (?, ?, ?)"
	_, err := db.Exec(query, title, descriptions, content)
	if err != nil {
		return err
	}
	return nil
}
