
# Cobraって？

## 概要

シンプルなインタフェースで、協力で現代的なCLIをgoで作成するためのライブラリ

コブラによって以下を実現できる。

1. サブコマンドベースのCLIを容易に実現できる
2. POSIX準拠のフラグ(ショートロング両方のバージョンのフラグを含む)
3. ネストされたサブコマンド
4. グローバル、ローカルフラグそれぞれだけでなく、フラグのカスケーディングに対応
5. インテリジェントなサジェスチョン
6. 自動的なヘルプの生成とbashの自動補完をアプリケーションに提供
7. アプリケーションのmanページを自動生成
8. コマンドのエイリアスをコードを変更なしに実現
9. 独自のヘルプや使い方を定義できる柔軟性を提供
10. 12ファクターアプリケーションを作成するためにviperと密にインテグレーションできる

## コンセプト


Cobraは以下の3要素から構成される。

1. `Commands`
2. `Args`
3. `Flags`

最良のアプリケーションは使うときには、文章のように読むことができる。

`APPNAME VERB NOUN --ADJECTIVE`
または
`APPNAME COMMAND ARG --FLAG`


```bash
$ hugo server --port=1313
```

```bash
$ git clone
```

## インストール方法

```
$ go get -u github.com/spf13/cobra/cobra
```

```
import "github.com/spf13/cobra
```

## 始め方

```text
appName/
 ├main.go
 └cmd/
    ├ add.go
    ├ your.go
    ├ commands.go
    └ here.go

```

### Cobraジェネレータ

`main.go`と`rootCmd`ファイルを生成できる。

### Cobraライブラリの利用

手動でCobraを実装するには、`main.go`と`rootCmd`ファイルを実装する必要がある。
そこから追加でコマンドを実装する。

#### rootCmdの作成

Cobraは特別なコンストラクタを必要としない。
単純にコマンドを作成するだけである。


利用的には、これを`app/cmd/root.go`に配置する。

```go
var rootCmd = &cobra.Command{
  Use:   "hugo",
  Short: "Hugo is a very fast static site generator",
  Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
```

`init()`関数の中で、フラグを追加で定義する。


例えば`cmd/root.go`はこんな感じ。

```go
import (
  "fmt"
  "os"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

func init() {
  cobra.OnInitialize(initConfig)
  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
  rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
  rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
  rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
  rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
  viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
  viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
  viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
  viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
  viper.SetDefault("license", "apache")
}

func initConfig() {
  // Don't forget to read config either from cfgFile or from home directory!
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".cobra" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".cobra")
  }

  if err := viper.ReadInConfig(); err != nil {
    fmt.Println("Can't read config:", err)
    os.Exit(1)
  }
}
```

#### main.goを作成する

rootコマンドがある場合、main関数を実行する必要がある。
Executeはrootで実行されるべきである。

Cobra appでは、`main.go`は典型的に単純である。

```go
package main

import (
  "{pathToYourApp}/cmd"
)

func main() {
  cmd.Execute()
}
```

#### 追加のコマンドの作成


```go
package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
  },
}
```

### Flagを使う

フラグはアクションコマンドがどのように動作しているかを制御するための
修飾子を提供する。

#### フラグをコマンドに割り当てる

フラグは、別の場所で定義され利用されるので、
フラグが動作するための適切なスコープを割り当てるために外側で変数を定義する必要がある。

```go
var Verbose bool
var Source string
```

2種類のフラグが存在する。
1. 永続フラグ
2. ローカルフラグ

##### 永続フラグ

すべてのサブコマンドあいだで有効なフラグのこと。
利用するには、rootで永続フラグとしてフラグを割り得てる必要がある。

```go
rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
```

##### ローカルフラグ

特定のサブコマンド配下でのみ有効なフラグ

```go
rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
```

〜〜〜もろもろめんどいので省略〜〜〜

### ポジショナルかつカスタムなArguments

`Command`の`Args`フィールドで利用できるポジショナルなargumentsのバリデーション

- ビルトインされてるバリデータ
    - NoArgs
        - ポジショナルなargsがある場合はエラーを出力する
    - ArbitraryArgs
        - どんなものも受け入れまっせ
    - OnlyValidArgs
        - `Command`の`VaridArgs`フィールドに含まれたものしか受け付けませんぜ
    - MinimumNArgs(int)
        - 最低N個必要
    - MaximumArgs(int)
        - 最大N個必要
    - ExactArgs(int)
        - 正確にN個必要
    - ExactValidArgs(int)
        - 正確にN個必要
        - または、`Command`の`ValidArgs`に含まれないargsが存在しないこと
    - RangeArgs(min, max)
        - 指定した範囲内の個数でなければNG

### Usageメッセージ

#### 独自利用方法の定義

#### バージョンフラグ

### PreRun、PostRunフック

- PersistentPreRun
- PreRun
- Run
- PostRun
- PersistentPostRun

# 参考情報

https://github.com/spf13/cobra
https://www.slideshare.net/dcubeio/gogo-81497334
