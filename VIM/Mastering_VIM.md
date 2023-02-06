# Mastering VIM

---

## -[ COMMAND LINE ]-

```
vim <FILENAME>                               Open the file <FILENAME> with VIM
vim <FILENAME1> <FILENAME2> .. <FILENAMEn>   Open multiple files with VIM
ex <FILENAME>                                Open the file <FILENAME> with VIM in Ex mode

vim -c {n} <FILENAME>                        Open <FILENAME> at the line {n}
vim -c /{pt} <FILENAME>                      Open <FILENAME> at the first occurence of pattern {pt}
vim + <FILENAME>                             Open <FILENAME> at the last line
vim +{n} <FILENAME>                          Open <FILENAME> at the line {n}

vim -R <FILENAME>                            Open <FILENAME> in read only mode
view <FILENAME>                              Open <FILENAME> in read only mode

vim -r                                       List all saved buffer by VI(Used for recovery)
ex -r                                        List all saved buffer by VI(Used for recovery)
vim -r <BUFFER>                              Recover the edited <BUFFER>
```

## -[ Abreviations ]-

- **{n}** Number
- **{ch}** A character
- **{CH}** An uppercase character
- **{to}** Text Object
- **{rg}** Register
- **{pt}** Regular expression pattern
- **{rpt}** Replacement expression pattern
- **{rm}** Regular expression modifier
- **{ec}** EX command

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

/{pt}       To the first word that match pattern {pt} forward
?{pt}       To the first word that match pattern {pt} backward

f{ch}       To the next occurrence of character {ch} on the current line
F{ch}       To the previous occurrence of character {ch} on the current line
t{ch}       Before the next occurrence of character {ch} on the current line
T{ch}       After the previous occurrence of character {ch} on the current line
```

## -[ REGISTER ]-

```
"{n}        Numbered register[1-9], the last nine deletions, from most to least recent
"{ch}       Named register[a-z], use like user clipboard
"{CH}       Named register, but when you use uppercase character, you append the register(Accumulator)
```

## -[ Marker ]-

```
m{ch}       Mark the current position with {ch}
'{ch}       Goto the first character of the line marked by {ch}
''          Goto the first character of the previous mark or context
`{ch}       Goto the position of the mark {ch}
``          Goto the position of the previous mark or context
```

## -[ COMMAND MODE ] -

```
:   EX MODE
gQ  FULL EX MODE

i   INSERT MODE

ZZ  Save and exit
```

### /- Single Movements ->

```
h     Left
j     Down
k     Up
l     Right

0     Begining of the line
^     Move to the first nonblank character of the line
$     End of the line
{n}|  Move to the character {n} on the current line
```

### /- Block Movements ->

```
w   Forward one word(Special characters count one word)
W   Forward one word(Withespace separated)
b   Backward one word(Special characters count one word)
B   Backward one word(Withespace separated)
e   Forward to the end of the word(Special characters count one word)
E   Forward to the end of the word(Withespace separated)

(   Move to the begining of current sentence
)   Move to the begining of the next sentence
{   Move to the begining of current paragraph
}   Move to the begining of the next paragraph
[[   Move to the begining of current section
]]   Move to the begining of the next section

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

### /- Line Movements ->

```
[ENTER]       Move to the first character of the next line
+             Move to the first character of the next line
-             Move to the first character of the previous line
^             Move to the first nonblank character of the current line
{n}|          Move to the {n} character of the current line

```

### /- Screen Movements ->

```
[CTRL]+F      Scroll one screen forward
[CTRL]+B      Scroll one screen backward
[CTRL]+D      Scroll half screen forward(Down)
[CTRL]+U      Scroll half screen Backward(Up)
[CTRL]+E      Scroll the screen one line down
[CTRL]+Y      Scroll the screen one line up

z [ENTER]     Move the current line on the top of the screen
z.            Move the current line on the center of the screen
z-            Move the current line on the bottom of the screen
{n}z [ENTER]  Move the line {n} on top of the screen
{n}z.         Move the line {n} on the center of the screen
{n}z-         Move the line {n} on the bottom of the screen
```

### /- Search Movements ->

#### Text Search Movements

```
/{pt}       Search pattern {pt} forward
?{pt}       Search pattern {pt} backward

n           Repeat the search in forward direction
N           Repeat the search in backward direction
```

#### Line Search Movements

```
f{ch}       Find the next occurrence of character {ch} in the current line(Move cursor to)
F{ch}       Find the previous occurrence of character {ch} in the current line(Move cursor to)
t{ch}       Find the character before the next occurrence character {ch} in the current line(Move cursor to)
T{ch}       Find the character after the previous occurrence character {ch} in the current line(Move cursor to)

;           Repeat the previous find command in the same direction
,           Repeat the previous find command in the opposite direction
```

### /- Line number Movements ->

```
{n}G      Goto the line {n}
G         Goto the last line of the file

