package assets

import (
	"encoding/json"
	"io/fs"
	"log/slog"
	"os"
	"sync"
)

const manifestFile = "manifest.json"

var fsys fs.FS = os.DirFS("public")
var manifestCache map[string]string
var manifestMu sync.Mutex

func SetFS(f fs.FS) {
	fsys = f
}

func Manifest() (map[string]string, error) {
	if !IsDev {
		manifestMu.Lock()
		cached := manifestCache
		manifestMu.Unlock()
		if cached != nil {
			return cached, nil
		}
	}

	data, err := fs.ReadFile(fsys, manifestFile)
	if err != nil {
		return nil, err
	}

	var manifest map[string]string
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	if !IsDev {
		manifestMu.Lock()
		manifestCache = manifest
		manifestMu.Unlock()
	}

	return manifest, nil
}

func AssetPath(entry string) string {
	manifest, err := Manifest()
	if err != nil {
		slog.Warn("assets manifest unavailable", "err", err)
		return "/public/assets/" + entry + ".js"
	}

	path, ok := manifest[entry]
	if !ok {
		slog.Warn("asset entry not found", "entry", entry)
		return "/public/assets/" + entry + ".js"
	}

	return "/public/" + path
}
