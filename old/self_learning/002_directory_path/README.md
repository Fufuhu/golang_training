
# 実行ファイルのパスを取得

例えば実行ファイルから相対パスで設定ファイルの配置先などが
定まっている場合には実行ファイルのパスが必要となる。

```bash
$ go build main.go
$ ./main
フルパス:/home/fujiwara/repositories/golang_training/self_learning/002_directory_path/main
ディレクトリ:/home/fujiwara/repositories/golang_training/self_learning/002_directory_path
```


```go
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	directory := filepath.Dir(exe)
```

# 参考情報

https://qiita.com/summer/items/2674a178ffe6acb895fb