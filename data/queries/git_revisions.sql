-- file: git_revisions.sql
-- url: github.com/conneroisu/dblogger/data/queries/git_revisions.sql
-- description: git_revisions.sql is

-- name: CountGitRevisions :one
SELECT
    COUNT(*)
FROM
    git_revisions;

-- name: InsertGitRevision :one
INSERT INTO
    git_revisions (git_revision)
VALUES
    (?) RETURNING *;

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

-- name: GetGitRevisionsByName :one
SELECT
    *
FROM
    git_revisions
WHERE
    git_revision = ?;

-- name: ListGitRevisionsPaginated :many
SELECT
    *
FROM
    git_revisions
LIMIT
    ? OFFSET ?;

-- name: InsertGitRevisionReturningID :one
INSERT INTO
    git_revisions (git_revision)
VALUES
    (?) RETURNING id;
