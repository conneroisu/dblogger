package dblogger

import (
	"context"
	"time"
)

const countBuildSums = `-- name: CountBuildSums :one
SELECT
    COUNT(*)
FROM
    build_sums
`

// CountBuildSums
//
//	SELECT
//	    COUNT(*)
//	FROM
//	    build_sums
func (q *Queries) CountBuildSums(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countBuildSums)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countDeployments = `-- name: CountDeployments :one
SELECT
    COUNT(*)
FROM
    deployments
`

// CountDeployments
//
//	SELECT
//	    COUNT(*)
//	FROM
//	    deployments
func (q *Queries) CountDeployments(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countDeployments)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countGitRevisions = `-- name: CountGitRevisions :one
SELECT
    COUNT(*)
FROM
    git_revisions
`

// CountGitRevisions
//
//	SELECT
//	    COUNT(*)
//	FROM
//	    git_revisions
func (q *Queries) CountGitRevisions(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countGitRevisions)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countGoVersions = `-- name: CountGoVersions :one
SELECT
    COUNT(*)
FROM
    go_versions
`

// CountGoVersions
//
//	SELECT
//	    COUNT(*)
//	FROM
//	    go_versions
func (q *Queries) CountGoVersions(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countGoVersions)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countLogLevels = `-- name: CountLogLevels :one
SELECT
    COUNT(*)
FROM
    log_levels
`

// CountLogLevels
//
//	SELECT
//	    COUNT(*)
//	FROM
//	    log_levels
func (q *Queries) CountLogLevels(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countLogLevels)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countLogs = `-- name: CountLogs :one
SELECT
    COUNT(*)
FROM
    api_logs
`

// CountLogs
//
//	SELECT
//	    COUNT(*)
//	FROM
//	    api_logs
func (q *Queries) CountLogs(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countLogs)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countURLs = `-- name: CountURLs :one
SELECT
    COUNT(*)
FROM
    urls
`

// CountURLs
//
//	SELECT
//	    COUNT(*)
//	FROM
//	    urls
func (q *Queries) CountURLs(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countURLs)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteBuildSumByID = `-- name: DeleteBuildSumByID :exec
DELETE FROM
    build_sums
WHERE
    id = ?
`

type DeleteBuildSumByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// DeleteBuildSumByID
//
//	DELETE FROM
//	    build_sums
//	WHERE
//	    id = ?
func (q *Queries) DeleteBuildSumByID(ctx context.Context, arg DeleteBuildSumByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteBuildSumByID, arg.ID)
	return err
}

const deleteDeploymentByID = `-- name: DeleteDeploymentByID :exec
DELETE FROM
    deployments
WHERE
    id = ?
`

type DeleteDeploymentByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// DeleteDeploymentByID
//
//	DELETE FROM
//	    deployments
//	WHERE
//	    id = ?
func (q *Queries) DeleteDeploymentByID(ctx context.Context, arg DeleteDeploymentByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteDeploymentByID, arg.ID)
	return err
}

const deleteGitRevisionByID = `-- name: DeleteGitRevisionByID :exec
DELETE FROM
    git_revisions
WHERE
    id = ?
`

type DeleteGitRevisionByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// DeleteGitRevisionByID
//
//	DELETE FROM
//	    git_revisions
//	WHERE
//	    id = ?
func (q *Queries) DeleteGitRevisionByID(ctx context.Context, arg DeleteGitRevisionByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteGitRevisionByID, arg.ID)
	return err
}

const deleteGoVersionByID = `-- name: DeleteGoVersionByID :exec
DELETE FROM
    go_versions
WHERE
    id = 1
`

// DeleteGoVersionByID
//
//	DELETE FROM
//	    go_versions
//	WHERE
//	    id = 1
func (q *Queries) DeleteGoVersionByID(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteGoVersionByID)
	return err
}

const deleteLogByID = `-- name: DeleteLogByID :exec
DELETE FROM
    api_logs
WHERE
    id = ?
`

type DeleteLogByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// DeleteLogByID
//
//	DELETE FROM
//	    api_logs
//	WHERE
//	    id = ?
func (q *Queries) DeleteLogByID(ctx context.Context, arg DeleteLogByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteLogByID, arg.ID)
	return err
}

const deleteLogLevelByID = `-- name: DeleteLogLevelByID :exec
DELETE FROM
    log_levels
WHERE
    id = ?
`

type DeleteLogLevelByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// DeleteLogLevelByID
//
//	DELETE FROM
//	    log_levels
//	WHERE
//	    id = ?
func (q *Queries) DeleteLogLevelByID(ctx context.Context, arg DeleteLogLevelByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteLogLevelByID, arg.ID)
	return err
}

const deleteURLByID = `-- name: DeleteURLByID :exec
DELETE FROM
    urls
WHERE
    id = ?
`

type DeleteURLByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// DeleteURLByID
//
//	DELETE FROM
//	    urls
//	WHERE
//	    id = ?
func (q *Queries) DeleteURLByID(ctx context.Context, arg DeleteURLByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteURLByID, arg.ID)
	return err
}

const getBuildSumByID = `-- name: GetBuildSumByID :one
SELECT
    id, build_sum, created_at
FROM
    build_sums
WHERE
    id = ?
`

type GetBuildSumByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// GetBuildSumByID
//
//	SELECT
//	    id, build_sum, created_at
//	FROM
//	    build_sums
//	WHERE
//	    id = ?
func (q *Queries) GetBuildSumByID(ctx context.Context, arg GetBuildSumByIDParams) (BuildSum, error) {
	row := q.db.QueryRowContext(ctx, getBuildSumByID, arg.ID)
	var i BuildSum
	err := row.Scan(&i.ID, &i.BuildSum, &i.CreatedAt)
	return i, err
}

const getBuildSumsByDate = `-- name: GetBuildSumsByDate :many
SELECT
    id, build_sum, created_at
FROM
    build_sums
WHERE
    created_at >= ?
`

