package config

import (
	"context"
)

type rowScanner interface {
	Scan(dest ...interface{}) error
}

type Querier interface {
	QueryRow(ctx context.Context, query string, args ...interface{}) rowScanner
}

var DB Querier
