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

rvim <FILENAME>                              Open <FILENAME> in restrictive mode(no shell command allowed)
rview <FILENAME>                             Open <FILENAME> in restrictive mode(no shell command allowed) and read only mode

evim <FILENAME>                              Open <FILENAME> with easy mode(VIM for beginner)
eview <FILENAME>                             Open <FILENAME> with easy mode(VIM for beginner) and read only mode

vimdiff <FILENAME1> <FILENAME2>              Open VIM on DIFF mode to compare <FILENAME1> and <FILENAME2>

vim -r                                       List all saved buffer by VI(Used for recovery)
ex -r                                        List all saved buffer by VI(Used for recovery)
vim -r <BUFFER>                              Recover the edited <BUFFER>

ex -s <FILENAME> < <SCRIPT_FILENAME>         Execute the vim script <SCRIPT_FILENAME> on the file <FILENAME>
```

## -[ Abbreviations ]-

- **{n}** Number
- **{ch}** A character
- **{CH}** An uppercase character
- **{to}** Text Object
- **{rg}** Register
- **{nrg}** Named Register
- **{pt}** Regular expression pattern
- **{rpt}** Replacement expression pattern
- **{rm}** Regular expression modifier
- **{ec}** EX command
- **{fn}** Function name

## -[ General Form of VI commands ]-

- (command)(number)(text object)
- (number)(command)(text object)

---

## -[ TEXT OBJECT ]-

For more information about _Text objects_ check `:help text-objects`

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

## -[ Regular expression metacharacters ]-

For more information check `:help regexp`

| Metacharacter | Description                                                                                                                   |
| ------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------- | ----- |
| .             | Match any _single_ character except a new line                                                                                |
| \*            | Match 0 or more of the single character that immediately precedes it                                                          |
| ^             | Match only line that begin with pattern after '^'. If '^' is not a the begining of the expression it's just the '^' character |
| $             | Match only line that end with pattern before '$'. If '$' is not a the end of the expression it's just the '$' character       |
| \             | Used to escape a character like '\.' to match a point and not any _single_ character                                          |
| \\{n}         | Recall a _subpattern_ , {n} is between 1 to 9                                                                                 |
| [ ]           | Match any _one_ of the characters enclosed between the brackets                                                               |
| [^ ]          | Match any _one_ of the characters that is **NOT** enclosed between the brackets                                               |
| [: :]         | Match any character which is part of the character classes                                                                    |
| [. .]         | Match multicharacter sequence that should be treated as a unit                                                                |
| [= =]         | Match an equivalence class list a set of characters that should be considered equivalent (ex: 'e' and 'é')                    |
| \\( \\)       | Save the subpattern enclosed between \( and \) into a special holding space (\1 .. \9)                                        |
| \\<           | Match only character at the begining of a word                                                                                |
| \\>           | Match only character at the end of a word                                                                                     |
| &             | Replace the '&' with the entire text matched by the search pattern(**ONLY FOR REPLACEMENT PATTERN {rpt}**)                    |
| ~             | Replace the '~' with the last used replacement pattern(**ONLY FOR REPLACEMENT PATTERN {rpt}**)                                |
| \\u           | Force the next characters to be on uppercase (**ONLY FOR REPLACEMENT PATTERN {rpt}**)                                         |
| \\U           | Force all next characters to be on uppercase (**ONLY FOR REPLACEMENT PATTERN {rpt}**)                                         |
| \\l           | Force the next characters to be on lowercase (**ONLY FOR REPLACEMENT PATTERN {rpt}**)                                         |
| \\L           | Force all next characters to be on lowercase (**ONLY FOR REPLACEMENT PATTERN {rpt}**)                                         |
| \\            |                                                                                                                               | String choice (ex car\\ | moto) |
| \\&           | If the pattern before the \\& match the pattern after is evaluated (ex .*Tom\\&.*Jerry)                                       |
| \\+           | Match 1 or more                                                                                                               |
| \\=           | Match 0 or 1                                                                                                                  |
| \\?           | Match 0 or 1                                                                                                                  |
| \\{...}       | Repeat the match {n} times or {n,m} in acceptable range                                                                       |
| ~             | Match the last given replacement string                                                                                       |
| \\(...\\)     | Grouping                                                                                                                      |
| \\{n}         | Call group {n} capture                                                                                                        |

### /- Regular expression character class -\

| Character class | Description                                                                                               |
| --------------- | --------------------------------------------------------------------------------------------------------- |
| \\a             | Alphabetic character: same as \[A-Za-z]                                                                   |
| \\A             | Nonalphabetic character: same as \[^A-Za-z]                                                               |
| \\b             | Backspace                                                                                                 |
| \\d             | Digit: same as \[0-9]                                                                                     |
| \\D             | Nondigit: same as \[^0-9]                                                                                 |
| \\e             | Escape                                                                                                    |
| \\f             | Matches any filename character, as defined by the isfname option                                          |
| \\F             | Like \\f, but excluding digits                                                                            |
| \\h             | Head of word character: same as \[A-Za-z\_]                                                               |
| \\H             | Non-head-of-word character: same as \[^A-Za-z\_]                                                          |
| \\i             | Matches any identifier character, as defined by the isident option                                        |
| \\I             | Like \\i, but excluding digits                                                                            |
| \\k             | Matches any keyword character, as defined by the iskeyword option                                         |
| \\K             | Like \\k, but excluding digits                                                                            |
| \\l             | Lowercase character: same as \[a-z]                                                                       |
| \\L             | Nonlowercase character: same as \[^a-z]                                                                   |
| \\n             | Matches a newline Can be used to match multiline patterns                                                 |
| \\o             | Octal digit: same as \[0-7]                                                                               |
| \\O             | Non-octal digit: same as \[^0-7]                                                                          |
| \\p             | Matches any printable character, as defined by the isprint option                                         |
| \\P             | Like \\p, but excluding digits                                                                            |
| \\r             | Carriage return                                                                                           |
| \\s             | Matches a whitespace character (exactly a space or a tab)                                                 |
| \\S             | Matches anything that isn’t a space or a tab                                                              |
| \\t             | Matches a tab                                                                                             |
| \\u             | Uppercase character: same as \[A-Z]                                                                       |
| \\U             | Nonuppercase character: same as \[^A-Z]                                                                   |
| \\w             | Word character: same as \[0-9A-Za-z\_]                                                                    |
| \\W             | Nonword character: same as \[^0-9A-Za-z\_]                                                                |
| \\x             | Hexadecimal digit: same as \[0-9A-Fa-f]                                                                   |
| \\X             | Nonhexadecimal digit: same as \[^0-9A-Fa-f]                                                               |
| \\\_x           | Where x is any of the previous characters above: match the same character class but with newline included |

### /- Regular expression delimiter -\

> Besides the **/** character, you may use any nonalphanumeric, nonspace character as your delimiter.
>
> EXCEPT **\\**, **"** or **\|**

### /- POSIX character classes -\

| Class      | Matching characters                                                 |
| ---------- | ------------------------------------------------------------------- |
| [:alnum:]  | Alphanumeric characters                                             |
| [:alpha:]  | Alphabetic characters                                               |
| [:blank:]  | Space and Tab characters only                                       |
| [:cntrl:]  | Control characters                                                  |
| [:digit:]  | Numeric characters                                                  |
| [:graph:]  | Printable and visible (nonspace) characters                         |
| [:lower:]  | Lowercase characters                                                |
| [:print:]  | Printable characters (includes whitespace)                          |
| [:punct:]  | Punctuation characters                                              |
| [:space:]  | All whitespace characters (space, tab, newline, vertical tab, etc.) |
| [:upper:]  | Uppercase characters                                                |
| [:xdigit:] | Hexadecimal digits                                                  |

## -[ Regular expression modifier ]-

```
:help s_flags   VIM documentation about all Regular expression modifier
```

```
g               Global replacement, replace all occurence that match
c               Confirm each replacement
i               Ignore case

