package dblogger

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"sync"

	"github.com/rs/zerolog"
	_ "modernc.org/sqlite"
)

// DBTX is the interface for the database/sql.Tx type
// and is used to simplify the queries interface by
// allowing the queries to be run within a transaction.
//
// Example:
//
//	tx, err := db.Begin()
//	if err != nil {
//	    return err
//	}
//	q := data.New(tx)
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// New creates a new instance of the Queries with the provided
// transaction. The provided transaction must be committed and
// rolled back when the instance is no longer needed.
//
// Example:
//
//	        db, err := sql.Open("sqlite", "my.db")
//		if err != nil {
//			log.Fatal(err)
//		}
//		tx, err := db.Begin()
//		if err != nil {
//		    return err
//		}
//		q := data.New(tx)
//		defer q.Close()
func New(db DBTX) *Queries {
	return &Queries{database: db, mutex: sync.Mutex{}}
}

// Queries defines the queries that are run against the database.
// These queries can be executed by the query method on the Queries
// type, or can be used by the *sql.Rows methods.
//
// Example:
//
//	     out, err := queries.CountGitRevisions(ctx context.Context)
//		if err != nil {
//		    return err
//		}
//	     println(out) // number of git revisions
type Queries struct {
	database DBTX
	mutex    sync.Mutex
}

// WithTx returns a new Queries instance that is already
// associated with the provided transaction.
//
// Example:
//
//	tx, err := db.Begin()
//	if err != nil {
//	    return err
//	}
//	q := dblogger.New(tx)
func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		database: tx,
		mutex:    sync.Mutex{},
	}
}

// Write writes the data to the database.
// This makes the Queries type implement the io.Writer interface.
func (q *Queries) Write(p []byte) (n int, err error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	type JsonLog struct {
		Level       string  `json:"level"`
		GitRevision string  `json:"git_revision"`
		Deployment  string  `json:"deployment"`
		GoVersion   string  `json:"go_version"`
		BuildSum    string  `json:"build_sum"`
		Method      string  `json:"method"`
		Url         string  `json:"url"`
		UserAgent   string  `json:"user_agent"`
		ElapsedMs   float64 `json:"elapsed_ms"`
		Time        string  `json:"time"`
		Message     string  `json:"message"`
	}
	var log JsonLog
	err = json.Unmarshal(p, &log)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal log: %w", err)
	}
	err = q.InsertLogEntry(context.Background(), InsertLogEntryParams{
		GoVersionID:   goVersionId,
		BuildSumID:    buildSumId,
		GitRevisionID: gitRevisionId,
		UserAgent:     log.UserAgent,
		Method:        log.Method,
		Url:           log.Url,
		ElapsedNs:     int64(math.Round(log.ElapsedMs * 1000000)),
		DeploymentID:  deploymentId,
		LevelID:       0,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to insert log entry: %w", err)
	}
	return len(p), nil
}

// WriteLevel writes the data to the database with the provided level.
// This makes the Queries type implement the zerolog.LevelWriter interface.
func (q *Queries) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	var levelID int64
	switch level {
	case zerolog.DebugLevel:
		levelID = 1
	case zerolog.InfoLevel:
		levelID = 2
	case zerolog.WarnLevel:
		levelID = 3
	case zerolog.ErrorLevel:
		levelID = 4
	case zerolog.FatalLevel:
		levelID = 5
	case zerolog.PanicLevel:
		levelID = 6
	}
	type JsonLog struct {
		Level       string  `json:"level"`
		GitRevision string  `json:"git_revision"`
		Deployment  string  `json:"deployment"`
		GoVersion   string  `json:"go_version"`
		BuildSum    string  `json:"build_sum"`
		Method      string  `json:"method"`
		Url         string  `json:"url"`
		UserAgent   string  `json:"user_agent"`
		ElapsedMs   float64 `json:"elapsed_ms"`
		Time        string  `json:"time"`
		Message     string  `json:"message"`
	}
	var log JsonLog
	err = json.Unmarshal(p, &log)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal log: %w", err)
	}
	err = q.InsertLogEntry(context.Background(), InsertLogEntryParams{
		GoVersionID:   goVersionId,
		BuildSumID:    buildSumId,
		GitRevisionID: gitRevisionId,
		UserAgent:     log.UserAgent,
		Method:        log.Method,
		Url:           log.Url,
		ElapsedNs:     int64(math.Round(log.ElapsedMs * 1000000)),
		DeploymentID:  deploymentId,
		LevelID:       levelID,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to insert log entry: %w", err)
	}
	return len(p), nil
}
