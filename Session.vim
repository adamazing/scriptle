let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/Documents/myprojects/scriptle
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
set shortmess=aoO
badd +9 README.md
badd +3 ~/Documents/myprojects/scriptle/go.mod
badd +7 ~/Documents/myprojects/scriptle/main.go
badd +43 ~/Documents/myprojects/scriptle/cmd/root.go
badd +1 ~/Documents/myprojects/scriptle/.gitignore
badd +18 ~/Documents/myprojects/scriptle/cmd/version.go
badd +1 ~/Documents/myprojects/scriptle/LICENSE
badd +476 ~/Documents/myprojects/scriptle/go.sum
badd +52 ~/Documents/myprojects/scriptle/pkg/scriptle/scriptle.go
badd +13 ~/Documents/myprojects/scriptle/pkg/model/model.go
badd +4 ~/Documents/myprojects/scriptle/pkg/utility/utility.go
argglobal
%argdel
$argadd README.md
edit ~/Documents/myprojects/scriptle/cmd/root.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 111 + 83) / 166)
exe 'vert 2resize ' . ((&columns * 54 + 83) / 166)
argglobal
balt ~/Documents/myprojects/scriptle/pkg/scriptle/scriptle.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 43 - ((42 * winheight(0) + 29) / 58)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 43
normal! 038|
wincmd w
argglobal
if bufexists(fnamemodify("~/Documents/myprojects/scriptle/pkg/scriptle/scriptle.go", ":p")) | buffer ~/Documents/myprojects/scriptle/pkg/scriptle/scriptle.go | else | edit ~/Documents/myprojects/scriptle/pkg/scriptle/scriptle.go | endif
if &buftype ==# 'terminal'
  silent file ~/Documents/myprojects/scriptle/pkg/scriptle/scriptle.go
endif
balt ~/Documents/myprojects/scriptle/pkg/utility/utility.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 57 - ((48 * winheight(0) + 29) / 58)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 57
normal! 024|
wincmd w
exe 'vert 1resize ' . ((&columns * 111 + 83) / 166)
exe 'vert 2resize ' . ((&columns * 54 + 83) / 166)
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
nohlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
