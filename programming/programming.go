package programming

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type LanguagesHandler struct {
	db *sql.DB
}

func NewLanguageHandler(db *sql.DB) LanguagesHandler {
	return LanguagesHandler{db: db}
}

func (handler LanguagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	type pgLanguage struct {
		Name     string `json:"name"`
		ImageURL string `json:"imageUrl"`
	}

	var pgLanguages []pgLanguage

	rows, err := handler.db.Query("select name, imageUrl from languages")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var imageUrl string
		err = rows.Scan(&name, &imageUrl)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		pgLanguages = append(pgLanguages, pgLanguage{Name: name, ImageURL: imageUrl})
	}
	err = rows.Err()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(pgLanguages); err != nil {
		fmt.Fprintln(w, err)
		return
	}
}
