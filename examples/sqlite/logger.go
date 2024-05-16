package main

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
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
func Get(getenv func(string) (string, error)) (*zerolog.Logger, error) {
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
		if os.Getenv("APP_ENV") != "development" {
			fileLogger := &lumberjack.Logger{
				Filename:   "dressingai.json",
				MaxSize:    5,    // megabytes
				MaxBackups: 10,   // number of backup files
				MaxAge:     14,   // days
				Compress:   true, // disabled by default
			}
			output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
		}
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
