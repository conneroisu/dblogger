-- file: urls.sql
-- url: github.com/conneroisu/dblogger/data/queries/urls.sql
-- description: urls.sql is

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
