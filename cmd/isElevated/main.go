package main

import (
	"fmt"
	"os"

	"github.com/zetamatta/go-windows-su"
)

func main() {
	flag, err := su.IsElevated()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	} else if flag {
		fmt.Println("Administrator mode")
		os.Exit(0)
	} else {
		fmt.Println("Not administrator mode")
		os.Exit(1)
	}
}
