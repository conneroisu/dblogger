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
