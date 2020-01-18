// +build example

package main

import (
	"github.com/zetamatta/go-windows-su"
)

func main() {
	su.ShellExecute(su.RUNAS,
		"notepad",
		`C:\windows\system32\drivers\etc\hosts`,
		`C:\`)
}
