<h1 align="center">
                                 Mastering Neovim
</h1>

<p align="center">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/07/Neovim-mark-flat.svg/489px-Neovim-mark-flat.svg.png" alt="Neovim logo" title="Neovim logo">
</p>

---
<!-- vim-markdown-toc GFM -->

* [References](#references)
    * [-[ Links ]-](#--links--)
    * [-[ Abbreviations ]-](#--abbreviations--)
    * [-[ COMMAND LINE ]-](#--command-line--)
        * [-| Options |-](#--options--)
    * [-[ General Form of VI commands ]-](#--general-form-of-vi-commands--)
    * [-[ TEXT OBJECT ]-](#--text-object--)
    * [-[ REGISTER ]-](#--register--)
    * [-[ Marker ]-](#--marker--)
    * [-[ Regular expression metacharacters ]-](#--regular-expression-metacharacters--)
        * [-| Regular expression character class |-](#--regular-expression-character-class--)
        * [-| Regular expression delimiter |-](#--regular-expression-delimiter--)
        * [-| POSIX character classes |-](#--posix-character-classes--)
    * [-[ Regular expression modifier ]-](#--regular-expression-modifier--)
* [Usage](#usage)
    * [-[ NORMAL MODE ]-](#--normal-mode--)
        * [-| Edit commands |-](#--edit-commands--)
        * [-| Copying commands |-](#--copying-commands--)
        * [-| Single Movements |-](#--single-movements--)
        * [-| Buffer Movements |-](#--buffer-movements--)
        * [-| Block Movements |-](#--block-movements--)
        * [-| Line Movements |-](#--line-movements--)
        * [-| Screen line Movements |-](#--screen-line-movements--)
        * [-| Screen Movements |-](#--screen-movements--)
        * [-| Search Movements |-](#--search-movements--)
            * [Text Search Movements](#text-search-movements)
            * [Line Search Movements](#line-search-movements)
        * [-| Line number Movements |-](#--line-number-movements--)
    * [-[ EX MODE ]-](#--ex-mode--)
        * [-| Filename shortcut |-](#--filename-shortcut--)
        * [-| The :set Command |-](#--the-set-command--)
        * [-| Shell command |-](#--shell-command--)
    * [-[ FULL EX MODE ]-](#--full-ex-mode--)
        * [-| FULL EX MODE Examples |-](#--full-ex-mode-examples--)
        * [-| Saving Commands |-](#--saving-commands--)
            * [Abbreviation](#abbreviation)
            * [Map](#map)
            * [Macro](#macro)
    * [-[ INSERT MODE ]-](#--insert-mode--)
        * [-| Insert / Delete |-](#--insert--delete--)
        * [-| Motions |-](#--motions--)
* [VIM FOR DEVELOPMENT](#vim-for-development)
    * [-[ NORMAL MODE ]-](#--normal-mode---1)
    * [-[ EX MODE ]-](#--ex-mode---1)
        * [-| Tags |-](#--tags--)
        * [-| Folding and Outlining |-](#--folding-and-outlining--)
            * [NORMAL MODE](#normal-mode)
            * [EX MODE](#ex-mode)
            * [Options](#options)
        * [-| Auto and smart indenting |-](#--auto-and-smart-indenting--)
            * [Options](#options-1)
        * [-| Word completion |-](#--word-completion--)
            * [INSERT MODE](#insert-mode)
        * [-| Syntax Highlighting |-](#--syntax-highlighting--)
            * [EX MODE](#ex-mode-1)
        * [-| Compiling |-](#--compiling--)
            * [EX MODE](#ex-mode-2)
        * [-| Quickfix list window |-](#--quickfix-list-window--)
            * [EX MODE](#ex-mode-3)
* [VIM SPECIFIC](#vim-specific)
    * [-[ COMMAND LINE ]-](#--command-line---1)
        * [-| Specific command line VIM options |-](#--specific-command-line-vim-options--)
            * [Remote configuration example](#remote-configuration-example)
    * [-[ Environment variables ]-](#--environment-variables--)
    * [-[ NORMAL MODE ]-](#--normal-mode---2)
        * [-| Movements |-](#--movements--)
    * [-[ EX MODE ]-](#--ex-mode---2)
        * [-| Buffers |-](#--buffers--)
            * [Status flags](#status-flags)
            * [Special buffers](#special-buffers)
    * [-[ VISUAL MODE ]-](#--visual-mode--)
    * [-[ Multiple Windows ]-](#--multiple-windows--)
        * [-| NORMAL MODE |-](#--normal-mode---3)
            * [Opening and closing window](#opening-and-closing-window)
            * [Moving to other windows](#moving-to-other-windows)
            * [Moving windows](#moving-windows)
            * [Window resizing](#window-resizing)
            * [Windows and Tags](#windows-and-tags)
        * [-| EX MODE |-](#--ex-mode---3)
            * [Settings](#settings)
            * [Commands](#commands)
    * [-[ Tabs ]-](#--tabs--)
        * [-| NORMAL MODE |-](#--normal-mode---4)
        * [-| EX MODE |-](#--ex-mode---4)
* [VIM Script](#vim-script)
* [VIM Tools](#vim-tools)
    * [-[ Terminal ]-](#--terminal--)
    * [-[ Spellchecking ]-](#--spellchecking--)
        * [-| NORMAL MODE |-](#--normal-mode---5)
        * [-| EX MODE |-](#--ex-mode---5)
    * [-[ Binary files ]-](#--binary-files--)
        * [-| Command line |-](#--command-line---2)
        * [-| EX MODE |-](#--ex-mode---6)
    * [-[ Non ASCII Characters ]-](#--non-ascii-characters--)
        * [-[ Digraph metacharacters ]-](#--digraph-metacharacters--)
    * [-[ Edit files over network ]-](#--edit-files-over-network--)
        * [-| Command line |-](#--command-line---3)
        * [-| EX MODE |-](#--ex-mode---7)
    * [-[ Changing directories ]-](#--changing-directories--)
        * [-| EX MODE |-](#--ex-mode---8)
    * [-[ Backup ]-](#--backup--)
        * [-| EX MODE |-](#--ex-mode---9)
    * [-[ Convert to HTML ]-](#--convert-to-html--)
        * [-| EX MODE |-](#--ex-mode---10)
    * [-[ Environment backup ]-](#--environment-backup--)
        * [-| EX MODE |-](#--ex-mode---11)
    * [-[ Lines settings ]-](#--lines-settings--)
        * [-| EX MODE |-](#--ex-mode---12)
    * [-[ Tips ]-](#--tips--)
        * [-| NORMAL MODE |-](#--normal-mode---6)
        * [-| Shell |-](#--shell--)
            * [Shell VI NORMAL MODE](#shell-vi-normal-mode)
        * [-| EX MODE |-](#--ex-mode---13)
* [Extend VIM](#extend-vim)
    * [-[ Plug-ins ]-](#--plug-ins--)
        * [-| Populars Plug-ins |-](#--populars-plug-ins--)
            * [All-in-One IDE configurations](#all-in-one-ide-configurations)
            * [Writer plug-ins](#writer-plug-ins)
        * [-| Links |-](#--links---1)

<!-- vim-markdown-toc -->
---

# References

## -[ Links ]-

- [Neovim](https://neovim.io/)
- [VIM FAQ](https://vimhelp.org/vim_faq.txt.html)
- [VIM TIPS WIKI](https://vim.fandom.com/wiki/Vim_Tips_Wiki)

## -[ Abbreviations ]-

- **{n}** Number
- **{ch}** A character
- **{CH}** An uppercase character
- **{sk}** Special key (Like <BS>, <Return>, ..)
- **{to}** Text Object
- **{rg}** Register
- **{nrg}** Named Register
- **{pt}** Regular expression pattern
- **{rpt}** Replacement expression pattern
- **{rm}** Regular expression modifier
- **{ec}** EX command
- **{fn}** Function name
- **{tg}** Tag
- **{ra}** Range
- **{pl;}** Path list(with separator ';')
- **{pl,}** Path list(with separator ',')
- **{lg}** Language
- **{li}** Lang ISO CODE (ex: 'en_us')
- **{cs}** Color scheme
- **{cl}** Color
- **{hg}** Highlight group
- **{hs}** Highlight setting
- **{sc}** Shell command
- **{fp}** Files pattern
- **{wd}** Word

## -[ COMMAND LINE ]-

```
vim [options] <FILENAME>                     General form of VIM command line
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
vim -d <FILENAME1> <FILENAME2>               Open VIM on DIFF mode to compare <FILENAME1> and <FILENAME2>

vim -r                                       List all saved buffer by VI(Used for recovery)
ex -r                                        List all saved buffer by VI(Used for recovery)
vim -r <BUFFER>                              Recover the edited <BUFFER>

ex -s <FILENAME> < <SCRIPT_FILENAME>         Execute the vim script <SCRIPT_FILENAME> on the file <FILENAME>
```

### -| Options |-

For more information about command line arguments check `:help cli-arguments`

| Option        | Description                                                                                                                                                                                                                                                                                                                                |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| +[num]        | Start editing at line number num, or at the last line of the file if num is omitted                                                                                                                                                                                                                                                        |
| +/{pt}        | Start editing at the first line matching pattern {pt}                                                                                                                                                                                                                                                                                      |
| +?{pt}        | Start editing at the last line matching pattern {pt}                                                                                                                                                                                                                                                                                       |
| -b            | Edit the file in binary mode                                                                                                                                                                                                                                                                                                               |
| -c {ec}       | Run the given EX command {ec} upon startup                                                                                                                                                                                                                                                                                                 |
| --cmd {ec}    | Like -c, but execute the command before any configuration files are read                                                                                                                                                                                                                                                                   |
| -C            | Start the editor in vi-compatible mode                                                                                                                                                                                                                                                                                                     |
| -d            | Run in diff mode<br> Works like vimdiff                                                                                                                                                                                                                                                                                                    |
| -D            | Debugging mode for use with scripts                                                                                                                                                                                                                                                                                                        |
| -e            | Run as EX (line-editing rather than full-screen mode)                                                                                                                                                                                                                                                                                      |
| -h            | Print help message                                                                                                                                                                                                                                                                                                                         |
| -l            | Enter Lisp mode for running Lisp programs                                                                                                                                                                                                                                                                                                  |
| -L            | List files that were saved due to an aborted editor session or system crash this option is the same as -r                                                                                                                                                                                                                                  |
| -m            | Start the editor with the write option turned off                                                                                                                                                                                                                                                                                          |
| -M            | Do not allow text in files to be modified<br> This is similar to -m but additionally blocks any changes to the buffer                                                                                                                                                                                                                      |
| -n            | Do not use a swap file(record changes in memory only)                                                                                                                                                                                                                                                                                      |
| --noplugin    | Do not load any plug-ins                                                                                                                                                                                                                                                                                                                   |
| -N            | Run VIM in a non-vi-compatible mode                                                                                                                                                                                                                                                                                                        |
| -o{n}         | Start VIM with {n} open windows                                                                                                                                                                                                                                                                                                            |
| -O{n}         | Start VIM with {n} open windows(split vertically)                                                                                                                                                                                                                                                                                          |
| -r <FILENAME> | Recovery mode; recover and resume editing on file after an aborted editor session or system crash<br> Without file, list files available for recovery                                                                                                                                                                                      |
| -R            | Edit files in read-only mode                                                                                                                                                                                                                                                                                                               |
| -s <FILENAME> | Read and execute commands given in the specified scriptfile <FILENAME>                                                                                                                                                                                                                                                                     |
| -S <FILENAME> | Read and execute commands given in commandfile <FILENAME><br> Shorthand for `vim -c 'source <FILENAME>'`                                                                                                                                                                                                                                   |
| -t {tg}       | Edit the file containing tag {tg}, and position the cursor at its definition                                                                                                                                                                                                                                                               |
| -T type       | Set the term (terminal type) option<br> This value overrides the $TERM environment variable                                                                                                                                                                                                                                                |
| -u <FILENAME> | Read configuration information from the specified configuration file instead of the default _.vimrc_ configuration file<br> If the file argument is **NONE**, VIM reads no configuration files, loads no plug-ins, and runs in compatible mode<br> If the argument is **NORC**, it reads no configuration files, but it does load plug-ins |
| -v            | Run in full-screen mode                                                                                                                                                                                                                                                                                                                    |
| --version     | Print version information                                                                                                                                                                                                                                                                                                                  |
| -V{n}         | Verbose mode; print messages about what options are being set and what files are being read or written<br> You can set a level of verbosity {n} to increase or decrease the number of messages received<br> The default value is 10 for high verbosity                                                                                     |
| -W <FILENAME> | Write all typed commands from the current session to the specified scriptfile <FILENAME><br> The file thus created can be used with the -s option                                                                                                                                                                                          |
| -y            | Modeless vi; run VIM in insert mode only, without a NORMAL MODE<br> This is the same as invoking VIM as `evim`                                                                                                                                                                                                                             |
| -Z            | Start VIM in restricted mode<br> Do not allow shell commands or suspension of the editor                                                                                                                                                                                                                                                   |

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

(			sentences backward
)			sentences forward
{			paragraphs backward
}			paragraphs forward
]]			sections forward or to the next "{" in the first column. When used after an operator, then also stops below a "}" in the first column
][			sections forward or to the next '}' in the first column
[[			sections backward or to the previous "{" in the first column
[]			sections backward or to the previous "}" in the first column
```

## -[ REGISTER ]-

```
"{n}        Numbered register[1-9], the last nine deletions, from most to least recent
"{ch}       Named register[a-z], use like user clipboard
"{CH}       Named register, but when you use uppercase character, you append the register(Accumulator)
```

## -[ Marker ]-

**Tips**: To get the char '`' you have to press the key twice

```
m{ch}       Mark the current position with {ch}

'{ch}       Goto the first character of the line marked by {ch}
''          Goto the first character of the line marked by the previous mark or context

`{ch}       Goto the position of the mark {ch}
``          Goto the position of the previous mark or context

`[          Goto the begining of the previous text operation
`]          Goto the end of the previous text operation

']          Goto the line of the previous text operation

`.          Goto the last change in the buffer
'.          Goto the last line changed in the buffer

:marks      List active marks
```

## -[ Regular expression metacharacters ]-

For more information check `:help regexp`

| Metacharacter | Description                                                                                                                  |
| ------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| .             | Match any _single_ character except a new line                                                                               |
| \*            | Match 0 or more of the single character that immediately precedes it                                                         |
| ^             | Match only line that begin with pattern after '^'. If '^' is not the beginning of the expression it's just the '^' character |
| $             | Match only line that end with pattern before '$'. If '$' is not the end of the expression it's just the '$' character        |
| \             | Used to escape a character like '\.' to match a point and not any _single_ character                                         |
| \\{n}         | Recall a _subpattern_ , {n} is between 1 to 9                                                                                |
| [ ]           | Match any _one_ of the characters enclosed between the brackets                                                              |
| [^ ]          | Match any _one_ of the characters that is **NOT** enclosed between the brackets                                              |
| [: :]         | Match any character which is part of the character classes                                                                   |
| [. .]         | Match multicharacter sequence that should be treated as a unit                                                               |
| [= =]         | Match an equivalence class list a set of characters that should be considered equivalent (ex: 'e' and 'é')                   |
| \\( \\)       | Save the subpattern enclosed between \( and \) into a special holding space (\1 .. \9)                                       |
| \\<           | Match only character at the beginning of a word                                                                              |
| \\>           | Match only character at the end of a word                                                                                    |
| &             | Replace the '&' with the entire text matched by the search pattern(**ONLY FOR REPLACEMENT PATTERN {rpt}**)                   |
| ~             | Replace the '~' with the last used replacement pattern(**ONLY FOR REPLACEMENT PATTERN {rpt}**)                               |
| \\u           | Force the next characters to be on uppercase (**ONLY FOR REPLACEMENT PATTERN {rpt}**)                                        |
| \\U           | Force all next characters to be on uppercase (**ONLY FOR REPLACEMENT PATTERN {rpt}**)                                        |
| \\l           | Force the next characters to be on lowercase (**ONLY FOR REPLACEMENT PATTERN {rpt}**)                                        |
| \\L           | Force all next characters to be on lowercase (**ONLY FOR REPLACEMENT PATTERN {rpt}**)                                        |
| \\\|          | String choice (ex car\\\| moto)                                                                                              |
| \\&           | If the pattern before the \\& match the pattern after is evaluated (ex .*Tom\\&.*Jerry)                                      |
| \\+           | Match 1 or more                                                                                                              |
| \\=           | Match 0 or 1                                                                                                                 |
| \\?           | Match 0 or 1                                                                                                                 |
| \\{...}       | Repeat the match {n} times or {n,m} in acceptable range                                                                      |
| ~             | Match the last given replacement string                                                                                      |
| \\(...\\)     | Grouping                                                                                                                     |
| \\{n}         | Call group {n} capture                                                                                                       |

### -| Regular expression character class |-

| Character class | Description                                                                                               |
| --------------- | --------------------------------------------------------------------------------------------------------- |
| \\a             | Alphabetic character: same as \[A-Za-z]                                                                   |
| \\A             | Non alphabetic character: same as \[^A-Za-z]                                                              |
| \\b             | Backspace                                                                                                 |
| \\d             | Digit: same as \[0-9]                                                                                     |
| \\D             | Non digit: same as \[^0-9]                                                                                |
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
| \\L             | Non lowercase character: same as \[^a-z]                                                                  |
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
| \\U             | Non uppercase character: same as \[^A-Z]                                                                  |
| \\w             | Word character: same as \[0-9A-Za-z\_]                                                                    |
| \\W             | Non word character: same as \[^0-9A-Za-z\_]                                                               |
| \\x             | Hexadecimal digit: same as \[0-9A-Fa-f]                                                                   |
| \\X             | Non hexadecimal digit: same as \[^0-9A-Fa-f]                                                              |
| \\\_x           | Where x is any of the previous characters above: match the same character class but with newline included |

### -| Regular expression delimiter |-

> Besides the **/** character, you may use any non alphanumeric, non space character as your delimiter.
>
> EXCEPT **\\**, **"** or **\|**

### -| POSIX character classes |-

| Class      | Matching characters                                                 |
| ---------- | ------------------------------------------------------------------- |
| [:alnum:]  | Alphanumeric characters                                             |
| [:alpha:]  | Alphabetic characters                                               |
| [:blank:]  | Space and Tab characters only                                       |
| [:cntrl:]  | Control characters                                                  |
| [:digit:]  | Numeric characters                                                  |
| [:graph:]  | Printable and visible (non space) characters                        |
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
# Usage 

## -[ NORMAL MODE ]-

```
:   EX MODE
gQ  FULL EX MODE

i   INSERT MODE

ZZ  Save and exit

&   Repeat the last substitution
```

### -| Edit commands |-

For more information check `:help change.txt`

```
i          Insert (under the cursor)
gi         Insert at the last editing position
I          Insert at the begining of the line
gI         Insert at the begining of the line of the last editing position
a          Append (after the cursor)
A          Append at the end of the line

o          Open an empty line below the cursor
O          Open an empty line above the cursor

J          Join the current line and the line under

c{to}      Change the text object(Start at the cursor position)
cc         Change all the current line
C          Change from the cursor to the end of the current line

r {ch}     Replace the character under the cursor by {ch}
R          Enter replace mode (replace until ESC)

s          Replace one character and enter insert mode (Alias for 'c ')
S          Delete the entire line and enter insert mode (Alias for 'cc')

d{to}      Delete the text object(Start at the cursor position)
dd         Delete the current line
D          Delete characters from the cursor to the end of the line (Alias for 'd$')
d^         Delete characters from the cursor to the begining of the line
d/{pt}     Delete characters forward until match pattern {pt} is found(not inclusive)
dn         Delete characters forward until the next match pattern is found(not inclusive)
dL         Delete characters to the last line on the screen
dG         Delete characters from the cursor to the end of the buffer

x          Delete character under the cursor
X          Delete character before the cursor

.          Repeat the last command

u          Undo the last command
U          Undo all edit on the current line
[CTRL]+R   Redo the last undo command

xp         Swap 2 characters
~          Swap uppercase/lowercase
g~~        Swap uppercase/lowercase for the whole line
g~w        Switch the case a word
guw        Change word to lowercase
gUw        Change word to uppercase

>>         Indent line on right side
<<         Indent line on left side
4>>        Indent right the 4 lines under the cursor

[CTRL]+A   Increment the number under the cursor
[CTRL]+X   Decrement the number under the cursor
```

### -| Copying commands |-

```
{rg}p      Put the text from the register after the cursor(It's a PASTE)
p          Put the text fromt the register "0 after the cursor

{rg}P      Put the text from the register before the cursor(It's a PASTE)
P          Put the text from the register "0 before the cursor

{rg}y{to}  Yank(Copy) the text object to the register
y{to}      Yank(Copy) the text object to the register "0
yy         Yank(Copy) the current line(Alias for 'y$')
Y          Yank(Copy) the current line(Alias for 'y$')
```

### -| Single Movements |-

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

### -| Buffer Movements |-

```
gg    Move to the first line of the buffer
G     Move to the last line of the buffer

{n}G  Goto line {n}
:{n}  Goto line {n}
```

### -| Block Movements |-

```
w     Forward one word(Special characters count one word)
W     Forward one word(Withespace separated)
b     Backward one word(Special characters count one word)
B     Backward one word(Withespace separated)
e     Forward to the end of the word(Special characters count one word)
E     Forward to the end of the word(Withespace separated)
ge    Last character of previous word(Special characters count one word)
gE    Last character of previous word(Withespace separated)

(     Move to the begining of current sentence
)     Move to the begining of the next sentence
{     Move to the begining of current paragraph
}     Move to the begining of the next paragraph
[[    Move to the begining of current section
]]    Move to the begining of the next section
```

### -| Line Movements |-

```
0             Move to the start of the line
$             Move to the end of the line

^             Move to the first nonblank character of the current line
g_            Move to the last nonblank character of the current line

[ENTER]       Move to the first character of the next line
+             Move to the first character of the next line
-             Move to the first character of the previous line

{n}|          Move to the {n} character of the current line

```

### -| Screen line Movements |-

What we call _screen line_ is the line **displayed** on the screen

```
g0            Move to the start of the screen line
g$            Move to the end of the screen line

g^            Move to the first nonblank character of the screen line
gm            Move to the middle of the screen line


```

### -| Screen Movements |-

```
H             Move to the top line of the screen
M             Move to the middle line of the screen
L             Move to the last line of the screen

[CTRL]+F      Scroll one screen forward
[CTRL]+B      Scroll one screen backward
[CTRL]+D      Scroll half screen forward(Down)
[CTRL]+U      Scroll half screen Backward(Up)
[CTRL]+E      Scroll the screen one line down
[CTRL]+Y      Scroll the screen one line up

z [ENTER]     Move the current line on the top of the screen
z.            Move the current line on the center of the screen
zz            Move the current line on the center of the screen
z-            Move the current line on the bottom of the screen
{n}z [ENTER]  Move the line {n} on top of the screen
{n}z.         Move the line {n} on the center of the screen
{n}z-         Move the line {n} on the bottom of the screen

[CTRL]+L      Redraw the screen
```

### -| Search Movements |-

#### Text Search Movements

```
/{pt}         Search pattern {pt} forward
/{pt}/+{n}    Goto line {n} after the pattern {pt}
/{pt}/-{n}    Goto line {n} before the pattern {pt}

?{pt}         Search pattern {pt} backward
?{pt}?+{n}    Goto line {n} after the pattern {pt}
?{pt}?-{n}    Goto line {n} before the pattern {pt}

n             Repeat the search in forward direction
N             Repeat the search in backward direction

*             Search forward for the word under the cursor(Match only exact word)
g*            Search forward for the word under the cursor
#             Search backward for the word under the cursor(Match only exact word)
g#            Search backward for the word under the cursor

%             Find match of current parenthesis, brace or bracket
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

### -| Line number Movements |-

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
:e +<FILENAME>            Begin editing at the end of the file
:e +{n} <FILENAME>        Begin editing at line {n}
:q                        Close the current buffer, if it's the only buffer exit VIM
:qa                       Close all buffer and exit
:q!                       Force exit without saving
:w                        Save current buffer
:wa                       Save all buffers
:w <FILENAME>             Save current buffer in a new file
:{n},{n}w <FILENAME>      Save the current range in a new file
:{n},{n}w >> <FILENAME>   Save the current range to the end of <FILENAME> (APPEND)
:w! <FILENAME>            Save current buffer in an existing file
:w %.new                  Save current buffer on a file with the name of current buffer(%) + '.new'
:x                        Save current buffer and exit (LIKE 'ZZ' in NORMAL MODE)

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

### -| Filename shortcut |-

```
%         Current filename
#         Alternate filename (Previous file)
```

### -| The :set Command |-

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

### -| Shell command |-

```
// EX MODE
:!{sc}                     General form to send shell command {sc} to the system and display the result
:{ra}!                     Send the content of the range {ra} to the shell command {sc}(like |)

:!pwd                      Get the current directory
:read !date                Append the result of 'date' command on the buffer

// NORMAL MODE
!{to} {sc}                 Pass the text object {to} to the shell command {sc}
!{to}!                     Repeat last shell command on text object {to}

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

### -| FULL EX MODE Examples |-

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

### -| Saving Commands |-

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

> map is for **NORMAL MODE**
> map! is for **INSERT MODE**

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

For more information about special keys for INSERT MODE, check `:help ins-special-keys`

```
ESC                     Exit INSERT MODE
[CTRL]+C                Exit INSERT MODE

[CTRL]+O                Execute one command, return to INSERT MODE
```

### -| Insert / Delete |-

```
[CTRL]+H                Delete the character before the cursor(<BS>)
[CTRL]+J                Begin a new line(<Enter>)
[CTRL]+I                Insert a <Tab>

[CTRL]+K {ch} {ch}      Insert non-ASCII character, for more information check ':help digraph'
[CTRL]+V {sk}           Insert the litteral for the special key {sk}
[CTRL]+Q {sk}           Insert the litteral for the special key {sk}

[CTRL]+A                Insert previously inserted text

[CTRL]+T                Increment indentation level for the whole line
[CTRL]+D                Decrement indentation level for the whole line

[CTRL]+U                Delete all characters on the line before the cursor

[CTRL]+W                Delete the word before the cursor

[CTRL]+N                Find next keyword(completion)
[CTRL]+P                Find previous keyword(completion)
[CTRL]+X                Completion mode, check ':help i_CTRL-X' for more information

[CTRL]+R {rg}           Insert the content of the register {rg}
[CTRL]+R [CTRL]+R {rg}  Insert the content of the register {rg} literally(Without special character interpollation)
[CTRL]+R [CTRL]+O {rg}  Insert the content of the register {rg} literally but without indentation
[CTRL]+R [CTRL]+P {rg}  Insert the content of the register {rg} literally but with auto-indentation

[CTRL]+E                Insert the character of the line below the cursor
[CTRL]+Y                Insert the character of the line above the cursor
```

### -| Motions |-

```
[CTRL]+G <up>           Goto to the begining of the line up
[CTRL]+G <down>         Goto to the begining of the line down

[CTRL]+<left>           Goto one word back
[SHIFT]+<left>          Goto one word back
[CTRL]+<right>          Goto one word front
[SHIFT]+<right>         Goto one word front

[CTRL]+<home>           Goto the begining of the buffer
[CTRL]+<end>            Goto the end of the buffer

[SHIFT]+<up>            Goto one screen up
[SHIFT]+<down>          Goto one screen down
```

# VIM FOR DEVELOPMENT

**$VIMRUNTIME**=`/usr/share/nvim/runtime`

## -[ NORMAL MODE ]-

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

:set taglength={n}        Controls the number of significant characters in a tag that is to be looked up. The default value of 0 indicates that all characters are significant
:set tags={pl;}            List of the files{pl;} to look for the tags
:set tagrelative          If using a tags file in another directory, file names in that tags file are relative to the directory where the tags file is
```

### -| Tags |-

To use _tags_ you have to install **ctags** `sudo apt install exuberant-ctags`

```
:!ctags %                 Generate tag file for the current file
:!ctags *.c               Generate tag file for the current directory

:tag {tg}                 Move the cursor to the definition of function {fn}

[CTRL]+]                  Goto tag definition of the word under the cursor
[CTRL]+T                  Goto previous location before the tag jump([CTRL]+])

:tags                     List all tags
:tselect {tg}             List all tags corresponding to the tag{tg}, user can choose
:stselect {tg}            List all tags corresponding to the tag{tg}, user can choose in a new window
:tnext                    Goto the next matching tag
:tprev                    Goto the previous matching tag
:tfirst                   Goto the first matching tag
:tlast                    Goto the last matching tag
```

### -| Folding and Outlining |-

VIM provide `:mkview` to save fold configuration and `:loadview` to load previous configuration

For more information about **Fold** check `:help fold.txt`

#### NORMAL MODE

All fold command begin by **z**

> When a fold is close, you can operate on it like it was only one line(very powerfull feature)

```
zf{to}              Create a fold with the text object {to}
zf%                 Create a fold with brace block ('{' and '}' it's an example)
{n}zF               Create a fold covering count{n} lines, starting with the current line
zd                  Delete fold under the cursor
zE                  Delete all fold in the window

zo                  Open the fold under the cursor
zc                  Close the fold under the cursor
za                  Toggle the fold under the cursor

zO                  Open all folds recursively
zA                  Toggle all folds recursively
zC                  Close all folds recursively
zD                  Delete all folds recursively

zr                  Increment foldlevel
zm                  Decrement foldlevel
zM                  Set option foldlevel to 0
zn                  Unset the foldenable option
zN                  Set the foldenable option
```

#### EX MODE

```
:{ra}fold           Create a fold with the lines of the range {ra}
```

#### Options

```
:set foldcolumn={n}     Define the width of the folder column, to see folder indicator in the margin
:set foldlevel={n}      Define the fold level to open, displays only lines whose fold levels are less than or egal to {n}
```

### -| Auto and smart indenting |-

For more information about **indenting** check ':help indent.txt'

#### Options

```
:set autoindent         Activate auto indentation
:set smartindent        Activate smart indentation (more powerfull than 'autoindent')
:set cindent            Activate C style indentation, more C-type language specific
:set indentexpr         This lets you define your own expression, which VIM evaluates in the context of each new line you begin

:set paste              Activate paste compatibility to avoid autoindent issue during the action
:set nopaste
```

### -| Word completion |-

For more information check `:help ins-completion`

#### INSERT MODE

```
[CTRL]+N                Next completion
[CTRL]+P                Previous completion

All specifics completion commands start with [CTRL]+X

[CTRL]+X [CTRL]+L       Line
[CTRL]+X [CTRL]+F       Filename
[CTRL]+X [CTRL]+I       Included files
[CTRL]+X [CTRL]+D       Definition

[CTRL]+X [CTRL]+N       Current file forwards
[CTRL]+X [CTRL]+P       Current file backwards
[CTRL]+X [CTRL]+K       Dictionary
[CTRL]+X [CTRL]+T       Thesaurus
[CTRL]+X [CTRL]+]       Tag
[CTRL]+X [CTRL]+V       VIM commands
[CTRL]+X [CTRL]+U       User defined
[CTRL]+X [CTRL]+O       Omni
[CTRL]+X [CTRL]+S       Spelling suggestions

[CTRL]+X [CTRL]+Z       Stop completion
```

### -| Syntax Highlighting |-

All syntax files are on **$VIMRUNTIME** + `/syntax`  
All color scheme files are on **$VIMRUNTIME** + `/colors`

For list and description of all **Highlight groups** check `:help highlight-groups`

#### EX MODE

```
:syntax on              Enable syntax highlighting

:set syntax={lg}        Set syntax highlighting for a specific language{lg}

:colorscheme{cs}        Use the color scheme {cs}

:set background?        Get background color
:set background={cl}    Set background color with the color {cl}

:highlight              Get all highlight group configuration
:highlight {hg}         Get the highlight configuration for the highlight group {hg}
:highlight {hg} {hs}    Set the highlight setting{hs} for the highlight group {hg}
:help highlight         Get help about highlight for more information about all settings possibilities
```

### -| Compiling |-

#### EX MODE

```
:make <FILENAME>        Compile <FILENAME>
:set makeprg={sc}       Set the shell command used by make
```

### -| Quickfix list window |-

#### EX MODE

```
:copen                  Open QuickFix list window
:cnext                  Goto next quickfix occurence(default error)
:cprevious              Goto previous quickfix occurence(default error)

:vimgrep {pt} {fp}      Search pattern {pt} for all files in {fp}(VIM grep alternative). Results goes on quickfix
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

vim -o <FILENAME1>..<FILENAMEn>             Open all files in separate windows
vim -p <FILENAME1>..<FILENAMEn>             Open all files in separate tabs
```

### -| Specific command line VIM options |-

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

## -[ NORMAL MODE ]-

```
[CTRL]+]                    Goto mark under the cursor
[CTRL]+O                    Goto previous position

v                           VISUAL MODE
```

### -| Movements |-

```
[CTRL]+<end>                Goto the last character of the file
[CTRL]+<home>               Goto the first character of the file

{n}%                        Goto the {n} percentage of the file
:go {n}                     Goto the {n} byte in the file
```

## -[ EX MODE ]-

```
:set compatible       Remove all specific VIM feature(VI pure compatibility)
:set incsearch        Activate incremental search(Move directly on the buffer as you type on the keyboard)

:set undolevels={n}   Define the number of undoable changes you can make in an editing session

:help                 Help introduction to use VIM
:help <subject>       Help for the subject in parameter
:h <subject>          Help for the subject in parameter

:undo                 Undo the last command
:redo                 Redo the last command
:help usr_32.txt      More information about how to navigate changes as a tree
```

### -| Buffers |-

```
:ls                   List the buffers
:ls!                  List all buffers of all VIM instance
:buffers              List the buffers
:files                List the buffers

:cwindow              Open error window (quickfix)
:lwindow              Open location window

:windo {ec}           Execute EX command {ec} on all windows
:bufdo {ec}           Execute EX command {ec} on all buffers

:ball                 Edit all args or buffers
:sball                Edit all args or buffers and open them in new windows
:unhide               Edit all loaded buffer
:sunhide              Edit all loaded buffer and open them in new windows

:badd <FILENAME>      Add file to the buffer list
:bunload              Unload current buffer
:bdelete              Unload current buffer and delete from the buffer list
:buffer {n}           Load buffer {n}
:sbuffer {n}          Load buffer {n} in a new window
:bnext                Move to the next buffer
:bnext {n}            Move to the {n}th next buffer
:sbnext {n}           Load the {n}th next buffer in a new window
:bNext                Move to previous buffer
:bprevious            Move to previous buffer
:bNext {n}            Move to the {n}th previous buffer
:sbNext {n}           Load the {n}th previous buffer in a new window
:blast                Move to the last buffer
:sblast               Load the last buffer in a new window and split horizontally
:vertical sblast      Load the last buffer in a new window and split vertically
:bmod {n}             Move to the {n}th modified buffer
:sbmod {n}            Load to the {n}th modified buffer in a new window

:stag {tg}            Load the file that containt the tag{tg} definition in a new window
```

#### Status flags

| Code   | Description                                                                                                                                                                                                                                                                                        |
| ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| u      | Unlisted buffer. This buffer is not listed unless you use the ! modifier. To see an example of an unlisted buffer, type :help. Vim splits the current window to include a new window in which the built-in help appears. The plain :ls command will not show the help buffer, but :ls! includes it |
| % or # | % is the buffer for the current window. # is the buffer to which you would switch with the :edit # command. These are mutually exclusive                                                                                                                                                           |
| a or h | a indicates an active buffer. That means the buffer is loaded and visible. h indicates a hidden buffer. The hidden buffer exists but is not visible in any window. These are mutually exclusive                                                                                                    |
| - or = | - indicates a buffer has the modifiable option turned off. The file is read-only. = is a read-only buffer that cannot be made modifiable (for instance, because you don’t have filesystem privileges to write to the file). These are mutually exclusive                                           |
| + or x | + indicates a modified buffer. x is a buffer with read errors. These are mutually exclusive                                                                                                                                                                                                        |

#### Special buffers

- _directory_: List directory content (read-only buffer), when you type [ENTER], the file under the cursor is loaded
- _help_: Show help buffer (read-only buffer)
- _QuickFix_: Contains the list of errors created by your commands(View with `:cwindow` and for location `:lwindow`)
- _scratch_: These buffers contain text for general purposes

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

## -[ Multiple Windows ]-

For more information check `:help windows.txt`

### -| NORMAL MODE |-

All windows command are prefixed with **[CTRL]+W**

#### Opening and closing window

```
[CTRL]+W s                Split current window in 2, Horizontal split
[CTRL]+W v                Split current window in 2, Vertical split
[CTRL]+W n                Create a new window and split horizontally

[CTRL]+W ^                Split current window in 2 and edit the alternate file, Horizontal split

[CTRL]+W q                Quit the current window
[CTRL]+W c                Close the current window
[CTRL]+W o                Close all window except the current one
```

#### Moving to other windows

```
[CTRL]+W j                Move to the next bottom window
[CTRL]+W <down>           Move to the next bottom window

[CTRL]+W k                Move to the next top window
[CTRL]+W <up>             Move to the next top window

[CTRL]+W h                Move to the next left window
[CTRL]+W <left>           Move to the next left window

[CTRL]+W l                Move to the next right window
[CTRL]+W <right>          Move to the next right window

[CTRL]+W w                Move to the next window (cycle)
[CTRL]+W p                Move to the last accessed window

[CTRL]+W t                Move to the top left window
[CTRL]+W b                Move to the bottom right window

```

#### Moving windows

```
[CTRL]+W r                Rotate windows downwards/rightwards
[CTRL]+W R                Rotate windows upwards/leftwards

[CTRL]+W x                Exchange current window with the next one

[CTRL]+W K            	  Move the current window to be at the very top
[CTRL]+W J	              Move the current window to be at the very bottom
[CTRL]+W H	              Move the current window to be at the far left
[CTRL]+W L	              Move the current window to be at the far right
[CTRL]+W T	              Move the current window to a new tab page.
```

#### Window resizing

```
[CTRL]+W =	              Make all windows (almost) equally high and wide
[CTRL]+W _	              Set current window height to N (default: highest possible).
[CTRL]+W |	              Set current window width to N (default: widest possible).
[CTRL]+W -	              Decrease current window height by N (default 1)
[CTRL]+W +	              Increase current window height by N (default 1)
[CTRL]+W <	              Decrease current window width by N (default 1)
[CTRL]+W >	              Increase current window width by N (default 1)

{n}z[ENTER]               Set current window height to {n} lines
```

#### Windows and Tags

```
[CTRL]+W ]                Open a new window with the file that define the tag under the cursor (perform a ':tag')
[CTRL]+W g ]              Open a new window with the list of the corresponding tag under the cursor(perform a ':tselect')
[CTRL]+W g [CTRL]+]       Open a new window with the list of the corresponding tag under the cursor(perform a ':tjump')

[CTRL]+W f                Open(if exist) the filename under the cursor in a new window(like 'gf')
[CTRL]+W F                Open(if exist) the filename under the cursor and go to the line number (ex: '/a/b/myfile:82') in a new window(like 'gF')

[CTRL]+W gf               Open(if exist) the filename under the cursor in a new tab(like 'gf')
[CTRL]+W gF               Open(if exist) the filename under the cursor and go to the line number (ex: '/a/b/myfile:82') in a new tab(like 'gF')

[CTRL]+W gt               Goto next tab
[CTRL]+W gT               Goto previous tab
```

### -| EX MODE |-

#### Settings

```
:set winheight={n}      Minimal number of line for the current window
:set winwidth={n}       Minimal number of columns for the current window
:set winminheight={n}   Minimal number of line for all windows
:set winminwidth={n}    Minimal number of columns for all windows
```

#### Commands

```
:wincmd {ch}                          Send window command({ch}), exactly the same behaviour as [CTRL]+W in NORMAL MODE (ex: ':wincmd s' to split)

:{n}split [++opt] [+cmd] <FILENAME>   General split command format
:split                                Split current window in 2, Horizontal split
:vsplit                               Split current window in 2, Vertical split
:{n}split                             Split current window in 2 with {n} lines, Horizontal split
:{n}vsplit                            Split current window in 2 with {n} columns, Vertical split
:new                                  Create a new window and split horizontally
:new <FILENAME>                       Create a new window and load <FILENAME>
:vnew                                 Create a new window and split vertically
:sview                                Split current window in 2 but the new window is in read-only mode(view)
:sfind <FILENAME>                     Split horizontally only if the <FILENAME> exist

:close                                Close the current window
:hide                                 Hide the current window
:only                                 Close all window except the current one

:resize                               Set current window height to highest possible
:resize -{n}                          Decrease current window height by {n} lines
:resize +{n}                          Increase current window height by {n} lines
:resize {n}                           Set current window height to {n} lines
:vertical resize {n}                  Set current window width to {n} columns
```

## -[ Tabs ]-

### -| NORMAL MODE |-

```
[CTRL]+<pagedown>                     Goto the next tab
gt                                    Goto the next tab
[CTRL]+<pageup>                       Goto the previous tab
gT                                    Goto the previous tab

[CTRL]+<tab>                          Goto the last accessed tab
g<tab>                                Goto the last accessed tab
```

### -| EX MODE |-

```
:tabnew                               Open a new tab
:tabnew <FILENAME>                    Open the file <FILENAME> in a new tab
:tabedit <FILENAME>                   Open the file <FILENAME> in a new tab

:tabclose                             Close the current tab

:tabonly                              Close all tabs except the current

:tabnext                              Goto the next tab
:tabprevious                          Goto the previous tab
:tabNext                              Goto the previous tab
:tabfirst                             Goto the first tab
:tablast                              Goto the last tab

:tabs                                 List all the tabs
```

# VIM Script

For more information check all these help topics:

- **autocmd**
- **scripts**
- **variables**
- **functions**
- **usr_41.txt**

# VIM Tools

## -[ Terminal ]-

For more information about this new feature of _Neovim_ check `:help terminal_emulator.txt`

Don't forget **[CTRL]+\ [CTRL]+N** to return to NORMAL MODE(In default configuration you can not use <esc>)  
And **[CTRL]+\ [CTRL]+O** to return to NORMAL MODE for only one command

```
:terminal               Open the terminal emulator in the current window (Interactive)
:terminal {sc}          Open the terminal emulator in the current window and run the shell command {sc} (Non-interactive)

:split term://{sc}      Open the terminal emulator in an horizontal split window and run the shell command {sc} (Non-interactive)
:vsplit term://{sc}     Open the terminal emulator in a vertical split window and run the shell command {sc} (Non-interactive)

:split +terminal        Open the terminal emulator in an horizontal split window (Interactive)
:tabnew +terminal       Open the terminal emulator in a new tab
```

## -[ Spellchecking ]-

For more information check `:help spell` or `:help spell.txt`

### -| NORMAL MODE |-

By default spellchecking is not active on VIM, **you must must activate it on demand**

```
]s  Goto the next occurrence of a misspelled word
[s  Goto the previous occurrence of a misspelled word
zg  Add the word under the cursor to the list of good words
zG  Add the word under the cursor to the list of good words in the internal-wordlist
zw  Add the word under the cursor to the list of bad words
zW  Add the word under the cursor to the list of bad words in the internal-wordlist
z=  Display the list of suggestions for replacement of a bad word
```

### -| EX MODE |-

```
:setlocal spell spelllang={li}   Turn on spellchecking and set the spellchecking region with the language ISO CODE {li}
:setlocal nospell                Turn off spellchecking

:spellgood {wd}                  Add word {wd} to the good word list
:spellgood! {wd}                 Add word {wd} to the good word list in internal-list
:spellwrong {wd}                 Add word {wd} to the bad word list
:spellwrong! {wd}                Add word {wd} to the bad word list in internal-list
```

## -[ Binary files ]-

### -| Command line |-

```
vim -b <FILENAME>
```

### -| EX MODE |-

```
:set binary                     Enable Binary mode

:%!xxd                          Transform the file in hexadecimal representation
:%!xxd -r                       Revert to string representation
:set filetype=xxd               Syntax highlighting for hexadecimal representation
```

## -[ Non ASCII Characters ]-

AKA **Digraph**

Check the digraph table with: `:help digraph-table`

Digraph is used on _INSERT MODE_ with **[CTRL]+K** and a combination of one character and one Metacharacter

**Example:**

You want to write 'É', you have to use `[CTRL]+K E,`
You want to write 'π', you have to use `[CTRL]+K p*`

### -[ Digraph metacharacters ]-

| char name         | char | meaning                           |
| ----------------- | ---- | --------------------------------- |
| Exclamation mark  | !    | Grave                             |
| Apostrophe        | '    | Acute accent                      |
| Greater-Than sign | >    | Circumflex accent                 |
| Question mark     | ?    | Tilde                             |
| Hyphen-Minus      | -    | Macron                            |
| Left parenthesis  | (    | Breve                             |
| Full stop         | .    | Dot above                         |
| Colon             | :    | Diaeresis                         |
| Comma             | ,    | Cedilla                           |
| Underline         | \_   | Underline                         |
| Solidus           | /    | Stroke                            |
| Quotation mark    | "    | Double acute accent               |
| Semicolon         | ;    | Ogonek                            |
| Less-Than sign    | <    | Caron                             |
| Zero              | 0    | Ring above                        |
| Two               | 2    | Hook                              |
| Nine              | 9    | Horn                              |
| Equals            | =    | Cyrillic (= used as second char)  |
| Asterisk          | \*   | Greek                             |
| Percent sign      | %    | Greek/Cyrillic special            |
| Plus              | +    | smalls: Arabic, capitals: Hebrew  |
| Three             | 3    | some Latin/Greek/Cyrillic letters |
| Four              | 4    | Bopomofo                          |
| Five              | 5    | Hiragana                          |
| Six               | 6    | Katakana                          |

## -[ Edit files over network ]-

You can do this only if **netrw** ist _enabled_(On default VIM version it's the case)

### -| Command line |-

```
vim <PROTOCOL>://<USER>@<HOST>:<PORT>//<PATH>                   General form of netrw remote file accessing
vim scp://bignose@192.168.1.101//home/bignose/test.txt          Open a distant file over SSH, absolute version
vim scp://bignose@192.168.1.101/test.txt                        Open a distant file over SSH, relative to user home
vim scp://bignose@192.168.1.101/                                Open a distant directory
```

### -| EX MODE |-

```
:Nread <PROTOCOL>://<USER>@<HOST>:<PORT>//<PATH>                General form of netrw remote file accessing for reading
:Nwrite <PROTOCOL>://<USER>@<HOST>:<PORT>//<PATH>               General form of netrw remote file writing
```

## -[ Changing directories ]-

You can do this only if **netrw** ist _enabled_(On default VIM version it's the case)

### -| EX MODE |-

```
:Explore                        Open explorer(netrw)
:Lexplore                       Toggle left explorer window
:Sexplore                       Open explorer to current directory in an horizontal window
:Sexplore!                      Open explorer to current directory in an horizontal window
```

## -[ Backup ]-

VIM backup files for recovery crash.  
The user can change some settings

### -| EX MODE |-

```
:set backup                     Enable backup (Keep alternate backup during the session)
:set nobackup                   Disable backup
:set writebackup                Enable write backup (Save alternate before save the buffer, after remove the backup )
:set nowritebackup              Disable write backup

:set backupdir={pl,}            Define directories {pl,} where backup can be found and save
```

## -[ Convert to HTML ]-

### -| EX MODE |-

```
:runtime!syntax/2html.vim       Tranform the current buffer to HTML for visualization
:TOhtml                         Tranform the current buffer to HTML for visualization
```

## -[ Environment backup ]-

You can change general environment backup with the `viminfo` _option_.  
Check `:help 'viminfo'` for more information about the expected format.

You can also use **session** for specific environment backup.

### -| EX MODE |-

```
:mksession <FILENAME>           Save actual VIM session on the file <FILENAME>
:source <FILENAME>              Load the VIM session from the file <FILENAME>
```

## -[ Lines settings ]-

### -| EX MODE |-

```
:set wrap                       Enable line wraping if the line is larger than the screen
:set nowrap                     Disable line wraping

:set list                       Enable displaying of non visible characters
:set nolist                     Disable displaying of non visible characters
:help 'listchars'               More information about non visible characters and how to represent on screen
```

## -[ Tips ]-

### -| NORMAL MODE |-

```
q:                              Open window command history
q/                              Open window forward search history
q?                              Open window backward search history

g                               Start many multiple character commands in VIM(Check ':help g')

K                               Open the man page with the word under the cursor

[CTRL]+O                        Return to the previous jump
[CTRL]+T                        Return to the previous location in the tag stack
[CTRL]+]                        Goto the tag under the cursor
```

### -| Shell |-

Check `man 3 readline` to know how to edit **readline** configuration with **.inputrc** in your home directory  
With this you can use all commands that use **readline** with VI keystroke

`set editing-mode vi` to activate VI mode for readline

```
set -o vi                       Enable VI editing mode for the shell(Append to your _.bashrc_)
shopt -s lithist                Enable multiline history for Bash
```

#### Shell VI NORMAL MODE

```
[ESC]                           Go on NORMAL MODE but for the shell this time
v                               Edit the command line with VIM(Default system $EDITOR)

:cq                             To quit VIM without execute the command line
```

### -| EX MODE |-

```
[CTRL]+F                        Open window command/search history

:set cmdwinheight={n}           Define the height of the command history window
```

# Extend VIM

## -[ Plug-ins ]-

You can check the populars Neovim plug-in manager: [**Lazy.vim**](https://github.com/folke/lazy.nvim)

Use [Neovim Craft](https://neovimcraft.com) to found a _plug-in_ according to your needs  
Or If you prefer a list go to [**Awesome Neovim**](https://github.com/rockerBOO/awesome-neovim)

### -| Populars Plug-ins |-

| Plug-in                                                                   | Description                                                                       |
| ------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| [**Telescope**](https://github.com/nvim-telescope/telescope.nvim)         | The fuzzy finder                                                                  |
| [**Nvim-treesitter**](https://github.com/nvim-treesitter/nvim-treesitter) | Syntax tree for a better syntax coloring                                          |
| [**nvim-cmp**](https://github.com/hrsh7th/nvim-cmp)                       | Completion                                                                        |
| [**coq.nvim**](https://github.com/ms-jpq/coq_nvim)                        | Completion                                                                        |
| [**LSPSaga**](https://github.com/glepnir/lspsaga.nvim)                    | Lightweight LSP plugin based on Neovim's built-in LSP with a highly performant UI |
| [**null-ls.nvim**](https://github.com/jose-elias-alvarez/null-ls.nvim)    | Use Neovim as a language server to inject LSP                                     |
| [**mason.nvim**](https://github.com/williamboman/mason.nvim)              | Easily install and manage LSP servers, DAP servers, linters, and formatters       |
| [**LuaSnip**](https://github.com/L3MON4D3/LuaSnip)                        | Snippet Engine for Neovim                                                         |
| [**Trouble**](https://github.com/folke/trouble.nvim)                      | Help show diagnostics                                                             |
| [**gitsigns.nvim**](https://github.com/lewis6991/gitsigns.nvim)           | Git decoration                                                                    |
| [**Neogit**](https://github.com/TimUntersberger/neogit)                   | Git TUI(Text User Interface)                                                      |
| [**nvim-autopairs**](https://github.com/windwp/nvim-autopairs)            | Autopair                                                                          |
| [**lazy.nvim**](https://github.com/folke/lazy.nvim)                       | Modern plugin manager                                                             |
| [**lualine.nvim**](https://github.com/nvim-lualine/lualine.nvim)          | Status line                                                                       |
| [**Tokyo Night**](https://github.com/folke/tokyonight.nvim)               | Theme                                                                             |
| [**leap.nvim**](https://github.com/ggandor/leap.nvim)                     | Go anywhere on the screen with a maximum of 3 keystrokes                          |
| [**Which Key**](https://github.com/folke/which-key.nvim)                  | Displays a popup with possible key bindings of the command you started typing     |
| [**nvim-tree**](https://github.com/nvim-tree/nvim-tree.lua)               | File explorer                                                                     |
| [**Harpoon**](https://github.com/ThePrimeagen/harpoon)                    | File switcher, Faster navigation                                                  |
| [**Diffview.nvim**](https://github.com/sindrets/diffview.nvim)            | Git Merge tool and files history                                                  |
| [**mini.nvim**](https://github.com/echasnovski/mini.nvim)                 | Library of 20+ independent Lua modules improving overall Neovim                   |
| [**dashboard-nvim**](https://github.com/glepnir/dashboard-nvim)           | Dashboard                                                                         |

#### All-in-One IDE configurations

| Link                                                              | Description                                                                                                |
| ----------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| [**NvChad**](https://github.com/NvChad/NvChad)                    | An attempt to make neovim cli functional like an IDE while being very beautiful, blazing fast startuptime  |
| [**LunarVim**](https://github.com/LunarVim/LunarVim)              | An IDE layer for Neovim with sane defaults                                                                 |
| [**AstroNvim**](https://github.com/AstroNvim/AstroNvim)           | An aesthetic and feature-rich neovim config that is extensible and easy to use with a great set of plugins |
| [**LazyVim**](https://github.com/LazyVim/LazyVim)                 | The flexibility to tweak your config as needed, along with the convenience of a pre-configured setup       |
| [**kickstart.nvim**](https://github.com/nvim-lua/kickstart.nvim)  | A launch point for your personal Neovim configuration                                                      |
| [**CosmicNvim**](https://github.com/CosmicNvim/CosmicNvim)        | Lightweight and opinionated Neovim config for web development                                              |
| [**Doom Nvim**](https://github.com/doom-neovim/doom-nvim)         | Configuration for the advanced martian hacker                                                              |
| [**CodeArtart**](https://github.com/artart222/CodeArt)            | Use NeoVim as general purpose IDE                                                                          |
| [**Nyoom.nvim**](https://github.com/nyoom-engineering/nyoom.nvim) | Used as a framework config for users to extend and add upon, leading to a more unique editing experience   |
| [**NVIM-IDE**](https://github.com/ldelossa/nvim-ide)              | nvim-ide is a complete IDE layer for Neovim, heavily inspired by vscode                                    |

#### Writer plug-ins

Check [Vim/Neovim Plugins for Writing](https://alpha2phi.medium.com/vim-neovim-plugins-for-writing-d18414c7b21d) for more informations

| Plug-in                                                                      | Description                                                                                    |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| [**Zen Mode**](https://github.com/folke/zen-mode.nvim)                       | Distraction-free                                                                               |
| [**true-zen.nvim**](https://github.com/Pocco81/true-zen.nvim)                | Clean and elegant distraction-free writing                                                     |
| [**markdowny.nvim**](https://github.com/antonk52/markdowny.nvim)             | For markdown like keybindings                                                                  |
| [**vim-markdown-toc**](https://github.com/mzlogin/vim-markdown-toc)          | Generate table of contents for Markdown files                                                  |
| [**glow.nvim**](https://github.com/ellisonleao/glow.nvim)                    | Preview Markdown code directly in your Neovim terminal                                         |
| [**markdown-preview.nvim**](https://github.com/iamcco/markdown-preview.nvim) | Preview Markdown on your modern browser with synchronised scrolling and flexible configuration |
| [**peek.nvim**](https://github.com/toppair/peek.nvim)                        | Preview Markdown                                                                               |
| [**KNAP**](https://github.com/frabjous/knap)                                 | Preview for Latex and Markdown                                                                 |
| [**vim-grammarous**](https://github.com/rhysd/vim-grammarous)                | Grammar checker                                                                                |
| [**vim-LanguageTool**](https://github.com/dpelle/vim-LanguageTool)           | Integrates the LanguageTool grammar checker                                                    |
| [**VIM Table Mode**](https://github.com/dhruvasagar/vim-table-mode)          | Automatic table creator & formatter                                                            |
| [**vim-pandoc**](https://github.com/vim-pandoc/vim-pandoc)                   | Provides facilities to integrate Neovim with pandoc                                            |
| [**vim-pencil**](https://github.com/preservim/vim-pencil)                    | Make Vim as powerful a tool for writers                                                        |
| [**HighStr.nvim**](https://github.com/Pocco81/high-str.nvim)                 | Free highlighter                                                                               |

### -| Links |-

| Link                                                                                  | Description                                                                         |
| ------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| [**Neovimcraft**](https://neovimcraft.com/)                                           | Plugins search engine specific to Neovim                                            |
| [**VimAwesome**](https://vimawesome.com/)                                             | The VIM plugins search engine                                                       |
| [**Awesome Neovim**](https://github.com/rockerBOO/awesome-neovim)                     | List of Neovim plugins by category                                                  |
| [**Neovim plugins blog**](https://www.barbarianmeetscoding.com/notes/neovim-plugins/) | Tutorial for most populars neovim plugins with a good start for using Lua on Neovim |
| [**This Week In Neovim**](https://this-week-in-neovim.org/)                           | Weekly news about Neovim                                                            |
