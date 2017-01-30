[![wercker status](https://app.wercker.com/status/5f54f03a8f1c96de4ed20dcf6ad055ba/s/master "wercker status")](https://app.wercker.com/project/byKey/5f54f03a8f1c96de4ed20dcf6ad055ba)
[![GoDoc](https://godoc.org/github.com/kwmt/combine?status.svg)](http://godoc.org/github.com/kwmt/combine) 

あるディレクトリ以下のファイルを一つのファイルにまとめる


### サンプル
```
combine.Combine("./dir", "all.csv")
```

### TODO
- [x] まとめるファイルがなかったら作成する
- [x] テスト
- [x] CI