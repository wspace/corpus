# Whitespace assembly mnemonics

<!-- Generated by tools/generate_assembly.jq; DO NOT EDIT. -->

These are the mnemonics used by known Whitespace assembly dialects for
instructions, ranked by popularity.

- `push` (100), `append` (2), `psh` (2), `push_num` (2), `<number>`, `<number_char>`, `pus`, `push_ch`, `push_number`, `push_number_stack_top`, `push_to_stack`, `pushchar`, `pushstack`, `setstacktop`
- `dup` (69), `duplicate` (18), `copy` (5), `dupe` (5), `doub` (3), `dupl` (3), `sdupli` (3), `^`, `copystacktop`, `duplicate_last`, `duplicate_stack_top`, `duplicate_top`, `duplicatestack`
- `copy` (53), `copynth` (3), `ref` (3), `scopy` (3), `copyn` (2), `dup_at` (2), "", `^<unsigned>`, `copy_n`, `copynthstack`, `dup2`, `duplicate_num`, `duplicate_value`, `dupn`, `dupnum`, `ncopy`, `nthcopytosetstacktop`, `pick`, `pull`, `take`
- `swap` (97), `swp` (4), `sswap` (3), `exch`, `exchange`, `swa`, `swap_last`, `swap_two_stack_items`, `swapstack`, `swapstacktopsec`, `swicth`, `xchg`
- `discard` (46), `pop` (39), `drop` (14), `sdiscard` (3), `disc` (2), `away`, `del`, `discard_stack_top`, `discard_top`, `discardstack`, `dsc`, `pop_number`, `remove`, `throwstacktop`
- `slide` (55), `sslide` (3), `pop_x` (2), `slid` (2), "", `<unsigned>slide`, `away`, `discard_value`, `discardnum`, `move`, `remove`, `removenthinstack`, `skip`, `slde`, `sliden`, `slidenstack`, `slideoff`
- `add` (98), `addition` (4), `+` (3), `addtion` (2), `plus` (2), `adding`, `infixplus`, `op+`
- `sub` (87), `subtract` (11), `subtraction` (5), `-` (3), `minus` (2), `subt` (2), `infixminus`, `op-`, `substraction`
- `mul` (75), `multiply` (10), `mult` (8), `multiplication` (6), `multi` (4), `*` (3), `times` (2), `infixtimes`, `multiple`, `op*`
- `div` (86), `divide` (9), `division` (7), `/` (3), `intdiv` (2), `infixdivide`, `integerdivision`, `op/`
- `mod` (88), `modulo` (15), `%` (2), `division_part`, `infixmodulo`, `rem`
- `store` (86), `set` (4), `stor` (4), `save` (3), `sto` (3), `put` (2), `addtoheap`, `heapwrite`, `push`, `push_heap`, `st`, `storeheap`, `valuetoadress`, `write`
- `retrieve` (61), `load` (16), `get` (8), `retr` (7), `retrive` (4), `fetch` (2), `pop` (2), `read` (2), `ret` (2), `getfromheap`, `heapread`, `ld`, `lod`, `pop_heap`, `rcl`, `reti`, `retreive`, `retri`, `retrieveheap`, `retrv`, `valuetostacktop`
- `label` (44), `mark` (28), `<label>:` (14), `def` (4), `<number>:` (3), `label_<number>:` (3), `lbl` (3), `mark_sub` (2), `marklabel` (2), `mrk` (2), `setlabel` (2), `%<number>:`, `<<number>>:`, `@<label>`, `addlabel`, `deflabel`, `defun`, `l<number>:`, `label <label>:`, `label#_####`, `label_<number>`, `labl`, `mark_location`, `marks`, `part`, `register`, `set_label`
- `call` (91), `callsubroutine` (4), `call_sub` (3), `call_subroutine` (2), `gosub` (2), `jsr` (2), `call_label`, `calls`, `callsrtn`, `callsub`, `cas`, `cll`, `sub`, `subroutine`
- `jump` (69), `jmp` (32), `goto` (5), `jumplabel` (3), `jp` (2), `unconditionaljump` (2), `b`, `j`, `jumps`
- `jz` (38), `jumpz` (12), `jumpzero` (12), `jmpz` (5), `jump_if_zero` (5), `jzero` (5), `jez` (3), `jump-zero` (3), `jmp_eq` (2), `jump_zero` (2), `jumplabelwhenzero` (2), `branchz`, `branchzs`, `brz`, `bz`, `bzero`, `equal`, `gotoiz`, `if stack==0 then goto`, `if0goto`, `ifstacktopiszerothenjump`, `ifzero`, `iz_jump`, `jeq`, `jmp_if0`, `jmz`, `jnil`, `jp0`, `jpz`, `jump if zero`, `jump_ifzero`, `jump_stack`, `jumpnull`, `jze`, `jzer`, `zero`
- `jn` (32), `jneg` (10), `jumpn` (10), `jumpneg` (6), `jumpnegative` (5), `jlz` (4), `jmpn` (4), `jump_if_neg` (4), `jump-neg` (3), `jmp_lt` (2), `jump_negative` (2), `jumplabelwhennegative` (2), `jumplz` (2), `less` (2), `bltz`, `bmi`, `bneg`, `branchltz`, `branchltzs`, `gotoin`, `if stack<0 then goto`, `ifisnegativegoto`, `ifnegative`, `ifstacktopisnegthenjump`, `in_jump`, `jlt`, `jltz`, `jmn`, `jmp_neg`, `jmpneg`, `jne`, `jpl0`, `jpn`, `js`, `jump if negative`, `jump_if_negative`, `jump_nega`, `jump_stack_bzero`, `jumpde`, `jumpnega`
- `ret` (49), `return` (31), `endsubroutine` (5), `end` (4), `exit_sub` (3), `endofsubroutine` (2), `ends` (2), `endsub` (2), `back`, `control_back`, `end subroutine`, `end_sub`, `endfunc`, `endsrtn`, `ens`, `finishsubroutine`, `leave`, `rts`, `subroutine_end`
- `end` (50), `exit` (37), `endprogram` (5), `halt` (4), `endprog` (3), `endofprogram` (2), `endp` (2), `finish` (2), `quit` (2), `terminate` (2), `die`, `end program`, `end_prog`, `endle`, `eof`, `finishprogram`, `hlt`
- `outchar` (14), `printc` (14), `putc` (10), `outc` (8), `putchar` (8), `printchar` (6), `ochr` (4), `out-char` (3), `output_char` (3), `outputchar` (3), `outputcharacter` (3), `print_char` (3), `writec` (3), `out` (2), `pchr` (2), `write_ch` (2), `write_char` (2), `char_out`, `charout`, `cout`, `o_chr`, `ochar`, `otc`, `out_char`, `outch`, `outcha`, `outcharacter`, `output`, `output character`, `output_character`, `outputc`, `pc`, `pchar`, `pop_char`, `print`, `print_c`, `printaschar`, `prtc`, `stacktopoutchar`, `wchar`, `wrc`, `write_character`, `writechar`, `wtc`
- `outnum` (13), `printi` (11), `putn` (8), `outn` (7), `putnum` (5), `output_number` (4), `outputnumber` (4), `printn` (4), `printnum` (4), `write_num` (4), `onum` (3), `out-num` (3), `outputnum` (3), `pnum` (3), `writen` (3), `oint` (2), `print_num` (2), `putint` (2), `iout`, `nout`, `num_out`, `numout`, `o_int`, `otn`, `out`, `out_n`, `out_num`, `outi`, `outint`, `outinteger`, `outnumber`, `output number`, `outputn`, `pn`, `pop_num`, `print_i`, `print_number`, `printasint`, `printint`, `printnumber`, `prtn`, `putdigit`, `puti`, `stacktopoutint`, `wnum`, `write_number`, `writeint`, `wrn`, `wtn`
- `readchar` (25), `readc` (19), `getc` (10), `read_char` (8), `getchar` (6), `ichr` (5), `inc` (5), `in-char` (3), `readcharacter` (3), `in` (2), `inchar` (2), `rchar` (2), `rdc` (2), `read_chr` (2), `char_in`, `charin`, `chr`, `cin`, `ichar`, `inch`, `incha`, `inpc`, `rc`, `rchr`, `read character`, `read_c`, `read_ch`, `read_character`, `readchartoadress`, `rec`, `redc`
- `readnum` (19), `readi` (11), `getn` (9), `readn` (9), `read_num` (8), `readnumber` (5), `inn` (4), `inum` (4), `read_number` (4), `getnum` (3), `in-num` (3), `innum` (3), `readint` (3), `rnum` (3), `getint` (2), `iint` (2), `rdn` (2), `geti`, `iin`, `in_n`, `ini`, `inint`, `inpn`, `int`, `nin`, `num_in`, `numin`, `read`, `read number`, `read_i`, `readinteger`, `readinttoadress`, `redn`, `ren`, `rn`
- `permr`, `shuffle`
- `debug_printstack` (3), `dumpstack`
- `debug_printheap` (3), `dumpheap`
- `trace`

## Extensions

- `wsa` (24)
- `asm` (5)
- `txt` (3)
- `wsm` (2)
- ""
- `as`
- `bs`
- `hws`
- `unws`
- `wil`
- `ws.rb`
- `wsasm`
- `wsf`
- `wss`

## Need documentation

- haskell/burghard-wsa
- haskell/helvm-wsa
- javascript/susisu
