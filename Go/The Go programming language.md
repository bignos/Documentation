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

```go
var x uint8 = 1<<1 | 1<<5
var y uint8 = 1<<1 | 1<<2

fmt.Printf("%08b\n",x)      // "00100010", the set {1,5}
fmt.Printf("%08b\n",y)      // "00000110", the set {1,2}

fmt.Printf("%08b\n", x&y)   // "00000010", the intersection {1}
fmt.Printf("%08b\n", x|y)   // "00100110", the union {1, 2, 5}
fmt.Printf("%08b\n", x^y)   // "00100100", the symmetric difference {2, 5}
fmt.Printf("%08b\n", x&^y)  // "00100000", the difference {5}

for i := uint(0); i < 8; i++ {
    if x & (1<<i) != 0 {    // membership test
        fmt.Println(i)      // "1", "5"
    }
}

fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
```

- In the shift operations `x << n` and `x >> n`,  
    the `n` operand determines the number of bit positions to shift and must be unsigned;  
    the `x` operand may be unsigned or signed

- Arithmetically, a left shift `x<<n` is equivalent to *multiplication* by 2^n  
    and a right shift `x>>n` is equivalent to the floor of *division* by 2^n .

- It is important to use **unsigned** arithmetic when you’re treating an integer as bit pattern.

- You should avoid conversions in which the operand is out of range for the target type,  
    because the behavior depends on the implementation:

```go
f := 1e100      // a float64
i := int(f)     // result is implementation-dependent
```

- When printing numbers using the `fmt` package,  
    we can control the radix and format with the %d, %o, and %x verbs,  
    as shown in this example:

```go
o := 0666
fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"

x := int64(0xdeadbeef)
fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x) // Output: 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
```

- Note the use of two `fmt` tricks.  
    Usually a `Printf̀` format string containing multiple `%` verbs would require the same number of extra operands,  
    but the `[1]` “adverbs” after `%` tell `Printf` to use the first operand over and over again.
-  Second, the `#` adverb for `%o` or `%x` or `%X` tells Printf to emit a `0` or `0x` or `0X` prefix respectively.

- Rune literals are written as a character within single quotes.

Rune example:

```go
ascii   := 'a'
unicode := '€'
newline := '\n'

fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 € '€'"
fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
```

### 3.2 Floating-Point Numbers

- A `float32` provides approximately **6** decimal digits of precision,  
    whereas a `float64` provides about **15** digits; `float64` should be preferred for most purposes  
    because `float32` computations accumulate error rapidly unless one is quite careful,  
    and the smallest positive integer that cannot be exactly represented as a float32 is not large:

```go
var f float32 = 16777216  // 1 << 24
fmt.Println(f == f+1)     // "true"!
```

- Digits may be omitted before the decimal point `.707` or after it `1.`.
- Very small or very large numbers are better written in scientific notation, 
    with the letter `e` or `E` preceding the decimal exponent:

```go
const Avogadro = 6.02214129e23
const Planck = 6.62606957e-34
```

- Floating-point values are *conveniently* printed with Printf’s `%g` verb,  
    which chooses the most compact representation that has adequate precision,  
    but for tables of data, the `%e` (exponent) or `%f` (no exponent) forms may be more appropriate.  
    All 3 verbs allow field width and numeric precision to be controlled.

```go
for x := 0; x < 8; x++ {
    fmt.Printf("x = %d exp = %8.3f\n", x, math.Exp(float64(x)))
}
```

- The code above prints the powers of e with **3** decimal digits of precision, aligned in an **8**-character field:

```
x = 0 exp =    1.000
x = 1 exp =    2.718
x = 2 exp =    7.389
x = 3 exp =   20.086
x = 4 exp =   54.598
x = 5 exp =  148.413
x = 6 exp =  403.429
x = 7 exp = 1096.633
```

- The next program illustrates floating-point graphics computation.  
    It plots a function of two variables z = f(x, y) as a wire mesh 3-D surface,  
    using Scalable Vector Graphics (SVG), a standard XML notation for line drawings.

```go
// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 1200, 640           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke- width: 0.7' "+
		"width='%d' height='%d'>",
		width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) { // Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)                        // Compute surface height z.
	sx := width/2 + (x-y)*cos30*xyscale // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
```

### 3.3 Complex Numbers

