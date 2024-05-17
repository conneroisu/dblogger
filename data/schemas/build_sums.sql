-- file: build_sums.sql
-- url: github.com/conneroisu/dblogger/data/schemas/build_sums.sql
-- description: build_sums.sql is a table that stores the build sums of the application

CREATE TABLE IF NOT EXISTS build_sums (
    id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
    build_sum TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX build_sums_build_sum_index ON build_sums (build_sum);
