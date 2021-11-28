" show current working directory structure, so we could have paths to jump to.
function! Structure()
    execute "vert 100 new"
    execute "r ! structure -ignore=\"\\\\.git\" -tree"
    execute "w! .seekbuf"
    execute "e .seekbuf"
endfunction
nnoremap <S-w> :call Structure()<CR>
