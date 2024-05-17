-- file: log_levels.sql
-- url: github.com/conneroisu/dblogger/data/queries/log_levels.sql
-- description: log_levels.sql is

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
