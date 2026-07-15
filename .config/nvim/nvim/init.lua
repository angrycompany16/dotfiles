-- Set various settings
vim.o.number = true
vim.o.wrap = true
vim.o.tabstop = 2
vim.o.shiftwidth = 2
vim.o.smartcase = true
vim.o.ignorecase = true
vim.o.hlsearch = false
vim.o.signcolumn = 'yes'
vim.o.relativenumber = true

-- Global copy + paste
vim.keymap.set({'n', 'x'}, 'gy', '"+y', {desc = 'Copy to clipboard'})
vim.keymap.set({'n', 'x'}, 'gp', '"+p', {desc = 'Paste clipboard text'})

-- Use space as leader key
vim.g.mapleader = vim.keycode('<Space>')

-- Simple keybinds for write + close
vim.keymap.set('n', '<leader>w', '<cmd>write<cr>', {desc = 'Save file'})
vim.keymap.set('n', '<leader>q', '<cmd>quitall<cr>', {desc = 'Exit vim'})

-- Disable arrow keys
vim.keymap.set('n', '<Up>', '<Nop>', {desc = 'Disable'})
vim.keymap.set('n', '<Down>', '<Nop>', {desc = 'Disable'})
vim.keymap.set('n', '<Left>', '<Nop>', {desc = 'Disable'})
vim.keymap.set('n', '<Right>', '<Nop>', {desc = 'Disable'})

-- Install/find minideps
local ok, MiniDeps = pcall(require, 'mini.deps')
if not ok then 
  vim.notify('[WARN] mini.deps module not found', vim.log.levels.WARN)
	return
end

MiniDeps.setup({})

-- Set up color scheme
MiniDeps.add({
  source = 'sainnhe/gruvbox-material',
})

vim.g.gruvbox_material_background = 'hard'
vim.cmd.colorscheme('gruvbox-material')

-- A bunch of BS with mini?
require('mini.files').setup({})
MiniDeps.add('neovim/nvim-lspconfig')
vim.keymap.set('n', '<leader>e', '<cmd>lua MiniFiles.open()<cr>', {desc = 'File explorer'})
require('mini.icons').setup({style = 'ascii'})
require('mini.pick').setup({})
vim.keymap.set('n', '<leader><space>', '<cmd>Pick buffers<cr>', {desc = 'Search open files'})
vim.keymap.set('n', '<leader>ff', '<cmd>Pick files<cr>', {desc = 'Search all files'})
vim.keymap.set('n', '<leader>fh', '<cmd>Pick help<cr>', {desc = 'Search help tags'})

require('mini.snippets').setup({})
require('mini.completion').setup({})
vim.lsp.enable({'gopls', 'tinymist'})

-- Super basic autosave
vim.api.nvim_create_autocmd({ "FocusLost", "BufLeave", "InsertLeave", "TextChanged" }, {
  pattern = "*",
  command = "silent! update",
})

-- Enter in normal mode to add newline below / above
vim.keymap.set('n', '<CR>', 'm`o<Esc>``')
vim.keymap.set('n', '<S-CR>', 'm`O<Esc>``')

-- Automatic bracket opening / closing
vim.pack.add({
	'https://github.com/windwp/nvim-autopairs.git',
})

-- require('nvim-autopairs').setup({})
-- 
-- local Rule = require("nvim-autopairs.rule")
-- local npairs = require('nvim-autopairs')
-- npairs.add_rule(Rule("$", "$", "typst"))

-- Enable 'go to definition'
vim.keymap.set('n', 'gd', vim.lsp.buf.definition, { desc = "Go to definition" })

-- Show error messages
vim.keymap.set('n', 'gl', vim.diagnostic.open_float, { desc = "Open diagnostic float" })

-- Buffer navigation
vim.keymap.set('n', '<tab>', '<cmd>bnext<CR>', {desc = 'Next buffer'})
vim.keymap.set('n', '<S-tab>', '<cmd>bprevious<CR>', {desc = 'Previous buffer'})

-- Enclose in brackets in visual mode
vim.keymap.set('v', '(', 'c(<ESC>pa)')
vim.keymap.set('v', '{', 'c{<ESC>pa}')
vim.keymap.set('v', '[', 'c[<ESC>pa]')
vim.keymap.set('v', "'", "c'<ESC>pa'")
vim.keymap.set('v', '"', 'c"<ESC>pa"')










