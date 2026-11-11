<<<<<<< HEAD
-- Set various settings
vim.o.number = true
vim.o.wrap = true
=======
vim.o.number = true
vim.o.wrap = false
>>>>>>> 1d7973c3bab1bcfa1df74e6dc2dab87acb60e406
vim.o.tabstop = 2
vim.o.shiftwidth = 2
vim.o.smartcase = true
vim.o.ignorecase = true
vim.o.hlsearch = false
vim.o.signcolumn = 'yes'
<<<<<<< HEAD
vim.o.relativenumber = true

-- Global copy + paste
vim.keymap.set({'n', 'x'}, 'gy', '"+y', {desc = 'Copy to clipboard'})
vim.keymap.set({'n', 'x'}, 'gp', '"+p', {desc = 'Paste clipboard text'})

-- Use space as leader key
vim.g.mapleader = vim.keycode('<Space>')

-- Simple keybinds for write + close
=======

vim.keymap.set({'n', 'x'}, 'gy', '"+y', {desc = 'Copy to clipboard'})
vim.keymap.set({'n', 'x'}, 'gp', '"+p', {desc = 'Paste clipboard text'})

vim.g.mapleader = vim.keycode('<Space>')

>>>>>>> 1d7973c3bab1bcfa1df74e6dc2dab87acb60e406
vim.keymap.set('n', '<leader>w', '<cmd>write<cr>', {desc = 'Save file'})
vim.keymap.set('n', '<leader>q', '<cmd>quitall<cr>', {desc = 'Exit vim'})

-- Disable arrow keys
vim.keymap.set('n', '<Up>', '<Nop>', {desc = 'Disable'})
vim.keymap.set('n', '<Down>', '<Nop>', {desc = 'Disable'})
vim.keymap.set('n', '<Left>', '<Nop>', {desc = 'Disable'})
vim.keymap.set('n', '<Right>', '<Nop>', {desc = 'Disable'})

<<<<<<< HEAD
-- Disable Ctrl + I (it is annoying because I keep misclicking!)
vim.keymap.set({'n', 'v', 's'}, '<C-i>', '<Nop>', { noremap = true, silent = true })

-- Install/find minideps
=======
>>>>>>> 1d7973c3bab1bcfa1df74e6dc2dab87acb60e406
local ok, MiniDeps = pcall(require, 'mini.deps')
if not ok then 
  vim.notify('[WARN] mini.deps module not found', vim.log.levels.WARN)
	return
end

MiniDeps.setup({})
<<<<<<< HEAD

-- Set up color scheme
=======
MiniDeps.add('neovim/nvim-lspconfig')
>>>>>>> 1d7973c3bab1bcfa1df74e6dc2dab87acb60e406
MiniDeps.add({
  source = 'sainnhe/gruvbox-material',
})

<<<<<<< HEAD
vim.g.gruvbox_material_background = 'hard'
vim.cmd.colorscheme('gruvbox-material')

-- A bunch of BS with mini?
require('mini.files').setup({})
MiniDeps.add('neovim/nvim-lspconfig')
vim.keymap.set('n', '<leader>e', '<cmd>lua MiniFiles.open()<cr>', {desc = 'File explorer'})
require('mini.icons').setup({style = 'glyph'})
require('mini.pick').setup()
require('mini.extra').setup()
require('mini.basics').setup()
vim.keymap.set('n', '<leader><space>', '<cmd>Pick buffers<cr>', {desc = 'Search open files'})
vim.keymap.set('n', '<leader>ff', '<cmd>Pick files<cr>', {desc = 'Search all files'})
vim.keymap.set('n', '<leader>fh', '<cmd>Pick help<cr>', {desc = 'Search help tags'})
vim.keymap.set("n", "<leader>sg", function()
    require("mini.pick").builtin.grep_live()
end)
vim.keymap.set("n", "<leader>d", function()
    require("mini.extra").pickers.diagnostic()
end)

require('mini.snippets').setup({})
require('mini.completion').setup({})
vim.lsp.enable({'gopls', 'tinymist', 'jedi_language_server'})

-- Auto-lint go
MiniDeps.add({
  source = 'mfussenegger/nvim-lint',
})

local lint = require("lint")
lint.linters_by_ft = {
	go = { "golangcilint" }
}

vim.api.nvim_create_autocmd("BufWritePost", {
	pattern = "*.go",
	callback = function()
		lint.try_lint()
	end,
})

-- Auto-format go
local go_group = vim.api.nvim_create_augroup("GoAutoFormat", { clear = true })

vim.api.nvim_create_autocmd("BufWritePre", {
		group = go_group,
    pattern = "*.go",
    callback = function(args)
				vim.lsp.buf.code_action({
					context = { only = { "source.organizeImports" } },
					apply = true,
				})

        vim.lsp.buf.format({ 
					bufnr = args.buf,
					async = false,
				})
    end,
})

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

-- Enable 'Rename symbol'
vim.keymap.set('n', '<leader>rn', function()
	vim.ui.input({prompt = 'New name: ', default = '' }, function(input)
		if not input or input == '' then return end
		vim.lsp.buf.rename(input)

		vim.defer_fn(function()
			vim.cmd("wall")
		end, 100)
	end)
end, { desc =  "Rename symbol" })

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

-- SuperCollider in nvim
vim.pack.add({'https://github.com/davidgranstrom/scnvim'})
require('scnvim').setup{
    ensure_installed = true
}

local scnvim = require 'scnvim'
local map = scnvim.map
local map_expr = scnvim.map_expr

scnvim.setup({
  keymaps = {
    ['<M-e>'] = map('editor.send_line', {'i', 'n'}),
    ['<C-e>'] = {
      map('editor.send_block', {'i', 'n'}),
      map('editor.send_selection', 'x'),
    },
    -- ['<CR>'] = map('postwin.toggle'),
    ['<M-CR>'] = map('postwin.toggle', 'i'),
    ['<M-L>'] = map('postwin.clear', {'n', 'i'}),
    ['<C-k>'] = map('signature.show', {'n', 'i'}),
    ['<F12>'] = map('sclang.hard_stop', {'n', 'x', 'i'}),
    ['<leader>st'] = map('sclang.start'),
    ['<leader>sk'] = map('sclang.recompile'),
    ['<F1>'] = map_expr('s.boot'),
    ['<F2>'] = map_expr('s.meter'),
  },
  editor = {
    highlight = {
      color = 'IncSearch',
    },
  },
  postwin = {
    float = {
      enabled = true,
    },
  },
})

-- Typst preview in neovim
vim.pack.add({ 'https://github.com/chomosuke/typst-preview.nvim' })
require('typst-preview').setup()


=======
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
>>>>>>> 1d7973c3bab1bcfa1df74e6dc2dab87acb60e406
