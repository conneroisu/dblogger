-- file: urls.sql
-- url: github.com/conneroisu/dblogger/data/schemas/urls.sql
-- description: urls.sql is the schema for the urls table

CREATE TABLE IF NOT EXISTS urls (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);
