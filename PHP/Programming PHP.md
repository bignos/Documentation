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

#### Integers

- Use the `is_int()` function (or its ̀`is_integer()` alias) to test whether a value is an integer

```php
if (is_int($x)) {
// $x is an integer
}
```

#### Floating-Point Numbers

- Use the `is_float()` function (or its `is_real()` alias) to test whether a value is a floating-point number:

```php
if (is_float($x)) {
// $x is a floating-point number
}

```

#### Strings
    - Variables are expanded (interpolated) within double quotes, while within single quotes they are not.
    - Use the `is_string()` function to test whether a value is a string.

#### Booleans
    - PHP defines some values as true and others as false.
    - In PHP, the following values all evaluate to false:
        - The *keyword* `false`
        - The *integer* `0`
        - The *floating-point* value `0.0`
        - The empty *string* (`""`) and the *string* `"0"`
        - An *array* with zero elements
        - An *object* with no values or functions
        - The `NULL` value

#### Arrays
    - An array holds a group of values, which you can *identify by position* (a number, with 0 being the first position)  
        or some *identifying name* (a string), called an *associative index*.

```php
$person[0]                = "Edison";
$person[1]                = "Wankel";
$person[2]                = "Crapper";

$creator['Light bulb']    = "Edison";
$creator['Rotary Engine'] = "Wankel";
$creator['Toilet']        = "Crapper";
```

- The `array()` construct creates an array. Here are two examples:

```php
$person = array("Edison", "Wankel", "Crapper");
$creator = array('Light bulb'    => "Edison",
                 'Rotary Engine' => "Wankel",
                 'Toilet'        => "Crapper");
```

- There are several ways to *loop* through arrays, but the most common is a `foreach` *loop*:

```php
foreach ($person as $name) {
	echo "Hello, {$name}\n";
}

foreach ($creator as $invention => $inventor) {
	echo "{$inventor} created the {$invention}\n";
}
```

- You can *sort* the elements of an array with the various `sort` functions:

```php
sort($person);
// $person is now array("Crapper", "Edison", "Wankel")

asort($creator);
// $creator is now array('Toilet'        => "Crapper",
//                       'Light bulb'    => "Edison",
//                       'Rotary Engine' => "Wankel");
```

#### Objects

- Classes are the building blocks of object-oriented design.  
    A *class* is a definition of a structure that contains *properties* (variables)  
    and *methods* (functions).  
    Classes are defined with the `class` keyword:

```php
class Person
{
    public $name = '';

    function name ($newname = NULL)
    {
        if (!is_null($newname)) {
            $this->name = $newname;
        }

        return $this->name
    }
}
```

- Once a *class* is defined, any number of objects can be made from it with the `new` keyword,  
    and the object’s *properties* and *methods* can be accessed with the `->` construct:

```php
$ed = new Person;
$ed->name('Edison');
echo "Hello, {$ed->name}\n";

$tc = new Person;
$tc->name('Crapper');
echo "Look out below {$tc->name}\n";
```

- Use the `is_object()` function to test whether a value is an object:

```php
if (is_object($x)) {
// $x is an object
}
```

#### Resources

-  For example, every database extension has at least a function to *connect* to the database,  
    a function to *send* a query to the database,  
    and a function to *close* the connection to the database.  
    Because you can have multiple database connections open at once,  
    the `connect` function gives you something by which to identify that unique connection  
    when you call the `query` and ̀`close` functions: a *resource* (or a “handle”).

- Each active resource has a unique *identifier*.  
    Each *identifier* is a *numerical index* into an *internal PHP lookup table* that holds information about all the active resources.  
    PHP maintains information about each resource in this table, including   
    the number of references to (or uses of) the resource throughout the code.  
    When the last reference toresource value goes away,  
    the extension that created the resource is called to *free any memory*, *close any connection, etc*., for that resource:

```php
$res = database_connect();
database_query($res);
// fictitious database connect function

$res = "boo";
// database connection automatically closed because $res is redefined
```

- When there are no more references to the resource, it’s automatically shut down.

- Use the `is_resource()` function to test whether a value is a *resource*:

```php
if (is_resource($x)) {
// $x is a resource
}
```

#### Callbacks

- Callbacks are functions or object methods used by some functions, such as `call_user_func()`.  
- Callbacks can also be created by the `create_function()` method and through closures.

```php
$callback = function myCallbackFunction()
    {
        echo "callback achieved";
    }

call_user_func($callback);
```

#### NULL

- There’s only one value of the *NULL* data type.  
    That value is available through the caseinsensitive keyword `NULL`.  
    The `NULL` value represents a variable that has no value (similar to Perl’s `undef` or Python’s `None`):

```php
$aleph = "beta";

$aleph = null; // variable's value is gone
$aleph = Null; // same
$aleph = NULL; // same
```

- Use the is_null() function to test whether a value is NULL—for instance, to see whether a variable has a value:

```php
if (is_null($x)) {
// $x is NULL
}
```

#### Variables

- Variables in PHP are identifiers prefixed with a dollar sign (`$`).
- A variable may hold a value of any type.  
    There is no compile-time or runtime type checking on variables.  
    You can replace a variable’s value with another of a different type:

