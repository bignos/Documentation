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

### [1.1] Hello, World
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

### [1.2] Command-Line Arguments
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

### [1.3] Finding Duplicate Lines

Dup1
```go
// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)

    for input.Scan() {
        counts[input.Text()]++
    }

    // NOTE: ignoring potential errors from input.Err()

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```

Dup2
```go
// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n",err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
}
```
- **Notice** that the call to countLines precedes its declaration. Functions and other package-level entities may be declared in any order.
- A **map** is a reference to the data structure created by make. When a **map** is passed to a function, the function receives a copy of the reference,  
    so any changes the called function makes to the underlying data structure will be visible through the caller’s map reference too.

Dup3
```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```


### [1.4] Animated GIFs

Lissajous
```go
// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
)

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles = 5   // number of complete xoscillator revolutions
        res = 0.001  // angular resolution
        size = 100   // image canvas covers [-size..+size]
        nframes = 64 // number of animation frames
        delay = 8    // delay between frames in 10ms units
    )

    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference

    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

```

- The value of a constant must be a number, string, or boolean.

### [1.5] Fetching a URL

Fetch
```go
// Fetch prints the content found at a URL.
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }

        b, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()

        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s:%v\n", url, err)
            os.Exit(1)
        }

        fmt.Printf("%s", b)
    }
}
```

### [1.6] Fetching URLs Concurrently

Fetchall
```go
// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)

    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }

    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }

    fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)

    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() // don't leak resources
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
```
- When one goroutine attempts a send or receive on a channel,  
    it blocks until another goroutine attempts the corresponding receive or send operation,  
   at which point the value is transferred and both goroutines proceed.

### [1.7] A Web Server
