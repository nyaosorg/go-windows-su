package main

import (
	"fmt"
	"os"
	"reflect"
	"syscall"

	"github.com/zetamatta/go-windows-su"
)

func mains() (int, error) {
	if len(os.Args) >= 2 {
		return su.ShellExecute("open", os.Args[1], "", "")
	}
	return 0, nil
}

func main() {
	if errno, err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		fmt.Fprintf(os.Stderr, "errno=%d\n", errno)
		fmt.Fprintln(os.Stderr, reflect.TypeOf(err).String())
		if e, ok := err.(syscall.Errno); ok {
			fmt.Fprintf(os.Stderr, "syscall.Errno=%d\n", e)
		}
		os.Exit(1)
	}
}
