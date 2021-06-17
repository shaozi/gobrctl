# Go Library to List Bridges in Linux

This is a go library that gets a list of bridges. The same as `brctl show`

## Installation

To install this package, you need to install Go and set your Go workspace first.

The first need Go installed (version 1.16+ is required), then you can use the below Go command to install gobrctl.

```sh
$ go get -u github.com/shaozi/gobrctl
```

Import it in your code:

```go
import "github.com/shaozi/gobrctl"
```

## Usage

```go
package main

import (
    "fmt"
    "strings"
    "github.com/shaozi/gobrctl"
)

// simulate brctl show
func main() {
    bridges := gobrctl.GetAllBridges()

    fmt.Println("bridge name     bridge id           STP enabled     interfaces")
    for _, bridge := range bridges {
        var stpEnabled string
        if bridge.Stp {
            stpEnabled = "yes"
        } else {
            stpEnabled = "no"
        }
        fmt.Printf("%-16s%-20s%-16s%s\n", bridge.Name, bridge.Id, stpEnabled, strings.Join(bridge.Interfaces, ", "))
    }
}

```