- Go provides two sizes of complex numbers, `complex64` and `complex128`,  
    whose components are `float32` and `float64` respectively.  
    The built-in function `complex` creates a complex number from its real and imaginary components,  
    and the built-in `real` and `imag` functions extract those components:

```go
var x complex128 = complex(1, 2)    // 1+2i
var y complex128 = complex(3, 4)    // 3+4i

fmt.Println(x*y)                    // "(-5+10i)"
fmt.Println(real(x*y))              // "-5"
fmt.Println(imag(x*y))              // "10"
```

- If a floating-point literal or decimal integer literal is immediately followed by `i`, such as `3.141592i` or `2i`,  
    it becomes an imaginary literal, denoting a complex number with a zero real component:

```go
fmt.Println(1i * 1i) // "(-1+0i)", i² = -1
```

- The declarations of `x` and `y` above can be simplified:

```go
x := 1 + 2i
y := 3 + 4i
```

- Complex numbers may be compared for equality with == and !=.  
    Two complex numbers are equal if their real parts are equal and their imaginary parts are equal.

- The `math/cmplx` package provides library functions for working with complex numbers,  
    such as the complex square root and exponentiation functions.

```go
fmt.Println(cmplx.Sqrt(-1)) // "(0+1i)"
```

- The following program uses complex128 arithmetic to generate a Mandelbrot set.

```go
// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y) // Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
```

### 3.4 Booleans

- A value of type `bool`, or *boolean*, has only two possible values, `true` and `false`.
- simplify redundant boolean expressions like `x==true` to `x`.
- There is no implicit conversion from a boolean value to a numeric value like `0` or `1`, or vice versa.
- Conversion function example:

```go
// btoi returns 1 if b is true and 0 if false.
func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
}

// itob reports whether i is non-zero.
func itob(i int) bool {
    return i != 0
}
```

### 3.5 Strings

- A string is an immutable sequence of bytes.
- The built-in `len` function returns the number of bytes (not `runes`) in a string,  
    and the index operation `s[i]` retrieves the i-th byte of string `s`, where `0 ≤ i < len(s)`.

```go
s := "hello, world"
fmt.Println(len(s))     // "12"
fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
fmt.Println(s[0:5])     // "hello"

// Simplified version
fmt.Println(s[:5])      // "hello"
fmt.Println(s[7:])      // "world"
fmt.Println(s[:])       // "hello, world"

// The + operator makes a new string by concatenating 2 strings
fmt.Println("goobye" + s[5:])   // "goodbye, world"
```

- Strings may be compared with comparison operators like `==` and `<`;  
    the comparison is done byte by byte, so the result is the natural lexicographic ordering.
- String values are immutable: the byte sequence contained in a string value can never be changed,  
    though of course we can assign a new value to a string variable. To append one string to another, for instance, we can write

```go
s := "left foot"
t := s
s += ", right foot"

fmt.Println(s)          // "left foot, right foot"
fmt.Println(t)          // "left foot"
```

- This does not modify the string that `s` originally held  
    but causes `s` to hold the new string formed by the `+=` statement;  
    meanwhile, `t` still contains the old string.

#### 3.5.1 Strings Literals

- A string value can be written as a *string literal*, a sequence of bytes enclosed in **double quotes**
- Because Go source files are always encoded in UTF-8 and  
    Go text strings are conventionally interpreted as UTF-8, we can include Unicode code points in string literals.

- ASCII escape sequences:
    - `\a`          "alert" or bell
    - `\b`          backspace
    - `\f`          form feed
    - `\n`          newline
    - `\r`          carriage return
    - `\t`          tab
    - `\v`          vertical tab
    - `\'`          single quote (only in the rune literal '\'')
    - `\"`          double quote (only within "..." literals)
    - `\\`          backslash

