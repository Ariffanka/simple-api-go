package controller

import (
	config "api-go/database"
	"api-go/model"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	db = config.ConnectDB()
}

func Create(w http.ResponseWriter, r *http.Request) {
	var bio model.Bio
	json.NewDecoder(r.Body).Decode(&bio)
	_, err := db.Exec("INSERT INTO bio(username,email) VALUES($1,$2)", bio.Username, bio.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success"))
}

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var bio model.Bio
	err := db.QueryRow("SELECT id, username, email FROM bio WHERE id= $1", params["id"]).Scan(&bio.ID, &bio.Username, &bio.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer db.Close()

	json.NewEncoder(w).Encode(bio)
}
