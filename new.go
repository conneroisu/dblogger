package dblogger

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"runtime/debug"
	"sync"
)

//go:embed data/combined/schema.sql
var SQLSchema string

//go:embed data/combined/queries.sql
var SQLQueries string

// gitRevision is the git revision of the application
var gitRevision string

// buildInfo is the build information of the application
var buildInfo debug.BuildInfo

// gitRevision is the git revision of the application
var buildSum string

// buildSumId is the id of the build sum
var buildSumId int64

// goVersion is the version of the Go compiler used to build the application
var goVersion string

// goVersionId is the id of the go version
var goVersionId int64

// gitRevisionId is the id of the git revision
var gitRevisionId int64

// once is used to ensure that the package is only initialized once
var once sync.Once

var deployment string
var deploymentId int64

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
func NewLogsDatabase(getenv func(string) string, db *sql.DB) (*Queries, error) {
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
		if err != sql.ErrNoRows && !isAlreadyExistsError(err) {
			return nil, fmt.Errorf("failed to create database: %w", err)
		}
	}
	bs, err := q.InsertBuildSum(ctx, InsertBuildSumParams{
		BuildSum: buildSum,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return q, nil
		}
		return nil, err
	}
	buildSumId = bs.ID
	gv, err := q.InsertGoVersion(ctx, InsertGoVersionParams{
		Name:    goVersion,
		Version: goVersion,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return q, nil
		}
		return nil, err
	}
	goVersionId = gv.ID
	gr, err := q.InsertGitRevision(ctx, InsertGitRevisionParams{
		GitRevision: gitRevision,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return q, nil
		}
		return nil, err
	}
	gitRevisionId = gr.ID
	if deployment := getenv("DEPLOYMENT"); deployment == "" {
		return nil, fmt.Errorf("deployment is not specified")
	}
	dp, err := q.InsertDeployment(ctx, InsertDeploymentParams{
		Name: deployment,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return q, nil
		}
		return nil, err
	}
	deploymentId = dp.ID
	return q, nil
}