- A raw string literal is written \`...\`, using backquotes instead of double quotes.
- Raw string literals are a convenient way to write regular expressions,  
    which tend to have lots of backslashes.  
    They are also useful for HTML templates, JSON literals, command usage messages,  
    and the like, which often extend over multiple lines.
```go
const GoUsage = `Go is a tool for managing Go source code.
Usage:
go command [arguments]
`
```

#### 3.5.2 Unicode

- A standard number called a *Unicode code point* or, in Go terminology, a `rune`.
- The natural data type to hold a single `rune` is `int32`, and that’s what Go uses;  
    it has the synonym `rune` for precisely this purpose.

#### 3.5.3 UTF-8

- UTF-8 is a variable-length encoding of Unicode code points as bytes.  
    UTF-8 was invented by Ken Thompson and Rob Pike, two of the creators of Go,  
    and is now a Unicode standard.
- Go source files are always encoded in UTF-8,  
    and UTF-8 is the preferred encoding for text strings manipulated by Go programs.
- The `unicode` package provides functions for working with individual runes  
    (such as distinguishing letters from numbers, or converting an upper-case letter to a lower-case one),  
    and the `unicode/utf8` package provides functions for encoding and decoding runes as bytes using UTF-8.
- There are two forms, `\uhhhh` for a 16-bit value and `\Uhhhhhhhh` for a 32-bit value,  
    where each `h` is a hexadecimal digit.
- The need for the 32-bit form arises very infrequently.

```go
"\xe4\xb8\x96\xe7\x95\x8c"
"\u4e16\u754c"
"\U00004e16\U0000754c"
```

- A `rune` whose value is less than 256 may be written with a single hexadecimal escape,  
    such as '\x41' for 'A',  
    but for higher values, a \u or \U escape must be used.  
    Consequently, '\xe4\xb8\x96' is not a legal rune literal,  
    even though those 3 bytes are a valid UTF-8 encoding of a single code point.

```go
// Prefix
func HasPrefix(s, prefix string) bool {
    return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// Suffix
func HasSuffix(s, suffix string) bool {
    return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}

// Substring
func Contains(s, substr string) bool {
    for i := 0; i < len(s); i++ {
        if HasPrefix(s[i:], substr) {
            return true
        }
    }
    return false
}
```

- How to count `rune` in a string

```go
import "unicode/utf8"

s := "Hello, \u4e16\u754c"

fmt.Println(len(s))                     // "13"
fmt.Println(utf8.RuneCountInString(s))  // "9"

for i := 0; i < len(s); {
    r, size := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%d\t%c\n", i, r)
    i += size
}

// Simplified version to count UTF-8 characters
n := 0
for range s {
    n++
}

// The more simplified is to use utf8.DecodeRuneInString(s) function
```

- Each call to `DecodeRuneInString` returns `r`, the rune itself,  
    and size, the number of bytes occupied by the UTF-8 encoding of `r`.
- Each time a UTF-8 decoder, whether explicit in a call to `utf8.DecodeRuneInString`  
    or implicit in a range loop, consumes an unexpected input byte,  
    it generates a special Unicode replacement character, '\uFFFD'.

#### 3.5.4 Strings and Byte Slices

- Four standard packages are particularly important for manipulating strings: 
    - `bytes`
        - The *bytes* package has similar functions for manipulating slices of bytes, of type `[]byte`,  
            which share some properties with `strings`. Because strings are immutable,  
            building up strings incrementally can involve a lot of allocation and copying.  
            In such cases, it’s more efficient to use the `bytes.Buffer` type.
    - `strings`
        - The *strings* package provides many functions for: 
            - searching
            - replacing
            - comparing
            - trimming
            - splitting
            - joining
    - `strconv`
        - The *strconv* package provides functions for converting  
            boolean, integer, and floating-point values to and from their string representations,  
            and functions for quoting and unquoting strings.
    - `unicode`
        - The *unicode* package provides functions like `IsDigit`, `IsLetter`, `IsUpper`, and `IsLower` for classifying runes.  
            Each function takes a single rune argument and returns a boolean.

- The basename function below was inspired by the Unix shell utility of the same name.

```go
// basename removes directory components and a.suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename1(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

// Version with strings  library functions
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}
```

- The task is to take a string representation of an integer,  
    such as "12345", and insert commas every three places, as in "12,345".

```go
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]  // return s - 3 last char ',' s - 3 first char
}
```

- The *bytes package* provides the Buffer type for efficient manipulation of byte slices.  
    A `Buffer` starts out empty but grows as data of types like `string`, `byte`, and `[]byte` are written to it.  
    As the example below shows, a bytes.Buffer variable requires no initialization because its zero value is usable:

```go
// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
```

- When appending the UTF-8 encoding of an arbitrary `rune` to a `bytes.Buffer`,  
    it’s best to use bytes.Buffer’s `WriteRune` method, but `WriteByte` is fine for ASCII characters.

#### 3.5.5 Conversions between Strings and Numbers

- it’s often necessary to convert between numeric values and their string representations. 
    This is done with functions from the `strconv` package.  
    To convert an integer to a string, one option is to use `fmt.Sprintf`;  
    another is to use the function `strconv.Itoa` ("integer to ASCII")

```go
x := 123
y := fmt.Sprintf("%d", x)

