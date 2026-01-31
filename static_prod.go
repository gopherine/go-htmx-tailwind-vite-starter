//go:build !dev
// +build !dev

package main

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed public
var embeddedPublic embed.FS

var PublicFS fs.FS

func init() {
	sub, err := fs.Sub(embeddedPublic, "public")
	if err != nil {
		log.Fatalf("failed to create embedded public FS: %v", err)
	}
	PublicFS = sub
}