&               Use previous regular expression modifier, you have to use ':&&' before to save old modifier
```

## -[ COMMAND MODE ]-

```
:   EX MODE
gQ  FULL EX MODE

i   INSERT MODE

ZZ  Save and exit

&   Repeat the last substitution
```

### /- Single Movements -\

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

### /- Block Movements -\

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

### /- Simple Edit -\

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

>>         Indent line on right side
<<         Indent line on left side
4>>        Indent right the 4 lines under the cursor

```

### /- Line Movements -\

```
[ENTER]       Move to the first character of the next line
+             Move to the first character of the next line
-             Move to the first character of the previous line
^             Move to the first nonblank character of the current line
{n}|          Move to the {n} character of the current line

```

### /- Screen Movements -\

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

### /- Search Movements -\

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

### /- Line number Movements -\

```
{n}G      Goto the line {n}
G         Goto the last line of the file

``        Goto the line before you use the last 'G' command(Return at the start)
''        Goto the start of the line before you use the last 'G' command(Return at the start)
```

## -[ EX MODE ]-

Use **[CTRL]+V** to get special key:

- **[ENTER]**('^M' or '\<cr\>')
- **[ESC]**('^[')
- **[BACKSPACE]**('^H')
- **[DELETE]**('\<del\>')
- **[CTRL]+T**
- **[CTRL]+W**
- **[CTRL]+X**