fmt.Println(y, strconv.Itoa(x))             // "123 123"
```

- `FormatInt` and `FormatUint` can be used to format numbers in a different base:

```go
fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
```

- The `fmt.Printf` verbs `%b`, `%d`, `%u`, and `%x` are often more convenient than Format functions,  
    especially if we want to include additional information besides the number:

```go
s := fmt.Sprintf("x=%b", x)                 // "x=1111011"
```

- To parse a string representing an integer, use the `strconv` functions `Atoi` or `ParseInt`,  
    or `ParseUint`  for unsigned integers:

```go
x, err := strconv.Atoi("123")               // x is an int
y, err := strconv.ParseInt("123", 10, 64)   // base 10, up to 64 bits
```

### 3.6 Constants

- Constants are expressions whose value is known to the compiler  
    and whose evaluation is guaranteed to occur at compile time, not at run time.

- The underlying type of every constant is a basic type: 
    - boolean
    - string
    - number
- Many computations on constants can be completely evaluated at compile time,  
    reducing the work necessary at run time and enabling other compiler optimizations.

- A constant declaration may specify a type as well as a value,
    but in the absence of an explicit type,  
    the type is inferred from the expression on the right-hand side.

- When a sequence of constants is declared as a group,  
    the right-hand side expression may be omitted for all but the first of the group,  
    implying that the previous expression and its type should be used again. For example:

```go
const (
    a = 1
    b
    c = 2
    d
)

fmt.Println(a, b, c, d) // "1 1 2 2"
```

#### 3.6.1 The Constant Generator iota

- A `const` declaration may use the constant generator `iota`,  
    which is used to create a sequence of related values without spelling out each one explicitly.  
    In a `const` declaration, the value of `iota` begins at zero and increments by one for each item in the sequence.

```go
type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

- This declares Sunday to be 0, Monday to be 1, and so on.

- As iota increments, each constant is assigned the value of `1 << iota`,  
    which evaluates to successive powers of 2, each corresponding to a single bit.  
    We can use these constants within functions that test, set, or clear one or more of these bits:

```go
type Flags uint

const (
    FlagUp Flags = 1 << iota    // is up
    FlagBroadcast               // supports broadcast access capability
    FlagLoopback                // is a loopback interface
    FlagPointToPoint            // belongs to a point-to- point link
    FlagMulticast               // supports multicast access capability
)

func IsUp(v Flags) bool {
     return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
    *v &^= FlagUp
}

func SetBroadcast(v *Flags) {
    *v |= FlagBroadcast
}

func IsCast(v Flags) bool {
    return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
    var v Flags = FlagMulticast | FlagUp
    fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"

    TurnDown(&v)
    fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"

    SetBroadcast(&v)
    fmt.Printf("%b %t\n", v, IsUp(v)) // "10010 false"
    fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
```

#### 3.6.2 Untyped Constants

- Many constants are not committed to a particular type.  
    The compiler represents these uncommitted constants with **much greater numeric precision** than values of basic types,  
    and arithmetic on them is more precise than machine arithmetic;  
    you may assume at least 256 bits of precision.  
    There are six flavors of these uncommitted constants, called: 
    - `untyped boolean`
    - `untyped integer`
    - `untyped rune`
    - `untyped floating-point`
    - `untyped complex`
    - `untyped string`

```go
// Untyped constants example:
var x float32 = math.Pi
var y float64 = math.Pi
var z complex128 = math.Pi

var f float64 = 212
fmt.Println((f - 32) * 5 / 9)       // "100"; (f - 32) * 5 is a float64
fmt.Println(5 / 9 * (f - 32))       // "0"; 5/9 is an untyped integer, 0
fmt.Println(5.0 / 9.0 * (f - 32))   // "100"; 5.0/9.0 is an untyped float

var f float64 = 3 + 0i              // untyped complex -> float64
f = 2                               // untyped integer -> float64
f = 1e123                           // untyped floating-point -> float64
f = 'a'                             // untyped rune -> float64

i := 0                              // untyped integer; implicit int(0)
r := '\000'                         // untyped rune; implicit rune('\000')
f := 0.0                            // untyped floating-point; implicit float64(0.0)
c := 0i                             // untyped complex; implicit complex128(0i)

fmt.Printf("%T\n", 0)               // "int"
fmt.Printf("%T\n", 0.0)             // "float64"
fmt.Printf("%T\n", 0i)              // "complex128"
fmt.Printf("%T\n", '\000')          // "int32" (rune)
```

