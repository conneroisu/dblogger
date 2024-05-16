-- name: CountLogs :one
SELECT
    COUNT(*)
FROM
    api_logs;

-- name: InsertLogEntry :exec
INSERT INTO
    api_logs (
        go_version_id,
        build_sum_id,
        git_revision_id,
        user_agent,
        method,
        url,
        elapsed_ms,
        status_code,
        deployment_id,
        level_id
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListLogs :many
SELECT
    *
FROM
    api_logs;

-- name: ListLogsByDeploymentID :many
SELECT
    *
FROM
    api_logs
WHERE
    deployment_id = ?;

-- name: DeleteLogByID :exec
DELETE FROM
    api_logs
WHERE
    id = ?;

-- name: GetLogsByGoVersionID :many
SELECT
    *
FROM
    api_logs
WHERE
    go_version_id = ?;

-- name: GetLogsByBuildSumID :many
SELECT
    *
FROM
    api_logs
WHERE
    build_sum_id = ?;

-- name: GetLogsByGitRevisionID :many
SELECT
    *
FROM
    api_logs
WHERE
    git_revision_id = ?;

-- name: GetLogsByElapsedRange :many
SELECT
    *
FROM
    api_logs
WHERE
    elapsed_ms BETWEEN ? AND ?;

-- name: GetLogsByDeploymentIDAndStatus :many
SELECT
    *
FROM
    api_logs
WHERE
    deployment_id = ?
    AND status_code = ?;

-- name: GetLogsByDateRange :many
SELECT
    *
FROM
    api_logs
WHERE
    created_at BETWEEN ? AND ?;

-- name: ListLogsPaginated :many
SELECT
    *
FROM
    api_logs
LIMIT
    ? OFFSET ?;