Use the shell command `$ od -c` to get all special key code from the system

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

:source <FILENAME>        Run the script from <FILENAME>
:so <FILENAME>            Run the script from <FILENAME>

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
:s                        Repeat the last substitution
:g/{pt}/ {ec}             Global search, apply the EX command {ex} on all lines that match the pattern {pt}

:preserve                 Force the system to save the buffer(not the file)
```

### /- Filename shortcut -\

```
%         Current filename
#         Alternate filename (Previous file)
```

### /- The :set Command -\

```
:set option               Set general form to enable an option
:set nooption             Set general form to disable an option
:set option!              Set general form to toggle(on/off) an option
:set option?              Set general form to get the name of the option

:set                      Show all options that you have specifically changed
:set all                  Show all active options

:set number               Display line number
:set nu                   Display line number
:set nonu                 Hide line number
:set nu!                  Toggle display/hide line number

:set nowrapscan           Stop search at the bottom(/{pt} or n) or at the top(?{pt} or N)
:set edcompatible         Record last regular expression modifier and use it for the next substitution

:set ic                   Enable 'ignore case', search patern must ignore case
:set noic                 Disable 'ignore case', search patern are case sensitive
```

### /- Shell command -\

```
// EX MODE
:!command                  General form to send 'command' to the system and display the result

:!pwd                      Get the current directory
:read !date                Append the result of 'date' command on the buffer


// COMMAND MODE
!{to}                      Pass the text object {to} to a command
!{to}!                     Repeat last command on text object {to}

!!awk '<AWK_SCRIPT>'<cr>   Give the current line as argument to system command 'awk' with '!!'
```

## -[ FULL EX MODE ]-

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

### /- FULL EX MODE Examples -\

```
:4,15d                                   Delete lines 4 to 15(inclusive)
:100,120m20                              Move lines 100 to 120 on line 20
:100,120co20                             Copy lines 100 to 120 on line 20

:.,$d                                    Delete from current line '.' to the end of the buffer '$'
:5,.m$                                   Move lines 5 to current line '.' on the end of the buffer '$'
:%d                                      Delete all the buffer
:%t$                                     Copy the buffer to the end of the file (consecutive duplicate)

:.,.+10d                                 Delete from the current line '.' to the next 10 lines '.+10'
:100,$m.-4                               Move line 100 to the end of the buffer '$' on 4th line above '-4' current line '.'
:.,+10#                                  Display line number from current line to 10 lines below
:-,+t0                                   Copy 3 lines one above '-', the current line and one below '-' to the top of the buffer '0'
:10;+3d                                  Delete from line 10 to line 13 (10 + 3)

:1,5d | s/teh/the/                       Delete line 1 to 5 and substitute 'teh' for 'the' on current line (before the first command it was the line 6)

:%s/\(That\) or \(this\)/\2 or \1/       Substitute and swap 'That' and 'this' in all the file
:%s/\(That\) or \(this\)/\u\2 or \l\1/   Substitute and swap 'That' and 'this' in all the file
:s/\(abcd\)\1/alphabet-soup/             Substitute 'abcdabcd' by 'alphabet-soup'
:%s/Fortran/\U&/                         Substitute 'Fortran' by 'FORTRAN'
:%s/\<child\>/&ren/g                     Substitute 'child' and only 'child' by 'children' in the whole file
:%s:/home/tim:/home/tom/g                Substitute '/home/tim' by '/home/tom' (use of separator ':' for readability)
:%s/  */ /g                              Substitute 2 or more space by one space (TRIM)
:%s/./\U&/g                              Transform all the character of the buffer to uppercase(\U)
:%&g                                     Repeat the last substitution everywhere
:~                                       Repeat the last substitution used in any command

