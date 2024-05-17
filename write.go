package dblogger

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	"github.com/rs/zerolog"
)

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
	err = q.InsertLogEntry(context.Background(), &InsertLogEntryParams{
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
func (q *Queries) WriteLevel(
	level zerolog.Level,
	p []byte,
) (n int, err error) {
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
	err = q.InsertLogEntry(context.Background(), &InsertLogEntryParams{
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
