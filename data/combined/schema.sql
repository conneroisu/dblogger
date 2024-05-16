CREATE TABLE IF NOT EXISTS build_sums (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    build_sum TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE deployments (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE git_revisions (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    git_revision TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS go_versions (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    version TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE log_levels (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

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
    elapsed_ns INTEGER NOT NULL,
    deployment_id INTEGER NOT NULL,
    FOREIGN KEY (level_id) REFERENCES log_levels (id),
    FOREIGN KEY (go_version_id) REFERENCES go_versions (id),
    FOREIGN KEY (build_sum_id) REFERENCES build_sums (id),
    FOREIGN KEY (git_revision_id) REFERENCES git_revisions (id),
    FOREIGN KEY (deployment_id) REFERENCES deployments (id)
);

CREATE TABLE urls (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    log_levels (name)
VALUES
    ('debug'),
    ('info'),
    ('warn'),
    ('error'),
    ('fatal');