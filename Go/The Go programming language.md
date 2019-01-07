The Go programming language
===========================

More information
----------------

- [Go official web site](https://golang.org)
- [Go official blog](https://blog.golang.org)
- [Go Playground](https://play.golang.org)
- [Go Tour](https://tour.golang.org)
- [Go standard library](https://golang.org/pkg)


1- Tutorial
-----------

### 1.1 Hello, World
```go
package main

import "fmt"

func main() {
        fmt.Println("Hello, World")
}
```
#### Notes:
- Package `main` is special. It defines a standalone executable program, not a library.
- The `import` declarations must follow the `package` declaration.
- For instance, the opening brace `{` of the function must be on the same line as the end of the func declaration, not on a line by itself.

### 1.2 Command-Line Arguments
```go
// Echo1 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }

    fmt.Println(s)
}
```

Second version of Echo:
```go
// Echo2 prints its command-line arguments.
package main

import (
"fmt"
"os"
)

func main() {
    s, sep := "", ""

    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
```
Alternative variable declaration:
```go
s := ""
var s string
var s = ""
var s string = ""
```

Third version of Echo:
```go
func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}
```

Alternative version without formatting:
```go
func main() {
    fmt.Println(os.Args[1:])
}
```

