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