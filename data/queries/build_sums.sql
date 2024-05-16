-- name: CountBuildSums :one
SELECT
    COUNT(*)
FROM
    build_sums;

-- name: InsertBuildSum :one
INSERT INTO
    build_sums (build_sum)
VALUES
    (?) RETURNING *;

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

-- name: ListBuildSumsPaginated :many
SELECT
    *
FROM
    build_sums
LIMIT
    ? OFFSET ?;
