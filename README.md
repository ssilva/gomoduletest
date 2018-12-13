# Go Module Test
1. Have the following in `gomoduletest-cmd/main.go`:
    ```go
    package main

    import (
      "fmt"
      "github.com/ssilva/gomoduletest"
    )

    func main() {
      fmt.Println(gomoduletest.Greet(""))
      fmt.Println(gomoduletest.Greet("John"))
    }
    ```
1. Enable modules.
    ```
    $ cd gomoduletest-cmd
    $ go mod init cmd
    go: creating new go.mod: module cmd
    ```
1. Run it
    ```
    $ go run main.go
    Hello, World!
    Hello, John!
    ```
1. Get a specific version of the module
    ```
    $ go get github.com/ssilva/gomoduletest@v1.0.0
    $ go run main.go
    Hello, !
    Hello, John!
    ```

Source: https://roberto.selbach.ca/intro-to-go-modules/
