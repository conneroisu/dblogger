-- file: deployments.sql
-- url: github.com/conneroisu/dblogger/data/schemas/deployments.sql
-- description: deployments.sql is a table that stores deployment information

CREATE TABLE IF NOT EXISTS deployments (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);
