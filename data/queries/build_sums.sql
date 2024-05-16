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
