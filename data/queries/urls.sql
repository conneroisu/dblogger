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
