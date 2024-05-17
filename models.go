package dblogger

import (
	"time"
)

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
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	DeploymentID  int64  `db:"deployment_id" json:"deployment_id"`
}

type BuildSum struct {
	ID        int64     `db:"id" json:"id"`
	BuildSum  string    `db:"build_sum" json:"build_sum"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Deployment struct {
	ID        int64  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

type GitRevision struct {
	ID          int64  `db:"id" json:"id"`
	GitRevision string `db:"git_revision" json:"git_revision"`
	CreatedAt   string `db:"created_at" json:"created_at"`
}

type GoVersion struct {
	ID        int64  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Version   string `db:"version" json:"version"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type LogLevel struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Url struct {
	ID        int64  `db:"id" json:"id"`
	Url       string `db:"url" json:"url"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
