# StrIpol

An extremely simple & easy-to-use string interpolation package.

## Example

```go
package main

import (
    "fmt"

    "github.com/EricFrancis12/stripol"
)

func main() {
    i := stripol.New("{{", "}}")

    i.RegisterVar("NAME", "Mike Tyson")
    i.RegisterVar("PET", "tiger")

    str := "{{ NAME }} has a pet {{ PET }}."
    result := i.Eval(str)

    fmt.Println(result)
    // Mike Tyson has a pet tiger.
}
```

## Installation

```bash
go get github.com/EricFrancis12/stripol
```

## Testing

```bash
make test
```
