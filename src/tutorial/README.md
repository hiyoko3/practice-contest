# Tutorial

[Go Tutorial](https://go-tour-jp.appspot.com/welcome/1)

## 初心者ミスあるある

Q. go run: cannot run non-main package
A. package名が`main`になっていない。

Q. 相対パスが読み込めないんだけど？？
A. goでは相対パスによる import を推奨していないらしいです。
やり方もわかりません。教えてください。

Q. `"math"`と`"math/rand"`が同時にimportできないんですけど？
A. エイリアスを使いましょう。

```go
import(
 "fmt
 r "math/rand"
 m "math"
)

func main() {
 fmt.Println(r.Intn(10))
 fmt.Println(m.sqrt(4))
}
```
