//go:build !production && !server

package main

import (
	"io/fs"
	"os"
)

var assets fs.FS = os.DirFS(".")
