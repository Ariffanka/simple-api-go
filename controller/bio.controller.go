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

	data, err := json.MarshalIndent(bio, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Gets(w http.ResponseWriter, r *http.Request) {
	var bios []model.Bio
	rows, err := db.Query("SELECT * FROM bio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var bio model.Bio
		if err := rows.Scan(
			&bio.ID,
			&bio.Username,
			&bio.Email,
		); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		bios = append(bios, bio)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Indentasi JSON sebelum menulisnya ke ResponseWriter
	data, err := json.MarshalIndent(bios, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var bio model.Bio
	json.NewDecoder(r.Body).Decode(&bio)
	_, err := db.Exec("UPDATE bio SET username=$1, email=$2 WHERE id=$3", bio.Username, bio.Email, params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success update"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var bio model.Bio
	json.NewDecoder(r.Body).Decode(&bio)
	_, err := db.Exec("DELETE FROM bio WHERE id=$1", params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success delete"))
}
