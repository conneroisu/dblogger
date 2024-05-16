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
