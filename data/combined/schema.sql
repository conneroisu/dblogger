-- file: build_sums.sql
-- url: github.com/conneroisu/dblogger/data/schemas/build_sums.sql
-- description: build_sums.sql is a table that stores the build sums of the application
CREATE TABLE IF NOT EXISTS build_sums (
    id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
    build_sum TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX build_sums_build_sum_index ON build_sums (build_sum);

-- file: deployments.sql
-- url: github.com/conneroisu/dblogger/data/schemas/deployments.sql
-- description: deployments.sql is a table that stores deployment information
CREATE TABLE deployments (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- file: git_revisions.sql
-- url: github.com/conneroisu/dblogger/data/schemas/git_revisions.sql
-- description: git_revisions.sql is a table that stores git revisions
CREATE TABLE git_revisions (
    id INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT,
    git_revision TEXT UNIQUE NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX git_revisions_index ON git_revisions (created_at);

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

-- file: log_levels.sql
-- url: github.com/conneroisu/dblogger/data/schemas/log_levels.sql
-- description: log_levels.sql is a table of log levels
CREATE TABLE log_levels (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- file: logs.sql
-- url: github.com/conneroisu/dblogger/data/schemas/logs.sql
-- description: logs.sql is a sqlite schema for the logs table
CREATE TABLE IF NOT EXISTS api_logs (
    id INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT,
    level_id INTEGER NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    go_version_id INTEGER NOT NULL,
    build_sum_id INTEGER NOT NULL,
    git_revision_id INTEGER NOT NULL,
    user_agent TEXT NOT NULL,
    method TEXT NOT NULL,
    url TEXT NOT NULL,
    elapsed_ns INTEGER NOT NULL,
    deployment_id INTEGER NOT NULL,
    FOREIGN KEY (level_id) REFERENCES log_levels (id),
    FOREIGN KEY (go_version_id) REFERENCES go_versions (id),
    FOREIGN KEY (build_sum_id) REFERENCES build_sums (id),
    FOREIGN KEY (git_revision_id) REFERENCES git_revisions (id),
    FOREIGN KEY (deployment_id) REFERENCES deployments (id)
);

CREATE INDEX IF NOT EXISTS api_logs_level_id_index ON api_logs (level_id);

CREATE INDEX IF NOT EXISTS api_logs_go_version_id_index ON api_logs (go_version_id);

CREATE INDEX IF NOT EXISTS api_logs_build_sum_id_index ON api_logs (build_sum_id);

CREATE INDEX IF NOT EXISTS api_logs_git_revision_id_index ON api_logs (git_revision_id);

CREATE INDEX IF NOT EXISTS api_logs_deployment_id_index ON api_logs (deployment_id);

-- file: urls.sql
-- url: github.com/conneroisu/dblogger/data/schemas/urls.sql
-- description: urls.sql is the schema for the urls table
CREATE TABLE urls (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- file: log_levels.sql
-- url: github.com/conneroisu/dblogger/data/seeds/log_levels.sql
-- description: log_levels.sql is
INSERT INTO
    log_levels (name)
VALUES
    ('debug'),
    ('info'),
    ('warn'),
    ('error'),
    ('fatal');
