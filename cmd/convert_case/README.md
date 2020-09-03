Z80 Exerciserのアセンブラソース内のケースデータを
Goの `"internal/zex".Case` に変換・出力するツール

入力は STDIN で出力は STDOUT

接頭辞は `Doc` 固定なので出力後に置き換える

使用例:

```console
$ go run ./cmd/convert_case < _z80/zexdoc.asm > internal/zex/doc.go

$ go run ./cmd/convert_case < _z80/zexall.asm | sed -e 's/Doc/All/g' > internal/zex/all.go
```
