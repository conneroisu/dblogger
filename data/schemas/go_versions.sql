-- file: go_versions.sql
-- url: github.com/conneroisu/dblogger/data/schemas/go_versions.sql
-- description: go_versions.sql is a table that stores the versions of 
--              the go compiler used to build the application.

CREATE TABLE IF NOT EXISTS go_versions (
    id INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    version TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);
