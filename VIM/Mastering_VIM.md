# Mastering VIM

---

## -[ Abreviations ]-

{n} Number
{to} Text Object

## -[ General Form of VI commands ]-

(command)(number)(text object)
(number)(command)(text object)

## -[ COMMAND LINE ]-

```

```

---

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
w   Forward one word(In this case punctuation is a word)
W   Forward one word(In this case word must be separate a space)
b   Backward one word(In this case punctuation is a word)
B   Backward one word(In this case word must be separate a space)

G     End of the file
{n}G  Go to line {n}
```

### /- Simple Edit ->

```
i       Insert before
a       Insert after

c{to}   Change
cc      Change all the current line
C       Change from the cursor to the end of the current line

r       Replace one character
R       Enter replace mode (replace until ESC)
s       Replace one character and enter insert mode
S       Delete the entire line and enter insert mode
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
