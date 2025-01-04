package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func updateArticleHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	articleID, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	// Parse JSON data from the request body
	var updatedArticle Article
	json.NewDecoder(r.Body).Decode(&updatedArticle)
	// Call the GetUser function to fetch the user data from the database
	currentArticle, err := getArticle(db, articleID)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	if currentArticle.Title != updatedArticle.Title {
		exist, err := checkExist(db, updatedArticle.Title, updatedArticle.ID)
		if err != nil {
			http.Error(w, "Database error while checking title", http.StatusInternalServerError)
			return
		}

		if exist {
			http.Error(w, "Article Title Exists", http.StatusConflict)
			return
		}
	}

	err = updateArticle(db, updatedArticle.Title, updatedArticle.Desc, updatedArticle.Content, articleID)
	if err != nil {
		http.Error(w, "Failed to update Article", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Article updated successfully")
}

func checkExist(db *sql.DB, title string, id int) (bool, error) {
	// Check if the new title already exists in another article
	checkQuery := "SELECT COUNT(*) FROM articles WHERE title = ? AND id != ?"
	var count int
	err := db.QueryRow(checkQuery, title, id).Scan(&count)
	if err != nil {
		return false, err
	}
	// If count > 0, the title exists
	return count > 0, nil
}
func updateArticle(db *sql.DB, title string, descriptions string, content string, id int) error {
	query := "UPDATE articles SET title = ?, descriptions = ?, content = ? WHERE id = ?"
	_, err := db.Exec(query, title, descriptions, content, id)
	if err != nil {
		return err
	}
	return nil
}
