-- name: CountDeployments :one
SELECT
    COUNT(*)
FROM
    deployments;

-- name: InsertDeployment :one
INSERT INTO
    deployments (name)
VALUES
    (?) RETURNING *;

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

-- name: GetDeploymentsBySubstring :many
SELECT
    *
FROM
    deployments
WHERE
    name LIKE ?;

-- name: ListDeploymentsPaginated :many
SELECT
    *
FROM
    deployments
LIMIT
    ? OFFSET ?;
