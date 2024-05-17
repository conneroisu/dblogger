package dblogger

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

// getenvTest is the getenv function for the test environment
func getenvTest(key string) (string, error) {
	switch key {
	case "DEPLOYMENT":
		return "staging", nil
	default:
		return "", fmt.Errorf("invalid key: %s", key)
	}
}

// initDatabase creates a new database and returns the Queries type
// for the test environment.
func initDatabase() (*Queries, error) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	q, err := NewLogsDatabase(getenvTest, db)
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}
	return q, nil
}

// TestQueries_CountBuildSums tests the CountBuildSums method
// of the Queries type.
//
// This test case ensures that the method correctly counts
// the number of build sums in the database.
func TestQueries_CountBuildSums(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	q, err := initDatabase()
	if err != nil {
		t.Fatal(err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := q.CountBuildSums(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CountBuildSums() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queries.CountBuildSums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_CountDeployments(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	q, err := initDatabase()
	if err != nil {
		t.Fatal(err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := q.CountDeployments(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CountDeployments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queries.CountDeployments() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestQueries_CountGitRevisions tests the CountGitRevisions method
// of the Queries type upon initialization.
func TestQueries_CountGitRevisions(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	q, err := initDatabase()
	if err != nil {
		t.Fatal(err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := q.CountGitRevisions(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CountGitRevisions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queries.CountGitRevisions() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestQueries_CountGoVersions tests the CountGoVersions method
// of the Queries type upon initialization.
func TestQueries_CountGoVersions(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	q, err := initDatabase()
	if err != nil {
		t.Fatal(err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := q.CountGoVersions(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CountGoVersions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queries.CountGoVersions() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestQueries_CountLogLevels tests the CountLogLevels method
func TestQueries_CountLogLevels(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	q, err := initDatabase()
	if err != nil {
		t.Fatal(err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
			},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := q.CountLogLevels(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CountLogLevels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queries.CountLogLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestQueries_CountLogs tests the CountLogs method
func TestQueries_CountLogs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	q, err := initDatabase()
	if err != nil {
		t.Fatal(err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := q.CountLogs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CountLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queries.CountLogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestQueries_CountURLs(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    int64
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.CountURLs(tt.args.ctx)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.CountURLs() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if got != tt.want {
//                                 t.Errorf("Queries.CountURLs() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_DeleteBuildSumByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg DeleteBuildSumByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.DeleteBuildSumByID(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.DeleteBuildSumByID() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_DeleteDeploymentByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg DeleteDeploymentByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.DeleteDeploymentByID(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.DeleteDeploymentByID() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_DeleteGitRevisionByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg DeleteGitRevisionByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.DeleteGitRevisionByID(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.DeleteGitRevisionByID() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_DeleteGoVersionByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.DeleteGoVersionByID(tt.args.ctx); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.DeleteGoVersionByID() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_DeleteLogByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg DeleteLogByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.DeleteLogByID(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.DeleteLogByID() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_DeleteLogLevelByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg DeleteLogLevelByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.DeleteLogLevelByID(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.DeleteLogLevelByID() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_DeleteURLByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg DeleteURLByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.DeleteURLByID(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.DeleteURLByID() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetBuildSumByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetBuildSumByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    BuildSum
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetBuildSumByID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetBuildSumByID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetBuildSumByID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetBuildSumsByDate(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetBuildSumsByDateParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []BuildSum
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetBuildSumsByDate(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetBuildSumsByDate() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetBuildSumsByDate() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetBuildSumsBySubstring(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetBuildSumsBySubstringParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []BuildSum
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetBuildSumsBySubstring(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetBuildSumsBySubstring() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetBuildSumsBySubstring() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetDeploymentByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetDeploymentByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    Deployment
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetDeploymentByID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetDeploymentByID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetDeploymentByID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetDeploymentsByDate(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetDeploymentsByDateParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []Deployment
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetDeploymentsByDate(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetDeploymentsByDate() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetDeploymentsByDate() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetDeploymentsByDateRange(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetDeploymentsByDateRangeParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []Deployment
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetDeploymentsByDateRange(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetDeploymentsByDateRange() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetDeploymentsByDateRange() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetDeploymentsBySubstring(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetDeploymentsBySubstringParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []Deployment
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetDeploymentsBySubstring(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetDeploymentsBySubstring() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetDeploymentsBySubstring() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetGitRevisionByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetGitRevisionByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    GitRevision
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetGitRevisionByID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetGitRevisionByID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetGitRevisionByID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetGitRevisionsByDate(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetGitRevisionsByDateParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []GitRevision
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetGitRevisionsByDate(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetGitRevisionsByDate() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetGitRevisionsByDate() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetGitRevisionsByName(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetGitRevisionsByNameParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    GitRevision
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetGitRevisionsByName(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetGitRevisionsByName() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetGitRevisionsByName() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetGoVersionByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetGoVersionByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    GoVersion
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetGoVersionByID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetGoVersionByID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetGoVersionByID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetGoVersionIdByName(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetGoVersionIdByNameParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    int64
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetGoVersionIdByName(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetGoVersionIdByName() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if got != tt.want {
//                                 t.Errorf("Queries.GetGoVersionIdByName() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetGoVersionsByDate(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetGoVersionsByDateParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []GoVersion
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetGoVersionsByDate(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetGoVersionsByDate() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetGoVersionsByDate() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetGoVersionsBySubstring(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetGoVersionsBySubstringParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []GoVersion
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetGoVersionsBySubstring(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetGoVersionsBySubstring() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetGoVersionsBySubstring() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetLogByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetLogByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetLogByID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetLogByID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetLogByID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetLogLevelsBySubstring(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetLogLevelsBySubstringParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []LogLevel
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetLogLevelsBySubstring(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetLogLevelsBySubstring() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetLogLevelsBySubstring() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetLogsByBuildSumID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetLogsByBuildSumIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetLogsByBuildSumID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetLogsByBuildSumID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetLogsByBuildSumID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetLogsByDate(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetLogsByDateParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetLogsByDate(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetLogsByDate() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetLogsByDate() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetLogsByDateRange(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetLogsByDateRangeParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetLogsByDateRange(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetLogsByDateRange() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetLogsByDateRange() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetLogsByElapsedRange(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetLogsByElapsedRangeParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetLogsByElapsedRange(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetLogsByElapsedRange() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetLogsByElapsedRange() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetLogsByGitRevisionID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetLogsByGitRevisionIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetLogsByGitRevisionID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetLogsByGitRevisionID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetLogsByGitRevisionID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetLogsByGoVersionID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetLogsByGoVersionIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetLogsByGoVersionID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetLogsByGoVersionID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetLogsByGoVersionID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetLogsByURLSubstring(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetLogsByURLSubstringParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetLogsByURLSubstring(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetLogsByURLSubstring() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetLogsByURLSubstring() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetURLByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetURLByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    Url
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetURLByID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetURLByID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetURLByID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetURLsByDate(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetURLsByDateParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []Url
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetURLsByDate(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetURLsByDate() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetURLsByDate() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_GetURLsBySubstring(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg GetURLsBySubstringParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []Url
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.GetURLsBySubstring(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.GetURLsBySubstring() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.GetURLsBySubstring() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertBuildSum(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertBuildSumParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    BuildSum
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.InsertBuildSum(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertBuildSum() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.InsertBuildSum() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertBuildSumReturningID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertBuildSumReturningIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    int64
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.InsertBuildSumReturningID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertBuildSumReturningID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if got != tt.want {
//                                 t.Errorf("Queries.InsertBuildSumReturningID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertBuildSumWithParam(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertBuildSumWithParamParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.InsertBuildSumWithParam(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertBuildSumWithParam() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertDeployment(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertDeploymentParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    Deployment
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.InsertDeployment(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertDeployment() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.InsertDeployment() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertDeploymentReturningID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertDeploymentReturningIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    int64
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.InsertDeploymentReturningID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertDeploymentReturningID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if got != tt.want {
//                                 t.Errorf("Queries.InsertDeploymentReturningID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertGitRevision(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertGitRevisionParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    GitRevision
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.InsertGitRevision(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertGitRevision() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.InsertGitRevision() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertGitRevisionReturningID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertGitRevisionReturningIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    int64
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.InsertGitRevisionReturningID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertGitRevisionReturningID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if got != tt.want {
//                                 t.Errorf("Queries.InsertGitRevisionReturningID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertGitRevisionWithParam(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertGitRevisionWithParamParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.InsertGitRevisionWithParam(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertGitRevisionWithParam() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertGoVersion(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertGoVersionParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    GoVersion
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.InsertGoVersion(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertGoVersion() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.InsertGoVersion() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertGoVersionReturningID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertGoVersionReturningIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    int64
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.InsertGoVersionReturningID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertGoVersionReturningID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if got != tt.want {
//                                 t.Errorf("Queries.InsertGoVersionReturningID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertLogEntry(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertLogEntryParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.InsertLogEntry(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertLogEntry() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertLogLevel(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertLogLevelParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.InsertLogLevel(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertLogLevel() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertURL(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertURLParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.InsertURL(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertURL() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_InsertURLWithParam(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg InsertURLWithParamParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.InsertURLWithParam(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.InsertURLWithParam() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListBuildSums(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []BuildSum
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListBuildSums(tt.args.ctx)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListBuildSums() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListBuildSums() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListBuildSumsPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListBuildSumsPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []BuildSum
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListBuildSumsPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListBuildSumsPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListBuildSumsPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListDeployments(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []Deployment
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListDeployments(tt.args.ctx)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListDeployments() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListDeployments() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListDeploymentsPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListDeploymentsPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []Deployment
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListDeploymentsPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListDeploymentsPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListDeploymentsPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListGitRevisions(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []GitRevision
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListGitRevisions(tt.args.ctx)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListGitRevisions() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListGitRevisions() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListGitRevisionsPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListGitRevisionsPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []GitRevision
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListGitRevisionsPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListGitRevisionsPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListGitRevisionsPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListGoVersions(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []GoVersion
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListGoVersions(tt.args.ctx)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListGoVersions() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListGoVersions() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListGoVersionsPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListGoVersionsPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []GoVersion
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListGoVersionsPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListGoVersionsPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListGoVersionsPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogLevels(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []LogLevel
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogLevels(tt.args.ctx)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogLevels() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogLevels() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogLevelsPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogLevelsPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []LogLevel
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogLevelsPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogLevelsPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogLevelsPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogs(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogs(tt.args.ctx)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogs() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogs() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsByDeploymentID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsByDeploymentIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsByDeploymentID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsByDeploymentID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsByDeploymentID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsByMethod(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsByMethodParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsByMethod(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsByMethod() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsByMethod() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsByUserAgent(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsByUserAgentParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsByUserAgent(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsByUserAgent() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsByUserAgent() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ApiLog
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoin(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoin(tt.args.ctx)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoin() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoin() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByBuildSumID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByBuildSumIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByBuildSumIDRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByBuildSumID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByBuildSumID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByBuildSumID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByBuildSumIDPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByBuildSumIDPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByBuildSumIDPaginatedRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByBuildSumIDPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByBuildSumIDPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByBuildSumIDPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByDate(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByDateParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByDateRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByDate(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByDate() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByDate() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByDatePaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByDatePaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByDatePaginatedRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByDatePaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByDatePaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByDatePaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByDateRange(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByDateRangeParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByDateRangeRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByDateRange(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByDateRange() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByDateRange() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByDateRangePaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByDateRangePaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByDateRangePaginatedRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByDateRangePaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByDateRangePaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByDateRangePaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByElapsedRange(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByElapsedRangeParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByElapsedRangeRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByElapsedRange(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByElapsedRange() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByElapsedRange() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByElapsedRangePaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByElapsedRangePaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByElapsedRangePaginatedRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByElapsedRangePaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByElapsedRangePaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByElapsedRangePaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByGitRevisionID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByGitRevisionIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByGitRevisionIDRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByGitRevisionID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByGitRevisionID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByGitRevisionID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByGitRevisionIDPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByGitRevisionIDPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByGitRevisionIDPaginatedRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByGitRevisionIDPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByGitRevisionIDPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByGitRevisionIDPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByGoVersionID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByGoVersionIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByGoVersionIDRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByGoVersionID(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByGoVersionID() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByGoVersionID() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinByGoVersionIDPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinByGoVersionIDPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinByGoVersionIDPaginatedRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinByGoVersionIDPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinByGoVersionIDPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinByGoVersionIDPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListLogsWithJoinPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListLogsWithJoinPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []ListLogsWithJoinPaginatedRow
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListLogsWithJoinPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListLogsWithJoinPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListLogsWithJoinPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListURLs(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []Url
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListURLs(tt.args.ctx)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListURLs() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListURLs() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_ListURLsPaginated(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg ListURLsPaginatedParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 want    []Url
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         got, err := q.ListURLsPaginated(tt.args.ctx, tt.args.arg)
//                         if (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.ListURLsPaginated() error = %v, wantErr %v", err, tt.wantErr)
//                                 return
//                         }
//                         if !reflect.DeepEqual(got, tt.want) {
//                                 t.Errorf("Queries.ListURLsPaginated() = %v, want %v", got, tt.want)
//                         }
//                 })
//         }
// }
//
// func TestQueries_UpdateGoVersionByID(t *testing.T) {
//         }
//         type args struct {
//                 ctx context.Context
//                 arg UpdateGoVersionByIDParams
//         }
//         tests := []struct {
//                 name    string
//                 args    args
//                 wantErr bool
//         }{
//                 // TODO: Add test cases.
//         }
//         for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                         q := &Queries{
//                         }
//                         if err := q.UpdateGoVersionByID(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
//                                 t.Errorf("Queries.UpdateGoVersionByID() error = %v, wantErr %v", err, tt.wantErr)
//                         }
//                 })
//         }
// }
