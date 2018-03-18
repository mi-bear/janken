# janken season3 [wip]
## stringerを使おう
### stringerとは
定数を定数名の文字列に変換する関数を生成するコマンド

```
stringer -type={定数型名} {file|directory}
```

### Usage

```
$ go get golang.org/x/tools/cmd/stringer

$ stringer -type Result result.go
$ stringer -type Hand hand.go
```

## go generateについて
goが標準で持っているコマンドで, コードの自動生成を行えます。
`stringer` もそのひとつ。