## 4. Composite Types

- `Arrays` and `structs` are *aggregate types*;  
    their values are concatenations of other values in memory.
- `Arrays` are *homogeneous* their elements all have the same type  
-   whereas `structs` are *heterogeneous*.
- Both `arrays` and `structs` are *fixed size*.  
    In contrast, `slices` and `maps` are *dynamic data structures* that grow as values are added.

### 4.1 Arrays

- An `array` is a *fixed-length* sequence of zero or more elements of a particular *type*. 
    Because of their *fixed length*, `arrays` are rarely used directly in Go.  
    `Slices`, which can grow and shrink, are much more *versatile*.

```go
var a [3]int                        // array of 3 integers
fmt.Println(a[0])                   // print the first element
fmt.Println(a[len(a)-1])            // print the last element, a[2]

// Print the indices and elements.
for i, v := range a {
    fmt.Printf("%d %d\n", i, v)
}

// Print the elements only.
for _, v := range a {
    fmt.Printf("%d\n", v)
}
```

-  We can use an *array literal* to initialize an `array` with a list of values

```go
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
fmt.Println(r[2])                   // "0"
```

- In an *array literal*, if an ellipsis *"..."* appears in place of the length,  
    the array length is determined by the number of initializers. 
    The definition of q can be simplified to:

```go
q := [...]int{1, 2, 3}
fmt.Printf("%T\n", q)               // "[3]int"
```

- Is also possible to specify a list of index and value pairs, like this:

```go
type Currency int

const (
    USD Currency = iota
    EUR
    GBP
    RMB
)

symbol := [...]string{USD: "$", EUR: "€", GBP: "£",RMB: "¥"}
fmt.Println(RMB, symbol[RMB])       // "3 ¥"
```

- Defines an array r with 100 elements, all zero except for the last, which has value −1:

```go
r := [...]int{99: -1}
```
- Arrays are comparable:

```go
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}

fmt.Println(a == b, a == c, b == c) // "true false false"

d := [3]int{1, 2}
fmt.Println(a == d)                 // compile error: cannot compare [2]int == [3]int
```

```go
import "crypto/sha256"

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
    // Output:
    // 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db0225871792
    // 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df6
    // false
    // [32]uint8
}
```

- Notice the `Printf` verbs: `%x` to print all the elements of an array or slice of bytes in *hexadecimal*,  
    `%t` to show a boolean, and `%T` to display the type of value.

- We can explicitly pass a *pointer* to an `array` 
    so that any modifications the function makes to array elements will be visible to the caller.  
    This function zeroes the contents of a [32]byte array:

```go
func zero(ptr *[32]byte) {
    *ptr = [32]byte{}
}
```

### 4.2 Slices

- Slices represent variable-length sequences whose elements all have the same type.  
    A slice type is written `[]T`, where the elements have type `T`;  
    it looks like an array type without a size.
- A slice has three components: 
    - a pointer
    - a length
    - a capacity.
- The built-in functions `len` and `cap` return those values.

- Shows an array of strings for the months of the year,  
    and two overlapping slices of it.  
    The array is declared as:

```go
months := [...]string{1: "January", /* ... */, 12: "December"}
```

- So *January* is months[1] and *December* is months[12].

- Let’s define overlapping slices for the second quarter and the northern summer:

```go
Q2 := months[4:7]
summer := months[6:9]

fmt.Println(Q2)             // ["April" "May" "June"]
fmt.Println(summer)         // ["June" "July" "August"]
```

- June is included in each and is the sole output of this (inefficient) test for common elements:

```go
for _, s := range summer {
    for _, q := range Q2 {
        if s == q {
            fmt.Printf("%s appears in both\n", s)
        }
    }
}
```

- Slicing beyond `cap(s)` causes a panic, but slicing beyond `len(s)` extends the slice,  
    so the result may be longer than the original:

```go
fmt.Println(summer[:20])    // panic: out of range
endlessSummer := summer[:5] // extend a slice (within capacity)
fmt.Println(endlessSummer)  // "[June July August September October]"
```

- Since a slice contains a pointer to an element of an array,  
    passing a slice to a function permits the function to modify the underlying array elements.  
    In other words, copying a slice creates an alias for the underlying array.  
