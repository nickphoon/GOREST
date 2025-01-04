package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}
	// Fetch the article before deleting it
	article, err := getArticle(db, articleID)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	deleteArticle(db, article.ID)

	fmt.Fprintln(w, "Article deleted successfully")

	// Convert the user object to JSON and send it in the response
	w.WriteHeader(http.StatusNoContent)

}
func deleteArticle(db *sql.DB, id int) error {
	query := "DELETE FROM articles WHERE id = ?"
	db.Exec(query, id)

	return nil
}
