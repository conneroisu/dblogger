package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/conneroisu/dblogger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

// once is used to ensure that the logger is only initialized once
var once sync.Once

// log is the global logger object
var log zerolog.Logger

// Get returns the logger
// if the logger is not initialized it will initialize the logger
// and return the logger
//
// Example:
//
//	     func(w http.ResponseWriter, r *http.Request) {
//		l, err := logger.Get(os.Getenv)
//		if err != nil {
//			return fmt.Errorf("error getting logger: %w", err)
//		}
//
//		l.Info().
//		    Str("method", r.Method).
//		    Str("url", r.URL.RequestURI()).
//		    Str("user_agent", r.UserAgent()).
//		    Dur("elapsed_ms", time.Since(start)).
//		    Msg("incoming request")
//		return nil
//	     }
func Get(getenv func(string) (string, error), queries *dblogger.Queries) (*zerolog.Logger, error) {
	verbosity, err := getenv("VERBOSITY")
	if err != nil {
		return nil, fmt.Errorf("failed to get verbosity: %w", err)
	}
	verbosityInt, err := strconv.Atoi(verbosity)
	if err != nil {
		return nil, fmt.Errorf("invalid verbosity: %w", err)
	}
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano
		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		output = zerolog.MultiLevelWriter(os.Stderr, queries)
		var gitRevision string
		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}
		log = zerolog.New(output).
			Level(zerolog.Level(verbosityInt)).
			With().
			Timestamp().
			Str("git_revision", gitRevision).
			Str("deployment", os.Getenv("DEPLOYMENT")).
			Str("go_version", buildInfo.GoVersion).
			Str("build_sum", buildInfo.Main.Sum).
			Logger()
	})
	return &log, nil
}

// requestLogger logs the incoming requests.
// It logs the incoming requests to the the path of the request
// and the time it took to process the request
//
// Example:
//
//	     mux := http.NewServeMux()
//	     handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte("Hello, World!"))
//	     })
//	     mux.HandleFunc("/", requestLogger(handler))
//	     http.ListenAndServe(":8080", mux)
func requestLogger(
	queries *dblogger.Queries,
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		l, err := Get(GetErrEnv(os.Getenv), queries)
		if err != nil {
			panic(fmt.Errorf("error getting logger: %w", err))
		}
		next.ServeHTTP(w, r)
		l.
			Info().
			Str("method", r.Method).
			Str("url", r.URL.RequestURI()).
			Str("user_agent", r.UserAgent()).
			Dur("elapsed_ms", time.Since(start)).
			Msg("request-to: " + r.URL.Path)
	}
}

// GetErrEnv returns a function that returns the value of the given key
// from the environment.
// If the key is not set, the function returns an error.
//
// The function is useful for getting environment variables
// in a type-safe way.
//
// Example:
//
//	func main() {
//		// Get the value of the "MY_ENV" environment variable.
//		val, err := GetErrEnv(os.GetErrEnv)("MY_ENV")
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(val)
//	}
func GetErrEnv(fn func(string) string) func(string) (string, error) {
	return func(key string) (string, error) {
		val := fn(key)
		if val == "" {
			return "",
				fmt.Errorf("env %s not set", key)
		}
		return val, nil
	}
}
