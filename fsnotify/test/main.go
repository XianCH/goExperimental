package main

import (
	"path/filepath"

	testfsnotify "github.com/x14n/goExperimental/fsnotify"
)

func main() {
	fileToWatch := filepath.Join("test", "test.txt")
	testfsnotify.WatchFile(fileToWatch)
}
