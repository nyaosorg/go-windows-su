package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zetamatta/go-windows-netresource"
	"github.com/zetamatta/go-windows-su"
)

func mains(args []string) error {
	var buffer strings.Builder
	buffer.WriteString(`/s /c "`)

	if netDrives, err := netresource.GetNetDrives(); err == nil {
		for _, n := range netDrives {
			fmt.Fprintf(&buffer, `net use %c: "%s" 2>nul & `,
				n.Letter, n.Remote)
		}
	}

	if dir, err := os.Getwd(); err == nil {
		fmt.Fprintf(&buffer, `cd /D "%s" & `, dir)
	}

	buffer.WriteString("copy ")
	for _, s := range args {
		if s[0] == '/' {
			fmt.Fprintf(&buffer, `%s `, s)
		} else {
			fmt.Fprintf(&buffer, `"%s" `, s)
		}
	}
	buffer.WriteString(`& pause"`)

	param := buffer.String()

	fmt.Printf("cmd %s\n", param)
	_, err := su.ShellExecute(su.RUNAS, "cmd.exe", param, "")
	return err
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
