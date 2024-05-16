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
