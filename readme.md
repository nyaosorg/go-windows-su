go-windows-su
=============

Switch or test administrator's mode

[Switch](./cmd/chhosts/main.go)
------------------------------

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

[Test](./cmd/isElevated/main.go)
--------------------------------

```go
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
```

Sample Application
------------------

[source](./cmd/sucopy/main.go)

```
sucopy SRC DST
```

Execute `copy` built in CMD.EXE in Administrator-mode.
`sucopy` keeps network-drives.
