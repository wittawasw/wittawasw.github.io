---
layout: post
title: Switch back to vim again.
date: 2016-01-21T02:20:22+07:00
comments: true
categories: vim environment editor
cc: Switch back to vim again
keywords: "vim, environment, editor"
description: Unlike most of my programmer friends, my first editor in programming is vi and second is vim as a slightly powerful editor. (I don't know about plugins at that time). C# as a first programming language with mono as a compiler and vim as an editor. vim's like a best friend back then. But when I started working extensively on C#, I embrace VisualStudio as my new editor and tool to work on real application projects.

---

Unlike most of my programmer friends, my first editor in programming is **vi** and
second is **vim** as a slightly powerful editor. (I don't know about plugins at that
time). C# as a first programming language with mono as a compiler and vim as an editor.
vim's like a best friend back then. But when I started working extensively on C#,
I embrace VisualStudio as my new editor and tool to work on real application projects.

Some years after that I've started on the new path, Python and Ruby. That's when I went
to TextMate and eventually Sublime 2. Great experiences compare to big and slow
VisualStudio. But I still have to work with vim sometimes when I have to work on
server configurations.

Over the last few years, I've been working on a lot of server and deployment stuffs.
Swapping between vim and Sublime a lot in the process. Eventually, have some thoughts
that should I switch to using just vim to make thing easier.

Then the boss of my team just purpose that we should have a month of vim, a month that
we all have to rely on vim alone as a practice and also let others that normally do
everything in Sublime to try vim for certain period.

It's been a week since then. I've gained most of my speed in using Vim that I've
lost over years of Textmate and Sublime. One good thing that happened to me after I
switched back to vim is, My coding style have changed to the way that it will suit Vim
better than what I used to have in Sublime. Cleaner and mind a lot more when to use
spaces and newline. The markdown of this new post is also done using vim. Feel great
that it's worked with most of my works even with my blog posts.

As a reminder for myself and everyone else who stumble apon this page. Below is
my current **.vimrc** that I use in my OSX. I thought I will stick with this for a while
before I add anything new since this's sufficient for me now.

{% codeblock lang:bash .vimrc %}
set nu
set showcmd
set tabstop=2 shiftwidth=2 expandtab
set mouse=a
set ls=2
syntax on

autocmd BufWritePre * :%s/\s\+$//e
autocmd BufNewFile,BufRead Gemfile set filetype=ruby
autocmd BufNewFile,BufRead *.ru set filetype=ruby

let g:ctrlp_match_func = {'match' : 'matcher#cmatch' }

command -nargs=1 E execute('silent! !mkdir -p "$(dirname "<args>")"') <Bar> e <args>

nmap <silent> <leader>t :TestNearest<CR>
nmap <silent> <leader>T :TestFile<CR>
nmap <silent> <leader>a :TestSuite<CR>
nmap <silent> <leader>l :TestLast<CR>
nmap <silent> <leader>g :TestVisit<CR>


call plug#begin('~/.vim/plugged')
Plug 'tpope/vim-rails'
Plug 'pbrisbin/vim-mkdir'
Plug 'ntpeters/vim-better-whitespace'

Plug 'elixir-lang/vim-elixir'
Plug 'fatih/vim-go'
Plug 'ngmy/vim-rubocop'
Plug 'cakebaker/scss-syntax.vim'
Plug 'tpope/vim-haml'
Plug 'mxw/vim-jsx'
Plug 'ekalinin/Dockerfile.vim'
Plug 'rking/ag.vim'
Plug 'kien/ctrlp.vim'
Plug 'JazzCore/ctrlp-cmatcher'
Plug 'tomtom/tcomment_vim'
Plug 'airblade/vim-gitgutter'
Plug 'tpope/vim-fugitive'
Plug 'tpope/vim-bundler'
Plug 'janko-m/vim-test'

Plug 'kchmck/vim-coffee-script'
Plug 'isRuslan/vim-es6'

call plug#end()


{% endcodeblock %}
