package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/zetamatta/go-windows-netresource"
	"github.com/zetamatta/go-windows-su"
)

func mains(args []string) error {
	var buffer strings.Builder
	buffer.WriteString(`/s /c "`)

	if dir, err := os.Getwd(); err == nil {
		fmt.Fprintf(&buffer, `cd "%s" & `, dir)
	}

	if netDrives, err := netresource.GetNetDrives(); err == nil {
		for _, n := range netDrives {
			fmt.Fprintf(&buffer, `net use %c: "%s" 2>nul & `,
				n.Letter, n.Remote)
		}
	}

	buffer.WriteString("copy ")
	for _, s := range args {
		fmt.Fprintf(&buffer, `"%s" `, s)
	}
	buffer.WriteString(`& pause"`)

	param := buffer.String()

	fmt.Printf("cmd %s\n", param)
	_, err := su.ShellExecute(su.RUNAS, "cmd.exe", param, "")
	return err
}

func main() {
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
