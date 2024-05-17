-- file: api_logs.sql
-- url: github.com/conneroisu/dblogger/data/queries/api_logs.sql
-- description: api_logs.sql is

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
        elapsed_ns,
        deployment_id,
        level_id
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?);

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
    elapsed_ns BETWEEN ? AND ?;

-- name: GetLogsByDateRange :many
SELECT
    *
FROM
    api_logs
WHERE
    created_at BETWEEN ? AND ?;

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

-- name: ListLogsPaginated :many
SELECT
    *
FROM
    api_logs
LIMIT
    ? OFFSET ?;

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

-- name: ListLogsWithJoin :many
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
    JOIN git_revisions gr ON l.git_revision_id = gr.id;

-- name: ListLogsWithJoinPaginated :many
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
    ? OFFSET ?;

-- name: ListLogsWithJoinByDate :many
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
    created_at BETWEEN ? AND ?;

-- name: ListLogsWithJoinByDatePaginated :many
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
    ? OFFSET ?;

-- name: ListLogsWithJoinByDateRange :many
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
    created_at BETWEEN ? AND ?;

-- name: ListLogsWithJoinByDateRangePaginated :many
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
    ? OFFSET ?;

-- name: ListLogsWithJoinByElapsedRangePaginated :many
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
    ? OFFSET ?;

-- name: ListLogsWithJoinByGoVersionID :many
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
    gv.id = ?;

-- name: ListLogsWithJoinByGoVersionIDPaginated :many
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
    gv.id = ?;

-- name: ListLogsWithJoinByBuildSumID :many
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
    bs.id = ?;

-- name: ListLogsWithJoinByBuildSumIDPaginated :many
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
    bs.id = ?;

-- name: ListLogsWithJoinByGitRevisionID :many
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
    gr.id = ?;

-- name: ListLogsWithJoinByGitRevisionIDPaginated :many
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
    gr.id = ?;

-- name: ListLogsWithJoinByElapsedRange :many
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
    elapsed_ns BETWEEN ? AND ?;
