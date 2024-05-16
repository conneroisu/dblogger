package dblogger

import "strings"

// isAlreadyExistsError checks if the error is an error that
// indicates that the database already exists.
//
// Example:
//
//	if isAlreadyExistsError(err) {
//	    // handle error
//	}
func isAlreadyExistsError(err error) bool {
	return strings.Contains(err.Error(), "already exists") 
}