type GetBuildSumsByDateParams struct {
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// GetBuildSumsByDate
//
//	SELECT
//	    id, build_sum, created_at
//	FROM
//	    build_sums
//	WHERE
//	    created_at >= ?
func (q *Queries) GetBuildSumsByDate(ctx context.Context, arg GetBuildSumsByDateParams) ([]BuildSum, error) {
	rows, err := q.db.QueryContext(ctx, getBuildSumsByDate, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BuildSum
	for rows.Next() {
		var i BuildSum
		if err := rows.Scan(&i.ID, &i.BuildSum, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBuildSumsBySubstring = `-- name: GetBuildSumsBySubstring :many
SELECT
    id, build_sum, created_at
FROM
    build_sums
WHERE
    build_sum LIKE ?
`

type GetBuildSumsBySubstringParams struct {
	BuildSum string `db:"build_sum" json:"build_sum"`
}

// GetBuildSumsBySubstring
//
//	SELECT
//	    id, build_sum, created_at
//	FROM
//	    build_sums
//	WHERE
//	    build_sum LIKE ?
func (q *Queries) GetBuildSumsBySubstring(ctx context.Context, arg GetBuildSumsBySubstringParams) ([]BuildSum, error) {
	rows, err := q.db.QueryContext(ctx, getBuildSumsBySubstring, arg.BuildSum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BuildSum
	for rows.Next() {
		var i BuildSum
		if err := rows.Scan(&i.ID, &i.BuildSum, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDeploymentByID = `-- name: GetDeploymentByID :one
SELECT
    id, name, created_at, updated_at
FROM
    deployments
WHERE
    id = ?
`

type GetDeploymentByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// GetDeploymentByID
//
//	SELECT
//	    id, name, created_at, updated_at
//	FROM
//	    deployments
//	WHERE
//	    id = ?
func (q *Queries) GetDeploymentByID(ctx context.Context, arg GetDeploymentByIDParams) (Deployment, error) {
	row := q.db.QueryRowContext(ctx, getDeploymentByID, arg.ID)
	var i Deployment
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDeploymentsByDate = `-- name: GetDeploymentsByDate :many
SELECT
    id, name, created_at, updated_at
FROM
    deployments
WHERE
    created_at >= ?
`

type GetDeploymentsByDateParams struct {
	CreatedAt string `db:"created_at" json:"created_at"`
}

// GetDeploymentsByDate
//
//	SELECT
//	    id, name, created_at, updated_at
//	FROM
//	    deployments
//	WHERE
//	    created_at >= ?
func (q *Queries) GetDeploymentsByDate(ctx context.Context, arg GetDeploymentsByDateParams) ([]Deployment, error) {
	rows, err := q.db.QueryContext(ctx, getDeploymentsByDate, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Deployment
	for rows.Next() {
		var i Deployment
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDeploymentsByDateRange = `-- name: GetDeploymentsByDateRange :many
SELECT
    id, name, created_at, updated_at
FROM
    deployments
WHERE
    created_at BETWEEN ? AND ?
`

type GetDeploymentsByDateRangeParams struct {
	FromCreatedAt string `db:"from_created_at" json:"from_created_at"`
	ToCreatedAt   string `db:"to_created_at" json:"to_created_at"`
}

// GetDeploymentsByDateRange
//
//	SELECT
//	    id, name, created_at, updated_at
//	FROM
//	    deployments
//	WHERE
//	    created_at BETWEEN ? AND ?
func (q *Queries) GetDeploymentsByDateRange(ctx context.Context, arg GetDeploymentsByDateRangeParams) ([]Deployment, error) {
	rows, err := q.db.QueryContext(ctx, getDeploymentsByDateRange, arg.FromCreatedAt, arg.ToCreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Deployment
	for rows.Next() {
		var i Deployment
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDeploymentsBySubstring = `-- name: GetDeploymentsBySubstring :many
SELECT
    id, name, created_at, updated_at
FROM
    deployments
WHERE
    name LIKE ?
`

type GetDeploymentsBySubstringParams struct {
	Name string `db:"name" json:"name"`
}

// GetDeploymentsBySubstring
//
//	SELECT
//	    id, name, created_at, updated_at
//	FROM
//	    deployments
//	WHERE
//	    name LIKE ?
func (q *Queries) GetDeploymentsBySubstring(ctx context.Context, arg GetDeploymentsBySubstringParams) ([]Deployment, error) {
	rows, err := q.db.QueryContext(ctx, getDeploymentsBySubstring, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Deployment
	for rows.Next() {
		var i Deployment
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGitRevisionByID = `-- name: GetGitRevisionByID :one
SELECT
    id, git_revision, created_at
FROM
    git_revisions
WHERE
    id = ?
`

type GetGitRevisionByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// GetGitRevisionByID
//
//	SELECT
//	    id, git_revision, created_at
//	FROM
//	    git_revisions
//	WHERE
//	    id = ?
func (q *Queries) GetGitRevisionByID(ctx context.Context, arg GetGitRevisionByIDParams) (GitRevision, error) {
	row := q.db.QueryRowContext(ctx, getGitRevisionByID, arg.ID)
	var i GitRevision
	err := row.Scan(&i.ID, &i.GitRevision, &i.CreatedAt)
	return i, err
}

const getGitRevisionsByDate = `-- name: GetGitRevisionsByDate :many
SELECT
    id, git_revision, created_at
FROM
    git_revisions
WHERE
    created_at >= ?
`

type GetGitRevisionsByDateParams struct {
	CreatedAt string `db:"created_at" json:"created_at"`
}

// GetGitRevisionsByDate
//
//	SELECT
//	    id, git_revision, created_at
//	FROM
//	    git_revisions
//	WHERE
//	    created_at >= ?
func (q *Queries) GetGitRevisionsByDate(ctx context.Context, arg GetGitRevisionsByDateParams) ([]GitRevision, error) {
	rows, err := q.db.QueryContext(ctx, getGitRevisionsByDate, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GitRevision
	for rows.Next() {
		var i GitRevision
		if err := rows.Scan(&i.ID, &i.GitRevision, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGitRevisionsByName = `-- name: GetGitRevisionsByName :one
SELECT
    id, git_revision, created_at
FROM
    git_revisions
WHERE
    git_revision = ?
`

type GetGitRevisionsByNameParams struct {
	GitRevision string `db:"git_revision" json:"git_revision"`
}

// GetGitRevisionsByName
//
//	SELECT
//	    id, git_revision, created_at
//	FROM
//	    git_revisions
//	WHERE
//	    git_revision = ?
func (q *Queries) GetGitRevisionsByName(ctx context.Context, arg GetGitRevisionsByNameParams) (GitRevision, error) {
	row := q.db.QueryRowContext(ctx, getGitRevisionsByName, arg.GitRevision)
	var i GitRevision
	err := row.Scan(&i.ID, &i.GitRevision, &i.CreatedAt)
	return i, err
}

const getGoVersionByID = `-- name: GetGoVersionByID :one
SELECT
    id, name, version, created_at
FROM
    go_versions
WHERE
    id = ?
`

type GetGoVersionByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// GetGoVersionByID
//
//	SELECT
//	    id, name, version, created_at
//	FROM
//	    go_versions
//	WHERE
//	    id = ?
func (q *Queries) GetGoVersionByID(ctx context.Context, arg GetGoVersionByIDParams) (GoVersion, error) {
	row := q.db.QueryRowContext(ctx, getGoVersionByID, arg.ID)
	var i GoVersion
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Version,
		&i.CreatedAt,
	)
	return i, err
}

const getGoVersionIdByName = `-- name: GetGoVersionIdByName :one
SELECT
    id
FROM
    go_versions
WHERE
    name = ?
`

type GetGoVersionIdByNameParams struct {
	Name string `db:"name" json:"name"`
}

// GetGoVersionIdByName
//
//	SELECT
//	    id
//	FROM
//	    go_versions
//	WHERE
//	    name = ?
func (q *Queries) GetGoVersionIdByName(ctx context.Context, arg GetGoVersionIdByNameParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getGoVersionIdByName, arg.Name)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getGoVersionsByDate = `-- name: GetGoVersionsByDate :many
SELECT
    id, name, version, created_at
FROM
    go_versions
WHERE
    created_at >= ?
`

type GetGoVersionsByDateParams struct {
	CreatedAt string `db:"created_at" json:"created_at"`
}

// GetGoVersionsByDate
//
//	SELECT
//	    id, name, version, created_at
//	FROM
//	    go_versions
//	WHERE
//	    created_at >= ?
func (q *Queries) GetGoVersionsByDate(ctx context.Context, arg GetGoVersionsByDateParams) ([]GoVersion, error) {
	rows, err := q.db.QueryContext(ctx, getGoVersionsByDate, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoVersion
	for rows.Next() {
		var i GoVersion
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Version,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGoVersionsBySubstring = `-- name: GetGoVersionsBySubstring :many
SELECT
    id, name, version, created_at
FROM
    go_versions
WHERE
    name LIKE ?
    OR version LIKE ?
`

type GetGoVersionsBySubstringParams struct {
	Name    string `db:"name" json:"name"`
	Version string `db:"version" json:"version"`
}

// GetGoVersionsBySubstring
//
//	SELECT
//	    id, name, version, created_at
//	FROM
//	    go_versions
//	WHERE
//	    name LIKE ?
//	    OR version LIKE ?
func (q *Queries) GetGoVersionsBySubstring(ctx context.Context, arg GetGoVersionsBySubstringParams) ([]GoVersion, error) {
	rows, err := q.db.QueryContext(ctx, getGoVersionsBySubstring, arg.Name, arg.Version)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoVersion
	for rows.Next() {
		var i GoVersion
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Version,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogByID = `-- name: GetLogByID :one
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    id = ?
`

type GetLogByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// GetLogByID
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    id = ?
func (q *Queries) GetLogByID(ctx context.Context, arg GetLogByIDParams) (ApiLog, error) {
	row := q.db.QueryRowContext(ctx, getLogByID, arg.ID)
	var i ApiLog
	err := row.Scan(
		&i.ID,
		&i.LevelID,
		&i.CreatedAt,
		&i.GoVersionID,
		&i.BuildSumID,
		&i.GitRevisionID,
		&i.UserAgent,
		&i.Method,
		&i.Url,
		&i.ElapsedNs,
		&i.DeploymentID,
	)
	return i, err
}

const getLogLevelsBySubstring = `-- name: GetLogLevelsBySubstring :many
SELECT
    id, name
FROM
    log_levels
WHERE
    name LIKE ?
`

type GetLogLevelsBySubstringParams struct {
	Name string `db:"name" json:"name"`
}

// GetLogLevelsBySubstring
//
//	SELECT
//	    id, name
//	FROM
//	    log_levels
//	WHERE
//	    name LIKE ?
func (q *Queries) GetLogLevelsBySubstring(ctx context.Context, arg GetLogLevelsBySubstringParams) ([]LogLevel, error) {
	rows, err := q.db.QueryContext(ctx, getLogLevelsBySubstring, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LogLevel
	for rows.Next() {
		var i LogLevel
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsByBuildSumID = `-- name: GetLogsByBuildSumID :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    build_sum_id = ?
`

type GetLogsByBuildSumIDParams struct {
	BuildSumID int64 `db:"build_sum_id" json:"build_sum_id"`
}

// GetLogsByBuildSumID
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    build_sum_id = ?
func (q *Queries) GetLogsByBuildSumID(ctx context.Context, arg GetLogsByBuildSumIDParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, getLogsByBuildSumID, arg.BuildSumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsByDate = `-- name: GetLogsByDate :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    created_at >= ?
`

type GetLogsByDateParams struct {
	CreatedAt string `db:"created_at" json:"created_at"`
}

// GetLogsByDate
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    created_at >= ?
func (q *Queries) GetLogsByDate(ctx context.Context, arg GetLogsByDateParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, getLogsByDate, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsByDateRange = `-- name: GetLogsByDateRange :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    created_at BETWEEN ? AND ?
`

type GetLogsByDateRangeParams struct {
	FromCreatedAt string `db:"from_created_at" json:"from_created_at"`
	ToCreatedAt   string `db:"to_created_at" json:"to_created_at"`
}

// GetLogsByDateRange
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    created_at BETWEEN ? AND ?
func (q *Queries) GetLogsByDateRange(ctx context.Context, arg GetLogsByDateRangeParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, getLogsByDateRange, arg.FromCreatedAt, arg.ToCreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsByElapsedRange = `-- name: GetLogsByElapsedRange :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    elapsed_ns BETWEEN ? AND ?
`

type GetLogsByElapsedRangeParams struct {
	FromElapsedNs int64 `db:"from_elapsed_ns" json:"from_elapsed_ns"`
	ToElapsedNs   int64 `db:"to_elapsed_ns" json:"to_elapsed_ns"`
}

// GetLogsByElapsedRange
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    elapsed_ns BETWEEN ? AND ?
func (q *Queries) GetLogsByElapsedRange(ctx context.Context, arg GetLogsByElapsedRangeParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, getLogsByElapsedRange, arg.FromElapsedNs, arg.ToElapsedNs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsByGitRevisionID = `-- name: GetLogsByGitRevisionID :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    git_revision_id = ?
`

type GetLogsByGitRevisionIDParams struct {
	GitRevisionID int64 `db:"git_revision_id" json:"git_revision_id"`
}

// GetLogsByGitRevisionID
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    git_revision_id = ?
func (q *Queries) GetLogsByGitRevisionID(ctx context.Context, arg GetLogsByGitRevisionIDParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, getLogsByGitRevisionID, arg.GitRevisionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsByGoVersionID = `-- name: GetLogsByGoVersionID :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    go_version_id = ?
`

type GetLogsByGoVersionIDParams struct {
	GoVersionID int64 `db:"go_version_id" json:"go_version_id"`
}

// GetLogsByGoVersionID
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    go_version_id = ?
func (q *Queries) GetLogsByGoVersionID(ctx context.Context, arg GetLogsByGoVersionIDParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, getLogsByGoVersionID, arg.GoVersionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsByURLSubstring = `-- name: GetLogsByURLSubstring :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    url LIKE ?
`

type GetLogsByURLSubstringParams struct {
	Url string `db:"url" json:"url"`
}

// GetLogsByURLSubstring
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    url LIKE ?
func (q *Queries) GetLogsByURLSubstring(ctx context.Context, arg GetLogsByURLSubstringParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, getLogsByURLSubstring, arg.Url)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getURLByID = `-- name: GetURLByID :one
SELECT
    id, url, created_at
FROM
    urls
WHERE
    id = ?
`

type GetURLByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

// GetURLByID
//
//	SELECT
//	    id, url, created_at
//	FROM
//	    urls
//	WHERE
//	    id = ?
func (q *Queries) GetURLByID(ctx context.Context, arg GetURLByIDParams) (Url, error) {
	row := q.db.QueryRowContext(ctx, getURLByID, arg.ID)
	var i Url
	err := row.Scan(&i.ID, &i.Url, &i.CreatedAt)
	return i, err
}

const getURLsByDate = `-- name: GetURLsByDate :many
SELECT
    id, url, created_at
FROM
    urls
WHERE
    created_at >= ?
`

type GetURLsByDateParams struct {
	CreatedAt string `db:"created_at" json:"created_at"`
}

// GetURLsByDate
//
//	SELECT
//	    id, url, created_at
//	FROM
//	    urls
//	WHERE
//	    created_at >= ?
func (q *Queries) GetURLsByDate(ctx context.Context, arg GetURLsByDateParams) ([]Url, error) {
	rows, err := q.db.QueryContext(ctx, getURLsByDate, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Url
	for rows.Next() {
		var i Url
		if err := rows.Scan(&i.ID, &i.Url, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getURLsBySubstring = `-- name: GetURLsBySubstring :many
SELECT
    id, url, created_at
FROM
    urls
WHERE
    url LIKE ?
`

type GetURLsBySubstringParams struct {
	Url string `db:"url" json:"url"`
}

// GetURLsBySubstring
//
//	SELECT
//	    id, url, created_at
//	FROM
//	    urls
//	WHERE
//	    url LIKE ?
func (q *Queries) GetURLsBySubstring(ctx context.Context, arg GetURLsBySubstringParams) ([]Url, error) {
	rows, err := q.db.QueryContext(ctx, getURLsBySubstring, arg.Url)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Url
	for rows.Next() {
		var i Url
		if err := rows.Scan(&i.ID, &i.Url, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertBuildSum = `-- name: InsertBuildSum :one
INSERT INTO
    build_sums (build_sum)
VALUES
    (?) RETURNING id, build_sum, created_at
`

type InsertBuildSumParams struct {
	BuildSum string `db:"build_sum" json:"build_sum"`
}

// InsertBuildSum
//
//	INSERT INTO
//	    build_sums (build_sum)
//	VALUES
//	    (?) RETURNING id, build_sum, created_at
func (q *Queries) InsertBuildSum(ctx context.Context, arg InsertBuildSumParams) (BuildSum, error) {
	row := q.db.QueryRowContext(ctx, insertBuildSum, arg.BuildSum)
	var i BuildSum
	err := row.Scan(&i.ID, &i.BuildSum, &i.CreatedAt)
	return i, err
}

const insertBuildSumWithParam = `-- name: InsertBuildSumWithParam :exec
INSERT INTO
    build_sums (build_sum)
VALUES
    (?)
`

type InsertBuildSumWithParamParams struct {
	BuildSum string `db:"build_sum" json:"build_sum"`
}

// InsertBuildSumWithParam
//
//	INSERT INTO
//	    build_sums (build_sum)
//	VALUES
//	    (?)
func (q *Queries) InsertBuildSumWithParam(ctx context.Context, arg InsertBuildSumWithParamParams) error {
	_, err := q.db.ExecContext(ctx, insertBuildSumWithParam, arg.BuildSum)
	return err
}

const insertDeployment = `-- name: InsertDeployment :one
INSERT INTO
    deployments (name)
VALUES
    (?) RETURNING id, name, created_at, updated_at
`

type InsertDeploymentParams struct {
	Name string `db:"name" json:"name"`
}

// InsertDeployment
//
//	INSERT INTO
//	    deployments (name)
//	VALUES
//	    (?) RETURNING id, name, created_at, updated_at
func (q *Queries) InsertDeployment(ctx context.Context, arg InsertDeploymentParams) (Deployment, error) {
	row := q.db.QueryRowContext(ctx, insertDeployment, arg.Name)
	var i Deployment
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertGitRevision = `-- name: InsertGitRevision :one
INSERT INTO
    git_revisions (git_revision)
VALUES
    (?) RETURNING id, git_revision, created_at
`

type InsertGitRevisionParams struct {
	GitRevision string `db:"git_revision" json:"git_revision"`
}

// InsertGitRevision
//
//	INSERT INTO
//	    git_revisions (git_revision)
//	VALUES
//	    (?) RETURNING id, git_revision, created_at
func (q *Queries) InsertGitRevision(ctx context.Context, arg InsertGitRevisionParams) (GitRevision, error) {
	row := q.db.QueryRowContext(ctx, insertGitRevision, arg.GitRevision)
	var i GitRevision
	err := row.Scan(&i.ID, &i.GitRevision, &i.CreatedAt)
	return i, err
}

const insertGitRevisionWithParam = `-- name: InsertGitRevisionWithParam :exec
INSERT INTO
    git_revisions (git_revision)
VALUES
    (?)
`

type InsertGitRevisionWithParamParams struct {
	GitRevision string `db:"git_revision" json:"git_revision"`
}

// InsertGitRevisionWithParam
//
//	INSERT INTO
//	    git_revisions (git_revision)
//	VALUES
//	    (?)
func (q *Queries) InsertGitRevisionWithParam(ctx context.Context, arg InsertGitRevisionWithParamParams) error {
	_, err := q.db.ExecContext(ctx, insertGitRevisionWithParam, arg.GitRevision)
	return err
}

const insertGoVersion = `-- name: InsertGoVersion :one
INSERT INTO
    go_versions (name, version)
VALUES
    (?, ?) RETURNING id, name, version, created_at
`

type InsertGoVersionParams struct {
	Name    string `db:"name" json:"name"`
	Version string `db:"version" json:"version"`
}

// InsertGoVersion
//
//	INSERT INTO
//	    go_versions (name, version)
//	VALUES
//	    (?, ?) RETURNING id, name, version, created_at
func (q *Queries) InsertGoVersion(ctx context.Context, arg InsertGoVersionParams) (GoVersion, error) {
	row := q.db.QueryRowContext(ctx, insertGoVersion, arg.Name, arg.Version)
	var i GoVersion
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Version,
		&i.CreatedAt,
	)
	return i, err
}

const insertLogEntry = `-- name: InsertLogEntry :exec
INSERT INTO
    api_logs (
        go_version_id,
        build_sum_id,
        git_revision_id,
        user_agent,
        method,
        url,
        elapsed_ns,
        deployment_id,
        level_id
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type InsertLogEntryParams struct {
	GoVersionID   int64  `db:"go_version_id" json:"go_version_id"`
	BuildSumID    int64  `db:"build_sum_id" json:"build_sum_id"`
	GitRevisionID int64  `db:"git_revision_id" json:"git_revision_id"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	DeploymentID  int64  `db:"deployment_id" json:"deployment_id"`
	LevelID       int64  `db:"level_id" json:"level_id"`
}

// InsertLogEntry
//
//	INSERT INTO
//	    api_logs (
//	        go_version_id,
//	        build_sum_id,
//	        git_revision_id,
//	        user_agent,
//	        method,
//	        url,
//	        elapsed_ns,
//	        deployment_id,
//	        level_id
//	    )
//	VALUES
//	    (?, ?, ?, ?, ?, ?, ?, ?, ?)
func (q *Queries) InsertLogEntry(ctx context.Context, arg InsertLogEntryParams) error {
	_, err := q.db.ExecContext(ctx, insertLogEntry,
		arg.GoVersionID,
		arg.BuildSumID,
		arg.GitRevisionID,
		arg.UserAgent,
		arg.Method,
		arg.Url,
		arg.ElapsedNs,
		arg.DeploymentID,
		arg.LevelID,
	)
	return err
}

const insertLogLevel = `-- name: InsertLogLevel :exec
INSERT INTO
    log_levels (name)
VALUES
    (?)
`

type InsertLogLevelParams struct {
	Name string `db:"name" json:"name"`
}

// InsertLogLevel
//
//	INSERT INTO
//	    log_levels (name)
//	VALUES
//	    (?)
func (q *Queries) InsertLogLevel(ctx context.Context, arg InsertLogLevelParams) error {
	_, err := q.db.ExecContext(ctx, insertLogLevel, arg.Name)
	return err
}

const insertURL = `-- name: InsertURL :exec
INSERT INTO
    urls (url)
VALUES
    (?)
`

type InsertURLParams struct {
	Url string `db:"url" json:"url"`
}

// InsertURL
//
//	INSERT INTO
//	    urls (url)
//	VALUES
//	    (?)
func (q *Queries) InsertURL(ctx context.Context, arg InsertURLParams) error {
	_, err := q.db.ExecContext(ctx, insertURL, arg.Url)
	return err
}

const insertURLWithParam = `-- name: InsertURLWithParam :exec
INSERT INTO
    urls (url)
VALUES
    (?)
`

type InsertURLWithParamParams struct {
	Url string `db:"url" json:"url"`
}

// InsertURLWithParam
//
//	INSERT INTO
//	    urls (url)
//	VALUES
//	    (?)
func (q *Queries) InsertURLWithParam(ctx context.Context, arg InsertURLWithParamParams) error {
	_, err := q.db.ExecContext(ctx, insertURLWithParam, arg.Url)
	return err
}

const listBuildSums = `-- name: ListBuildSums :many
SELECT
    id, build_sum, created_at
FROM
    build_sums
`

// ListBuildSums
//
//	SELECT
//	    id, build_sum, created_at
//	FROM
//	    build_sums
func (q *Queries) ListBuildSums(ctx context.Context) ([]BuildSum, error) {
	rows, err := q.db.QueryContext(ctx, listBuildSums)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BuildSum
	for rows.Next() {
		var i BuildSum
		if err := rows.Scan(&i.ID, &i.BuildSum, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBuildSumsPaginated = `-- name: ListBuildSumsPaginated :many
SELECT
    id, build_sum, created_at
FROM
    build_sums
LIMIT
    ? OFFSET ?
`

type ListBuildSumsPaginatedParams struct {
	Limit  int64 `db:"limit" json:"limit"`
	Offset int64 `db:"offset" json:"offset"`
}

// ListBuildSumsPaginated
//
//	SELECT
//	    id, build_sum, created_at
//	FROM
//	    build_sums
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListBuildSumsPaginated(ctx context.Context, arg ListBuildSumsPaginatedParams) ([]BuildSum, error) {
	rows, err := q.db.QueryContext(ctx, listBuildSumsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BuildSum
	for rows.Next() {
		var i BuildSum
		if err := rows.Scan(&i.ID, &i.BuildSum, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDeployments = `-- name: ListDeployments :many
SELECT
    id, name, created_at, updated_at
FROM
    deployments
`

// ListDeployments
//
//	SELECT
//	    id, name, created_at, updated_at
//	FROM
//	    deployments
func (q *Queries) ListDeployments(ctx context.Context) ([]Deployment, error) {
	rows, err := q.db.QueryContext(ctx, listDeployments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Deployment
	for rows.Next() {
		var i Deployment
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDeploymentsPaginated = `-- name: ListDeploymentsPaginated :many
SELECT
    id, name, created_at, updated_at
FROM
    deployments
LIMIT
    ? OFFSET ?
`

type ListDeploymentsPaginatedParams struct {
	Limit  int64 `db:"limit" json:"limit"`
	Offset int64 `db:"offset" json:"offset"`
}

// ListDeploymentsPaginated
//
//	SELECT
//	    id, name, created_at, updated_at
//	FROM
//	    deployments
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListDeploymentsPaginated(ctx context.Context, arg ListDeploymentsPaginatedParams) ([]Deployment, error) {
	rows, err := q.db.QueryContext(ctx, listDeploymentsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Deployment
	for rows.Next() {
		var i Deployment
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGitRevisions = `-- name: ListGitRevisions :many
SELECT
    id, git_revision, created_at
FROM
    git_revisions
`

// ListGitRevisions
//
//	SELECT
//	    id, git_revision, created_at
//	FROM
//	    git_revisions
func (q *Queries) ListGitRevisions(ctx context.Context) ([]GitRevision, error) {
	rows, err := q.db.QueryContext(ctx, listGitRevisions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GitRevision
	for rows.Next() {
		var i GitRevision
		if err := rows.Scan(&i.ID, &i.GitRevision, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGitRevisionsPaginated = `-- name: ListGitRevisionsPaginated :many
SELECT
    id, git_revision, created_at
FROM
    git_revisions
LIMIT
    ? OFFSET ?
`

type ListGitRevisionsPaginatedParams struct {
	Limit  int64 `db:"limit" json:"limit"`
	Offset int64 `db:"offset" json:"offset"`
}

// ListGitRevisionsPaginated
//
//	SELECT
//	    id, git_revision, created_at
//	FROM
//	    git_revisions
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListGitRevisionsPaginated(ctx context.Context, arg ListGitRevisionsPaginatedParams) ([]GitRevision, error) {
	rows, err := q.db.QueryContext(ctx, listGitRevisionsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GitRevision
	for rows.Next() {
		var i GitRevision
		if err := rows.Scan(&i.ID, &i.GitRevision, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGoVersions = `-- name: ListGoVersions :many
SELECT
    id, name, version, created_at
FROM
    go_versions
`

// ListGoVersions
//
//	SELECT
//	    id, name, version, created_at
//	FROM
//	    go_versions
func (q *Queries) ListGoVersions(ctx context.Context) ([]GoVersion, error) {
	rows, err := q.db.QueryContext(ctx, listGoVersions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoVersion
	for rows.Next() {
		var i GoVersion
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Version,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGoVersionsPaginated = `-- name: ListGoVersionsPaginated :many
SELECT
    id, name, version, created_at
FROM
    go_versions
LIMIT
    ? OFFSET ?
`

type ListGoVersionsPaginatedParams struct {
	Limit  int64 `db:"limit" json:"limit"`
	Offset int64 `db:"offset" json:"offset"`
}

// ListGoVersionsPaginated
//
//	SELECT
//	    id, name, version, created_at
//	FROM
//	    go_versions
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListGoVersionsPaginated(ctx context.Context, arg ListGoVersionsPaginatedParams) ([]GoVersion, error) {
	rows, err := q.db.QueryContext(ctx, listGoVersionsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoVersion
	for rows.Next() {
		var i GoVersion
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Version,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogLevels = `-- name: ListLogLevels :many
SELECT
    id, name
FROM
    log_levels
`

// ListLogLevels
//
//	SELECT
//	    id, name
//	FROM
//	    log_levels
func (q *Queries) ListLogLevels(ctx context.Context) ([]LogLevel, error) {
	rows, err := q.db.QueryContext(ctx, listLogLevels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LogLevel
	for rows.Next() {
		var i LogLevel
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogLevelsPaginated = `-- name: ListLogLevelsPaginated :many
SELECT
    id, name
FROM
    log_levels
LIMIT
    ? OFFSET ?
`

type ListLogLevelsPaginatedParams struct {
	Limit  int64 `db:"limit" json:"limit"`
	Offset int64 `db:"offset" json:"offset"`
}

// ListLogLevelsPaginated
//
//	SELECT
//	    id, name
//	FROM
//	    log_levels
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListLogLevelsPaginated(ctx context.Context, arg ListLogLevelsPaginatedParams) ([]LogLevel, error) {
	rows, err := q.db.QueryContext(ctx, listLogLevelsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LogLevel
	for rows.Next() {
		var i LogLevel
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogs = `-- name: ListLogs :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
`

// ListLogs
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
func (q *Queries) ListLogs(ctx context.Context) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, listLogs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsByDeploymentID = `-- name: ListLogsByDeploymentID :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    deployment_id = ?
`

type ListLogsByDeploymentIDParams struct {
	DeploymentID int64 `db:"deployment_id" json:"deployment_id"`
}

// ListLogsByDeploymentID
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    deployment_id = ?
func (q *Queries) ListLogsByDeploymentID(ctx context.Context, arg ListLogsByDeploymentIDParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, listLogsByDeploymentID, arg.DeploymentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsByMethod = `-- name: ListLogsByMethod :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    method = ?
`

type ListLogsByMethodParams struct {
	Method string `db:"method" json:"method"`
}

// ListLogsByMethod
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    method = ?
func (q *Queries) ListLogsByMethod(ctx context.Context, arg ListLogsByMethodParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, listLogsByMethod, arg.Method)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsByUserAgent = `-- name: ListLogsByUserAgent :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
WHERE
    user_agent = ?
`

type ListLogsByUserAgentParams struct {
	UserAgent string `db:"user_agent" json:"user_agent"`
}

// ListLogsByUserAgent
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	WHERE
//	    user_agent = ?
func (q *Queries) ListLogsByUserAgent(ctx context.Context, arg ListLogsByUserAgentParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, listLogsByUserAgent, arg.UserAgent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsPaginated = `-- name: ListLogsPaginated :many
SELECT
    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
FROM
    api_logs
LIMIT
    ? OFFSET ?
`

type ListLogsPaginatedParams struct {
	Limit  int64 `db:"limit" json:"limit"`
	Offset int64 `db:"offset" json:"offset"`
}

// ListLogsPaginated
//
//	SELECT
//	    id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
//	FROM
//	    api_logs
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListLogsPaginated(ctx context.Context, arg ListLogsPaginatedParams) ([]ApiLog, error) {
	rows, err := q.db.QueryContext(ctx, listLogsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiLog
	for rows.Next() {
		var i ApiLog
		if err := rows.Scan(
			&i.ID,
			&i.LevelID,
			&i.CreatedAt,
			&i.GoVersionID,
			&i.BuildSumID,
			&i.GitRevisionID,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.DeploymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoin = `-- name: ListLogsWithJoin :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
`

type ListLogsWithJoinRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoin
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
func (q *Queries) ListLogsWithJoin(ctx context.Context) ([]ListLogsWithJoinRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinRow
	for rows.Next() {
		var i ListLogsWithJoinRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByBuildSumID = `-- name: ListLogsWithJoinByBuildSumID :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    bs.id = ?
`

type ListLogsWithJoinByBuildSumIDParams struct {
	ID int64 `db:"id" json:"id"`
}

type ListLogsWithJoinByBuildSumIDRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByBuildSumID
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    bs.id = ?
func (q *Queries) ListLogsWithJoinByBuildSumID(ctx context.Context, arg ListLogsWithJoinByBuildSumIDParams) ([]ListLogsWithJoinByBuildSumIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByBuildSumID, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByBuildSumIDRow
	for rows.Next() {
		var i ListLogsWithJoinByBuildSumIDRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByBuildSumIDPaginated = `-- name: ListLogsWithJoinByBuildSumIDPaginated :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    bs.id = ?
`

type ListLogsWithJoinByBuildSumIDPaginatedParams struct {
	ID int64 `db:"id" json:"id"`
}

type ListLogsWithJoinByBuildSumIDPaginatedRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByBuildSumIDPaginated
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    bs.id = ?
func (q *Queries) ListLogsWithJoinByBuildSumIDPaginated(ctx context.Context, arg ListLogsWithJoinByBuildSumIDPaginatedParams) ([]ListLogsWithJoinByBuildSumIDPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByBuildSumIDPaginated, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByBuildSumIDPaginatedRow
	for rows.Next() {
		var i ListLogsWithJoinByBuildSumIDPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByDate = `-- name: ListLogsWithJoinByDate :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    created_at BETWEEN ? AND ?
`

type ListLogsWithJoinByDateParams struct {
	FromCreatedAt   string    `db:"from_created_at" json:"from_created_at"`
	FromCreatedAt_2 string    `db:"from_created_at_2" json:"from_created_at_2"`
	FromCreatedAt_3 string    `db:"from_created_at_3" json:"from_created_at_3"`
	FromCreatedAt_4 time.Time `db:"from_created_at_4" json:"from_created_at_4"`
	FromCreatedAt_5 string    `db:"from_created_at_5" json:"from_created_at_5"`
	ToCreatedAt     string    `db:"to_created_at" json:"to_created_at"`
	ToCreatedAt_2   string    `db:"to_created_at_2" json:"to_created_at_2"`
	ToCreatedAt_3   string    `db:"to_created_at_3" json:"to_created_at_3"`
	ToCreatedAt_4   time.Time `db:"to_created_at_4" json:"to_created_at_4"`
	ToCreatedAt_5   string    `db:"to_created_at_5" json:"to_created_at_5"`
}

type ListLogsWithJoinByDateRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByDate
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    created_at BETWEEN ? AND ?
func (q *Queries) ListLogsWithJoinByDate(ctx context.Context, arg ListLogsWithJoinByDateParams) ([]ListLogsWithJoinByDateRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByDate,
		arg.FromCreatedAt,
		arg.FromCreatedAt_2,
		arg.FromCreatedAt_3,
		arg.FromCreatedAt_4,
		arg.FromCreatedAt_5,
		arg.ToCreatedAt,
		arg.ToCreatedAt_2,
		arg.ToCreatedAt_3,
		arg.ToCreatedAt_4,
		arg.ToCreatedAt_5,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByDateRow
	for rows.Next() {
		var i ListLogsWithJoinByDateRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByDatePaginated = `-- name: ListLogsWithJoinByDatePaginated :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    created_at BETWEEN ? AND ?
LIMIT
    ? OFFSET ?
`

type ListLogsWithJoinByDatePaginatedParams struct {
	FromCreatedAt   string    `db:"from_created_at" json:"from_created_at"`
	FromCreatedAt_2 string    `db:"from_created_at_2" json:"from_created_at_2"`
	FromCreatedAt_3 string    `db:"from_created_at_3" json:"from_created_at_3"`
	FromCreatedAt_4 time.Time `db:"from_created_at_4" json:"from_created_at_4"`
	FromCreatedAt_5 string    `db:"from_created_at_5" json:"from_created_at_5"`
	ToCreatedAt     string    `db:"to_created_at" json:"to_created_at"`
	ToCreatedAt_2   string    `db:"to_created_at_2" json:"to_created_at_2"`
	ToCreatedAt_3   string    `db:"to_created_at_3" json:"to_created_at_3"`
	ToCreatedAt_4   time.Time `db:"to_created_at_4" json:"to_created_at_4"`
	ToCreatedAt_5   string    `db:"to_created_at_5" json:"to_created_at_5"`
	Limit           int64     `db:"limit" json:"limit"`
	Offset          int64     `db:"offset" json:"offset"`
}

type ListLogsWithJoinByDatePaginatedRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByDatePaginated
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    created_at BETWEEN ? AND ?
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListLogsWithJoinByDatePaginated(ctx context.Context, arg ListLogsWithJoinByDatePaginatedParams) ([]ListLogsWithJoinByDatePaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByDatePaginated,
		arg.FromCreatedAt,
		arg.FromCreatedAt_2,
		arg.FromCreatedAt_3,
		arg.FromCreatedAt_4,
		arg.FromCreatedAt_5,
		arg.ToCreatedAt,
		arg.ToCreatedAt_2,
		arg.ToCreatedAt_3,
		arg.ToCreatedAt_4,
		arg.ToCreatedAt_5,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByDatePaginatedRow
	for rows.Next() {
		var i ListLogsWithJoinByDatePaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByDateRange = `-- name: ListLogsWithJoinByDateRange :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    created_at BETWEEN ? AND ?
`

type ListLogsWithJoinByDateRangeParams struct {
	FromCreatedAt   string    `db:"from_created_at" json:"from_created_at"`
	FromCreatedAt_2 string    `db:"from_created_at_2" json:"from_created_at_2"`
	FromCreatedAt_3 string    `db:"from_created_at_3" json:"from_created_at_3"`
	FromCreatedAt_4 time.Time `db:"from_created_at_4" json:"from_created_at_4"`
	FromCreatedAt_5 string    `db:"from_created_at_5" json:"from_created_at_5"`
	ToCreatedAt     string    `db:"to_created_at" json:"to_created_at"`
	ToCreatedAt_2   string    `db:"to_created_at_2" json:"to_created_at_2"`
	ToCreatedAt_3   string    `db:"to_created_at_3" json:"to_created_at_3"`
	ToCreatedAt_4   time.Time `db:"to_created_at_4" json:"to_created_at_4"`
	ToCreatedAt_5   string    `db:"to_created_at_5" json:"to_created_at_5"`
}

type ListLogsWithJoinByDateRangeRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByDateRange
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    created_at BETWEEN ? AND ?
func (q *Queries) ListLogsWithJoinByDateRange(ctx context.Context, arg ListLogsWithJoinByDateRangeParams) ([]ListLogsWithJoinByDateRangeRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByDateRange,
		arg.FromCreatedAt,
		arg.FromCreatedAt_2,
		arg.FromCreatedAt_3,
		arg.FromCreatedAt_4,
		arg.FromCreatedAt_5,
		arg.ToCreatedAt,
		arg.ToCreatedAt_2,
		arg.ToCreatedAt_3,
		arg.ToCreatedAt_4,
		arg.ToCreatedAt_5,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByDateRangeRow
	for rows.Next() {
		var i ListLogsWithJoinByDateRangeRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByDateRangePaginated = `-- name: ListLogsWithJoinByDateRangePaginated :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    created_at BETWEEN ? AND ?
LIMIT
    ? OFFSET ?
`

type ListLogsWithJoinByDateRangePaginatedParams struct {
	FromCreatedAt   string    `db:"from_created_at" json:"from_created_at"`
	FromCreatedAt_2 string    `db:"from_created_at_2" json:"from_created_at_2"`
	FromCreatedAt_3 string    `db:"from_created_at_3" json:"from_created_at_3"`
	FromCreatedAt_4 time.Time `db:"from_created_at_4" json:"from_created_at_4"`
	FromCreatedAt_5 string    `db:"from_created_at_5" json:"from_created_at_5"`
	ToCreatedAt     string    `db:"to_created_at" json:"to_created_at"`
	ToCreatedAt_2   string    `db:"to_created_at_2" json:"to_created_at_2"`
	ToCreatedAt_3   string    `db:"to_created_at_3" json:"to_created_at_3"`
	ToCreatedAt_4   time.Time `db:"to_created_at_4" json:"to_created_at_4"`
	ToCreatedAt_5   string    `db:"to_created_at_5" json:"to_created_at_5"`
	Limit           int64     `db:"limit" json:"limit"`
	Offset          int64     `db:"offset" json:"offset"`
}

type ListLogsWithJoinByDateRangePaginatedRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByDateRangePaginated
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    created_at BETWEEN ? AND ?
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListLogsWithJoinByDateRangePaginated(ctx context.Context, arg ListLogsWithJoinByDateRangePaginatedParams) ([]ListLogsWithJoinByDateRangePaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByDateRangePaginated,
		arg.FromCreatedAt,
		arg.FromCreatedAt_2,
		arg.FromCreatedAt_3,
		arg.FromCreatedAt_4,
		arg.FromCreatedAt_5,
		arg.ToCreatedAt,
		arg.ToCreatedAt_2,
		arg.ToCreatedAt_3,
		arg.ToCreatedAt_4,
		arg.ToCreatedAt_5,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByDateRangePaginatedRow
	for rows.Next() {
		var i ListLogsWithJoinByDateRangePaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByElapsedRange = `-- name: ListLogsWithJoinByElapsedRange :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    elapsed_ns BETWEEN ? AND ?
`

type ListLogsWithJoinByElapsedRangeParams struct {
	FromElapsedNs int64 `db:"from_elapsed_ns" json:"from_elapsed_ns"`
	ToElapsedNs   int64 `db:"to_elapsed_ns" json:"to_elapsed_ns"`
}

type ListLogsWithJoinByElapsedRangeRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByElapsedRange
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    elapsed_ns BETWEEN ? AND ?
func (q *Queries) ListLogsWithJoinByElapsedRange(ctx context.Context, arg ListLogsWithJoinByElapsedRangeParams) ([]ListLogsWithJoinByElapsedRangeRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByElapsedRange, arg.FromElapsedNs, arg.ToElapsedNs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByElapsedRangeRow
	for rows.Next() {
		var i ListLogsWithJoinByElapsedRangeRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByElapsedRangePaginated = `-- name: ListLogsWithJoinByElapsedRangePaginated :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    elapsed_ns BETWEEN ? AND ?
LIMIT
    ? OFFSET ?
`

type ListLogsWithJoinByElapsedRangePaginatedParams struct {
	FromElapsedNs int64 `db:"from_elapsed_ns" json:"from_elapsed_ns"`
	ToElapsedNs   int64 `db:"to_elapsed_ns" json:"to_elapsed_ns"`
	Limit         int64 `db:"limit" json:"limit"`
	Offset        int64 `db:"offset" json:"offset"`
}

type ListLogsWithJoinByElapsedRangePaginatedRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByElapsedRangePaginated
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    elapsed_ns BETWEEN ? AND ?
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListLogsWithJoinByElapsedRangePaginated(ctx context.Context, arg ListLogsWithJoinByElapsedRangePaginatedParams) ([]ListLogsWithJoinByElapsedRangePaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByElapsedRangePaginated,
		arg.FromElapsedNs,
		arg.ToElapsedNs,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByElapsedRangePaginatedRow
	for rows.Next() {
		var i ListLogsWithJoinByElapsedRangePaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByGitRevisionID = `-- name: ListLogsWithJoinByGitRevisionID :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    gr.id = ?
`

type ListLogsWithJoinByGitRevisionIDParams struct {
	ID int64 `db:"id" json:"id"`
}

type ListLogsWithJoinByGitRevisionIDRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByGitRevisionID
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    gr.id = ?
func (q *Queries) ListLogsWithJoinByGitRevisionID(ctx context.Context, arg ListLogsWithJoinByGitRevisionIDParams) ([]ListLogsWithJoinByGitRevisionIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByGitRevisionID, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByGitRevisionIDRow
	for rows.Next() {
		var i ListLogsWithJoinByGitRevisionIDRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByGitRevisionIDPaginated = `-- name: ListLogsWithJoinByGitRevisionIDPaginated :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    gr.id = ?
`

type ListLogsWithJoinByGitRevisionIDPaginatedParams struct {
	ID int64 `db:"id" json:"id"`
}

type ListLogsWithJoinByGitRevisionIDPaginatedRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByGitRevisionIDPaginated
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    gr.id = ?
func (q *Queries) ListLogsWithJoinByGitRevisionIDPaginated(ctx context.Context, arg ListLogsWithJoinByGitRevisionIDPaginatedParams) ([]ListLogsWithJoinByGitRevisionIDPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByGitRevisionIDPaginated, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByGitRevisionIDPaginatedRow
	for rows.Next() {
		var i ListLogsWithJoinByGitRevisionIDPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByGoVersionID = `-- name: ListLogsWithJoinByGoVersionID :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    gv.id = ?
`

type ListLogsWithJoinByGoVersionIDParams struct {
	ID int64 `db:"id" json:"id"`
}

type ListLogsWithJoinByGoVersionIDRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByGoVersionID
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    gv.id = ?
func (q *Queries) ListLogsWithJoinByGoVersionID(ctx context.Context, arg ListLogsWithJoinByGoVersionIDParams) ([]ListLogsWithJoinByGoVersionIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByGoVersionID, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByGoVersionIDRow
	for rows.Next() {
		var i ListLogsWithJoinByGoVersionIDRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinByGoVersionIDPaginated = `-- name: ListLogsWithJoinByGoVersionIDPaginated :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
WHERE
    gv.id = ?
`

type ListLogsWithJoinByGoVersionIDPaginatedParams struct {
	ID int64 `db:"id" json:"id"`
}

type ListLogsWithJoinByGoVersionIDPaginatedRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinByGoVersionIDPaginated
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	WHERE
//	    gv.id = ?
func (q *Queries) ListLogsWithJoinByGoVersionIDPaginated(ctx context.Context, arg ListLogsWithJoinByGoVersionIDPaginatedParams) ([]ListLogsWithJoinByGoVersionIDPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinByGoVersionIDPaginated, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinByGoVersionIDPaginatedRow
	for rows.Next() {
		var i ListLogsWithJoinByGoVersionIDPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLogsWithJoinPaginated = `-- name: ListLogsWithJoinPaginated :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ns,
    l.level_id,
    le.name AS level_name,
    gv.name AS go_version_name,
    gv.version AS go_version,
    bs.build_sum,
    gr.git_revision
FROM
    api_logs l
    JOIN log_levels le ON l.level_id = le.id
    JOIN deployments d ON l.deployment_id = d.id
    JOIN go_versions gv ON l.go_version_id = gv.id
    JOIN build_sums bs ON l.build_sum_id = bs.id
    JOIN git_revisions gr ON l.git_revision_id = gr.id
LIMIT
    ? OFFSET ?
`

type ListLogsWithJoinPaginatedParams struct {
	Limit  int64 `db:"limit" json:"limit"`
	Offset int64 `db:"offset" json:"offset"`
}

type ListLogsWithJoinPaginatedRow struct {
	ID            int64  `db:"id" json:"id"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	Method        string `db:"method" json:"method"`
	Url           string `db:"url" json:"url"`
	ElapsedNs     int64  `db:"elapsed_ns" json:"elapsed_ns"`
	LevelID       int64  `db:"level_id" json:"level_id"`
	LevelName     string `db:"level_name" json:"level_name"`
	GoVersionName string `db:"go_version_name" json:"go_version_name"`
	GoVersion     string `db:"go_version" json:"go_version"`
	BuildSum      string `db:"build_sum" json:"build_sum"`
	GitRevision   string `db:"git_revision" json:"git_revision"`
}

// ListLogsWithJoinPaginated
//
//	SELECT
//	    l.id,
//	    l.created_at,
//	    l.user_agent,
//	    l.method,
//	    l.url,
//	    l.elapsed_ns,
//	    l.level_id,
//	    le.name AS level_name,
//	    gv.name AS go_version_name,
//	    gv.version AS go_version,
//	    bs.build_sum,
//	    gr.git_revision
//	FROM
//	    api_logs l
//	    JOIN log_levels le ON l.level_id = le.id
//	    JOIN deployments d ON l.deployment_id = d.id
//	    JOIN go_versions gv ON l.go_version_id = gv.id
//	    JOIN build_sums bs ON l.build_sum_id = bs.id
//	    JOIN git_revisions gr ON l.git_revision_id = gr.id
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListLogsWithJoinPaginated(ctx context.Context, arg ListLogsWithJoinPaginatedParams) ([]ListLogsWithJoinPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsWithJoinPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsWithJoinPaginatedRow
	for rows.Next() {
		var i ListLogsWithJoinPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserAgent,
			&i.Method,
			&i.Url,
			&i.ElapsedNs,
			&i.LevelID,
			&i.LevelName,
			&i.GoVersionName,
			&i.GoVersion,
			&i.BuildSum,
			&i.GitRevision,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listURLs = `-- name: ListURLs :many
SELECT
    id, url, created_at
FROM
    urls
`

// ListURLs
//
//	SELECT
//	    id, url, created_at
//	FROM
//	    urls
func (q *Queries) ListURLs(ctx context.Context) ([]Url, error) {
	rows, err := q.db.QueryContext(ctx, listURLs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Url
	for rows.Next() {
		var i Url
		if err := rows.Scan(&i.ID, &i.Url, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listURLsPaginated = `-- name: ListURLsPaginated :many
SELECT
    id, url, created_at
FROM
    urls
LIMIT
    ? OFFSET ?
`

type ListURLsPaginatedParams struct {
	Limit  int64 `db:"limit" json:"limit"`
	Offset int64 `db:"offset" json:"offset"`
}

// ListURLsPaginated
//
//	SELECT
//	    id, url, created_at
//	FROM
//	    urls
//	LIMIT
//	    ? OFFSET ?
func (q *Queries) ListURLsPaginated(ctx context.Context, arg ListURLsPaginatedParams) ([]Url, error) {
	rows, err := q.db.QueryContext(ctx, listURLsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Url
	for rows.Next() {
		var i Url
		if err := rows.Scan(&i.ID, &i.Url, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateGoVersionByID = `-- name: UpdateGoVersionByID :exec
UPDATE
    go_versions
SET
    name = ?,
    version = ?
WHERE
    id = 1
`

type UpdateGoVersionByIDParams struct {
	Name    string `db:"name" json:"name"`
	Version string `db:"version" json:"version"`
}

// UpdateGoVersionByID
//
//	UPDATE
//	    go_versions
//	SET
//	    name = ?,
//	    version = ?
//	WHERE
//	    id = 1
func (q *Queries) UpdateGoVersionByID(ctx context.Context, arg UpdateGoVersionByIDParams) error {
	_, err := q.db.ExecContext(ctx, updateGoVersionByID, arg.Name, arg.Version)
	return err
}
