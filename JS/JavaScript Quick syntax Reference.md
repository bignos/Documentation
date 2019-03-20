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

>> P. 10
