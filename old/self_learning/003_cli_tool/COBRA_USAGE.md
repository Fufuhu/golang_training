# Cobraの使い方


## 最小限?のファイル構成

```text
/
├ main.go
└ cmd/
    ├ root.go
    └ version.go
```


**`main.go`**

```go
package main

import (
	"./cmd"
)

func main() {
	cmd.Execute()
}
```


**`cmd/root.go`**

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "cli-sample",
    Short: `(Short) Description of Long Version of this application`,
    Long:  `(Long) Description of Long Version of this application`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```


**`cmd/version.go`**
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
	Short: "(Short) Print the version number of this application",
	Long:  "(Long) Print the version number of this application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This Application version is v0.0.1")
	},
}
```