# JavaScript Quick syntax Reference

## Using JavaScript

### Embedding Javascript

`<script> ... </script>`

Or

`<script src="mycode.js"></script>`

### Displaying Text

```html
<script>
    document.write("Hello World");
</script>
```

### Browser Compatibility

```html
<noscript>
    Please enable JavaScript for full functionality of this site.
</noscript>
```

### Console Window

```html
<script>
    console.log("Hello Console");
</script>
```

### Comments

```html
<script>
    // single-line comment
    /* multi-line
       comment    */
</script>
```

## Variables

### Declaring Variables

```html
<script>
    var myVar;

    myVar = 10;

    var myVar = 10;

    var myVar = 10, myVar2 = 20, myVar3;

    document.write(myVar);  // "10"
</script>
```

### Dynamic Typing

```html
<script>
    var myType = "Hi";      // string type
    myType = 1.5;           // number type

    console.log(myType);    // "1.5"
</script>
```

### Number Type

```html
<script>
    var dec = 10;           // decimal notation
    var oct = 012;          // octal notation
    var hex = 0xA;          // hexadecimal notation

    var num = 1.23;
    var exp = 3e2;          // 3*10^2 = 300
</script>
```

### Bool Type

```html
<script>
    var myBool = true;
</script>
```

### Undefined Type

```html
<script>
    var myUndefined;
    console.log(myUndefined);           // "undefined"
    console.log(typeof myUndefined);    // "undefined"

    console.log(myUndeclared);          // throws a ReferenceError
</script>
```

### Null Type

```html
<script>
    var myNull = null;
    console.log(myNull);                // "null"

    console.log(typeof myNull);         // "object"

    // Boolean context
    console.log(!!null);                // false
    console.log(!!undefined);           // false

    // Numeric context
    console.log(null * 5);              // "0"
    console.log(undefined * 5)          // "NaN"
</script>
```

### Special Numeric Values

```html
<script>
    console.log(1 / 0);                 // "Infinity"
    console.log(-1 / 0);                // "-Infinity"
    console.log(0 / 0);                 // "NaN"

    var myNaN = Math.sqrt(-1);
    console.log(myNaN);                 // "NaN"
    console.log(typeof myNaN);          // "number"

    console.log("Hi" * 3);              // "NaN"

    console.log(NaN == NaN);            // "false"
    console.log(isNaN(myNaN));          // "true"
</script>
```

## Operators

### Arithmetic Operators

```html
<script>
    x = 3 + 2;              // 5   # Addition
    x = 3 - 2;              // 1   # Substraction
    x = 3 * 2;              // 6   # Multiplication
    x = 3 / 2;              // 1.5 # division
    x = 3 % 2;              // 1   # modulus (division remainder)
</script>
```

### Assignment Operators

```html
<script>
    x = 0;                  // assignment
</script>
```

### Combined Assignment Operators

```html
<script>
    x += 5;                 // x = x + 5;
    x -= 5;                 // x = x - 5;
    x *= 5;                 // x = x * 5;
    x /= 5;                 // x = x / 5;
    x %= 5;                 // x = x % 5;
</script>
```

### Increment and Decrement Operators

```html
<script>
    x++;                    // x = x + 1;
    x--;                    // x = x - 1;

    x = 5; y = x++;         // y=5, x=6
    x = 5; y = ++x;         // y=6, x=6
</script>
```

### Comparison Operators

```html
<script>
    x = (2 == 3);           // false # equal to
    x = (2 === 3);          // false # identical
    x = (2 !== 3);          // true  # not identical
    x = (2 != 3);           // true  # not equal to
    x = (2 > 3);            // false # greater than
    x = (2 < 3);            // true  # less than
    x = (2 >= 3);           // false # greater than or equal to
    x = (2 <= 3);           // true  # less than or equal to

    x = (1 == "1");         // true (same value)
    x = (1 === "1");        // false (different types)
</script>
```

- It is considered good practice to use strict comparison  
    when the type conversion feature of the equal to operation is not needed.

### Logical Operators

```html
<script>
    x = (true && false);    // false # logical and
    x = (true || false);    // true  # logical or
    x = !(true);            // false # logical not
</script>
```

### Bitwise Operators

```html
<script>
    x = 5 & 4;              // 101 & 100 = 100      (4)  # And
    x = 5 | 4;              // 101 | 100 = 101      (5)  # Or
    x = 5 ^ 4;              // 101 ^ 100 = 001      (1)  # XOr
    x = 4 << 1;             // 100 << 1  = 1000     (8)  # Left shift
    x = 4 >> 1;             // 100 >> 1  = 010      (2)  # Right shift
    x = 4 >>> 1;            // 100 >>> 1 = 010      (2)  # Zero-fill right shift
    x = ~4;                 // ~00000100 = 11111011 (-5) # Invert
</script>
```

- The bitwise operators also have combined assignment operators:
    - `&=`
    - `|=`
    - `^=`
    - `<<=`
    - `>>=`
    - `>>>=`

- JavaScript numbers are stored as double precision floating-point numbers.  
    However, the bitwise operations need to operate on integers and  
    therefore numbers are temporarily converted to 32-bit signed integers when bitwise operations are performed.

## Arrays

### Numeric Arrays

```html
<script>
    var a = new Array();    // empty array

    a[0] = 1;
    a[1] = 2;
    a[2] = 3;

    // The initial capacity of an array can be specified by passing a single numeric parameter
    var b = new Array(3);

    // Passing more than one argument, or a non-numeric argument, to the array constructor
    var c = new Array(1, 2, 3);

    // Literal declaration
    var d = [1, 2, 3];

    var e = [];             // empty array

    var f =  [1, 2, 3];
    document.write(f[0] + f[1] + f[2]);     // "6"

    document.write(f[3]);                   // "undefined"

    // An array can store any data type or combination
    var mixed = [0, 3.14, "string", true];
</script>
```

### Associative Arrays

```html
<script>
    var g     = new Array();
    g["name"] = "Peter";
    g["age"]  = 25;

    document.write(g["name"] + " is " + g["age"]);      // "Peter is 25"

    // Arrays in Javascript are objects and their elements are object properties.
    // Therefore, elements of associative arrays can alternatively be referenced using the dot notation.
    var h = new Array();
    h.name = "Peter";
    h.age = 25;
    document.write(h.name + " is " + h.age);            // "Peter is 25"
</script>
```

> P. 14