:g/# FIXME/ d                            Delete all lines with 'FIXME' comments on them
:g/# FIXME/ s/FIXME/DONE/                Substitute all lines with 'FIXME' by 'DONE'
:g/editer/s//editor/g                    Substitute all line with 'editer' and replace by editor (== ':%s/editer/editor/g')
:g/<description>/,/<parameters>/-1 d     Delete a block from <description> to <parameters> not include ('-1')
:g/^/ move 0                             Reverse the lines in the buffer
:1,10g/^/ 3,4 t $                        Repeat 10 times the copy (t) of lines 3 and 4 at the end of the buffer
```

### /- Saving Commands -\

#### Abbreviation

> Abbreviations are for **INSERT MODE**

```
:ab abbreviation string    General form to declare an abbreviation.

:ab mov Master of VIM      Now in Insert mode if you tape 'mov ' VIM change it for 'Master of VIM'
:unab mov                  Remove the 'mov' abbreviation

:ab 123 One^MTwo^MTree     Use of cariage return(^M) into an abbreviation

:ab                        List all abbreviations
```

#### Map

> map are for **NORMAL MODE**
> map! are for **INSERT MODE**

```
:map x sequence                                  General form to declare a map(shortcut). Define character 'x' as a sequence of editing commands.
:map #1 sequence                                 General form to declare a map(shortcut). Define [F1] as a sequence of editing commands.
:unmap x                                         General form to unset the map(shortcut) for 'x'

