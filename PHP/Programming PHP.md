# Programming PHP

- Use php command line to have a webserver, usefull for development environment:

```
php -S localhost:8000
```

## Language Basics

### Case Sensitivity

- The names of user-defined classes and functions, as well as built-in constructs and keywords  
    such as `echo`, `while`, `class`, etc., are case-insensitive. Thus, these three lines are equivalent:

```php
echo("hello,world");
ECHO("hello,world");
EcHo("hello,world");
```

- Variables, on the other hand, are case-sensitive.  
    That is, `$name`, `$NAME`, and `$NaME` are 3 different variables.

### Comments

- You can use 3 forms of comments in PHP

```php
# Shell style comment
// C++ style comment
/* C style comment */
```

### Identifiers

#### Constants

- A constant is an identifier for a simple value;  
    only scalar values—Boolean, integer, double, and string—can be constants.  
    Once set, the value of a constant cannot change.  
    Constants are referred to by their identifiers and are set using the `define()` function:

```php
define('PUBLISHER', "Adobe");
echo PUBLISHER;
```

### Keyword

- The keywords in PHP, which are case-insensitive

```
__CLASS__           echo           insteadof
__DIR__             else           interface
__FILE__            elseif         isset()
__FUNCTION__        empty()        list()
__LINE__            enddeclare     namespace
__METHOD__          endfor         new
__NAMESPACE__       endforeach     or
__TRAIT__           endif          print
__halt_compiler()   endswitch      private
abstract            endwhile       protected
and                 eval()         public
array()             exit()         require
as                  extends        require_once
break               final          return
callable            for            static
case                foreach        switch
catch               function       throw
class               global         trait
clone               goto           try
const               if             unset()
continue            implements     use
declare             include        var
default             include_once   while
die()               instanceof     xor
do
```

### Data types

- Integers
    - Use the `is_int()` function (or its ̀`is_integer()` alias) to test whether a value is an integer

```php
if (is_int($x)) {
// $x is an integer
}
```

- Floating-Point Numbers
    - Use the `is_float()` function (or its `is_real()` alias) to test whether a value is a floating-point number:

```php
if (is_float($x)) {
// $x is a floating-point number
}

```

- Strings
    - Variables are expanded (interpolated) within double quotes, while within single quotes they are not.
    - Use the `is_string()` function to test whether a value is a string.

- Booleans
    - PHP defines some values as true and others as false.
    - In PHP, the following values all evaluate to false:
        - The *keyword* `false`
        - The *integer* `0`
        - The *floating-point* value `0.0`
        - The empty *string* (`""`) and the *string* `"0"`
        - An *array* with zero elements
        - An *object* with no values or functions
        - The `NULL` value

