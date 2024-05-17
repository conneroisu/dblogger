package dblogger

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"runtime/debug"
	"strings"
)

//go:embed data/combined/schema.sql
var SQLSchema string

//go:embed data/combined/queries.sql
var SQLQueries string

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
func NewLogsDatabase(getenv func(string) (string, error), db *sql.DB) (*Queries, error) {
	ctx := context.Background()
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, v := range buildInfo.Settings {
			if v.Key == "vcs.revision" {
				gitRevision = v.Value
				break
			}
		}
	}
	goVersion = buildInfo.GoVersion
	buildSum = buildInfo.Main.Sum
	_, err := db.Exec(SQLSchema)
	q := New(db)
	if err != nil {
		if err != sql.ErrNoRows && !strings.Contains(err.Error(), "already exists") {
			return nil, fmt.Errorf("failed to create database: %w", err)
		}
	}
	deployment, err = getenv("DEPLOYMENT")
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment: %w", err)
	}
	buildSumId, err = q.InsertBuildSumReturningID(ctx, InsertBuildSumReturningIDParams{
		BuildSum: buildSum,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return q, nil
		}
		return nil, fmt.Errorf("failed to insert build sum: %w", err)
	}
	goVersionId, err = q.InsertGoVersionReturningID(ctx, InsertGoVersionReturningIDParams{
		Name:    goVersion,
		Version: goVersion,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return q, nil
		}
		return nil, fmt.Errorf("failed to insert go version: %w", err)
	}
	gitRevisionId, err = q.InsertGitRevisionReturningID(ctx, InsertGitRevisionReturningIDParams{
		GitRevision: gitRevision,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return q, nil
		}
		return nil, fmt.Errorf("failed to insert git revision: %w", err)
	}
	deploymentId, err = q.InsertDeploymentReturningID(ctx, InsertDeploymentReturningIDParams{
		Name: deployment,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return q, nil
		}
		return nil, fmt.Errorf("failed to insert deployment: %w", err)
	}
	return q, nil
}
