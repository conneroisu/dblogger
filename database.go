package dblogger

import (
	"context"
	"database/sql"
	"sync"

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
//
// Deprecated: use NewLogsDatabase instead
func New(db DBTX) *Queries {
	return &Queries{db: db, mutex: sync.Mutex{}}
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
	db    DBTX
	mutex sync.Mutex
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
		db:    tx,
		mutex: sync.Mutex{},
	}
}
