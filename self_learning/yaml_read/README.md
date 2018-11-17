
```bash
$ go get gopkg.in/yaml.v2
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

```yaml
users:
  - name: John
    description: HogeHoge
    age: 24
  - name: Smith
    description: PiyoPiyo
    age: 30
```