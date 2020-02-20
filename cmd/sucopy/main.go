package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/zetamatta/go-windows-su"
)

func mains(args []string) error {
	var buffer strings.Builder

	buffer.WriteString(`/s /c "`)
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintf(&buffer, `cd "%s" && copy`, dir)
	for _, s := range args {
		fmt.Fprintf(&buffer, ` "%s"`, s)
	}
	buffer.WriteString(` & pause"`)

	param := buffer.String()

	fmt.Println(param)
	_, err = su.ShellExecute(su.RUNAS, "cmd.exe", param, "")
	return err
}

func main() {
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
