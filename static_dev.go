//go:build dev
// +build dev

package main

import (
	"io/fs"
	"os"
)

var PublicFS fs.FS = os.DirFS("public")
