package dblogger

import (
	"database/sql"
	"fmt"
	"testing"

	_ "modernc.org/sqlite"

	"github.com/ztrue/tracerr"
)

// TestQueries tests the database queries by writing a log message to the database
func TestQueries(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	getenv := func(key string) (string, error) {
		switch key {
		case "DEPLOYMENT":
			return "staging", nil
		default:
			return "", fmt.Errorf("invalid key: %s", key)
		}
	}
	q, err := NewLogsDatabase(getenv, db)
	if err != nil {
		tracerr.PrintSourceColor(err)
	}
	bytesVers := []byte(`{"level":"info","git_revision":"e8ff7107e13808170243634f7e4256245c3236be","deployment":"","go_version":"go1.22.0","build_sum":"","method":"GET","url":"/home","user_agent":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36","elapsed_ms":0.135385,"time":"2024-05-16T06:45:36.630788595-05:00","message":"incoming request to //home"}`)
	n, err := q.Write(bytesVers)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("t: %v\n", n)
}

func TestSchema(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(SQLSchema)
	if err != nil {
		t.Fatal(err)
	}
}
