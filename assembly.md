# Whitespace assembly instructions

These are the names used by known Whitespace assembly dialects for
instructions, ranked by popularity.

- `push` (15), `mod.push`, `psh`, `push <address>`, `push <number>`, `stack push`
- `dup` (13), `copy`, `dupl`, `duplicate`, `mod.dupe`, `stack dup`
- `copy` (11), `copynth` (2), `mod.copy`, `pull`, `ref`, `stack copy`, `take`
- `swap` (14), `swp` (2), `mod.swap`, `stack swap`, `swa`, `swicth`, `xchg`
- `discard` (8), `pop` (5), `drop` (2), `disc`, `dsc`, `mod.pop`, `stack discard`
- `slide` (13), `mod.slde`, `skip`, `stack slide`
- `add` (15), `add <address>`, `add <address> <address>`, `add <address> <number>`, `add <number>`, `add <number> <address>`, `arith add`, `math.add`, `plus`
- `sub` (14), `arith sub`, `math.sub`, `minus`, `sub <address>`, `sub <address> <address>`, `sub <address> <number>`, `sub <number>`, `sub <number> <address>`, `subtract`
- `mul` (12), `mult` (2), `arith mul`, `math.mult`, `mul <address>`, `mul <address> <address>`, `mul <address> <number>`, `mul <number>`, `mul <number> <address>`, `multiply`, `times`
- `div` (14), `divide` (2), `arith div`, `div <address>`, `div <address> <address>`, `div <address> <number>`, `div <number>`, `div <number> <address>`, `math.div`
- `mod` (14), `modulo` (2), `arith mod`, `math.mod`, `mod <address>`, `mod <address> <address>`, `mod <address> <number>`, `mod <number>`, `mod <number> <address>`, `rem`
- `store` (13), `heap store`, `heap.stor`, `set`, `sto`, `stor`, `store <number>`, `store <number> <number>`
- `retrieve` (12), `get` (2), `heap retrieve`, `heap.ret`, `rcl`, `retr`, `retrieve <number>`
- `mark` (6), `label` (5), `<label>:` (3), `@<label>`, `@l<number>`, `flow label`, `flow.labl`, `lbl .<label>`, `lbl <label>`, `setlabel`
- `call` (14), `call .<label>`, `call_subroutine`, `flow call`, `flow.sub`, `gosub`, `jsr`
- `jump` (10), `jmp` (6), `b`, `flow jump`, `flow.jump`, `goto`, `j`, `jmp .<label>`
- `jz` (5), `jumpz` (4), `jumpzero` (3), `jez` (2), `bz`, `flow jz`, `flow.jmpz`, `jmp_if0`, `jpz .<label>`, `jzero`
- `jn` (3), `jumpn` (3), `jlz` (2), `jumpneg` (2), `bltz`, `flow jneg`, `flow.jmpn`, `jltz`, `jmp_neg`, `jneg`, `jpn .<label>`, `js`, `jumplz`, `jumpnegative`
- `ret` (9), `return` (4), `endofsubroutine`, `ends`, `endsub`, `flow return`, `flow.ret`
- `exit` (7), `end` (6), `endofprogram`, `endp`, `endprog`, `finish`, `flow end`, `flow.halt`
- `putchar` (3), `pchr` (2), `putc` (2), `io outchar`, `io.out`, `ochar`, `outc`, `outchar`, `outputchar`, `print_c`, `print_char`, `printc`, `write_char`, `writec`, `writechar`
- `putnum` (3), `pnum` (2), `putn` (2), `io outnumber`, `io.nout`, `onum`, `outn`, `outnum`, `outputnum`, `print_i`, `print_number`, `printi`, `write_num`, `writeint`, `writen`
- `readchar` (4), `getc` (2), `getchar` (2), `read_char` (2), `readc` (2), `ichar`, `ichr`, `inpc`, `io readchar`, `io.in`, `rchr`, `read_c`
- `readnum` (3), `getn` (2), `getnum` (2), `inum` (2), `inpn`, `io readnumber`, `io.nin`, `read_i`, `read_num`, `read_number`, `readi`, `readint`, `readn`, `rnum`
