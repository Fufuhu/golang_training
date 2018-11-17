# yamlファイルの読み込み

```bash
$ go get gopkg.in/yaml.v2
```


- **`file.yaml(読み込み対象のyamlファイル)`**
```yaml
users:
  - name: John
    description: HogeHoge
    age: 24
  - name: Smith
    description: PiyoPiyo
    age: 30
```

```go
type Data struct {
	Users []User `yaml:"users"`
}

type User struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Age         int    `yaml:"age"`
}
```

ファイル読み込みから、yamlとしてparseするまで。

```go
	buf, err := ioutil.ReadFile("./file.yaml")

	if err != nil {
		panic(err)
	}

	var d Data
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}
	fmt.Printf("d: %v\n", d)
```

```bash
$ go run sample.go
d: {[{John HogeHoge 24} {Smith PiyoPiyo 30}]}
{John HogeHoge 24}
{Smith PiyoPiyo 30}
```


# 参考情報

https://github.com/go-yaml/yaml

http://tweeeety.hateblo.jp/entry/2017/06/04/231043