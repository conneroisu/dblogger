package dblogger

import (
	"context"
)

// Querier is the interface for the Queries type
// and is used to simplify the queries interface by
// defining a common interface for a type that can
// be used to run the defined queries.
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
//
//	// use the count
type Querier interface {
	//CountBuildSums
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      build_sums
	CountBuildSums(ctx context.Context) (int64, error)
	//CountDeployments
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      deployments
	CountDeployments(ctx context.Context) (int64, error)
	//CountGitRevisions
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      git_revisions
	CountGitRevisions(ctx context.Context) (int64, error)
	//CountGoVersions
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      go_versions
	CountGoVersions(ctx context.Context) (int64, error)
	//CountLogs
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      api_logs
	CountLogs(ctx context.Context) (int64, error)
	//CountURLs
	//
	//  SELECT
	//      COUNT(*)
	//  FROM
	//      urls
	CountURLs(ctx context.Context) (int64, error)
	//DeleteBuildSumByID
	//
	//  DELETE FROM
	//      build_sums
	//  WHERE
	//      id = ?
	DeleteBuildSumByID(ctx context.Context, arg DeleteBuildSumByIDParams) error
	//DeleteDeploymentByID
	//
	//  DELETE FROM
	//      deployments
	//  WHERE
	//      id = ?
	DeleteDeploymentByID(ctx context.Context, arg DeleteDeploymentByIDParams) error
	//DeleteGitRevisionByID
	//
	//  DELETE FROM
	//      git_revisions
	//  WHERE
	//      id = ?
	DeleteGitRevisionByID(ctx context.Context, arg DeleteGitRevisionByIDParams) error
	//DeleteGoVersionByID
	//
	//  DELETE FROM
	//      go_versions
	//  WHERE
	//      id = 1
	DeleteGoVersionByID(ctx context.Context) error
	//DeleteLogByID
	//
	//  DELETE FROM
	//      api_logs
	//  WHERE
	//      id = ?
	DeleteLogByID(ctx context.Context, arg DeleteLogByIDParams) error
	//DeleteURLByID
	//
	//  DELETE FROM
	//      urls
	//  WHERE
	//      id = ?
	DeleteURLByID(ctx context.Context, arg DeleteURLByIDParams) error
	//GetBuildSumByID
	//
	//  SELECT
	//      id, build_sum, created_at
	//  FROM
	//      build_sums
	//  WHERE
	//      id = ?
	GetBuildSumByID(ctx context.Context, arg GetBuildSumByIDParams) (BuildSum, error)
	//GetBuildSumsByDate
	//
	//  SELECT
	//      id, build_sum, created_at
	//  FROM
	//      build_sums
	//  WHERE
	//      created_at >= ?
	GetBuildSumsByDate(ctx context.Context, arg GetBuildSumsByDateParams) ([]BuildSum, error)
	//GetBuildSumsBySubstring
	//
	//  SELECT
	//      id, build_sum, created_at
	//  FROM
	//      build_sums
	//  WHERE
	//      build_sum LIKE ?
	GetBuildSumsBySubstring(ctx context.Context, arg GetBuildSumsBySubstringParams) ([]BuildSum, error)
	//GetDeploymentByID
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	//  WHERE
	//      id = ?
	GetDeploymentByID(ctx context.Context, arg GetDeploymentByIDParams) (Deployment, error)
	//GetDeploymentsByDate
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	//  WHERE
	//      created_at >= ?
	GetDeploymentsByDate(ctx context.Context, arg GetDeploymentsByDateParams) ([]Deployment, error)
	//GetDeploymentsByDateRange
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	//  WHERE
	//      created_at BETWEEN ? AND ?
	GetDeploymentsByDateRange(ctx context.Context, arg GetDeploymentsByDateRangeParams) ([]Deployment, error)
	//GetDeploymentsBySubstring
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	//  WHERE
	//      name LIKE ?
	GetDeploymentsBySubstring(ctx context.Context, arg GetDeploymentsBySubstringParams) ([]Deployment, error)
	//GetGitRevisionByID
	//
	//  SELECT
	//      id, git_revision, created_at
	//  FROM
	//      git_revisions
	//  WHERE
	//      id = ?
	GetGitRevisionByID(ctx context.Context, arg GetGitRevisionByIDParams) (GitRevision, error)
	//GetGitRevisionsByDate
	//
	//  SELECT
	//      id, git_revision, created_at
	//  FROM
	//      git_revisions
	//  WHERE
	//      created_at >= ?
	GetGitRevisionsByDate(ctx context.Context, arg GetGitRevisionsByDateParams) ([]GitRevision, error)
	//GetGitRevisionsBySubstring
	//
	//  SELECT
	//      id, git_revision, created_at
	//  FROM
	//      git_revisions
	//  WHERE
	//      git_revision LIKE ?
	GetGitRevisionsBySubstring(ctx context.Context, arg GetGitRevisionsBySubstringParams) ([]GitRevision, error)
	//GetGoVersionByID
	//
	//  SELECT
	//      id, name, version, created_at
	//  FROM
	//      go_versions
	//  WHERE
	//      id = ?
	GetGoVersionByID(ctx context.Context, arg GetGoVersionByIDParams) (GoVersion, error)
	//GetGoVersionsByDate
	//
	//  SELECT
	//      id, name, version, created_at
	//  FROM
	//      go_versions
	//  WHERE
	//      created_at >= ?
	GetGoVersionsByDate(ctx context.Context, arg GetGoVersionsByDateParams) ([]GoVersion, error)
	//GetGoVersionsBySubstring
	//
	//  SELECT
	//      id, name, version, created_at
	//  FROM
	//      go_versions
	//  WHERE
	//      name LIKE ?
	//      OR version LIKE ?
	GetGoVersionsBySubstring(ctx context.Context, arg GetGoVersionsBySubstringParams) ([]GoVersion, error)
	//GetLogByID
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      id = ?
	GetLogByID(ctx context.Context, arg GetLogByIDParams) (ApiLog, error)
	//GetLogsByBuildSumID
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      build_sum_id = ?
	GetLogsByBuildSumID(ctx context.Context, arg GetLogsByBuildSumIDParams) ([]ApiLog, error)
	//GetLogsByDate
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      created_at >= ?
	GetLogsByDate(ctx context.Context, arg GetLogsByDateParams) ([]ApiLog, error)
	//GetLogsByDateRange
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      created_at BETWEEN ? AND ?
	GetLogsByDateRange(ctx context.Context, arg GetLogsByDateRangeParams) ([]ApiLog, error)
	//GetLogsByDeploymentIDAndStatus
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      deployment_id = ?
	//      AND status_code = ?
	GetLogsByDeploymentIDAndStatus(ctx context.Context, arg GetLogsByDeploymentIDAndStatusParams) ([]ApiLog, error)
	//GetLogsByElapsedRange
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      elapsed_ms BETWEEN ? AND ?
	GetLogsByElapsedRange(ctx context.Context, arg GetLogsByElapsedRangeParams) ([]ApiLog, error)
	//GetLogsByGitRevisionID
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      git_revision_id = ?
	GetLogsByGitRevisionID(ctx context.Context, arg GetLogsByGitRevisionIDParams) ([]ApiLog, error)
	//GetLogsByGoVersionID
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      go_version_id = ?
	GetLogsByGoVersionID(ctx context.Context, arg GetLogsByGoVersionIDParams) ([]ApiLog, error)
	//GetLogsByURLSubstring
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      url LIKE ?
	GetLogsByURLSubstring(ctx context.Context, arg GetLogsByURLSubstringParams) ([]ApiLog, error)
	//GetURLByID
	//
	//  SELECT
	//      id, url, created_at
	//  FROM
	//      urls
	//  WHERE
	//      id = ?
	GetURLByID(ctx context.Context, arg GetURLByIDParams) (Url, error)
	//GetURLsByDate
	//
	//  SELECT
	//      id, url, created_at
	//  FROM
	//      urls
	//  WHERE
	//      created_at >= ?
	GetURLsByDate(ctx context.Context, arg GetURLsByDateParams) ([]Url, error)
	//GetURLsBySubstring
	//
	//  SELECT
	//      id, url, created_at
	//  FROM
	//      urls
	//  WHERE
	//      url LIKE ?
	GetURLsBySubstring(ctx context.Context, arg GetURLsBySubstringParams) ([]Url, error)
	//InsertBuildSum
	//
	//  INSERT INTO
	//      build_sums (build_sum)
	//  VALUES
	//      (?)
	InsertBuildSum(ctx context.Context, arg InsertBuildSumParams) error
	//InsertBuildSumWithParam
	//
	//  INSERT INTO
	//      build_sums (build_sum)
	//  VALUES
	//      (?)
	InsertBuildSumWithParam(ctx context.Context, arg InsertBuildSumWithParamParams) error
	//InsertDeployment
	//
	//  INSERT INTO
	//      deployments (name)
	//  VALUES
	//      (?)
	InsertDeployment(ctx context.Context, arg InsertDeploymentParams) error
	//InsertDeploymentWithParam
	//
	//  INSERT INTO
	//      deployments (name)
	//  VALUES
	//      (?)
	InsertDeploymentWithParam(ctx context.Context, arg InsertDeploymentWithParamParams) error
	//InsertGitRevision
	//
	//  INSERT INTO
	//      git_revisions (git_revision)
	//  VALUES
	//      (?)
	InsertGitRevision(ctx context.Context, arg InsertGitRevisionParams) error
	//InsertGitRevisionWithParam
	//
	//  INSERT INTO
	//      git_revisions (git_revision)
	//  VALUES
	//      (?)
	InsertGitRevisionWithParam(ctx context.Context, arg InsertGitRevisionWithParamParams) error
	//InsertGoVersion
	//
	//  INSERT INTO
	//      go_versions (name, version)
	//  VALUES
	//      (?, ?)
	InsertGoVersion(ctx context.Context, arg InsertGoVersionParams) error
	//InsertGoVersionWithParams
	//
	//  INSERT INTO
	//      go_versions (name, version)
	//  VALUES
	//      (?, ?)
	InsertGoVersionWithParams(ctx context.Context, arg InsertGoVersionWithParamsParams) error
	//InsertLogEntry
	//
	//  INSERT INTO
	//      api_logs (
	//          go_version_id,
	//          build_sum_id,
	//          git_revision_id,
	//          user_agent,
	//          method,
	//          url,
	//          elapsed_ms,
	//          status_code,
	//          deployment_id,
	//          level_id
	//      )
	//  VALUES
	//      (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	InsertLogEntry(ctx context.Context, arg InsertLogEntryParams) error
	//InsertLogWithParams
	//
	//  INSERT INTO
	//      api_logs (
	//          level_id,
	//          go_version_id,
	//          build_sum_id,
	//          git_revision_id,
	//          user_agent,
	//          method,
	//          url,
	//          elapsed_ms,
	//          status_code,
	//          deployment_id
	//      )
	//  VALUES
	//      (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	InsertLogWithParams(ctx context.Context, arg InsertLogWithParamsParams) error
	//InsertURL
	//
	//  INSERT INTO
	//      urls (url)
	//  VALUES
	//      (?)
	InsertURL(ctx context.Context, arg InsertURLParams) error
	//InsertURLWithParam
	//
	//  INSERT INTO
	//      urls (url)
	//  VALUES
	//      (?)
	InsertURLWithParam(ctx context.Context, arg InsertURLWithParamParams) error
	//ListBuildSums
	//
	//  SELECT
	//      id, build_sum, created_at
	//  FROM
	//      build_sums
	ListBuildSums(ctx context.Context) ([]BuildSum, error)
	//ListBuildSumsPaginated
	//
	//  SELECT
	//      id, build_sum, created_at
	//  FROM
	//      build_sums
	//  LIMIT
	//      ? OFFSET ?
	ListBuildSumsPaginated(ctx context.Context, arg ListBuildSumsPaginatedParams) ([]BuildSum, error)
	//ListDeployments
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	ListDeployments(ctx context.Context) ([]Deployment, error)
	//ListDeploymentsPaginated
	//
	//  SELECT
	//      id, name, created_at, updated_at
	//  FROM
	//      deployments
	//  LIMIT
	//      ? OFFSET ?
	ListDeploymentsPaginated(ctx context.Context, arg ListDeploymentsPaginatedParams) ([]Deployment, error)
	//ListGitRevisions
	//
	//  SELECT
	//      id, git_revision, created_at
	//  FROM
	//      git_revisions
	ListGitRevisions(ctx context.Context) ([]GitRevision, error)
	//ListGitRevisionsPaginated
	//
	//  SELECT
	//      id, git_revision, created_at
	//  FROM
	//      git_revisions
	//  LIMIT
	//      ? OFFSET ?
	ListGitRevisionsPaginated(ctx context.Context, arg ListGitRevisionsPaginatedParams) ([]GitRevision, error)
	//ListGoVersions
	//
	//  SELECT
	//      id, name, version, created_at
	//  FROM
	//      go_versions
	ListGoVersions(ctx context.Context) ([]GoVersion, error)
	//ListGoVersionsPaginated
	//
	//  SELECT
	//      id, name, version, created_at
	//  FROM
	//      go_versions
	//  LIMIT
	//      ? OFFSET ?
	ListGoVersionsPaginated(ctx context.Context, arg ListGoVersionsPaginatedParams) ([]GoVersion, error)
	//ListLogs
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	ListLogs(ctx context.Context) ([]ApiLog, error)
	//ListLogsByDeploymentID
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      deployment_id = ?
	ListLogsByDeploymentID(ctx context.Context, arg ListLogsByDeploymentIDParams) ([]ApiLog, error)
	//ListLogsByMethod
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      method = ?
	ListLogsByMethod(ctx context.Context, arg ListLogsByMethodParams) ([]ApiLog, error)
	//ListLogsByStatusCode
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      status_code = ?
	ListLogsByStatusCode(ctx context.Context, arg ListLogsByStatusCodeParams) ([]ApiLog, error)
	//ListLogsByUserAgent
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  WHERE
	//      user_agent = ?
	ListLogsByUserAgent(ctx context.Context, arg ListLogsByUserAgentParams) ([]ApiLog, error)
	//ListLogsPaginated
	//
	//  SELECT
	//      id, level_id, created_at, go_version_id, build_sum_id, git_revision_id, user_agent, method, url, elapsed_ms, status_code, deployment_id
	//  FROM
	//      api_logs
	//  LIMIT
	//      ? OFFSET ?
	ListLogsPaginated(ctx context.Context, arg ListLogsPaginatedParams) ([]ApiLog, error)
	//ListLogsWithJoin
	//
	//  SELECT
	//      l.id,
	//      l.created_at,
	//      l.user_agent,
	//      l.method,
	//      l.url,
	//      l.elapsed_ms,
	//      l.status_code,
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
	//ListURLs
	//
	//  SELECT
	//      id, url, created_at
	//  FROM
	//      urls
	ListURLs(ctx context.Context) ([]Url, error)
	//ListURLsPaginated
	//
	//  SELECT
	//      id, url, created_at
	//  FROM
	//      urls
	//  LIMIT
	//      ? OFFSET ?
	ListURLsPaginated(ctx context.Context, arg ListURLsPaginatedParams) ([]Url, error)
}

var _ Querier = (*Queries)(nil)
