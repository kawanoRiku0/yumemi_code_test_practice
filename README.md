# 概要
https://www.yumemi.co.jp/serverside_recruit

上記のゆめみさんのコーディングテストを200分の制限で解いてみた

# 利用方法
```
go run main.go ../csv/playlog.go // at src
```
OR
```src
go build main.go // at src
```
```
./main csv/playlog.go // コンパイルされたmainがあるディレクトリで
```
# コメント
・データソースやインターフェースが変わってもランキングを算出するシステムは転用できるようにしたかったため、ドメイン層が他のレイヤーに依存しないようにした。

# 懸念点
・仕様を勘違いしていて、ランキングの結果をcsvファイルで保存するようにしている。
→標準出力されるよりも汎用性が高いし、明らかに要件としてcsvファイルで保存することが追加されそうなので、まぁ良いかも

・時間がタイトでドメイン周りがリファクタできなかった。

