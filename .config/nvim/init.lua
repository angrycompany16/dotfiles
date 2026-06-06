vim.o.number = true
vim.o.wrap = false
vim.o.tabstop = 2
vim.o.shiftwidth = 2
vim.o.smartcase = true
vim.o.ignorecase = true
vim.o.hlsearch = false
vim.o.signcolumn = 'yes'

vim.keymap.set({'n', 'x'}, 'gy', '"+y', {desc = 'Copy to clipboard'})
vim.keymap.set({'n', 'x'}, 'gp', '"+p', {desc = 'Paste clipboard text'})

vim.g.mapleader = vim.keycode('<Space>')

vim.keymap.set('n', '<leader>w', '<cmd>write<cr>', {desc = 'Save file'})
vim.keymap.set('n', '<leader>q', '<cmd>quitall<cr>', {desc = 'Exit vim'})

-- Disable arrow keys
vim.keymap.set('n', '<Up>', '<Nop>', {desc = 'Disable'})
vim.keymap.set('n', '<Down>', '<Nop>', {desc = 'Disable'})
vim.keymap.set('n', '<Left>', '<Nop>', {desc = 'Disable'})
vim.keymap.set('n', '<Right>', '<Nop>', {desc = 'Disable'})

local ok, MiniDeps = pcall(require, 'mini.deps')
if not ok then 
  vim.notify('[WARN] mini.deps module not found', vim.log.levels.WARN)
	return
end

MiniDeps.setup({})
MiniDeps.add('neovim/nvim-lspconfig')
MiniDeps.add({
  source = 'sainnhe/gruvbox-material',
})

require('mini.files').setup({})
vim.keymap.set('n', '<leader>e', '<cmd>lua MiniFiles.open()<cr>', {desc = 'File explorer'})
require('mini.icons').setup({style = 'ascii'})
require('mini.pick').setup({})
vim.keymap.set('n', '<leader><space>', '<cmd>Pick buffers<cr>', {desc = 'Search open files'})
vim.keymap.set('n', '<leader>ff', '<cmd>Pick files<cr>', {desc = 'Search all files'})
vim.keymap.set('n', '<leader>fh', '<cmd>Pick help<cr>', {desc = 'Search help tags'})

require('mini.snippets').setup({})
require('mini.completion').setup({})
vim.lsp.enable({'gopls'})

vim.g.gruvbox_material_background = 'hard' -- optional
vim.cmd.colorscheme('gruvbox-material')
