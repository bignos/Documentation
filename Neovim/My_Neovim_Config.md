<h1 align="center">
    My Neovim Configuration with Lazy.vim
</h1>

Here I will describe my installation **step by step**.

The objective is not necessarily to copy everything but rather to understand how the configuration of **Neovim** with **Lazy** works.
This will also allow me to present my _choice of plugins_ and to explain why they are useful.

The goal of this configuration is to have a practical editor for programming but still light to avoid visual information overload.  
The configuration is _not oriented to the aesthetic_ even if this point is not forgotten.
I will focus on the **practical** and **fast** side of the editor.
The configuration will have to be flexible to adapt to different types of files or languages, but without ever doing too much.

# Kitty

You have to install a fast, featured, GPU based terminal emulator: [**Kitty**](https://sw.kovidgoyal.net/kitty/) to use [**Nerdfont**](https://www.nerdfonts.com/)  
[**Nerdfont**](https://www.nerdfonts.com/) is required to display icons

Install **Kitty** on [Ubuntu](https://ubuntu.com/)
``` bash
sudo apt update
sudo apt install kitty
```

Now you can use Kitty with Neovim

## Install Nerd fonts for kitty

Go to [**Nerd font release page**](https://github.com/ryanoasis/nerd-fonts/releases), and download **NerdFontsSymbolsOnly.zip**
Unzip the file in `~/.fonts` and run the command `fc-cache -fv` to rebuild the cache  
Edit the file `~/.config/kitty/kitty.conf` and add this line:

```
symbol_map U+23FB-U+23FE,U+2665,U+26A1,U+2B58,U+E000-U+E00A,U+E0A0-U+E0A3,U+E0B0-U+E0C8,U+E0CA,U+E0CC-U+E0D2,U+E0D4,U+E200-U+E2A9,U+E300-U+E3E3,U+E5FA-U+E634,U+E700-U+E7C5,U+EA60-U+EBEB,U+F000-U+F2E0,U+F300-U+F32F,U+F400-U+F4A9,U+F500-U+F8FF Symbols Nerd Font Mono
```

Reload _Kitty_ with <CTRL>+<F5>  
Now _Kitty_ use Nerd fonts symbols with the current font

# Lazy.vim

[Lazy.vim](https://github.com/folke/lazy.nvim) is a modern plugin manager for Neovim

## Initialize

The goal here is to install a minimal configuration startup to use [Lazy.vim](https://github.com/folke/lazy.nvim) with a clear directory structure

Create the file `~/.config/nvim/init.lua`

``` lua
require("config.lazy")
```

This file is only a boostrap to centralize the configuration in `lua/config` directory  

After create directories architecture:

``` bash
mkdir -p ~/.config/nvim/lua/config ~/.config/nvim/lua/plugins ~/.config/nvim/lua/helpers
```

Create all skeleton files

``` bash
touch ~/.config/nvim/lua/config/autocmds.lua
touch ~/.config/nvim/lua/config/keymaps.lua
touch ~/.config/nvim/lua/config/options.lua
```

Create the theme plugin [**tokyonight**](https://github.com/folke/tokyonight.nvim) file `~/.config/nvim/lua/plugins/tokyonight.lua`

``` lua
return { {
    "folke/tokyonight.nvim",
    lazy = false,
    priority = 1000,
    config = function()
        vim.cmd([[colorscheme tokyonight]])
    end
} }
```

Create the file `~/.config/nvim/lua/config/lazy.lua`  
This file is the **main entry point** for the neovim configuration start.
_It should not need to be modified frequently_

``` lua
local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"
if not vim.loop.fs_stat(lazypath) then
    vim.fn.system({
        "git",
        "clone",
        "--filter=blob:none",
        "https://github.com/folke/lazy.nvim.git",
        "--branch=stable", -- latest stable release
        lazypath,
    })
end
vim.opt.rtp:prepend(lazypath)

require("config.options")
require("lazy").setup({
    spec = {
        { import = "plugins" },
    },
    defaults = {
        version = false, -- always use the latest git commit
    },
    install = { colorscheme = { "tokyonight" } },
    checker = { enabled = true, notify = false }, -- automatically check for plugin updates
})
require("config.keymaps")
require("config.autocmds")
```

- Now you have [Lazy.vim](https://github.com/folke/lazy.nvim) installed
- And you have directories structure(config, plugins, helpers) with minimal skeleton files(autocmds.lua, keymaps.lua, lazy.lua, options.lua)  
- And you have add the plugin [**tokyonight**](https://github.com/folke/tokyonight.nvim) to avoid errors with **Lazy** configuration file

# Neovim options configuration

The files `options.lua` have to be used for configuration options of Neovim.  
I'm trying to avoid plugins dependencies, but in this case we have:
- [_treesitter_](https://github.com/nvim-treesitter/nvim-treesitter)
- [_conceal_](https://github.com/Jxstxs/conceal.nvim)
- [_undotree_](https://github.com/mbbill/undotree)

In here you have **Globals settings** and **Options** for Neovim

## The Leader Key

In this configuration **The Leader Key** is _very important_ because I use it for **personal key mapping**

Neovim already uses a big part of the keyboard, the purpose of this key will be to be able to add all the **extra shortcuts** I would need  
It is important that the key is **easily accessible** because it will be **heavily solicited**  
That's why like many I chose to use **Space** as **The Leader Key**, _one key to control them all_

``` lua
vim.g.mapleader          = " "
vim.g.maplocalleader     = " "
```

## The Neovim Native file explorer

**Netrw** is the native file explorer for Neovim, fore more informations check `:help netrw`

At first sight it seems particularly **not user-friendly**.
It is really not pretty but is **very usable** and does not require the installation of any plugin.
It requires a little time of adaptation to know the useful keys but it does its job very well.

For more precise information about **Netrw** key mapping  `:help netrw-quickmap`

In addition, unlike the classic file explorer, it can also connect remotely and allow **remote file editing**.

The goal here is to change these default settings to make it **less ugly** by positioning it vertically on the left

``` lua
-- netrw
vim.g.netrw_winsize      = 16
vim.g.netrw_liststyle    = 3
vim.g.netrw_banner       = 0
vim.g.netrw_browse_split = 4
vim.g.netrw_altv         = 1
```

## Better Folding with Treesitter

The objective here is to use the power of the syntax tree produced by [**Treesitter**](https://github.com/nvim-treesitter/nvim-treesitter) to help Neovim to build folds 

``` lua
-- Folding linked with treesitter
opt.foldlevel            = 20
opt.foldmethod           = "expr"
opt.foldexpr             = "nvim_treesitter#foldexpr()"
```

## Enable Conceal

Set the [_conceal_](https://github.com/Jxstxs/conceal.nvim) level to 3
Concealment is mainly used to have a more pleasant reading of structured content

``` lua
-- Conceal
opt.conceallevel         = 3
```

## Neovim settings

### Line number

No one will argue with me when I say that having **line numbers** to edit code is essential
But in my case I will choose a relative numbering to be able to navigate on the screen more easily

``` lua
opt.nu                   = true
opt.relativenumber       = true
```

### Indentation

These parameters will guarantee an indentation of 4 with spaces and not tabs

``` lua
opt.tabstop              = 4
opt.softtabstop          = 4
opt.shiftwidth           = 4
opt.expandtab            = true
```

Activate Smart indentation

``` lua
opt.smartindent          = true
```

Don't wrap line please my eyes burn

``` lua
opt.wrap                 = false
```

## Backup? Undotree to the rescue

The use of [**Undotree**](https://github.com/mbbill/undotree) allows not to use swap files, because even in case of crash [**Undotree**](https://github.com/mbbill/undotree) will be able to find the last changes.
That's why we disable swap files and backups, they have no more use

``` lua
opt.swapfile             = false
opt.backup               = false
opt.undodir              = os.getenv("HOME") .. "/.vim/undodir"
opt.undofile             = true
```

## Highlight search

By default in Neovim **Highlight search** is enabled  
As a consequence, after each search you have to do `:noh`, to remove the highlights  
To avoid this **horror** you can activate the **highlights during the search only**

``` lua
opt.hlsearch             = false
opt.incsearch            = true
```

## Better colors

We are in 2023, I think we can use more than 16 colors in a terminal.

``` lua
opt.termguicolors        = true
```

## Complete options.lua file

``` lua
vim.g.mapleader          = " "
vim.g.maplocalleader     = " "

-- netrw
vim.g.netrw_winsize      = 16
vim.g.netrw_liststyle    = 3
vim.g.netrw_banner       = 0
vim.g.netrw_browse_split = 4
vim.g.netrw_altv         = 1

local opt                = vim.opt

-- Folding linked with treesitter
opt.foldlevel            = 20
opt.foldmethod           = "expr"
opt.foldexpr             = "nvim_treesitter#foldexpr()"

-- Conceal
opt.conceallevel         = 3

-- Neovim
opt.nu                   = true
opt.relativenumber       = true

opt.tabstop              = 4
opt.softtabstop          = 4
opt.shiftwidth           = 4
opt.expandtab            = true

opt.smartindent          = true
opt.wrap                 = false

opt.swapfile             = false
opt.backup               = false
opt.undodir              = os.getenv("HOME") .. "/.vim/undodir"
opt.undofile             = true

opt.hlsearch             = false
opt.incsearch            = true

opt.termguicolors        = true
```

# Personal Key mapping

The configuration of the personal key mapping must be done in the file provided for this purpose: `~/.config/nvim/lua/config/keymaps.lua`

**Key mapping** here are only for Neovim **not for the Plugins**  
For plugins the key mapping is described in their configuration file  
It is important to _decouple_ this because it allows not to be polluted by the additions and deletions of the different plugins

My idea here is to minimize the number of personal key mappings to avoid overloading the brain with new shortcuts to learn  
So I will orient my configuration to have only the essential shortcuts

I would use _**lower case** characters for common operations_ and _**upper case** for less common operations_
I'll try to stick to it as much as I can

It is also important to indicate a description with 'desc' for each shortcut, it will be used with the [which-key](https://github.com/folke/which-key.nvim) plugin

## Lazy

The configuration being oriented around Lazy VIM,so the first thing to do is to be able to call the Lazy panel

So I have chose `<leader>L`

``` lua
vim.keymap.set("n", "<leader>L", vim.cmd.Lazy, { desc = "Lazy" })

```
## Buffers

In this configuration we will **not use tabs**, but **only buffers** which will also act as tabs  
To do this there is a dependency with the [**bufferline**](https://github.com/akinsho/bufferline.nvim) plugin

So we want to be able to navigate easily between the buffers  
The idea here is to avoid moving the hands as much as possible, so I'm going to use the `h` and `l` keys with a **meta key**, in this case **Shift**

``` lua
vim.keymap.set("n", "H", "<cmd>BufferLineCyclePrev<cr>", { desc = "Prev buffer" })
vim.keymap.set("n", "L", "<cmd>BufferLineCycleNext<cr>", { desc = "Next buffer" })
```

## Netrw

It is important to be able to move in a file system easily to have a visualization of the hierarchy and the files present  
Setting up a shortcut for [**Netrw**](https://vonheikemen.github.io/devlog/tools/using-netrw-vim-builtin-file-explorer/) is a bit redundant in the sense that I will use another explorer to have a nicer presentation.
But it will be very useful while waiting to have installed an explorer with a nicer presentation

I chose a capital letter for the action because it will be less used in the future
So I will use `<leader>`+`E`

``` lua
vim.keymap.set("n", "<leader>E", vim.cmd.Lexplore, { desc = "File Explorer" })
```

## Insert to Normal Mode with the Terminal

One of the most interesting features of **Neovim** is the integration of the **Terminal**  
But once in **Insert Mode** to return to **Normal Mode** you have to use the `<CTRL-\>` shortcut so clearly it's **not natural** at all

I chose to use `<ESC><ESC>` which seems to me more logical

``` lua
vim.keymap.set("t", "<esc><esc>", "<C-\\><C-n>", { desc = "Normal mode for Terminal" })
```

## Move Selected lines in Visual Mode

There is obviously _no need_ for a shortcut to move lines with **Neovim**  
But doing it **visually** with an **automatic indentation** is still very convenient to move code blocks.

So as for the **buffers** I will use one of the Neovim Movement keys and a meta key to avoid moving the hands 

Don't forget that these shortcuts work **Only on Visual Mode**

I'm going to use the `j` and `k` with **Shift**

``` lua
vim.keymap.set("v", "J", ":m '>+1<CR>gv=gv", { desc = "Move selected line down" })
vim.keymap.set("v", "K", ":m '<-2<CR>gv=gv", { desc = "Move selected line up" })
```

## The Missing Yank

For some reason there is no shortcut to **Yank from the cursor to the end of the line**.
So we have to create it with the shortcut `Y`

``` lua
vim.keymap.set("n", "Y", "yg$", { desc = "The missing yank" })
```

## Join without moving the cursor

The idea here is to redo the 'J' shortcut to **join 2 lines without moving the cursor**.
This avoids unnecessary backtracking

``` lua
vim.keymap.set("n", "J", "mzJ`z", { desc = "Join lines" })
```

## Center far cursor move

The objective is to make the **screen movements centered on the cursor**, this _avoids wasting time looking for where it is_ on the screen

``` lua
vim.keymap.set("n", "<C-d>", "<C-d>zz", { desc = "Page-down" })
vim.keymap.set("n", "<C-u>", "<C-u>zz", { desc = "Page-up" })
vim.keymap.set("n", "n", "nzzzv", { desc = "Next search" })
vim.keymap.set("n", "N", "Nzzzv", { desc = "Previous search" })
```

## Replace with the Black Hole

One of the behaviors of **Neovim** that can be particularly **annoying** is the behavior that consists in putting what is erased in the unnamed register(AKA `""`) what replaces it  
In most cases this is not a problem but when you want to paste the same thing several times with replacement **it becomes hell**  

To avoid this we will add an additional shortcut to copy what is deleted in the **Black Hole**(AKA `"_`)
For this shortcut we use `<leader>p`

``` lua
vim.keymap.set("x", "<leader>p", "\"_dP", { desc = "Better Past for replace" })
```
**This shortcut is a mistake because Neovim already has one for this purpose: `P` in Visual Mode**

## The hating Q

Who want to use a command to _repeat the last recorded register_ ?
Here is  **very personal**, I will make a **Neovim** shortcut **inactive** because I do not use it but I often press the key by mistake.

``` lua
vim.keymap.set("n", "Q", "<nop>")
```

## Yank to the system clipboard

Let's add a shortcut to make a **yank** to the **system clipboard**(AKA `"+`) with **<leader>**

``` lua
vim.keymap.set("n", "<leader>y", "\"+y", { desc = "Yank to clipboard" })
vim.keymap.set("v", "<leader>y", "\"+y", { desc = "Yank to clipboard" })
vim.keymap.set("n", "<leader>Y", "\"+Y", { desc = "Yank to clipboard" })
```

## Delete with the Black Hole

Some more shortcuts to **delete without replace** the unnamed register(`""`)

``` lua
vim.keymap.set("n", "<leader>d", "\"_d", { desc = "Delete to the blackhole" })
vim.keymap.set("v", "<leader>d", "\"_d", { desc = "Delete to the blackhole" })
```

## Quickfix and Location list fast navigation

We will also add navigation shortcuts for **Quickfix** and **Location** list navigation to be able to move quickly.

- <CTRL> for quickfix  
- <Leader> for location

``` lua
vim.keymap.set("n", "<C-k>", "<cmd>cnext<CR>zz", { desc = "Next quick list" })
vim.keymap.set("n", "<C-j>", "<cmd>cprev<CR>zz", { desc = "Previous quick list" })
vim.keymap.set("n", "<leader>k", "<cmd>lnext<CR>zz", { desc = "Next location list" })
vim.keymap.set("n", "<leader>j", "<cmd>lprev<CR>zz", { desc = "Previous location list" })
```

## The Live Replace

This shortcut is clearly **not essential**, but it allows you to see a the replacement in a visual way

I would say that it is mostly a visual comfort

I'm going to use `<leader>s` (s for search) for this shortcut

``` lua
vim.keymap.set(
    "n", "<leader>s",
    [[:%s/\<<C-r><C-w>\>/<C-r><C-w>/gI<Left><Left><Left>]],
    { desc = "Rename all occurrence of the word under the cursor" }
)
```

## Source your configuration files

The latter is clearly **the least essential**  
But it will make it easier to configure Neovim, by allowing the configuration to be reloaded to verify changes

``` lua
vim.keymap.set(
    "n", "<leader><leader>",
    function()
        vim.cmd("so")
    end,
    { desc = "Source the current file" }
)
```

## The complete keymaps.lua file

``` lua
-- lazy
vim.keymap.set("n", "<leader>L", vim.cmd.Lazy, { desc = "Lazy" })

-- Buffers
vim.keymap.set("n", "H", "<cmd>BufferLineCyclePrev<cr>", { desc = "Prev buffer" })
vim.keymap.set("n", "L", "<cmd>BufferLineCycleNext<cr>", { desc = "Next buffer" })

-- Netrw
vim.keymap.set("n", "<leader>E", vim.cmd.Lexplore, { desc = "File Explorer" })

-- Terminal
vim.keymap.set("t", "<esc><esc>", "<C-\\><C-n>", { desc = "Normal mode for Terminal" })

-- Move line in visual mode
vim.keymap.set("v", "J", ":m '>+1<CR>gv=gv", { desc = "Move selected line down" })
vim.keymap.set("v", "K", ":m '<-2<CR>gv=gv", { desc = "Move selected line up" })

-- The missing yank
vim.keymap.set("n", "Y", "yg$", { desc = "The missing yank" })

-- Join without the cursor motion
vim.keymap.set("n", "J", "mzJ`z", { desc = "Join lines" })

-- Centered Page-down
vim.keymap.set("n", "<C-d>", "<C-d>zz", { desc = "Page-down" })

-- Centered Page-up
vim.keymap.set("n", "<C-u>", "<C-u>zz", { desc = "Page-up" })

-- Centered Next search
vim.keymap.set("n", "n", "nzzzv", { desc = "Next search" })

-- Centered Previous search
vim.keymap.set("n", "N", "Nzzzv", { desc = "Previous search" })

-- No Q hole
vim.keymap.set("n", "Q", "<nop>")

-- Better system integration with system clipboard
vim.keymap.set("n", "<leader>y", "\"+y", { desc = "Yank to clipboard" })
vim.keymap.set("v", "<leader>y", "\"+y", { desc = "Yank to clipboard" })
vim.keymap.set("n", "<leader>Y", "\"+Y", { desc = "Yank to clipboard" })

vim.keymap.set("n", "<leader>d", "\"_d", { desc = "Delete to the blackhole" })
vim.keymap.set("v", "<leader>d", "\"_d", { desc = "Delete to the blackhole" })

-- Quick and location list navigation
vim.keymap.set("n", "<C-k>", "<cmd>cnext<CR>zz", { desc = "Next quick list" })
vim.keymap.set("n", "<C-j>", "<cmd>cprev<CR>zz", { desc = "Previous quick list" })
vim.keymap.set("n", "<leader>k", "<cmd>lnext<CR>zz", { desc = "Next location list" })
vim.keymap.set("n", "<leader>j", "<cmd>lprev<CR>zz", { desc = "Previous location list" })

-- Live replace
vim.keymap.set(
    "n", "<leader>s",
    [[:%s/\<<C-r><C-w>\>/<C-r><C-w>/gI<Left><Left><Left>]],
    { desc = "Rename all occurrence of the word under the cursor" }
)

-- Reload a config file
vim.keymap.set(
    "n", "<leader><leader>",
    function()
        vim.cmd("so")
    end,
    { desc = "Source the current file" }
)
```

# Personal Auto Commands

In the `~/.config/nvim/lua/config/autocmds.lua` file, we will put the actions to be performed **when Neovim is started**

In my case here everything is **about comfort**

## Lazy auto group for which-key

I use this auto group to generate groups with [which-key](https://github.com/folke/which-key.nvim)

``` lua
local function augroup(name)
    return vim.api.nvim_create_augroup("lazyvim_" .. name, { clear = true })
end
```

## Close some file type buffer with q

The objective here is to be able to easily **close information or tool buffers** with only one key: `q`

``` lua
vim.api.nvim_create_autocmd("FileType", {
    group = augroup("close_with_q"),
    pattern = {
        "qf",
        "help",
        "man",
        "notify",
        "lspinfo",
        "tsplayground",
        "netrw",
    },
    callback = function(event)
        vim.bo[event.buf].buflisted = false
        vim.keymap.set("n", "q", "<cmd>close<cr>", { buffer = event.buf, silent = true })
    end,
})
```

## Activate spell checking for some file type

The objective here is to activate **spell checking** automatically for editorial files like markdown

``` lua
vim.api.nvim_create_autocmd("FileType", {
    group = augroup("auto_spell"),
    pattern = { "gitcommit", "markdown" },
    callback = function()
        vim.opt_local.spell = true
    end,
})
```

## Restore the cursor position

To keep track of the files I edit, I like the cursor to return to the last editing position
I wanted something simple so I didn't install any plugin for that  
I have integrated a solution given on [nvim-lastplace](https://github.com/ethanholz/nvim-lastplace/blob/main/lua/nvim-lastplace/init.lua)

``` lua
local ignore_buftype = { "quickfix", "nofile", "help" }
local ignore_filetype = { "gitcommit", "gitrebase", "svn", "hgcommit" }

local function run()
    if vim.tbl_contains(ignore_buftype, vim.bo.buftype) then
        return
    end

    if vim.tbl_contains(ignore_filetype, vim.bo.filetype) then
        -- reset cursor to first line
        vim.cmd [[normal! gg]]
        return
    end

    -- If a line has already been specified on the command line, we are done
    --   nvim file +num
    if vim.fn.line(".") > 1 then
        return
    end

    local last_line = vim.fn.line([['"]])
    local buff_last_line = vim.fn.line("$")

    -- If the last line is set and the less than the last line in the buffer
    if last_line > 0 and last_line <= buff_last_line then
        local win_last_line = vim.fn.line("w$")
        local win_first_line = vim.fn.line("w0")
        -- Check if the last line of the buffer is the same as the win
        if win_last_line == buff_last_line then
            -- Set line to last line edited
            vim.cmd [[normal! g`"]]
            -- Try to center
        elseif buff_last_line - last_line > ((win_last_line - win_first_line) / 2) - 1 then
            vim.cmd [[normal! g`"zz]]
        else
            vim.cmd [[normal! G'"<c-e>]]
        end
    end
end

vim.api.nvim_create_autocmd({ 'BufWinEnter', 'FileType' }, {
    group    = vim.api.nvim_create_augroup('nvim-lastplace', {}),
    callback = run
})
```
## The complete autocmds.lua file

``` lua
local function augroup(name)
    return vim.api.nvim_create_augroup("lazyvim_" .. name, { clear = true })
end

-- Close some filetypes with <q>
vim.api.nvim_create_autocmd("FileType", {
    group = augroup("close_with_q"),
    pattern = {
        "qf",
        "help",
        "man",
        "notify",
        "lspinfo",
        "tsplayground",
        "netrw",
    },
    callback = function(event)
        vim.bo[event.buf].buflisted = false
        vim.keymap.set("n", "q", "<cmd>close<cr>", { buffer = event.buf, silent = true })
    end,
})

-- Check for spell in text filetypes
vim.api.nvim_create_autocmd("FileType", {
    group = augroup("auto_spell"),
    pattern = { "gitcommit", "markdown" },
    callback = function()
        vim.opt_local.spell = true
    end,
})

-- Restore the cursor position
-- adapted from https://github.com/ethanholz/nvim-lastplace/blob/main/lua/nvim-lastplace/init.lua
--
-- More clear version https://this-week-in-neovim.org/2023/Jan/2#tips
local ignore_buftype = { "quickfix", "nofile", "help" }
local ignore_filetype = { "gitcommit", "gitrebase", "svn", "hgcommit" }

local function run()
    if vim.tbl_contains(ignore_buftype, vim.bo.buftype) then
        return
    end

    if vim.tbl_contains(ignore_filetype, vim.bo.filetype) then
        -- reset cursor to first line
        vim.cmd [[normal! gg]]
        return
    end

    -- If a line has already been specified on the command line, we are done
    --   nvim file +num
    if vim.fn.line(".") > 1 then
        return
    end

    local last_line = vim.fn.line([['"]])
    local buff_last_line = vim.fn.line("$")

    -- If the last line is set and the less than the last line in the buffer
    if last_line > 0 and last_line <= buff_last_line then
        local win_last_line = vim.fn.line("w$")
        local win_first_line = vim.fn.line("w0")
        -- Check if the last line of the buffer is the same as the win
        if win_last_line == buff_last_line then
            -- Set line to last line edited
            vim.cmd [[normal! g`"]]
            -- Try to center
        elseif buff_last_line - last_line > ((win_last_line - win_first_line) / 2) - 1 then
            vim.cmd [[normal! g`"zz]]
        else
            vim.cmd [[normal! G'"<c-e>]]
        end
    end
end

vim.api.nvim_create_autocmd({ 'BufWinEnter', 'FileType' }, {
    group    = vim.api.nvim_create_augroup('nvim-lastplace', {}),
    callback = run
})
```

Now we have finish Neovim main configuration
We can finally **install Plugins** comfortably

# Plug-ins installation

Now we can install all the plugins to make **Neovim Magic**

The idea is to make **one file per plugin** in the plugins directory `~/.config/nvim/lua/plugins/`
The plugin configuration will contain the **settings** and the **specific shortcuts** of the plugin

I have sorted the plugins from the most useful to the least  
This is obviously a subjective ranking, according to my use

## General

As I mentioned above, one plugin = one file

Lazy expects a return with an array of plugins  
Here is the minimal file for the declaration of a plugin

``` lua
return {
    {}
}
```

The first element of the table is considered as the Github tag of the plugin, it is the essential part  
Without this information nothing can happen

Just a little code example to illustrate:

``` lua
return {
    { "bignos/myvimplugin" }
}
```

You can find a precise description of the elements expected by Lazy [here](https://github.com/folke/lazy.nvim#-plugin-spec)

On my side I will use a lot the following elements:

- `dependencies` : To declare dependencies with other plugins
- `init` : Allows to initialize the plugin with a Lua function
- `keys` : For specific plugin key maps
- `config` : Like `init` but more options oriented 

## Telescope

The Ultimate search plugins, after using it you will wonder how we did it before.

[Telescope](https://github.com/nvim-telescope/telescope.nvim) allows you to do extremely fast search in almost all type of list  
From files to buffers, including Git commits  
Moreover it can integrate extensions to give itself new powers(treesitter).

While using it I even wondered about the usefulness of a file explorer, when you have everything at hand with a file viewer.

In short, if you don't have this plugin installed you missed your life, **go back to VScode !!!**

With **Telescope** there is a coupling with [Treesitter](https://github.com/nvim-treesitter/nvim-treesitter) and [Trouble](https://github.com/folke/trouble.nvim)

Initialization:

``` lua
{
    "nvim-telescope/telescope.nvim",
    tag = "0.1.1",
    dependencies = { "nvim-lua/plenary.nvim" },
}
```

Key mapping:

``` lua
{
    keys = {
        -- Find Files
        {
            "<leader>ff",
            function() require("telescope.builtin").find_files() end,
            desc = "Find Files",
        },

        -- Find Recent files
        {
            "<leader>fr",
            function() require("telescope.builtin").oldfiles() end,
            desc = "Find Recent files",
        },


        -- Find Buffers
        {
            "<leader>fb",
            function() require("telescope.builtin").buffers() end,
            desc = "Find Buffers",
        },

        -- Find Git files
        {
            "<leader>fg",
            function() require("telescope.builtin").git_files() end,
            desc = "Find Git File",
        },

        -- Find Strings on files
        {
            "<leader>fs",
            function() require("telescope.builtin").grep_string({ search = vim.fn.input("Grep > ") }) end,
            desc = "Find Strings on files",
        },

        -- Find Symbols from Treesiter
        {
            "<leader>fS",
            function() require("telescope.builtin").treesitter() end,
            desc = "Find Symbols",
        },

        -- Find Diagnostics lsp
        {
            "<leader>fd",
            function() require("telescope.builtin").diagnostics() end,
            desc = "Find Diagnostics lsp",
        },

        -- Find Plugin files (Lazy specific)
        {
            "<leader>fp",
            function() require("telescope.builtin").find_files({ cwd = require("lazy.core.config").options.root }) end,
            desc = "Find Plugin File",
        },

    },
}
```

Init function:

``` lua
{
    init = function()
        local trouble = require("trouble.providers.telescope")
        local telescope = require("telescope")
        telescope.setup {
            defaults = {
                mappings = {
                    i = { ["<c-t>"] = trouble.open_with_trouble },
                    n = { ["<c-t>"] = trouble.open_with_trouble },
                },
            },
        }
    end
}
```

The complete `~/.config/nvim/lua/plugins/telescope.lua` file:

``` lua
return {
    {
        "nvim-telescope/telescope.nvim",
        tag = "0.1.1",
        dependencies = { "nvim-lua/plenary.nvim" },
        keys = {
            -- Find Files
            {
                "<leader>ff",
                function() require("telescope.builtin").find_files() end,
                desc = "Find Files",
            },

            -- Find Recent files
            {
                "<leader>fr",
                function() require("telescope.builtin").oldfiles() end,
                desc = "Find Recent files",
            },


            -- Find Buffers
            {
                "<leader>fb",
                function() require("telescope.builtin").buffers() end,
                desc = "Find Buffers",
            },

            -- Find Git files
            {
                "<leader>fg",
                function() require("telescope.builtin").git_files() end,
                desc = "Find Git File",
            },

            -- Find Strings on files
            {
                "<leader>fs",
                function() require("telescope.builtin").grep_string({ search = vim.fn.input("Grep > ") }) end,
                desc = "Find Strings on files",
            },

            -- Find Symbols from Treesiter
            {
                "<leader>fS",
                function() require("telescope.builtin").treesitter() end,
                desc = "Find Symbols",
            },

            -- Find Diagnostics lsp
            {
                "<leader>fd",
                function() require("telescope.builtin").diagnostics() end,
                desc = "Find Diagnostics lsp",
            },

            -- Find Plugin files (Lazy specific)
            {
                "<leader>fp",
                function() require("telescope.builtin").find_files({ cwd = require("lazy.core.config").options.root }) end,
                desc = "Find Plugin File",
            },

        },
        init = function()
            local trouble = require("trouble.providers.telescope")
            local telescope = require("telescope")
            telescope.setup {
                defaults = {
                    mappings = {
                        i = { ["<c-t>"] = trouble.open_with_trouble },
                        n = { ["<c-t>"] = trouble.open_with_trouble },
                    },
                },
            }
        end
    },
}
```

## Undotree

[Undotree](https://github.com/mbbill/undotree) is extremely useful because it allows you to **visualize the changes** in Neovim

You can **go back to any state of the file**, and it saves all the changes  
That's why **swap** or **backup** files are not really interesting anymore

The complete `~/.config/nvim/lua/plugins/undotree.lua` file:

``` lua
return {
    {
        "mbbill/undotree",
        keys = {
            -- UndoTree
            {
                "<leader>u",
                vim.cmd.UndotreeToggle,
                desc = "UndoTree",
            },

        }
    }
}
```

## Which-Key

[Which-key](https://github.com/folke/which-key.nvim) displays a popup with possible key bindings of the command you started typing

At first I thought it would _help mostly beginners_, but in reality it's a tool that helps you to go **faster** and to **remember existing shortcuts**

It will not be invasive either because it will be triggered after a timeout  
So if you type the commands fast enough it doesn't display but when you are hesitant it helps you

Clearly it is a very helpful tool, it will also allow you to see the content of the **registers**, **marks** and even suggestions for **spelling corrections**  
It integrates perfectly with the **Neovim** environment

I adapted the basic configuration to have **borders**, but especially to **configure the shortcut groups**

The complete `~/.config/nvim/lua/plugins/which-key.lua` file

``` lua
return {
    {
        "folke/which-key.nvim",
        event = "VeryLazy",
        config = function(_, _)
            vim.o.timeout = true
            vim.o.timeoutlen = 500

            local wk = require("which-key")
            wk.setup({
                plugins = {
                    spelling = {
                        enabled = true
                    }
                },
                window = {
                    border = "single",
                },
            })
            local keymaps = {
                mode = { "n", "v" },
                f = { name = ">Find" },
                g = { name = ">Git" },
                v = { name = ">Code" },
                r = { name = ">Refactor" },
                x = { name = ">Trouble" },
            }
            wk.register(keymaps, { prefix = "<leader>" })
        end,
    },
}
```

## Treesitter

[**Treesitter**](https://github.com/nvim-treesitter/nvim-treesitter) which will allow first of all to have an even **more precise syntactic coloring**

But the great strength of **Treesitter** is that it transforms a file into an abbreviated form that can be browsed with the help of **query**  
This is not going to be useful for the user but rather for other plugins that can use the tree provided by **Treesitter** to provide services.

I do anything strange it's a standard configuration of **Treesitter** with Lazy

The complete `~/.config/nvim/lua/plugins/treesitter.lua`

``` lua
return {
    {
        "nvim-treesitter/nvim-treesitter",
        version = false, -- last release is way too old and doesn't work on Windows
        build = ":TSUpdate",
        ---@type TSConfig
        opts = {
            highlight = { enable = true },
            indent = { enable = true },
            context_commentstring = { enable = true, enable_autocmd = false },
            ensure_installed = {
                "bash",
                "c",
                "help",
                "html",
                "javascript",
                "json",
                "lua",
                "markdown",
                "markdown_inline",
                "python",
                "query",
                "regex",
                "ruby",
                "rust",
                "tsx",
                "typescript",
                "vim",
                "yaml",
            },
            auto_install = true,
        },
        ---@param opts TSConfig
        config = function(_, opts)
            require("nvim-treesitter.configs").setup(opts)
        end,
    }
}
```

## LSP Zero

[**LSP Zero**](https://github.com/VonHeikemen/lsp-zero.nvim) The purpose of this plugin is to bundle all the "boilerplate code" necessary to have:  
- [nvim-cmp](https://github.com/hrsh7th/nvim-cmp) (A popular autocompletion plugin)
- [nvim-lspconfig](https://github.com/neovim/nvim-lspconfig)
- [mason.nvim](https://github.com/williamboman/mason.nvim) (Install language servers from inside Neovim)

I use the basic configuration provided by the plugin documentation, I added some changes: 
- No automatic completion popup, I like to be able to decide when I want to activate the completion
- Added shortcuts to make it easier to navigate in the code with the LSP

The complete `~/.config/nvim/lua/plugins/lsp-zero.lua` file

``` lua
return {
    {
        'VonHeikemen/lsp-zero.nvim',
        branch = 'v1.x',
        dependencies = {
            -- LSP Support
            { 'neovim/nvim-lspconfig' }, -- Required
            { 'williamboman/mason.nvim' }, -- Optional
            { 'williamboman/mason-lspconfig.nvim' }, -- Optional

            -- Autocompletion
            { 'hrsh7th/nvim-cmp' }, -- Required
            { 'hrsh7th/cmp-nvim-lsp' }, -- Required
            { 'hrsh7th/cmp-buffer' }, -- Optional
            { 'hrsh7th/cmp-path' }, -- Optional
            { 'saadparwaiz1/cmp_luasnip' }, -- Optional
            { 'hrsh7th/cmp-nvim-lua' }, -- Optional

            -- Snippets
            { 'L3MON4D3/LuaSnip' }, -- Required
            { 'rafamadriz/friendly-snippets' }, -- Optional
        },
        init = function()
            local lsp = require('lsp-zero').preset({
                name = 'minimal',
                set_lsp_keymaps = true,
                manage_nvim_cmp = true,
                suggest_lsp_servers = false,
            })

            -- (Optional) Configure lua language server for neovim
            lsp.nvim_workspace()

            local cmp = require('cmp')
            cmp.setup({
                completion = {
                    autocomplete = false,  -- No toggle auto the completion menu
                },
            })

            local cmp_select = { behavior = cmp.SelectBehavior.Select }
            local cmp_mappings = lsp.defaults.cmp_mappings({
                ['<C-p>'] = cmp.mapping.select_prev_item(cmp_select),
                ['<C-n>'] = cmp.mapping.select_next_item(cmp_select),
                ['<C-y>'] = cmp.mapping.confirm({ select = true }),
                ['<cr>'] = cmp.mapping.confirm({ select = true }),
                ['<C-Space>'] = cmp.mapping.complete(),
            })

            lsp.setup_nvim_cmp({
                mapping = cmp_mappings
            })

            lsp.on_attach(function(_, bufnr)
                local opts = { buffer = bufnr, remap = false }
                vim.keymap.set("n", "<leader>F", function() vim.lsp.buf.format() end, { desc = "Format buffer" })

                vim.keymap.set('n', 'gd', function() vim.lsp.buf.definition() end, opts)
                vim.keymap.set('n', 'K', function() vim.lsp.buf.hover() end, opts)
                vim.keymap.set('n', '<leader>vs', function() vim.lsp.buf.workspace_symbol() end, opts)
                vim.keymap.set('n', '<leader>vd', function() vim.lsp.buf.diagnostic.open_float() end, opts)
                vim.keymap.set('n', '[d', function() vim.lsp.buf.diagnostic.goto_next() end, opts)
                vim.keymap.set('n', ']d', function() vim.lsp.buf.diagnostic.goto_prev() end, opts)
                vim.keymap.set('n', '<leader>vca', function() vim.lsp.buf.code_action() end, opts)
                vim.keymap.set('n', '<leader>vrr', function() vim.lsp.buf.references() end, opts)
                vim.keymap.set('n', '<leader>vrn', function() vim.lsp.buf.rename() end, opts)
                vim.keymap.set('i', '<C-h>', function() vim.lsp.buf.signature_help() end, opts)
            end)

            lsp.setup()
        end
    }
}
```

Now if you need a new LSP, DAP, Linter or formatter use `:Mason`
