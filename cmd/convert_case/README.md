# convert\_case

Convert test case data in assembler source of Z80 Exerciser to `"internal/zes".Case` for Go.

Z80 Exerciserのアセンブラソース内のケースデータを
Goの `"internal/zex".Case` に変換・出力するツール

## How to use

```console
$ go run ./cmd/convert_case < _z80/zexdoc.asm > internal/zex/doc.go

$ go run ./cmd/convert_case < _z80/zexall.asm | sed -e 's/Doc/All/g' > internal/zex/all.go
```

## Notes

Read from STDIN, Write to STDOUT.

The data names in the test cases are prefixed with `Doc`.
For zexall, replace these with `All`.
