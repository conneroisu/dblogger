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
-- name: ListLogsWithJoin :many
SELECT
    l.id,
    l.created_at,
    l.user_agent,
    l.method,
    l.url,
    l.elapsed_ms,
    l.status_code,
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
    JOIN git_revisions gr ON l.git_revision_id = gr.id;

-- name: InsertBuildSumWithParam :exec
INSERT INTO
    build_sums (build_sum)
VALUES
    (?);

-- name: GetBuildSumByID :one
SELECT
    *
FROM
    build_sums
WHERE
    id = ?;

-- name: GetBuildSumsByDate :many
SELECT
    *
FROM
    build_sums
WHERE
    created_at >= ?;

-- name: InsertDeploymentWithParam :exec
INSERT INTO
    deployments (name)
VALUES
    (?);

-- name: GetDeploymentByID :one
SELECT
    *
FROM
    deployments
WHERE
    id = ?;

-- name: GetDeploymentsByDate :many
SELECT
    *
FROM
    deployments
WHERE
    created_at >= ?;

-- name: InsertGoVersionWithParams :exec
INSERT INTO
    go_versions (name, version)
VALUES
    (?, ?);

-- name: GetGoVersionByID :one
SELECT
    *
FROM
    go_versions
WHERE
    id = ?;

-- name: GetGoVersionsByDate :many
SELECT
    *
FROM
    go_versions
WHERE
    created_at >= ?;

-- name: GetLogByID :one
SELECT
    *
FROM
    api_logs
WHERE
    id = ?;

-- name: GetLogsByDate :many
SELECT
    *
FROM
    api_logs
WHERE
    created_at >= ?;

-- name: InsertURLWithParam :exec
INSERT INTO
    urls (url)
VALUES
    (?);

-- name: GetURLByID :one
SELECT
    *
FROM
    urls
WHERE
    id = ?;

-- name: GetURLsByDate :many
SELECT
    *
FROM
    urls
WHERE
    created_at >= ?;

-- name: ListLogsByUserAgent :many
SELECT
    *
FROM
    api_logs
WHERE
    user_agent = ?;

-- name: ListLogsByMethod :many
SELECT
    *
FROM
    api_logs
WHERE
    method = ?;

-- name: ListLogsByStatusCode :many
SELECT
    *
FROM
    api_logs
WHERE
    status_code = ?;

-- name: ListBuildSumsPaginated :many
SELECT
    *
FROM
    build_sums
LIMIT
    ? OFFSET ?;

-- name: ListDeploymentsPaginated :many
SELECT
    *
FROM
    deployments
LIMIT
    ? OFFSET ?;

-- name: ListGitRevisionsPaginated :many
SELECT
    *
FROM
    git_revisions
LIMIT
    ? OFFSET ?;

-- name: CountGoVersions :one
SELECT
    COUNT(*)
FROM
    go_versions;

-- name: GetDeploymentsBySubstring :many
SELECT
    *
FROM
    deployments
WHERE
    name LIKE ?;
-- name: CountBuildSums :one
SELECT
    COUNT(*)
FROM
    build_sums;

-- name: InsertBuildSum :exec
INSERT INTO
    build_sums (build_sum)
VALUES
    (?);

-- name: ListBuildSums :many
SELECT
    *
FROM
    build_sums;

-- name: DeleteBuildSumByID :exec
DELETE FROM
    build_sums
WHERE
    id = ?;

-- name: GetBuildSumsBySubstring :many
SELECT
    *
FROM
    build_sums
WHERE
    build_sum LIKE ?;
-- name: CountDeployments :one
SELECT
    COUNT(*)
FROM
    deployments;

-- name: InsertDeployment :exec
INSERT INTO
    deployments (name)
VALUES
    (?);

-- name: ListDeployments :many
SELECT
    *
FROM
    deployments;

-- name: DeleteDeploymentByID :exec
DELETE FROM
    deployments
WHERE
    id = ?;

-- name: GetDeploymentsByDateRange :many
SELECT
    *
FROM
    deployments
WHERE
    created_at BETWEEN ? AND ?;
-- name: CountGitRevisions :one
SELECT
    COUNT(*)
FROM
    git_revisions;

-- name: InsertGitRevision :exec
INSERT INTO
    git_revisions (git_revision)
VALUES
    (?);

-- name: ListGitRevisions :many
SELECT
    *
FROM
    git_revisions;

-- name: DeleteGitRevisionByID :exec
DELETE FROM
    git_revisions
WHERE
    id = ?;

-- name: InsertGitRevisionWithParam :exec
INSERT INTO
    git_revisions (git_revision)
VALUES
    (?);

-- name: GetGitRevisionByID :one
SELECT
    *
FROM
    git_revisions
WHERE
    id = ?;

-- name: GetGitRevisionsByDate :many
SELECT
    *
FROM
    git_revisions
WHERE
    created_at >= ?;

-- name: GetGitRevisionsBySubstring :many
SELECT
    *
FROM
    git_revisions
WHERE
    git_revision LIKE ?;
-- name: InsertGoVersion :exec
INSERT INTO
    go_versions (name, version)
VALUES
    (?, ?);

-- name: ListGoVersions :many
SELECT
    *
FROM
    go_versions;

-- name: UpdateGoVersionByID :exec
UPDATE
    go_versions
SET
    name = ?,
    version = ?
WHERE
    id = 1;

-- name: DeleteGoVersionByID :exec
DELETE FROM
    go_versions
WHERE
    id = 1;

-- name: GetGoVersionsBySubstring :many
SELECT
    *
FROM
    go_versions
WHERE
    name LIKE ?
    OR version LIKE ?;

-- name: ListGoVersionsPaginated :many
SELECT
    *
FROM
    go_versions
LIMIT
    ? OFFSET ?;
-- name: CountLogLevels :one
SELECT
    COUNT(*)
FROM
    log_levels;

-- name: InsertLogLevel :exec
INSERT INTO
    log_levels (name)
VALUES
    (?);

-- name: ListLogLevels :many
SELECT
    *
FROM
    log_levels;

-- name: DeleteLogLevelByID :exec
DELETE FROM
    log_levels
WHERE
    id = ?;

-- name: GetLogLevelsBySubstring :many
SELECT
    *
FROM
    log_levels
WHERE
    name LIKE ?;

-- name: ListLogLevelsPaginated :many
SELECT
    *
FROM
    log_levels
LIMIT
    ? OFFSET ?;
-- name: CountURLs :one
SELECT
    COUNT(*)
FROM
    urls;

-- name: InsertURL :exec
INSERT INTO
    urls (url)
VALUES
    (?);

-- name: ListURLs :many
SELECT
    *
FROM
    urls;

-- name: DeleteURLByID :exec
DELETE FROM
    urls
WHERE
    id = ?;

-- name: GetURLsBySubstring :many
SELECT
    *
FROM
    urls
WHERE
    url LIKE ?;

-- name: GetLogsByURLSubstring :many
SELECT
    *
FROM
    api_logs
WHERE
    url LIKE ?;

-- name: ListURLsPaginated :many
SELECT
    *
FROM
    urls
LIMIT
    ? OFFSET ?;
