# The Go programming language

## More information

- [Go official web site](https://golang.org)
- [Go official blog](https://blog.golang.org)
- [Go Playground](https://play.golang.org)
- [Go Tour](https://tour.golang.org)
- [Go standard library](https://golang.org/pkg)

## 1. Tutorial

### 1.1 Hello, World

```go
package main

import "fmt"

func main() {
        fmt.Println("Hello, World")
}
```

- Package `main` is special. It defines a standalone executable program, not a library.
- The `import` declarations must follow the `package` declaration.
- For instance, the opening brace `{` of the function must be on the same line as the end of the func declaration, not on a line by itself.

### 1.2 Command-Line Arguments

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

### 1.3 Finding Duplicate Lines

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

### 1.4 Animated GIFs

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

### 1.5 Fetching a URL

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

### 1.6 Fetching URLs Concurrently

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

### 1.7 A Web Server

Server1

```go
// Server1 is a minimal "echo" server.
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler) // each request calls handler
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

```

Server2

```go
// Server2 is a minimal "echo" and counter server.
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var mu sync.Mutex
var count int

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/count", counter)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count++
    mu.Unlock()
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}
```

- A handler pattern that ends with a slash matches any URL that has the pattern as a prefix.  
    Behind the scenes, the server runs the handler for each incoming request in a separate goroutine  
    so that it can serve multiple requests simultaneously.

Server2 handler_debug function

```go
// handler echoes the HTTP request.
func handlerDebug(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

    for k, v := range r.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }

    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }

    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }
}
```

Server2 handler_gif function

```go
// handler animated GIF
func handlerGif(w http.ResponseWriter, r *http.Request) {
    lissajous.Lissajous(w)
}
```

### 1.8 Loose Ends

Switch example:

```go
switch coinflip() {
    case "heads":
        heads++
    case "tails":
        tails++
    default:
        fmt.Println("landed on edge!")
}
```

- Cases do not fall through from one to the next as in C-like languages  
    (though there is a rarely used `fallthrough` statement that overrides this behavior).

Tagless switch example:

```go
func Signum(x int) int {
    switch {
        case x > 0:
            return +1
        default:
            return 0
        case x < 0:
            return -1
    }
}
```

Type declaration example:

```go
type Point struct {
    X, Y int
}

var p Point
```

## 2. Program Structure

### 2.1 Names

Go as **25** reserved keywords: 

```go
break
case
chan
const
continue
default
defer
else
fallthrough
for
func
go
goto
if
import
interface
map
package
range
return
select
struct
switch
type
var
```

Go predeclared names:

```go
// Constants
true
false
iota
nil

// Types
int
int8
int16
int32
int64

uint
uint8
uint16
uint32
uint64
uintptr

float32
float64

complex128
complex64

bool

byte

rune

string

error

// Functions
append
cap
close
complex
copy
delete
imag
len
make
new
panic
real
recover
```

- These names are not reserved, so you may use them in declarations.  
    We’ll see a handful of places where redeclaring one of them makes sense, but beware of the potential for confusion.

- If an entity is declared within a function, it is local to that function.  
    If declared outside of a function, however, it is visible in all files of the package to which it belongs.

- The case of the first letter of a name determines its visibility across package boundaries.  
    If the name **begins with an upper-case letter**, it is exported,  
    which means that it is visible and accessible outside of its own package and may be referred to by other parts of the program,  
    as with Printf in the fmt package. **Package names themselves are always in lower case.**

- Go programmers use **camel case** when forming names by combining words

### 2.2 Declarations

- There are four major kind of declarations:

```go
var
const
type
func
```

- A Go program is stored in one or more files whose names end in `.go`.  
    Each file begins with a `package` declaration that says what package the file is part of.  
    The `package` declaration is followed by any `import` declarations,  
    and then a sequence of package-level declarations of types, variables, constants, and functions, in any order.

Example 'boiling' (src/ch2/boiling.go):

```go
// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0

func main() {
    var f = boilingF
    var c = (f - 32) * 5 / 9
    fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
    // Output:
    // boiling point = 212°F or 100°C
}
```

Example 'ftoc' (src/ch2/ftoc.go):  
- The function fToC below encapsulates the temperature conversion logic so that it is defined only once but may be used from multiple places.  
    Here main calls it twice, using the values of two different local constants:

```go
// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
    const freezingF, boilingF = 32.0, 212.0
    fmt.Printf("%g°F = %g°C\n", freezingF,
    fToC(freezingF)) // "32°F = 0°C"
    fmt.Printf("%g°F = %g°C\n", boilingF,
    fToC(boilingF)) // "212°F = 100°C"
}

func fToC(f float64) float64 {
    return (f - 32) * 5 / 9
}
```

### 2.3 Variables

- Variables declaration: `var name type = expression`
- Either the type or the = expression part may be omitted, but not both. If the type is omitted, it is determined by the initializer expression.  
    If the `expression` is omitted,  
    the initial value is the *zero value* for the type, which is `0` for numbers, `false` for booleans, "" for strings,  
    and `nil` for interfaces and reference types (slice, pointer, map, channel, function).  
    The zero value of an aggregate type like an array or a struct has the zero value of all of its elements or fields.

One line multiple variable declaration example:

```go
var i, j, k int                     // int, int, int
var b, f, s = true, 2.3, "four"     // bool, float64, string
```

A set of variables can also be initialized by calling a function that retrurns multiple values:

```go
var f, err = os.Open(name)          // os.Open returns a file and an error
```

#### 2.3.1 Short Variable Declarations

Short variable declaration example:

```go
anim := gif.GIF{LoopCount: nframes}
freq := rand.Float64() * 3.0
t := 0.0
```

- Because of their brevity and flexibility, short variable declarations are used to declare and initialize the majority of local variables.  
    A `var` declaration tends to be reserved for local variables that need an explicit type that differs from that of the initializer expression,  
    or for when the variable will be assigned a value later and its initial value is unimportant.

Mix of variable declaration example:

```go
i := 100
var boiling float64 = 100
var names []string
var err error
var p Point
```

- declarations with multiple initializer expressions should be used only when they help readability,  
    such as for short and natural groupings like the initialization part of a `for` loop.

- Keep in mind that := is a declaration, whereas = is an assignment.  
    A multi-variable declaration should not be confused with a tuple assignment,  
    in which each variable on the left-hand side is assigned the corresponding value from the right-hand side:

```go
i, j = j, i             // swap values of i and j
```

- A short variable declaration must declare at least one new variable, however, so this code will not compile:

```go
f, err := os.Open(infile)
// ...
f, err := os.Create(outfile)        // compile error: no new variables
```

#### 2.3.2 Pointers

- A *pointer* value is the *address* of a variable.  
    A pointer is thus the location at which a value is stored.  
    Not every value has an address, but every variable does.  
    With a pointer, we can read or update the value of a variable *indirectly*,  
    without using or even knowing the name of the variable, if indeed it has a name.

Pointer declaration example:

```go
x := 1
p := &x                 // p, of type *int, points to x

fmt.Println(*p)         // "1"

*p = 2

fmt.Println(x)          // "2"

```

- Variables are sometimes described as *addressable* values.  
    Expressions that denote variables are the only expressions to which the *address-of* operator `&` may be applied.  
    The zero value for a pointer of any type is `nil`.  
    The test `p != nil` is true if `p` points to a variable.  
    Pointers are comparable; two pointers are equal if and only if they point to the same variable or both are nil.

Pointer comparaison example:

```go
var x, y int

fmt.Println(&x == &x, &x == &y, &x == nil)      // "true false false"
```

- It is perfectly safe for a function to return the address of a local variable.  
    For instance, in the code below, the local variable `v` created by this particular call to `f` will remain in existence  
    even after the call has returned, and the pointer `p` will still refer to it:

```go
var p = f()

func f() *int {
    v := 1
    return &v
}
```

- Each call of `f` returns a distinct value

```go
fmt.Println(f() == f())                         // "false"
```

- Because a pointer contains the address of a variable,  
    passing a pointer argument to a function makes it possible for the function to update the variable that was indirectly passed.  
    For example, this function increments the variable that its argument points to and  
    returns the new value of the variable so it may be used in an expression:

```go
func incr(p *int) int {
    *p++                                        // increment what p points to; does not change p
    return *p
}

v := 1
incr(&v)                                        // side effect: v is now 2

fmt.Println(incr(&v))                           // "3" (and v is 3)
```

Echo4 code sample:

```go
// Echo4 prints its command-line arguments.
package main

import (
    "flag"
    "fmt"
    "strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
    flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *sep))
    if !*n {
        fmt.Println()
    }
}
```

#### 2.3.3 The new Function

- Another way to create a variable is to use the built-in function `new`.  
    The expression `new(T)` creates an unnamed variable of type T, initializes it to the zero value of T,  
    and returns its address, which is a value of type `*T.`

```go
p := new(int)                                   // p, of type *int, points to an unnamed int variable
fmt.Println(p)                                  // "0"
*p = 2                                          // sets the unnamed int to 2
fmt.Println(p)                                  // "2"
```

- Thus `new` is only a syntactic convenience, not a fundamental notion: the two newInt functions below have identical behaviors.

```go
func newInt() *int {
    return new(int)
}

func newInt() *int {
    var dummy int
    return &dummy
}
```

- Each call to new returns a distinct variable with a unique address:

```go
p := new(int)
q := new(int)
fmt.Println(p == q)                             // "false"
```

- The `new` function is relatively rarely used because the most common unnamed variables are of struct types,  
    for which the struct literal syntax is more flexible.

- Since `new` is a predeclared function, not a keyword, it’s possible to redefine the name for something else within a function, for example:

```go
func delta(old, new int) int {
    return new - old
}
```

- Of course, within `delta`, the built-in new function is unavailable.

#### 2.3.4 Lifetime of Variables

```go
var global *int

func f() {
    var x int
    x = 1
    global = &x
}

func g() {
    y := new(int)
    *y = 1
}
```

- Here, `x` must be heap-allocated because it is still reachable from the variable `global` after `f` has returned,  
    despite being declared as a local variable; we say `x` *escapes from* `f`.  
    Conversely, when `g` returns, the variable `*y` becomes unreachable and can be recycled.  
    Since `*y` does not escape from `g`, it’s safe for the compiler to allocate `*y` on the stack, even though it was allocated with `new`.
- It’s good to keep in mind during performance optimization, since each variable that escapes requires an **extra memory allocation**.

- For example, keeping unnecessary pointers to short-lived objects within long-lived objects,  
    especially global variables, will prevent the garbage collector from reclaiming the short-lived objects.

### 2.4 Assignments

- The value held by a variable is updated by an assignment statement,  
    which in its simplest form has a variable on the left of the = sign and an expression on the right.

```go
x = 1                                       // named variable
*p = true                                   // indirect variable
person.name = "bob"                         // Struct field
count[x] = count[x] * scale                 // array or slice or map element
```

Assignation with assignment operator example:

```go
count[x] *= scale
v := 1
v++                                         // same as v = v + 1; v becomes 2
v--                                         // same as v = v - 1; v becomes 1 again
```

#### 2.4.1 Tuple Assignment

- Another form of assignment, known as tuple assignment, allows several variables to be assigned at once.  
    All of the right-hand side expressions are evaluated before any of the variables are updated,  
    making this form most useful when some of the variables appear on both sides of the assignment, as happens,  
    for example, when swapping the values of two variables:

```go
x, y = y, x
a[i], a[j] = a[j], a[i]
i, j, k = 2, 3, 5

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}

func fib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x
}

f, err = os.Open("foo.txt")                 // function call returns two values

v, ok = m[key]                              // map lookup
v, ok = x.(T)                               // type assertion
v, ok = <-ch                                // channel receive

_, err = io.Copy(dst, src)                  // discard byte count
_, ok  = x.(T)                              // check type but discard result
```

#### 2.4.2 Assignability

```go
medals := []string{"gold", "silver", "bronze"}

// Equivalent
medals[0] = "gold"
medals[1] = "silver"
medals[2] = "bronze"
```

- The elements of maps and channels, though not ordinary variables, are also subject to similar implicit assignments.
- Whether two values may be compared with == and != is related to assignability:  
    in any comparison, the first operand must be assignable to the type of the second operand, or vice versa.  
    As with assignability, we’ll explain the relevant cases for comparability when we present each new type.

### 2.5 Type Declarations

- Type declarations format: `type name underlying-type`
- Type declarations most often appear at package level, where the named type is visible throughout the package,  
    and if the name is exported (it starts with an upper-case letter), it’s accessible from other packages as well.

```go
// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit { 
    return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius { 
    return Celsius((f - 32) * 5 / 9) 
}
```
- For every type `T`, there is a corresponding conversion operation `T(x)` that converts the value `x` to type `T`.  
    A conversion from one type to another is allowed if both have the same underlying type,  
    or if both are unnamed pointer types that point to variables of the same underlying type;  
    these conversions change the type but not the representation of the value.
- The underlying type of a named type determines its structure and representation,  
    and also the set of intrinsic operations it supports, which are the same as if the underlying type had been used directly.

```go
fmt.Printf("%g\n", BoilingC - FreezingC)                    // "100" °C

boilingF := CToF(BoilingC)
fmt.Printf("%g\n", boilingF - CToF(FreezingC))              // "180" °F
fmt.Printf("%g\n", boilingF - FreezingC)                    // Compile error: type mismatch
```

Two values of different named types cannot be compared directly:

```go
var c Celsius
var f Fahrenheit

fmt.Println(c == 0)                                         // "true"
fmt.Println(f >= 0)                                         // "true"
fmt.Println(c == f)                                         // Compile error: type mismatch
fmt.Println(c == Celsius(f))                                // "true"
```

- A named type may provide notational convenience if it helps avoid writing out complex types over and over again.  
    The advantage is small when the underlying type is simple like `float64`, but big for complicated types.

Function for Celsius convertion to string

```go
func (c Celsius) String() string {
    fmt.Sprintf("%g°C", c)
}

c := FToC(212.0)
fmt.Println(c.String())                                     // "100°C"
fmt.Printf("%v\n", c)                                       // "100°C"; no need to call String explicitly
fmt.Printf("%s\n", c)                                       // "100°C"
fmt.Println(c)                                              // "100°C"
fmt.Printf("%g\n", c)                                       // "100"; does not call String
fmt.Println(float64(c))                                     // "100"; does not call String
```

### 2.6 Packages and Files

- Packages in Go serve the same purposes as libraries or modules in other languages,  
    supporting modularity, encapsulation, separate compilation, and reuse.
-  In Go, a simple rule governs which identifiers are exported and which are not:  
    exported identifiers start with an upper-case letter.
- Each file starts with a package declaration that defines the package name.

Package example:

- First create a directory 'tempconv'
- Create 'tempconv/tempconv.go'

```go
// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
```

- Create 'tempconv/conv.go'

```go
package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
```

- Package-level names like the types and constants declared in one file of a package  
    are **visible to all the other files of the package**, as if the source code were all insingle file.
- Because the package-level const names begin with upper-case letters,  
    they too are accessible with qualified names like `tempconv.AbsoluteZeroC`

#### 2.6.1 Imports

- Within a Go program, every package is identified by a unique string called its import path.  
    These are the strings that appear in an import declaration like "gopl.io/ch2/tempconv"
- An import path denotes a directory containing one or more Go source files that together make up the package.

cf.go

```go
// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"./tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
```

- For this example I use a relative path to import `tempconv` package but it's far more clean to use $GOPATH/src/[website]/[author]/[package_name].

#### 2.6.2 Package Initialization

- Within each file, init functions are automatically executed when the program starts, in the order in which they are declared.

```go
func init() { /* ... */ }
```

- One package is initialized at a time, in the order of imports in the program, dependencies first,  
    so a package `p` importing `q` can be sure that `q` is fully initialized before `p`’s initialization begins.
- The `main` package is the last to be initialized.  
    In this manner, all packages are fully initialized before the application’s `main` function begins.

Popcount package with init() example

```go
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
```

- Note that the range loop in init uses only the index;  
    the value is unnecessary and thus need not be included. The loop could also have been written as

### 2.7 Scope

- The *scope* of a declaration is the part of the source code where a use of the declared name refers to that declaration.
- The declarations of built-in types, functions, and constants like `int`, `len`, and `true` are in **the universe block**  
    and can be referred to throughout the entire program.
- Declarations outside any function, that is, at **package level**, can be referred to from any file in the same package.

Shadow variable example:

```go
func f() {}

var g = "g"

func main() {
    f := "f"
    fmt.Println(f) // "f"; local var f shadows package-level func f
    fmt.Println(g) // "g"; package-level var
    fmt.Println(h) // compile error: undefined: h
}
```

Scope rules example (Ugly Style)

```go
func main() {
    x := "hello!"
    for i := 0; i < len(x); i++ {
        x := x[i]
        if x != '!' {
            x := x + 'A' - 'a'
            fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
        }
    }
}
```

- Like `for` loops, `if` statements and `switch` statements also create implicit blocks in addition to their body blocks.

```go
if x := f(); x == 0 {
    fmt.Println(x)
} else if y := g(x); x == y {
    fmt.Println(x, y)
} else {
    fmt.Println(x, y)
}

fmt.Println(x, y) // compile error: x and y are not visible here
```

- Beware of this:

```go
if f, err := os.Open(fname); err != nil { // compile error: unused: f
    return err
}

f.ReadByte()    // compile error: undefined f
f.Close()       // compile error: undefined f
```

```go
var cwd string

if

func init() {
    cwd, err := os.Getwd() // compile error: unused: cwd
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}
```

Solution: Declare `err` before

```go
var cwd string

func init() {
    var err error
	if cwd, err = os.Getwd(); err != nil {
		log.Fatal("os.Getwd failed: %v", err)
	}
}
```

## 3. Basic Data Types

- Go’s types fall into four categories: 
    - `basic types`: 
        - numbers
        - strings
        - booleans
    - `aggregate types`:
        - arrays
        - structs
    - `reference types`:
        - pointers
        - slices
        - maps
        - functions
        - channels
    - `interface types`

### 3.1 Integers

- Go integer types:
    - Signed:
        - int8
        - int16
        - int32
        - int64
    - Unsigned:
        - uint8
        - uint16
        - uint32
        - uint64

- The type `rune` is a synonym for `int32` and conventionally indicates that a value is a Unicode code point.  
    The two names may be used interchangeably. Similarly, the type `byte` is a synonym for `uint8`,  
    and emphasizes that the value is a piece of raw data rather than a small numeric quantity.

- Finally, there is an unsigned integer type `uintptr`,  
    whose width is not specified but is sufficient to hold all the bits of a pointer value.  
    The `uintptr` type is used only for low-level programming, such as at the boundary of a Go program with a C library or an operating system.

- An explicit conversion is required to use an `int` value where an `int32` is needed, and vice versa.
- For instance, the range of `int8` is −128 to 127, whereas the range of `uint8` is 0 to 255.

- Go binary operators:
    - `*` `/` `%` `<<` `>>` `&` `&^`
    - `+` `-` `|` `^`
    - `==` `!=` `<` `<=` `>` `>=`
    - `&&`
    - `||`

- If the result of an arithmetic operation, whether signed or unsigned, has more bits than can be represented in the result type,  
    it is said to *overflow*.  
    The high-order bits that do not fit are silently discarded.  
    If the original number is a signed type, the result could be negative if the leftmost bit is a 1, as in the `int8`  
    example here:

```go
var u uint8 = 255
fmt.Println(u, u+1, u*u) // "255 0 1"

var i int8 = 127
fmt.Println(i, i+1, i*i) // "127 -128 1"
```

- Go comparison operators:
    - `==`          equal to
    - `!=`          not equal to
    - `<`           less than
    - `<=`          less than or equal to
    - `>`           greater than
    - `>=`          greater than or equal to

- Go unary operators:
    - `+`           unary positive(no effect)
    - `-`           unary negation

- Go bitwise binary operators:
    - `&`           bitwise AND
    - `|`           bitwise OR
    - `^`           bitwise XOR
    - `&^`          bit clear(AND NOT)
    - `<<`          left shift
    - `>>`          right shift

- The operator `^` is bitwise exclusive OR (XOR) when used as a binary operator,  
    but when used as a unary prefix operator it is bitwise negation or complement;  
    that is, it returns a value with each bit in its operand inverted.
