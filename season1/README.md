# janken season1
## じゃんけんゲームを作ろう

1. `user`, `computer`から与えられた値を比較(計算)してじゃんけんの結果を出力してください。
1. コマンドラインからの入力は必要ないので, 変数に入力した値を利用してください。(後のステップでやります)
1. じゃんけんの手, 結果はそれぞれ, Enumで定義してください。

## GoのEnumについて
Go言語には列挙型がありません。  
そのかわり, `iota` を利用することで列挙型と同じような振る舞いを実装することができます。

### Example: 熊の種類を列挙型で定義したい

```
1: アメリカクロクマ
2: ホッキョクグマ
3: ツキノワグマ
4: マレーグマ
```

```go
type Bear int

const (
  AmericanBlackBear Bear = iota + 1
  PolarBear
  AsianBlackBear
  MalayanSunBear
)
```

`iota` は, 定数で利用される, 型ナシの連続する整数定数です。  
最初は `0` で初期化され, 1つずつインクリメントされます。2行目移行の `iota` は省略が可能です。


```go
type Bear int

const (
  _ Bear = iota
  AmericanBlackBear
  PolarBear
  AsianBlackBear
  MalayanSunBear
)
```
