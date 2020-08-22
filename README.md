# koron-go/z80

[![GoDoc](https://godoc.org/github.com/koron-go/z80?status.svg)](https://godoc.org/github.com/koron-go/z80)
[![Actions/Go](https://github.com/koron-go/z80/workflows/Go/badge.svg)](https://github.com/koron-go/z80/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/z80)](https://goreportcard.com/report/github.com/koron-go/z80)

Z80 emulation in Go.

## Progress

Z80 instruction set exerciser, passed tests:

* [x] adc16   - `<adc,sbc> hl,<bc,de,hl,sp> (38,912 cycles)`
* [x] add16   - `add hl,<bc,de,hl,sp> (19,456 cycles)`
* [x] add16x  - `add ix,<bc,de,ix,sp> (19,456 cycles)`
* [x] add16y  - `add iy,<bc,de,iy,sp> (19,456 cycles)`
* [x] alu8i   - `aluop a,nn (28,672 cycles)`
* [ ] alu8r   - `aluop a,<b,c,d,e,h,l,(hl),a> (753,664 cycles)` CRC
* [ ] alu8rx  - `aluop a,<ixh,ixl,iyh,iyl> (376,832 cycles)` DECODE
    * ADD/ADC/SUB/SBC/AND/XOR/OR/CP IXH/IXL/IYH/IYL
* [ ] alu8x   - `aluop a,(<ix,iy>+1) (229,376 cycles)` CRC
* [x] bitx    - `bit n,(<ix,iy>+1) (2048 cycles)`
* [x] bitz80  - `bit n,<b,c,d,e,h,l,(hl),a> (49,152 cycles)`
* [x] cpd1    - `cpd<r> (1) (6144 cycles)`
* [x] cpi1    - `cpi<r> (1) (6144 cycles)`
* [x] daaop   - `<daa,cpl,scf,ccf> (65,536 cycles)`
* [x] inca    - `<inc,dec> a (3072 cycles)`
* [x] incb    - `<inc,dec> b (3072 cycles)`
* [x] incbc   - `<inc,dec> bc (1536 cycles)`
* [x] incc    - `<inc,dec> c (3072 cycles)`
* [x] incd    - `<inc,dec> d (3072 cycles)`
* [x] incde   - `<inc,dec> de (1536 cycles)`
* [x] ince    - `<inc,dec> e (3072 cycles)`
* [x] inch    - `<inc,dec> h (3072 cycles)`
* [x] inchl   - `<inc,dec> hl (1536 cycles)`
* [x] incix   - `<inc,dec> ix (1536 cycles)`
* [x] inciy   - `<inc,dec> iy (1536 cycles)`
* [x] incl    - `<inc,dec> l (3072 cycles)`
* [x] incm    - `<inc,dec> (hl) (3072 cycles)`
* [x] incsp   - `<inc,dec> sp (1536 cycles)`
* [x] incx    - `<inc,dec> (<ix,iy>+1) (6144 cycles)`
* [x] incxh   - `<inc,dec> ixh (3072 cycles)`
* [x] incxl   - `<inc,dec> ixl (3072 cycles)`
* [x] incyh   - `<inc,dec> iyh (3072 cycles)`
* [x] incyl   - `<inc,dec> iyl (3072 cycles)`
* [x] ld161   - `ld <bc,de>,(nnnn) (32 cycles)`
* [x] ld162   - `ld hl,(nnnn) (16 cycles)`
* [x] ld163   - `ld sp,(nnnn) (16 cycles)`
* [x] ld164   - `ld <ix,iy>,(nnnn) (32 cycles)`
* [x] ld165   - `ld (nnnn),<bc,de> (64 cycles)`
* [x] ld166   - `ld (nnnn),hl (16 cycles)`
* [x] ld167   - `ld (nnnn),sp (16 cycles)`
* [x] ld168   - `ld (nnnn),<ix,iy> (64 cycles)`
* [x] ld16im  - `ld <bc,de,hl,sp>,nnnn (64 cycles)`
* [x] ld16ix  - `ld <ix,iy>,nnnn (32 cycles)`
* [x] ld8bd   - `ld a,<(bc),(de)> (44 cycles)`
* [x] ld8im   - `ld <b,c,d,e,h,l,(hl),a>,nn (64 cycles)`
* [x] ld8imx  - `ld (<ix,iy>+1),nn (32 cycles)`
* [x] ld8ix1  - `ld <b,c,d,e>,(<ix,iy>+1) (512 cycles)`
* [x] ld8ix2  - `ld <h,l>,(<ix,iy>+1) (256 cycles)`
* [x] ld8ix3  - `ld a,(<ix,iy>+1) (128 cycles)`
* [x] ld8ixy  - `ld <ixh,ixl,iyh,iyl>,nn (32 cycles)`
* [x] ld8rr   - `ld <b,c,d,e,h,l,a>,<b,c,d,e,h,l,a> (3456 cycles)`
* [x] ld8rrx  - `ld <b,c,d,e,ixy,a>,<b,c,d,e,ixy,a> (6912 cycles)`
* [x] lda     - `ld a,(nnnn) / ld (nnnn),a (44 cycles)`
* [x] ldd1    - `ldd<r> (1) (44 cycles)`
* [x] ldd2    - `ldd<r> (2) (44 cycles)`
* [x] ldi1    - `ldi<r> (1) (44 cycles)`
* [x] ldi2    - `ldi<r> (2) (44 cycles)`
* [x] negop   - `neg (16,384 cycles)`
* [x] rldop   - `<rld,rrd> (7168 cycles)`
* [x] rot8080 - `<rlca,rrca,rla,rra> (6144 cycles)`
* [x] rotxy   - `shift/rotate (<ix,iy>+1) (416 cycles)`
* [x] rotz80  - `shift/rotate <b,c,d,e,h,l,(hl),a> (6784 cycles)`
* [x] srz80   - `<set,res> n,<b,c,d,e,h,l,(hl),a> (7936 cycles)`
* [x] srzx    - `<set,res> n,(<ix,iy>+1) (1792 cycles)`
* [x] st8ix1  - `ld (<ix,iy>+1),<b,c,d,e> (1024 cycles)`
* [x] st8ix2  - `ld (<ix,iy>+1),<h,l> (256 cycles)`
* [x] st8ix3  - `ld (<ix,iy>+1),a (64 cycles)`
* [x] stabd   - `ld (<bc,de>),a (96 cycles)`

## References

* [Zilog Z80 DAA実行結果](http://ver0.sakura.ne.jp/doc/daa.html)

    LZ8514(SHARPのZ80互換CPU)で実際に実行して得た結果

* [Z80 DAA 内部キャリーフラグは実在するのか?](https://uniabis.net/pico/msx/z80daa/)

* [8ビット CPU Z80命令セット](http://www.yamamo10.jp/yamamoto/comp/Z80/instructions/index.php)

* [zmac - Z-80 Macro Cross Assembler](http://48k.ca/zmac.html)

* [Z80 Instruction Exerciser](http://mdfs.net/Software/Z80/Exerciser/)

    * [zmac port](https://github.com/obiwanjacobi/Zim80/tree/master/Source/Code/Jacobi.Zim80.IntegrationTests/CpuZ80/Zexlax)
    * [another zmac port](https://github.com/DavidDiPaola/esp32_crimbus-lights-z80/blob/master/src/z80/roms/zexdoc.src)

* [HOME OF THE Z80 CPU](http://www.z80.info/)

    * [Z80 instruction set (tables)](http://clrhome.org/table/)

* [WebMSX's Z80 implementation](https://github.com/ppeccin/WebMSX/blob/master/src/main/msx/cpu/CPU.js)
