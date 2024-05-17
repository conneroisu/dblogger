# dblogger

[![Go Test](https://github.com/conneroisu/dblogger/actions/workflows/test.yaml/badge.svg)](https://github.com/conneroisu/dblogger/actions/workflows/test.yaml)
[![golangci-lint](https://github.com/conneroisu/dblogger/actions/workflows/ci.yaml/badge.svg)](https://github.com/conneroisu/dblogger/actions/workflows/ci.yaml)

Zerolog Log Writer for Sqlite, Turso, RqLite, and other sqlite-like dialected databases.

A zerolog-based logger for the sqlite3 dialect.

## Installation

```bash
go get github.com/conneroisu/dblogger
```

## Usage

## Development

### SQL

Each sql file 

```sql
-- file: {file_name}
-- url: github.com/conneroisu/dblogger/data/{file_path}.sql
-- description: {description}
```

### Makefile

```bash
make dev.requirements
make database
make lint
make test
```

### Testing

```bash
make test
```

### Linting

```bash
make lint
```

### Database

Combine all the sql files (shcemas, queries, seeds) into a single sql file each. (ex: [`data/combined/schema.sql`](https://github.com/conneroisu/dblogger/blob/main/data/combined/schema.sql))
```bash
make database
```