:map x dwElp                                     Define 'x' to swap 2 words (not perfect example)
:map x I<Root>^M^I<Node>^[ea</Node>^M</Root>^[   Define 'x' to encapsulate a word with a Root/Node XML structure
:map x I/* ^[A */^[                              Define 'x' to add '/*' '*/' around a line
:map x :s;.*;/* & */;^M                          Define 'x' to add '/*' '*/' around a line

:let mapleader="`"                               Define '`' as the leader key
:map <leader>a :q<cr>                            Define leader + 'a' to execute :q[ENTER] (quit)

:map! x sequence                                 General form to declare a map(shortcut) but for INSERT MODE
:unmap! x                                        General form to unset the map(shortcut) for INSERT MODE
:map! + ^[lbi<U>^[ea</U>                         Define '+' to surround a word with <U> </U> on INSERT MODE

:map                                             List all maps for NORMAL MODE
:map!                                            List all maps for INSERT MODE

:help :map-mode                                  Help about all map mode (map, noremap, map!, unmap)
```

#### Macro

To use macro you have to use **Named register**.  
Save the commands sequence on a **Named register**.  
Call the sequence with **@{nrg}**.  
You can repeat the last macro with **@@**.

```
q{nrg}          Start macro recording for named register {nrg}
q               Stop macro recording
@{nrg}          Execute macro from named register {nrg}

```

## -[ INSERT MODE ]-

```
ESC        Exit INSERT MODE

[CTRL]+T   Increment indentation level for the whole line
[CTRL]+D   Decrement indentation level for the whole line

[CTRL]+U   Erase all characters on the line before the cursor
```

# VIM FOR DEVELOPMENT

## -[ COMMAND MODE ]-

```
%             Move the cursor to the other bracket(Useful to find where you forgot a '}')
```

## -[ EX MODE ]-

```
:set autoindent           Automatic indentation control
:set shiftwidth=4         Indentation width is 4 spaces
:set tabstop=4            Indentation width of a [TAB]('\t') is 4 spaces
:set expandtab            Tabs is write with ' ' and not '\t' (Convert TABS to SPACES)
:set list                 Show hidden characters

:10,20 l                  Show hidden characters from line 10 to 20(Useful to verify some lines)

:set showmatch            Show pair of brackets
```

### /- Tags -\

To use _tags_ you have to install **ctags** `sudo apt install exuberant-ctags`

```
:!ctags %                 Generate tag file for the current file
:!ctags *.c               Generate tag file for the current directory

:tag {fn}                 Move the cursor to the definition of function {fn}

[CTRL]+]                  Goto tag definition of the word under the cursor
[CTRL]+T                  Goto previous location before the tag jump([CTRL]+])
```

# VIM SPECIFIC

In this section you will get all specific **VIM** features over **VI**  
`help vi_diff`

**Popular VIM features:**

- Initialization
- Infinite undo
- GUI
- Multiple windows
- Programmer assistance
- Keyword completion
- Syntax extensions
- Scripting and plug-ins
- Postprocessing
- Arbitrary length lines and binary data
- Session context
- Transitions
- Transparent editing
- Meta-information
- The black hole register

**VIM Tutorial:**

- [OpenVIM](https://www.openvim.com/)
- [VIM adventure](https://vim-adventures.com/)

**VIM Startup**

- Check `:help startup` for up to date information about VIM Startup

## -[ COMMAND LINE ]-

```
gvim <FILENAME>                             Open GUI version of VIM
vim -g <FILENAME>                           Open GUI version of VIM

evim <FILENAME>                             Open easy VIM(a more beginner friendly version of VIM)
vim -y <FILENAME>                           Open easy VIM(a more beginner friendly version of VIM)

vimtutor                                    A VIM tutorial
```

### /- Specific command line VIM options -\

| Option       | Description                                                                                             |
| ------------ | ------------------------------------------------------------------------------------------------------- |
| -b           | Edit file in binary mode                                                                                |
| -c _command_ | Execute _command_ as an EX command                                                                      |
| -C           | Run VIM in compatible VI mode                                                                           |
| -d           | Start VIM in DIFF mode                                                                                  |
| -E           | Start VIM in improved EX mode                                                                           |
| -F or -A     | Start VIM in Farsi or Arabic modes                                                                      |
| -g           | Start VIM in GUI mode                                                                                   |
| -M           | Turn off the write option (Read Only mode)                                                              |
| -o[{n}]      | Open all files in separate windows. The optionally integer {n} is used to specify the number of windows |
| -O[{n}]      | Like `-o` but with _vertical split_                                                                     |
| -y           | Start VIM in EASY mode                                                                                  |
| -Z           | Start VIM in restricted mode                                                                            |

#### Remote configuration example

**Server VIM instance**

To get more information: `:help client-server`

```
vim --listen 192.168.0.101:6666                               Run a server instance of vim on host 192.168.0.101 port 6666
```

**Remote client**

```
vim --server 192.168.0.101:6666 --remote <FILENAME>           Connect to the VIM instance 192.168.0.101:6666 and open the file <FILENAME> on the remote host

vim --server 192.168.0.101:6666 --remote-send <COMMAND>       Send <COMMAND> to the VIM instance 192.168.0.101:6666
```

## -[ Environment variables ]-

| Variable | Description                                                                                                                                                                                                                             |
| -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| VIMINIT  | VIM execute its content as an EX command(_This is the first configuration entry_)                                                                                                                                                       |
| EXINIT   | VIM execute its content as an EX command(_This is the second configuration entry, executed only if VIMINIT is empty_)                                                                                                                   |
| MYVIMRC  | Overrides Vim’s search for initialization files. If MYVIMRC has a value when starting, Vim assumes the value is the name of an initialization file and, if the file exists, takes initial settings from it. No other file is consulted. |
| SHELL    | Define the shell or external command interpreter vim had to use                                                                                                                                                                         |

## -[ COMMAND MODE ]-

```
[CTRL]+]                    Goto mark under the cursor
[CTRL]+O                    Goto previous position

v                           VISUAL MODE
```

### /- Movements -\

```
[CTRL]+<end>                Goto the last character of the file
[CTRL]+<home>               Goto the first character of the file

{n}%                        Goto the {n} percentage of the file
:go {n}                     Goto the {n} byte in the file
```

## -[ EX MODE ]-

```
:set compatible       Remove all specific VIM feature(VI pure compatibility)

:help                 Help introduction to use VIM
:help <subject>       Help for the subject in parameter
:h <subject>          Help for the subject in parameter
```

## -[ VISUAL MODE ]-

**a** for Around  
**i** for Inner

Movement general form:

> `{n} a|i {to}` Add {n} {to} around or inner

```
{n}aw | {n}aW   Add {n} word on the selection
{n}iw | {n}iW   Add {n} inner word on the selection(White space count as a word)
as    | is      Add a sentence(a) or an inner(i) sentence on the selection
ap    | ip      Add a paragraph or an inner paragraph on the selection

a'    | i'    Add the content of the ' block on the selection
a"    | i"    Add the content of the " block on the selection
a`    | i`    Add the content of the ` block on the selection
a{    | i{    Add the content of the { block on the selection
a[    | i[    Add the content of the [ block on the selection
a(    | i(    Add the content of the ( block on the selection
a<    | i<    Add the content of the < block on the selection
```
