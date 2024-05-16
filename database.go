package dblogger

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/rs/zerolog"
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
	return &Queries{db: db}
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
	db DBTX
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
//	q := data.New(tx)
//	defer q.Close()
func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}

// NewLogsDatabase creates a new logs database queries for an
// logs database. It takes as input a sql database and returns the
// queries struct for a logs database. additionaly, it will execute the
// sql schema for the logs database.
//
// Example:
//
//	db, err := sql.Open("sqlite", "my.db")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//	q := NewLogsDatabase(db)
func NewLogsDatabase(db *sql.DB) (*Queries, error) {
	if _, err := db.Exec(SQLSchema); err != nil {
		return nil, err
	}
	q := New(db)
	return q, nil
}

// Write writes the data to the database.
// This makes the Queries type implement the io.Writer interface.
func (q Queries) Write(p []byte) (n int, err error) {
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
		return 0, err
	}
	err = q.InsertLogEntry(context.Background(), InsertLogEntryParams{
		GoVersionID:   ,
		BuildSumID:    apiLog.BuildSumID,
		GitRevisionID: apiLog.GitRevisionID,
		UserAgent:     apiLog.UserAgent,
		Method:        apiLog.Method,
		Url:           apiLog.Url,
		ElapsedMs:     apiLog.ElapsedMs,
		StatusCode:    apiLog.StatusCode,
		DeploymentID:  apiLog.DeploymentID,
		LevelID:       apiLog.LevelID,
	})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// WriteLevel writes the data to the database with the provided level.
// This makes the Queries type implement the zerolog.LevelWriter interface.
func (q Queries) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	var apiLog ApiLog
	err = json.Unmarshal(p, &apiLog)
	if err != nil {
		return 0, err
	}
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
	// insert into database
	err = q.InsertLogEntry(context.Background(), InsertLogEntryParams{
		GoVersionID:   apiLog.GoVersionID,
		BuildSumID:    apiLog.BuildSumID,
		GitRevisionID: apiLog.GitRevisionID,
		UserAgent:     apiLog.UserAgent,
		Method:        apiLog.Method,
		Url:           apiLog.Url,
		ElapsedMs:     apiLog.ElapsedMs,
		StatusCode:    apiLog.StatusCode,
		DeploymentID:  apiLog.DeploymentID,
		LevelID:       levelID,
	})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
