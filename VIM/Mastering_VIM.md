# Mastering VIM

---

## -[ COMMAND LINE ]-

```

```

## -[ Abreviations ]-

- **{n}** Number
- **{to}** Text Object
- **{rg}** Register

## -[ General Form of VI commands ]-

- (command)(number)(text object)
- (number)(command)(text object)

---

## -[ TEXT OBJECT ]-

```
{n}w        {n} word under the cursor
{n}W        {n} Word under the cursor(Withespace separated)
{n}b        {n} Word before the cursor
{n}B        {n} Word before the cursor(Withespace separated)
{n}l        {n} character(Only for Yank)
{n}h        {n} character(Only for Yank)
$           To the end of the line
0           To the begining of the line
```

## -[ COMMAND MODE ] -

```
:   EX MODE
i   INSERT MODE
ZZ  Save and exit
```

### /- Single Movements ->

```
h   Left
j   Down
k   Up
l   Right

0   Begining of the line
$   End of the line
```

### /- Block Movements ->

```
w   Forward one word(Special characters count one word)
W   Forward one word(Withespace separated)
b   Backward one word(Special characters count one word)
B   Backward one word(Withespace separated)
e   Forward to the end of the word(Special characters count one word)
E   Forward to the end of the word(Withespace separated)

G     End of the file
{n}G  Go to line {n}
```

### /- Simple Edit ->

```
i          Insert (under the cursor)
I          Insert at the begining of the line
a          Append (after the cursor)
A          Append at the end of the line

o          Open an empty line below the cursor
O          Open an empty line above the cursor

J          Join the current line and the line under

c{to}      Change the text object(Start at the cursor position)
cc         Change all the current line
C          Change from the cursor to the end of the current line

r          Replace one character
R          Enter replace mode (replace until ESC)

s          Replace one character and enter insert mode (Alias for 'c ')
S          Delete the entire line and enter insert mode (Alias for 'cc')

d{to}      Delete the text object(Start at the cursor position)
dd         Delete the current line
D          Delete characters from the cursor to the end of the line (Alias for 'd$')

x          Delete character under the cursor
X          Delete character before the cursor

{rg}p      Put the text from the register after the cursor(It's a PASTE)
p          Put the text fromt the register "0 after the cursor

{rg}P      Put the text from the register before the cursor(It's a PASTE)
P          Put the text from the register "0 before the cursor

{rg}y{to}  Yank(Copy) the text object to the register
y{to}      Yank(Copy) the text object to the register "0
yy         Yank(Copy) the current line(Alias for 'y$')
Y          Yank(Copy) the current line(Alias for 'y$')

.          Repeat the last command

u          Undo the last command
U          Undo all edit on the current line
[CTRL]+R   Redo the last undo command

xp         Swap 2 characters
~          Swap uppercase/lowercase
```

## -[ EX MODE ] -

```
:e <FILENAME>   Open/Edit a file
:e!             Reload Current file
:q              Exit VIM
:q!             Force exit without saving
:w              Save current buffer
:w <FILENAME>   Save current buffer in a new file
:w! <FILENAME>  Save current buffer in an existing file
```

## -[ INSERT MODE ] -

```
ESC   Exit INSERT MODE
```

## -[ VISUAL MODE ] -
