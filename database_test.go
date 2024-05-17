package dblogger

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "modernc.org/sqlite"

	"github.com/rs/zerolog"
	"github.com/ztrue/tracerr"
)

// TestSchema tests the database schema to see that a database can be created
// with the schema.
func TestSchema(t *testing.T) {
	t.Parallel()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to open database: %w", err))
	}
	_, err = db.Exec(SQLSchema)
	if err != nil {
		t.Fatal(fmt.Errorf("failed to create database: %w", err))
	}
}

// TestQueries tests the database queries by writing a log message to the database
func TestQueries(t *testing.T) {
	t.Parallel()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to open database: %w", err))
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
	bytesVers, err := os.ReadFile("./testdata/valid_log.json")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to read test data: %w", err))
	}
	n, err := q.Write(bytesVers)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("t: %v\n", n)
}

// TestTx tests the database transactions
func TestTx(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to open database: %w", err))
	}
	getenv := func(key string) (string, error) {
		switch key {
		case "DEPLOYMENT":
			return "staging", nil
		default:
			return "", fmt.Errorf("invalid key: %s", key)
		}
	}
	qdb, err := NewLogsDatabase(getenv, db)
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(fmt.Errorf("failed to begin transaction: %w", err))
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			t.Fatal(fmt.Errorf("failed to rollback transaction: %w", err))
		}
	}()
	qtx := qdb.WithTx(tx)
	if err != nil {
		t.Fatal(err)
	}
	_, err = qtx.InsertBuildSumReturningID(context.Background(), &InsertBuildSumReturningIDParams{
		BuildSum: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = qtx.InsertDeploymentReturningID(context.Background(), &InsertDeploymentReturningIDParams{
		Name: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = qtx.InsertGitRevisionReturningID(context.Background(), &InsertGitRevisionReturningIDParams{
		GitRevision: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

}

// TestQueries tests the database queries by writing a log message to the database
func TestQueriesParallel(t *testing.T) {
	t.Parallel()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to open database: %w", err))
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
	// bytesVers := []byte(``)
	// read from ./testdata/valid_log.json
	bytesVers, err := os.ReadFile("./testdata/valid_log.json")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to read test data: %w", err))
	}
	t.Run("write1", func(t *testing.T) {
		t.Parallel()
		n, err := q.WriteLevel(zerolog.DebugLevel, bytesVers)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("t: %v\n", n)
	})
	t.Run("write2", func(t *testing.T) {
		t.Parallel()
		n, err := q.Write(bytesVers)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("t: %v\n", n)
	})
	t.Run("write3", func(t *testing.T) {
		t.Parallel()
		n, err := q.Write(bytesVers)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("t: %v\n", n)
	})
	t.Run("write4", func(t *testing.T) {
		t.Parallel()
		n, err := q.Write(bytesVers)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("t: %v\n", n)
	})
	t.Run("write5", func(t *testing.T) {
		t.Parallel()
		n, err := q.Write(bytesVers)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("t: %v\n", n)
	})
	t.Run("write6", func(t *testing.T) {
		t.Parallel()
		n, err := q.Write(bytesVers)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("t: %v\n", n)
	})
	n, err := q.Write(bytesVers)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("t: %v\n", n)
}

// TestQueries tests the database queries by writing a log message to the database
func TestQueriesEmptyUrl(t *testing.T) {
	t.Parallel()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to open database: %w", err))
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
	// bytesVers := []byte(``)
	// read from ./testdata/valid_log.json
	bytesVers, err := os.ReadFile("./testdata/valid_log_empty_url.json")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to read test data: %w", err))
	}
	n, err := q.Write(bytesVers)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("t: %v\n", n)
}

// TestQueries tests the database queries by writing a log message to the database
func TestQueriesNoUrl(t *testing.T) {
	t.Parallel()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to open database: %w", err))
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
	// bytesVers := []byte(``)
	// read from ./testdata/valid_log.json
	bytesVers, err := os.ReadFile("./testdata/valid_log_no_url.json")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to read test data: %w", err))
	}
	n, err := q.Write(bytesVers)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("t: %v\n", n)
}

// TestQueries tests the database queries by writing a log message to the database
func TestQueriesInvalid(t *testing.T) {
	t.Parallel()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to open database: %w", err))
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
	bytesVers, err := os.ReadFile("./testdata/invalid_log.json")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to read test data: %w", err))
	}
	_, err = q.Write(bytesVers)
	if err == nil {
		t.Fatal(err)
	}
}
