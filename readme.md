go-windows-su
=============

Switch or test administrator's mode

Switch
------

```go
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
