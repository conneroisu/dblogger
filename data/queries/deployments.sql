-- name: CountDeployments :one
SELECT
    COUNT(*)
FROM
    deployments;

-- name: InsertDeployment :exec
INSERT INTO
    deployments (name)
VALUES
    (?);

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
