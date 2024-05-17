-- file: log_levels.sql
-- url: github.com/conneroisu/dblogger/data/schemas/log_levels.sql
-- description: log_levels.sql is a table of log levels

CREATE TABLE IF NOT EXISTS log_levels (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);