- The function reverse reverses the elements of an []int slice in place, and it may be applied to slices of any length.

```go
// reverse reverses a slice of ints in place.
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

a := [...]int{0, 1, 2, 3, 4, 5}
reverse(a[:])
fmt.Println(a) // "[5 4 3 2 1 0]"
```

- A simple way to *rotate* a slice left by n elements is to apply the `reverse` function three times,  
    first to the leading `n` elements, then to the remaining elements,  
    and finally to the whole slice. (To rotate to the right, make the third call first.)

```go
s := []int{0, 1, 2, 3, 4, 5}

// Rotate s left by two positions.
reverse(s[:2])
reverse(s[2:])
reverse(s)

fmt.Println(s) // "[2 3 4 5 0 1]"
```

- Unlike arrays, slices are not comparable,  
    so we cannot use `==` to test whether two slices contain the same elements.  
    The standard library provides the highly optimized `bytes.Equal` function for comparing two slices of bytes ([]byte),  
    but for other types of slice, we must do the comparison ourselves:

```go
func equal(x, y []string) bool {
    if len(x) != len(y) {
        return false
    }
    for i := range x {
        if x[i] != y[i] {
            return false
        }
    }

    return true
}
```

- The only legal slice comparison is against `nil`

```go
if summer == nil { /* ... */ }
```

```go
var s []int     // len(s) == 0, s == nil
s = nil         // len(s) == 0, s == nil
s = []int(nil)  // len(s) == 0, s == nil
s = []int{}     // len(s) == 0, s != nil
```

- So, if you need to test whether a slice is empty, use `len(s) == 0`, not `s == nil`.

- The built-in function `make` creates a slice of a specified element type, length, and capacity.  
    The capacity argument may be omitted, in which case the capacity equals the length.

```go
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
```

#### 4.2.1 The append Function

- The built-in append function appends items to slices:

- The `append` function is crucial to understanding how slices work,  
    so let’s take a look at what is going on.  
    Here’s a version called `appendInt` that is specialized for `[]int slices`:

```go
func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        // There is room to grow. Extend the slice.
        z = x[:zlen]
    } else {
        // There is insufficient space. Allocate a new array.
        // Grow by doubling, for amortized linear complexity.
        zcap := zlen

        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }

        z = make([]int, zlen, zcap)
        copy(z, x) // a built-in function; see text
    }

    z[len(x)] = y
    return z
}
```

- For efficiency, the new array is usually somewhat larger than the minimum needed to hold `x` and `y`.  
    Expanding the array by doubling its size at each expansion  
    avoids an excessive number of allocations and ensures that appending a single element takes constant time on average.  
    This program demonstrates the effect:

```go
func main() {
    var x, y []int
    for i := 0; i < 10; i++ {
        y = appendInt(x, i)
        fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
        x = y
    }
}
```

- Output:

```
0 cap=1     [0]
1 cap=2     [0 1]
2 cap=4     [0 1 2]
3 cap=4     [0 1 2 3]
4 cap=8     [0 1 2 3 4]
5 cap=8     [0 1 2 3 4 5]
6 cap=8     [0 1 2 3 4 5 6]
7 cap=8     [0 1 2 3 4 5 6 7]
8 cap=16    [0 1 2 3 4 5 6 7 8]
9 cap=16    [0 1 2 3 4 5 6 7 8 9]
```

- The built-in `append` function may use a more sophisticated growth strategy than `appendInt`’s simplistic one.

- In this respect, slices are not “pure” reference types but resemble an aggregate type such as this struct:

```go
type IntSlice struct {
    ptr *int
    len, cap int
}
```

- Our `appendInt` function adds a single element to a slice,  
    but the built-in `append` lets us add more than one new element,  
    or even a whole slice of them.

```go
var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, 4, 5, 6)
x = append(x, x...)     // append the slice x
fmt.Println(x)          // "[1 2 3 4 5 6 1 2 3 4 5 6]"
```

- With the small modification shown below, we can match the behavior of the built-in `append`.  
    The ellipsis "..." in the declaration of appendInt makes the function variadic:  
    it accepts any number of final arguments.  
    The corresponding ellipsis in the call above to `append` shows how to supply a list of arguments from a slice.

```go
func appendInt(x []int, y ...int) []int {
    var z []int
    zlen := len(x) + len(y)

    // ...expand z to at least zlen...
    copy(z[len(x):], y)
    return z
}
```

