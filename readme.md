go-windows-su
=============

Switch or test administrator's mode

Switch
------

```go
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
```

Test
----

```go

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
```

Sample Application
------------------

[source](./cmd/sucopy/main.go)

```
sucopy SRC DST
```

Execute `copy` built in CMD.EXE in Administrator-mode.
`sucopy` keeps network-drives.
