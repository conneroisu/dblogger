CREATE TABLE IF NOT EXISTS api_logs (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    level_id INTEGER NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    go_version_id INTEGER NOT NULL,
    build_sum_id INTEGER NOT NULL,
    git_revision_id INTEGER NOT NULL,
    user_agent TEXT NOT NULL,
    method TEXT NOT NULL,
    url TEXT NOT NULL,
    elapsed_ms INTEGER NOT NULL,
    status_code INTEGER NOT NULL,
    deployment_id INTEGER NOT NULL,
    FOREIGN KEY (level_id) REFERENCES log_levels (id),
    FOREIGN KEY (go_version_id) REFERENCES go_versions (id),
    FOREIGN KEY (build_sum_id) REFERENCES build_sums (id),
    FOREIGN KEY (git_revision_id) REFERENCES git_revisions (id),
    FOREIGN KEY (deployment_id) REFERENCES deployments (id)
);
