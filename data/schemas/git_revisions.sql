-- file: git_revisions.sql
-- url: github.com/conneroisu/dblogger/data/schemas/git_revisions.sql
-- description: git_revisions.sql is a table that stores git revisions

CREATE TABLE IF NOT EXISTS git_revisions (
    id INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT,
    git_revision TEXT UNIQUE NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX git_revisions_index ON git_revisions (created_at);
