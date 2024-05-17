package dblogger

import (
	"context"
)

// Querier is the interface for the Queries type
// and is used to simplify the queries interface by
// defining a common interface for a type that can
// be used to run the defined queries.
//
// // vim regex to fix the commenting
// // :%s#//\([a-z]\)#// \1#g
//
// Example:
//
//	db, err := sql.Open("sqlite", "my.db")
//	if err != nil {
//		log.Fatal(err)
//	}
//	q := data.New(db)
//	var count int64
//	err := q.SelectCount(ctx, &count)
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("count: %v\n", count)
type Querier interface {
	// file: build_sums.sql
	// url: github.com/conneroisu/dblogger/data/queries/build_sums.sql
	// description: build_sums.sql is
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      build_sums
	CountBuildSums(ctx context.Context) (int64, error)
	// file: deployments.sql
	// url: github.com/conneroisu/dblogger/data/queries/deployments.sql
	// description: deployments.sql is
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      deployments
	CountDeployments(ctx context.Context) (int64, error)
	// file: git_revisions.sql
	// url: github.com/conneroisu/dblogger/data/queries/git_revisions.sql
	// description: git_revisions.sql is
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      git_revisions
	CountGitRevisions(ctx context.Context) (int64, error)
	// file: go_versions.sql
	// url: github.com/conneroisu/dblogger/data/queries/go_versions.sql
	// description: go_versions.sql is
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      go_versions
	CountGoVersions(ctx context.Context) (int64, error)
	// file: log_levels.sql
	// url: github.com/conneroisu/dblogger/data/queries/log_levels.sql
	// description: log_levels.sql is
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      log_levels
	CountLogLevels(ctx context.Context) (int64, error)
	// file: api_logs.sql
	// url: github.com/conneroisu/dblogger/data/queries/api_logs.sql
	// description: api_logs.sql is
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      api_logs
	CountLogs(ctx context.Context) (int64, error)
	// file: urls.sql
	// url: github.com/conneroisu/dblogger/data/queries/urls.sql
	// description: urls.sql is
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      urls
	CountURLs(ctx context.Context) (int64, error)
	// DeleteBuildSumByID
	//
	//  DELETE FROM
	//      build_sums
	//  WHERE
	//      id = ?
	DeleteBuildSumByID(ctx context.Context, arg *DeleteBuildSumByIDParams) error
	// DeleteDeploymentByID
	//
	//  DELETE FROM
	//      deployments
	//  WHERE
	//      id = ?
	DeleteDeploymentByID(ctx context.Context, arg *DeleteDeploymentByIDParams) error
	// DeleteGitRevisionByID
	//
	//  DELETE FROM
	//      git_revisions
	//  WHERE
	//      id = ?
	DeleteGitRevisionByID(ctx context.Context, arg *DeleteGitRevisionByIDParams) error
	// DeleteGoVersionByID
	//
	//  DELETE FROM
	//      go_versions
	//  WHERE
	//      id = 1
	DeleteGoVersionByID(ctx context.Context) error
	// DeleteLogByID
	//
	//  DELETE FROM
	//      api_logs
	//  WHERE
	//      id = ?
	DeleteLogByID(ctx context.Context, arg *DeleteLogByIDParams) error
	// DeleteLogLevelByID
	//
	//  DELETE FROM
	//      log_levels
	//  WHERE
	//      id = ?
	DeleteLogLevelByID(ctx context.Context, arg *DeleteLogLevelByIDParams) error
	// DeleteURLByID
	//
	//  DELETE FROM
	//      urls
	//  WHERE
	//      id = ?
	DeleteURLByID(ctx context.Context, arg *DeleteURLByIDParams) error
	// GetBuildSumByID
	//
	//  SELECT
	//      id, build_sum, created_at
	//  FROM
	//      build_sums
	//  WHERE
	//      id = ?
	GetBuildSumByID(ctx context.Context, arg *GetBuildSumByIDParams) (BuildSum, error)
	// GetBuildSumsByDate
	//
	//  SELECT
	//      id, build_sum, created_at
	//  FROM
	//      build_sums
	//  WHERE
	//      created_at >= ?
	GetBuildSumsByDate(ctx context.Context, arg *GetBuildSumsByDateParams) ([]BuildSum, error)
	// GetBuildSumsBySubstring
	//
	//  SELECT
	//      id, build_sum, created_at
	//  FROM
	//      build_sums
	//  WHERE
	//      build_sum LIKE ?
	GetBuildSumsBySubstring(ctx context.Context, arg *GetBuildSumsBySubstringParams) ([]BuildSum, error)
	// GetDeploymentByID
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	//  WHERE
	//      id = ?
	GetDeploymentByID(ctx context.Context, arg *GetDeploymentByIDParams) (Deployment, error)
	// GetDeploymentsByDate
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	//  WHERE
	//      created_at >= ?
	GetDeploymentsByDate(ctx context.Context, arg *GetDeploymentsByDateParams) ([]Deployment, error)
	// GetDeploymentsByDateRange
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	//  WHERE
	//      created_at BETWEEN ? AND ?
	GetDeploymentsByDateRange(ctx context.Context, arg *GetDeploymentsByDateRangeParams) ([]Deployment, error)
	// GetDeploymentsBySubstring
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	//  WHERE
	//      name LIKE ?
	GetDeploymentsBySubstring(ctx context.Context, arg *GetDeploymentsBySubstringParams) ([]Deployment, error)
	// GetGitRevisionByID
	//
	//  SELECT
	//      id, git_revision, created_at
	//  FROM
	//      git_revisions
	//  WHERE
	//      id = ?
	GetGitRevisionByID(ctx context.Context, arg *GetGitRevisionByIDParams) (GitRevision, error)
	// GetGitRevisionsByDate
	//
	//  SELECT
	//      id, git_revision, created_at
	//  FROM
	//      git_revisions
	//  WHERE
	//      created_at >= ?
	GetGitRevisionsByDate(ctx context.Context, arg *GetGitRevisionsByDateParams) ([]GitRevision, error)
	// GetGitRevisionsByName
	//
	//  SELECT
	//      id, git_revision, created_at
	//  FROM
	//      git_revisions
	//  WHERE
	//      git_revision = ?
	GetGitRevisionsByName(ctx context.Context, arg *GetGitRevisionsByNameParams) (GitRevision, error)
	// GetGoVersionByID
	//
	//  SELECT
	//      id, name, version, created_at
	//  FROM
	//      go_versions
	//  WHERE
	//      id = ?
	GetGoVersionByID(ctx context.Context, arg *GetGoVersionByIDParams) (GoVersion, error)
	// GetGoVersionIdByName
	//
	//  SELECT
	//      id
	//  FROM
	//      go_versions
	//  WHERE
	//      name = ?
	GetGoVersionIdByName(ctx context.Context, arg *GetGoVersionIdByNameParams) (int64, error)
	// GetGoVersionsByDate
	//
	//  SELECT
	//      id, name, version, created_at
	//  FROM
	//      go_versions
	//  WHERE
	//      created_at >= ?
	GetGoVersionsByDate(ctx context.Context, arg *GetGoVersionsByDateParams) ([]GoVersion, error)
	// GetGoVersionsBySubstring
	//
	//  SELECT
	//      id, name, version, created_at
	//  FROM
	//      go_versions
	//  WHERE
	//      name LIKE ?
	//      OR version LIKE ?
	GetGoVersionsBySubstring(ctx context.Context, arg *GetGoVersionsBySubstringParams) ([]GoVersion, error)
	// GetLogByID
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      id = ?
	GetLogByID(ctx context.Context, arg *GetLogByIDParams) (ApiLog, error)
	// GetLogLevelsBySubstring
	//
	//  SELECT
	//      id, name
	//  FROM
	//      log_levels
	//  WHERE
	//      name LIKE ?
	GetLogLevelsBySubstring(ctx context.Context, arg *GetLogLevelsBySubstringParams) ([]LogLevel, error)
	// GetLogsByBuildSumID
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      build_sum_id = ?
	GetLogsByBuildSumID(ctx context.Context, arg *GetLogsByBuildSumIDParams) ([]ApiLog, error)
	// GetLogsByDate
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      created_at >= ?
	GetLogsByDate(ctx context.Context, arg *GetLogsByDateParams) ([]ApiLog, error)
	// GetLogsByDateRange
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      created_at BETWEEN ? AND ?
	GetLogsByDateRange(ctx context.Context, arg *GetLogsByDateRangeParams) ([]ApiLog, error)
	// GetLogsByElapsedRange
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      elapsed_ns BETWEEN ? AND ?
	GetLogsByElapsedRange(ctx context.Context, arg *GetLogsByElapsedRangeParams) ([]ApiLog, error)
	// GetLogsByGitRevisionID
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      git_revision_id = ?
	GetLogsByGitRevisionID(ctx context.Context, arg *GetLogsByGitRevisionIDParams) ([]ApiLog, error)
	// GetLogsByGoVersionID
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      go_version_id = ?
	GetLogsByGoVersionID(ctx context.Context, arg *GetLogsByGoVersionIDParams) ([]ApiLog, error)
	// GetLogsByURLSubstring
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      url LIKE ?
	GetLogsByURLSubstring(ctx context.Context, arg *GetLogsByURLSubstringParams) ([]ApiLog, error)
	// GetURLByID
	//
	//  SELECT
	//      id, url, created_at
	//  FROM
	//      urls
	//  WHERE
	//      id = ?
	GetURLByID(ctx context.Context, arg *GetURLByIDParams) (Url, error)
	// GetURLsByDate
	//
	//  SELECT
	//      id, url, created_at
	//  FROM
	//      urls
	//  WHERE
	//      created_at >= ?
	GetURLsByDate(ctx context.Context, arg *GetURLsByDateParams) ([]Url, error)
	// GetURLsBySubstring
	//
	//  SELECT
	//      id, url, created_at
	//  FROM
	//      urls
	//  WHERE
	//      url LIKE ?
	GetURLsBySubstring(ctx context.Context, arg *GetURLsBySubstringParams) ([]Url, error)
	// InsertBuildSum
	//
	//  INSERT INTO
	//      build_sums (build_sum)
	//  VALUES
	//      (?) RETURNING id, build_sum, created_at
	InsertBuildSum(ctx context.Context, arg *InsertBuildSumParams) (BuildSum, error)
	// InsertBuildSumReturningID
	//
	//  INSERT INTO
	//      build_sums (build_sum)
	//  VALUES
	//      (?) RETURNING id
	InsertBuildSumReturningID(ctx context.Context, arg *InsertBuildSumReturningIDParams) (int64, error)
	// InsertBuildSumWithParam
	//
	//  INSERT INTO
	//      build_sums (build_sum)
	//  VALUES
	//      (?)
	InsertBuildSumWithParam(ctx context.Context, arg *InsertBuildSumWithParamParams) error
	// InsertDeployment
	//
	//  INSERT INTO
	//      deployments (name)
	//  VALUES
	//      (?) RETURNING id, name, created_at, updated_at
	InsertDeployment(ctx context.Context, arg *InsertDeploymentParams) (Deployment, error)
	// InsertDeploymentReturningID
	//
	//  INSERT INTO
	//      deployments (name)
	//  VALUES
	//      (?) RETURNING id
	InsertDeploymentReturningID(ctx context.Context, arg *InsertDeploymentReturningIDParams) (int64, error)
	// InsertGitRevision
	//
	//  INSERT INTO
	//      git_revisions (git_revision)
	//  VALUES
	//      (?) RETURNING id, git_revision, created_at
	InsertGitRevision(ctx context.Context, arg *InsertGitRevisionParams) (GitRevision, error)
	// InsertGitRevisionReturningID
	//
	//  INSERT INTO
	//      git_revisions (git_revision)
	//  VALUES
	//      (?) RETURNING id
	InsertGitRevisionReturningID(ctx context.Context, arg *InsertGitRevisionReturningIDParams) (int64, error)
	// InsertGitRevisionWithParam
	//
	//  INSERT INTO
	//      git_revisions (git_revision)
	//  VALUES
	//      (?)
	InsertGitRevisionWithParam(ctx context.Context, arg *InsertGitRevisionWithParamParams) error
	// InsertGoVersion
	//
	//  INSERT INTO
	//      go_versions (name, version)
	//  VALUES
	//      (?, ?) RETURNING id, name, version, created_at
	InsertGoVersion(ctx context.Context, arg *InsertGoVersionParams) (GoVersion, error)
	// InsertGoVersionReturningID
	//
	//  INSERT INTO
	//      go_versions (name, version)
	//  VALUES
	//      (?, ?) RETURNING id
	InsertGoVersionReturningID(ctx context.Context, arg *InsertGoVersionReturningIDParams) (int64, error)
	// InsertLogEntry
	//
	//  INSERT INTO
	//      api_logs (
	//          go_version_id,
	//          build_sum_id,
	//          git_revision_id,
	//          user_agent,
	//          method,
	//          url,
	//          elapsed_ns,
	//          deployment_id,
	//          level_id
	//      )
	//  VALUES
	//      (?, ?, ?, ?, ?, ?, ?, ?, ?)
	InsertLogEntry(ctx context.Context, arg *InsertLogEntryParams) error
	// InsertLogLevel
	//
	//  INSERT INTO
	//      log_levels (name)
	//  VALUES
	//      (?)
	InsertLogLevel(ctx context.Context, arg *InsertLogLevelParams) error
	// InsertURL
	//
	//  INSERT INTO
	//      urls (url)
	//  VALUES
	//      (?)
	InsertURL(ctx context.Context, arg *InsertURLParams) error
	// InsertURLWithParam
	//
	//  INSERT INTO
	//      urls (url)
	//  VALUES
	//      (?)
	InsertURLWithParam(ctx context.Context, arg *InsertURLWithParamParams) error
	// ListBuildSums
	//
	//  SELECT
	//      id, build_sum, created_at
	//  FROM
	//      build_sums
	ListBuildSums(ctx context.Context) ([]BuildSum, error)
	// ListBuildSumsPaginated
	//
	//  SELECT
	//      id, build_sum, created_at
	//  FROM
	//      build_sums
	//  LIMIT
	//      ? OFFSET ?
	ListBuildSumsPaginated(ctx context.Context, arg *ListBuildSumsPaginatedParams) ([]BuildSum, error)
	// ListDeployments
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	ListDeployments(ctx context.Context) ([]Deployment, error)
	// ListDeploymentsPaginated
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	//  LIMIT
	//      ? OFFSET ?
	ListDeploymentsPaginated(ctx context.Context, arg *ListDeploymentsPaginatedParams) ([]Deployment, error)
	// ListGitRevisions
	//
	//  SELECT
	//      id, git_revision, created_at
	//  FROM
	//      git_revisions
	ListGitRevisions(ctx context.Context) ([]GitRevision, error)
	// ListGitRevisionsPaginated
	//
	//  SELECT
	//      id, git_revision, created_at
	//  FROM
	//      git_revisions
	//  LIMIT
	//      ? OFFSET ?
	ListGitRevisionsPaginated(ctx context.Context, arg *ListGitRevisionsPaginatedParams) ([]GitRevision, error)
	// ListGoVersions
	//
	//  SELECT
	//      id, name, version, created_at
	//  FROM
	//      go_versions
	ListGoVersions(ctx context.Context) ([]GoVersion, error)
	// ListGoVersionsPaginated
	//
	//  SELECT
	//      id, name, version, created_at
	//  FROM
	//      go_versions
	//  LIMIT
	//      ? OFFSET ?
	ListGoVersionsPaginated(ctx context.Context, arg *ListGoVersionsPaginatedParams) ([]GoVersion, error)
	// ListLogLevels
	//
	//  SELECT
	//      id, name
	//  FROM
	//      log_levels
	ListLogLevels(ctx context.Context) ([]LogLevel, error)
	// ListLogLevelsPaginated
	//
	//  SELECT
	//      id, name
	//  FROM
	//      log_levels
	//  LIMIT
	//      ? OFFSET ?
	ListLogLevelsPaginated(ctx context.Context, arg *ListLogLevelsPaginatedParams) ([]LogLevel, error)
	// ListLogs
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	ListLogs(ctx context.Context) ([]ApiLog, error)
	// ListLogsByDeploymentID
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      deployment_id = ?
	ListLogsByDeploymentID(ctx context.Context, arg *ListLogsByDeploymentIDParams) ([]ApiLog, error)
	// ListLogsByMethod
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      method = ?
	ListLogsByMethod(ctx context.Context, arg *ListLogsByMethodParams) ([]ApiLog, error)
	// ListLogsByUserAgent
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      user_agent = ?
	ListLogsByUserAgent(ctx context.Context, arg *ListLogsByUserAgentParams) ([]ApiLog, error)
	// ListLogsPaginated
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ns, deployment_id
	//  FROM
	//      api_logs
	//  LIMIT
	//      ? OFFSET ?
	ListLogsPaginated(ctx context.Context, arg *ListLogsPaginatedParams) ([]ApiLog, error)
	// ListLogsWithJoin
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	ListLogsWithJoin(ctx context.Context) ([]ListLogsWithJoinRow, error)
	// ListLogsWithJoinByBuildSumID
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      bs.id = ?
	ListLogsWithJoinByBuildSumID(ctx context.Context, arg *ListLogsWithJoinByBuildSumIDParams) ([]ListLogsWithJoinByBuildSumIDRow, error)
	// ListLogsWithJoinByBuildSumIDPaginated
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      bs.id = ?
	ListLogsWithJoinByBuildSumIDPaginated(ctx context.Context, arg *ListLogsWithJoinByBuildSumIDPaginatedParams) ([]ListLogsWithJoinByBuildSumIDPaginatedRow, error)
	// ListLogsWithJoinByDate
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      created_at BETWEEN ? AND ?
	ListLogsWithJoinByDate(ctx context.Context, arg *ListLogsWithJoinByDateParams) ([]ListLogsWithJoinByDateRow, error)
	// ListLogsWithJoinByDatePaginated
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      created_at BETWEEN ? AND ?
	//  LIMIT
	//      ? OFFSET ?
	ListLogsWithJoinByDatePaginated(ctx context.Context, arg *ListLogsWithJoinByDatePaginatedParams) ([]ListLogsWithJoinByDatePaginatedRow, error)
	// ListLogsWithJoinByDateRange
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      created_at BETWEEN ? AND ?
	ListLogsWithJoinByDateRange(ctx context.Context, arg *ListLogsWithJoinByDateRangeParams) ([]ListLogsWithJoinByDateRangeRow, error)
	// ListLogsWithJoinByDateRangePaginated
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      created_at BETWEEN ? AND ?
	//  LIMIT
	//      ? OFFSET ?
	ListLogsWithJoinByDateRangePaginated(ctx context.Context, arg *ListLogsWithJoinByDateRangePaginatedParams) ([]ListLogsWithJoinByDateRangePaginatedRow, error)
	// ListLogsWithJoinByElapsedRange
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      elapsed_ns BETWEEN ? AND ?
	ListLogsWithJoinByElapsedRange(ctx context.Context, arg *ListLogsWithJoinByElapsedRangeParams) ([]ListLogsWithJoinByElapsedRangeRow, error)
	// ListLogsWithJoinByElapsedRangePaginated
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      elapsed_ns BETWEEN ? AND ?
	//  LIMIT
	//      ? OFFSET ?
	ListLogsWithJoinByElapsedRangePaginated(ctx context.Context, arg *ListLogsWithJoinByElapsedRangePaginatedParams) ([]ListLogsWithJoinByElapsedRangePaginatedRow, error)
	// ListLogsWithJoinByGitRevisionID
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      gr.id = ?
	ListLogsWithJoinByGitRevisionID(ctx context.Context, arg *ListLogsWithJoinByGitRevisionIDParams) ([]ListLogsWithJoinByGitRevisionIDRow, error)
	// ListLogsWithJoinByGitRevisionIDPaginated
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      gr.id = ?
	ListLogsWithJoinByGitRevisionIDPaginated(ctx context.Context, arg *ListLogsWithJoinByGitRevisionIDPaginatedParams) ([]ListLogsWithJoinByGitRevisionIDPaginatedRow, error)
	// ListLogsWithJoinByGoVersionID
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      gv.id = ?
	ListLogsWithJoinByGoVersionID(ctx context.Context, arg *ListLogsWithJoinByGoVersionIDParams) ([]ListLogsWithJoinByGoVersionIDRow, error)
	// ListLogsWithJoinByGoVersionIDPaginated
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  WHERE
	//      gv.id = ?
	ListLogsWithJoinByGoVersionIDPaginated(ctx context.Context, arg *ListLogsWithJoinByGoVersionIDPaginatedParams) ([]ListLogsWithJoinByGoVersionIDPaginatedRow, error)
	// ListLogsWithJoinPaginated
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ns,
	//      l.level_id,
	//      le.name AS level_name,
	//      gv.name AS go_version_name,
	//      gv.version AS go_version,
	//      bs.build_sum,
	//      gr.git_revision
	//  FROM
	//      api_logs l
	//      JOIN log_levels le ON l.level_id = le.id
	//      JOIN deployments d ON l.deployment_id = d.id
	//      JOIN go_versions gv ON l.go_version_id = gv.id
	//      JOIN build_sums bs ON l.build_sum_id = bs.id
	//      JOIN git_revisions gr ON l.git_revision_id = gr.id
	//  LIMIT
	//      ? OFFSET ?
	ListLogsWithJoinPaginated(ctx context.Context, arg *ListLogsWithJoinPaginatedParams) ([]ListLogsWithJoinPaginatedRow, error)
	// ListURLs
	//
	//  SELECT
	//      id, url, created_at
	//  FROM
	//      urls
	ListURLs(ctx context.Context) ([]Url, error)
	// ListURLsPaginated
	//
	//  SELECT
	//      id, url, created_at
	//  FROM
	//      urls
	//  LIMIT
	//      ? OFFSET ?
	ListURLsPaginated(ctx context.Context, arg *ListURLsPaginatedParams) ([]Url, error)
	// UpdateGoVersionByID
	//
	//  UPDATE
	//      go_versions
	//  SET
	//      name = ?,
	//      version = ?
	//  WHERE
	//      id = 1
	UpdateGoVersionByID(ctx context.Context, arg *UpdateGoVersionByIDParams) error
}

var _ Querier = (*Queries)(nil)