``        Goto the line before you use the last 'G' command(Return at the start)
''        Goto the start of the line before you use the last 'G' command(Return at the start)
```

## -[ EX MODE ] -

```
:e <FILENAME>             Open/Edit a file
:e!                       Reload Current file
:q                        Exit VIM
:q!                       Force exit without saving
:w                        Save current buffer
:w <FILENAME>             Save current buffer in a new file
:{n},{n}w <FILENAME>      Save the current range in a new file
:{n},{n}w >> <FILENAME>   Save the current range to the end of <FILENAME> (APPEND)
:w! <FILENAME>            Save current buffer in an existing file
:x                        Save current buffer and exit (LIKE 'ZZ' in COMMAND MODE)

:read <FILENAME>          Append current buffer with the content of <FILENAME>
:r <FILENAME>             Append current buffer with the content of <FILENAME>

:args                     List the files from the command line
:ar                       List the files from the command line
:n                        Edit next files from the command Line
:previous                 Edit previous files from the command line
:prev                     Edit previous files from the command line
:N                        Edit previous files from the command line
:rewind                   Edit the first file from the command line
:rew                      Edit the first file from the command line
:last                     Edit the last file from the command line


:{n}                      Goto the line {n}
:{n},{n}                  Range of lines

:s/{pt}/{rpt}/{rm}        Search the match pattern {pt} and replace with the replacement expression {rpt} and use regular expression modifier {rm}

:preserve                 Force the system to save the buffer(not the file)

:set number               Display line number
:set nu                   Display line number
:set nonumber             Hide line number
:set nonu                 Hide line number
:set nu!                  Toggle display/hide line number

:set nowrapscan           Stop search at the bottom(/{pt} or n) or at the top(?{pt} or N)
```

### /- Filename shortcut ->

```
%         Current filename
#         Alternate filename (Previous file)
```

## -[ FULL EX MODE ] -

```
:p                  Print current line
:[ENTER]            Goto next line
:{n}                Goto line {n}
:{n}p               Goto line {n} and print
:{n},{n}            Goto last line of the range and print the range
:{n};{n}            Goto last line of the range and print the range but the second number {n} is relative to the first
:{n},{n}#           Goto last line of the range and print the range with lines number
:$                  Goto last line of the buffer
:%                  Goto last line of the buffer and print all the buffer(%='All file' like 1,$)
:/{pt}              Goto the line that match the pattern {pt} forward
:?{pt}              Goto the line that match the pattern {pt} backward

:delete             Delete the current line
:d                  Delete the current line
:move{n}            Move current line to line {n}
:m{n}               Move current line to line {n}
:copy{n}            Copy current line and paste to line {n}
:co{n}              Copy current line and paste to line {n}
:t{n}               Copy current line and paste to line {n}

:=                  Print total numbers of lines
:.=                 Print current line number
:/{pt}/=            Print the line number of the first line that match the pattern {pt} from the current line

:/{pt}/d            Delete the next line containing pattern {pt}
:/{pt}/+d           Delete the line below the next line containing pattern {pt}
:/{pt1}/,/{pt2}/d   Delete from the first line containing pattern {pt1} to the first line containing pattern {pt2}

:g/{pt}             Print all lines that match the pattern {pt} (GLOBAL SEARCH)
:g!/{pt}            Print all lines that not match the pattern {pt} (GLOBAL SEARCH)
:v/{pt}             Print all lines that not match the pattern {pt} (GLOBAL SEARCH)
:g/{pt}/{ec}        Print all lines that match the pattern {pt} (GLOBAL SEARCH) and execute EX command {ec}

:ya {rg}            Yank(Copy) current line to the register {rg}
:pu {rg}            Put the content of register {rg} after the current line

:visual             Exit Full EX MODE (Return to visual editor[vi])
:vi                 Exit Full EX MODE (Return to visual editor[vi])
```

### /- FULL EX MODE Examples ->

```
:4,15d               Delete lines 4 to 15(inclusive)
:100,120m20          Move lines 100 to 120 on line 20
:100,120co20         Copy lines 100 to 120 on line 20

:.,$d                Delete from current line '.' to the end of the buffer '$'
:5,.m$               Move lines 5 to current line '.' on the end of the buffer '$'
:%d                  Delete all the buffer
:%t$                 Copy the buffer to the end of the file (consecutive duplicate)

:.,.+10d             Delete from the current line '.' to the next 10 lines '.+10'
:100,$m.-4           Move line 100 to the end of the buffer '$' on 4th line above '-4' current line '.'
:.,+10#              Display line number from current line to 10 lines below
:-,+t0               Copy 3 lines one above '-', the current line and one below '-' to the top of the buffer '0'
:10;+3d              Delete from line 10 to line 13 (10 + 3)

:1,5d | s/teh/the/   Delete line 1 to 5 and substitute 'teh' for 'the' on current line (before the first command it was the line 6)
```

## -[ INSERT MODE ] -

```
ESC   Exit INSERT MODE
```

## -[ VISUAL MODE ] -
