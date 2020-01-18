// +build run

package main

import (
	"github.com/zetamatta/go-windows-su"
)

func main() {
	flag, err := su.IsElevated()
	if err != nil {
		println(err.Error())
		return
	}
	if flag {
		println("Administrator mode")
	} else {
		println("Not administrator mode")
	}
}
