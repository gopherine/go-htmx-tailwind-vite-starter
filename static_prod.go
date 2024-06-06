//go:build !dev
// +build !dev

package main

import (
	"embed"
	"net/http"
)

//go:embed public
var PublicFS embed.FS

func public() http.Handler {
	return http.FileServerFS(PublicFS)
}
