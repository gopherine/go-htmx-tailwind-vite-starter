//go:build tools
// +build tools

package tools

import (
	_ "github.com/a-h/templ/cmd/templ"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)
