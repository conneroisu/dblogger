package dblogger

import (
	"time"
)

// ApiLog represents a log entry from the api
// that is stored in the database.
//
// Example:
//
//	{
//	    "id": 1,
//	    "level_id": 1,
//	    "created_at": "2022-01-01T00:00:00Z",
//	    "go_version_id": 1,
//	    "build_sum_id": 1,
//	    "git_revision_id": 1,
//	    "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
//	    "method": "GET",
//	    "url": "https://example.com",
//	    "elapsed_ms": 123,
//	    "status_code": 200,
//	    "deployment_id": 1
//	}
type ApiLog struct {
	ID            int64  `db:"id" json:"id"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	GoVersionID   int64  `db:"go_version_id" json:"go_version_id"`
	BuildSumID    int64  `db:"build_sum_id" json:"build_sum_id"`
	GitRevisionID int64  `db:"git_revision_id" json:"git_revision_id"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedMs     int64  `db:"elapsed_ms" json:"elapsed_ms"`
	StatusCode    int64  `db:"status_code" json:"status_code"`
	DeploymentID  int64  `db:"deployment_id" json:"deployment_id"`
}

// BuildSum represents a build sum that is stored in the database.
//
// Example:
//
//	{
//	    "id": 1,
//	    "build_sum": "1234567890",
//	    "created_at": "2022-01-01T00:00:00Z"
//	}
type BuildSum struct {
	ID        int64     `db:"id" json:"id"`
	BuildSum  string    `db:"build_sum" json:"build_sum"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// Deployment represents a deployment that is stored in the database.
//
// Example:
//
//	{
//	    "id": 1,
//	    "name": "production",
//	    "created_at": "2022-01-01T00:00:00Z",
//	    "updated_at": "2022-01-01T00:00:00Z"
//	}
type Deployment struct {
	ID        int64  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

// GitRevision represents a git revision that is stored in the database.
//
// Example:
//
//	{
//	    "id": 1,
//	    "git_revision": "1234567890",
//	    "created_at": "2022-01-01T00:00:00Z"
//	}
type GitRevision struct {
	ID          int64  `db:"id" json:"id"`
	GitRevision string `db:"git_revision" json:"git_revision"`
	CreatedAt   string `db:"created_at" json:"created_at"`
}

// GoVersion represents a go version that is stored in the database.
//
// Example:
//
//	{
//	    "id": 1,
//	    "name": "go1.17",
//	    "version": "1.17",
//	    "created_at": "2022-01-01T00:00:00Z"
//	}
type GoVersion struct {
	ID        int64  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Version   string `db:"version" json:"version"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

// LogLevel represents a log level that is stored in the database.
//
// Example:
//
//	{
//	    "id": 1,
//	    "name": "info",
//	}
type LogLevel struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// Url represents a url that is stored in the database.
//
// Example:
//
//	{
//	    "id": 1,
//	    "url": "https://example.com",
//	    "created_at": "2022-01-01T00:00:00Z"
//	}
type Url struct {
	ID        int64  `db:"id" json:"id"`
	Url       string `db:"url" json:"url"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
