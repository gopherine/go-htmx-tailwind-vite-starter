//go:build dev
// +build dev

package main

import (
	"fmt"
	"net/http"
	"os"
)

func public() http.Handler {
	fmt.Println("Building static files for development")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}