```php
$what = "Fred";
$what = 35;
$what = array("Fred", 35, "Wilma");
```

- There is *no explicit syntax* for declaring variables in PHP.  
    The first time the value of a variable is set, the variable is created.  
    In other words, setting a value to a variable also functions as a declaration.

- A variable whose value has not been set behaves like the NULL value:

```php
if ($uninitializedVariable === NULL) {
    echo "Undefined variable";
}
```

#### Variable Variables

- You can reference the value of a variable whose name is stored in another variable c
    by prefacing the variable reference with an additional dollar sign (`$`). For example:

```php
$foo  = "bar";
$$foo = "baz";
```

- After the second statement executes, the variable `$bar` has the value "baz".

#### Variable References

- In PHP, *references* are how you create *variable aliases*.  
    To make `$black` an alias for the variable `$white`, use:

```php
$black =& $white;
```

- After the assignment, the two variables are alternate names for the same value.
- Unsetting a variable that is aliased does not affect other names for that variable’s value, however:

```php
$white = "snow";
$black =& $white;

unset($white);
print $black; // "snow"
```

- Functions can return values by reference

```php
function &retRef() // note the &
{
    $var = "PHP";
    return $var;
}

$v =& retRef(); // note the &
```

### Variable Scope

- There are 4 types of variable scope in PHP: 
    - local
    - global
    - static
    - function parameters

#### Local scope

- A variable declared in a function is *local* to that function.
- By default, variables defined outside a function (called *global variables*) are not accessible inside the function.

```php
function updateCounter()
{
    $counter++;
}

$counter = 10;

updateCounter();
echo $counter; // 10
```

- **Only functions can provide local scope**.  
    Unlike in other languages, in PHP you **can’t create** a variable whose scope is: 
    - a loop, 
    - conditional branch 
    - other type of block

#### Global scope

- Variables declared outside a function are global.
- To allow a function to access a global variable,  
    you can use the `global` keyword inside the function to declare the variable within the function.

```php
function updateCounter()
{
    global $counter;
    $counter++;
}

$counter = 10;
updateCounter();

echo $counter; // 11
```

#### Static variables

- A *static* variable retains its value between calls to a function but is visible only within that function. 
- You declare a variable *static* with the `static` keyword.

```php
function updateCounter()
{
    static $counter = 0;
    $counter++;
    echo "Static counter is now {$counter}\n";
}

$counter = 10;
updateCounter(); // Static counter is now 1
updateCounter(); // Static counter is now 2
echo "Global counter is {$counter}\n"; // Global counter is 10
```

#### Function parameters

- Function parameters are *local*, meaning that they are available only inside their functions.

### Garbage Collection

- When you copy a value from one variable to another, PHP doesn’t get more memory for a copy of the value.  
    Instead, it updates the symbol table to indicate that “both of these variables are names for the same chunk of memory.”  
    So the following code doesn’t actually create a new array:

```php
$worker = array("Fred", 35, "Wilma");
$other = $worker; // array isn't copied
```

- If you subsequently modify either copy, PHP allocates the required memory and makes the copy:

```php
$worker[1] = 36; // array is copied, value changed
```

- By delaying the allocation and copying, PHP saves time and memory in a lot of situations.
- This is *copy-on-write*.

- When the reference count of a value reaches 0, its memory is released.
- This is reference counting.

-  If you do insist on trying to get a little more information or control over freeing a variable’s value,  
    use the `isset()` and `unset()` functions

- To see if a variable has been set to something—even the empty string—use `isset()`:

```php
$s1 = isset($name); // $s1 is false
$name = "Fred";
$s2 = isset($name); // $s2 is true
```

- Use `unset()` to remove a variable’s value:

```php
$name = "Fred";
unset($name); // $name is NULL
```

### Expressions and Operators

```
[Associativity] 	 [Operators]                                              	 [Additional Information]
non-associative 	 clone new                                                	 clone and new
left            	 [                                                        	 array()
right           	 **                                                       	 arithmetic
right           	 ++ -- ~ (int) (float) (string) (array) (object) (bool) @ 	 types and increment/decrement
non-associative 	 instanceof                                               	 types
right           	 !                                                        	 logical
left            	 * / %                                                    	 arithmetic
left            	 + - .                                                    	 arithmetic and string
left            	 << >>                                                    	 bitwise
non-associative 	 < <= > >=                                                	 comparison
non-associative 	 == != === !== <> <=>                                     	 comparison
left            	 &                                                        	 bitwise and references
left            	 ^                                                        	 bitwise
left            	 |                                                        	 bitwise
left            	 &&                                                       	 logical
left            	 ||                                                       	 logical
right           	 ??                                                       	 comparison
left            	 ? :                                                      	 ternary
right           	 = += -= *= **= /= .= %= &= |= ^= <<= >>=                 	 assignment
right           	 yield from                                               	 yield from
right           	 yield                                                    	 yield
left            	 and                                                      	 logical
left            	 xor                                                      	 logical
left            	 or                                                       	 logical
```

#### Implicit Casting


