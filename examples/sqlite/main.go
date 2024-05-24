package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/conneroisu/dblogger"
)

func main() {
	router := http.NewServeMux()
	srv := &http.Server{Addr: "8080", Handler: router}
	getenv := func(key string) (string, error) {
		switch key {
		case "DEPLOYMENT":
			return "staging", nil
		default:
			return "", fmt.Errorf("invalid key: %s", key)
		}
	}
	db, err := sql.Open("sqlite", "my.db")
	if err != nil {
		panic(err)
	}
	q, err := dblogger.NewLogsDatabase(getenv, db)
	if err != nil {
		panic(err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})
	// register the routes
	router.HandleFunc("/", requestLogger(q, handler))

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
