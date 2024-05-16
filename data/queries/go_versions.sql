-- name: CountGoVersions :one
SELECT
    COUNT(*)
FROM
    go_versions;

-- name: InsertGoVersion :one
INSERT INTO
    go_versions (name, version)
VALUES
    (?, ?) RETURNING *;

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

-- name: GetGoVersionIdByName :one
SELECT
    id
FROM
    go_versions
WHERE
    name = ?;
