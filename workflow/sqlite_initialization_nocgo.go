//go:build !cgo

package workflow

import "context"

// Without cgo, mattn/go-sqlite3 cannot open databases, so there are no
// SQLite lock errors to retry.
func retrySQLiteInitialization(ctx context.Context, initialize func(context.Context) error) error {
	return initialize(ctx)
}
