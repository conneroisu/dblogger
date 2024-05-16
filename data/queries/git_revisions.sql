-- name: CountGitRevisions :one
SELECT
    COUNT(*)
FROM
    git_revisions;

-- name: InsertGitRevision :exec
INSERT INTO
    git_revisions (git_revision)
VALUES
    (?);

-- name: ListGitRevisions :many
SELECT
    *
FROM
    git_revisions;

-- name: DeleteGitRevisionByID :exec
DELETE FROM
    git_revisions
WHERE
    id = ?;

-- name: InsertGitRevisionWithParam :exec
INSERT INTO
    git_revisions (git_revision)
VALUES
    (?);

-- name: GetGitRevisionByID :one
SELECT
    *
FROM
    git_revisions
WHERE
    id = ?;

-- name: GetGitRevisionsByDate :many
SELECT
    *
FROM
    git_revisions
WHERE
    created_at >= ?;

-- name: GetGitRevisionsBySubstring :many
SELECT
    *
FROM
    git_revisions
WHERE
    git_revision LIKE ?;
