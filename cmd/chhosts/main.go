package main

import (
	"os"
	"path/filepath"

	"github.com/zetamatta/go-windows-su"
)

func main() {
	windir := os.Getenv("SystemRoot")
	hosts := filepath.Join(windir, `system32\drivers\etc\hosts`)

	su.ShellExecute(su.RUNAS,
		"notepad",
		hosts,
		`C:\`)
}
