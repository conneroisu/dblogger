CREATE TABLE urls (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);
