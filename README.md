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
    s := stripol.New("{{", "}}")

    s.RegisterVar("NAME", "Mike Tyson")
    s.RegisterVar("PET", "tiger")

    str := "{{ NAME }} has a pet {{ PET }}."
    result := s.Eval(str)

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