#### 4.2.2 In-Place Slice Techniques

- Let’s see more examples of functions that, like `rotate` and `reverse`,  
    modify the elements of a slice in place.  
    Given a list of strings, the nonempty function returns the non-empty ones:

```go
// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }

    return strings[:i]
}
```

- The subtle part is that the input slice and the output slice share the same underlying array.  
    This avoids the need to allocate another array, though of course the contents of `data` are partly overwritten,  
    as evidenced by the second print statement:

```go
data := []string{"one", "", "three"}
fmt.Printf("%q\n", nonempty(data))  // `["one" "three"]`
fmt.Printf("%q\n", data)            // `["one" "three" "three"]`
```

- The `nonempty` function can also be written using `append`:

```go
func nonempty2(strings []string) []string {
    out := strings[:0] // zero-length slice of original
    for _, s := range strings {
        if s != "" {
            out = append(out, s)
        }
    }

    return out
}
```

- A slice can be used to implement a *stack*.  
    Given an initially empty slice stack,  
    we can push a new value onto the end of the slice with `append`:

```go
stack = append(stack, v)        // push v

// The top of the stack is the last element:
top := stack[len(stack)-1]      // top of stack

// shrinking the stack by popping that element is
stack = stack[:len(stack)-1]    // pop
```

- To remove an element from the middle of a slice,  
    preserving the order of the remaining elements,  
    use `copy` to slide the higher-numbered elements down by one to fill the gap:

```go
func remove(slice []int, i int) []int {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}

func main() {
    s := []int{5, 6, 7, 8, 9}
    fmt.Println(remove(s, 2))   // "[5 6 8 9]"
}
```

- And if we don’t need to preserve the order,  
    we can just move the last element into the gap:

```go
func remove(slice []int, i int) []int {
    slice[i] = slice[len(slice)-1]
    return slice[:len(slice)-1]
}
func main() {
    s := []int{5, 6, 7, 8, 9}
    fmt.Println(remove(s, 2)) // "[5 6 9 8]
}
```

### 4.3 Maps

- In Go, a `map` is a reference to a hash table,  
    and a map type is written `map[K]V`, where `K` and `V` are the **types** of its *keys* and *values*.
-  All of the **keys** in a given `map` are of the *same type*,  
    and all of the **values** are of the *same type*.
- The key *type* `K` must be comparable using `==`.

```go
ages := make(map[string]int) // mapping from strings to ints
```

- We can also use a map literal to create a new map populated with some initial key/value pairs:

```go
ages := map[string]int{
    "alice":    31,
    "charlie":  34,
}

// ==

ages := make(map[string]int)
ages["alice"]   = 31
ages["charlie"] = 34

// so an alternative expression for a new empty map is map[string]int{}.

// Map elements are accessed through the usual subscript notation:
ages["alice"] = 32
fmt.Println(ages["alice"]) // "32"

// Removed with the built-in function delete:
delete(ages, "alice") // remove element ages["alice"]
```

- map element is not a variable, and we cannot take its address:

```go
_ = &ages["bob"] // compile error: cannot take address of map element

```

- To enumerate all the key/value pairs in the map, we use a range-based for loop

```go
for name, age := range ages {
    fmt.Printf("%s\t%d\n", name, age)
}
```

- To enumerate the key/value pairs in order, we must sort the keys explicitly,  
    for instance, using the *Strings* function from the `sort` package if the keys are strings.  
    This is a common pattern:

```go
import "sort"

var names []string
for name := range ages {
    names = append(names, name)
}

sort.Strings(names)
for _, name := range names {
    fmt.Printf("%s\t%d\n", name, ages[name])
}
```

- Storing to a `nil` map causes a panic:
 
```go
ages["carol"] = 21 // panic: assignment to entry in nil map
```

- If the element type is numeric,  
    you might have to distinguish between a nonexistent element  
    and an element that happens to have the value zero,  
    using a test like this:

```go
if age, ok := ages["bob"]; !ok { /* ... */ }
```

- Subscripting a `map` in this context yields 2 values;  
    the 2nd is a *boolean* that reports whether the element was present.

- The *dedup* program uses a `map` whose keys represent  
    the set of lines that have already appeared  
    to ensure that subsequent occurrences are not printed.

```go
func main() {
    seen := make(map[string]bool) // a set of strings
    input := bufio.NewScanner(os.Stdin)

    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }

    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }
}
```

P. 163
