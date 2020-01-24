package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
}

var db *Database

type Config struct {
	DatabaseFile string
}

func InitConfig() (*Config, error) {
	config := &Config{
		DatabaseFile: "live.db",
	}
	if config.DatabaseFile == "" {
		return nil, fmt.Errorf("DatabaseFile must be set")
	}
	return config, nil
}

func NewDB(config *Config) (*Database, error) {
	db, err := sql.Open("sqlite3", config.DatabaseFile)
	if err != nil {
		return nil, fmt.Errorf("error %s: unable to connect to database", err)
	}
	return &Database{db}, nil
}

func (db *Database) CheckKey(key string) (bool, error) {
	stmt, err := db.Prepare("SELECT count(1) FROM publishers WHERE key=?")
	if err != nil {
		return false, fmt.Errorf("error %s: unable to prepare SQL statement", err)
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(key).Scan(&count)
	if err != nil || count == 0 {
		return false, err
	}

	return true, nil
}

func authorize(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	key := r.Form.Get("name")

	if key == "" {
		log.Print(err)
		w.WriteHeader(403)
		w.Write([]byte(`{"message": "Stream not authorized. Check your URL/credentials"}`))
		return
	}

	found, err := db.CheckKey(key)
	if err != nil {
		log.Print(err)
		w.WriteHeader(403)
		w.Write([]byte(`{"message": "Stream not authorized. Check your URL/credentials"}`))
		return
	}

	if found {
		w.WriteHeader(200)
		w.Write([]byte(`{"message": "Stream authorized"}`))
		return
	}
	w.WriteHeader(403)
	w.Write([]byte(`{"message": "Stream not authorized. Check your URL/credentials"}`))
}

func main() {
	config, err := InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err = NewDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/auth", authorize).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", router))
}
