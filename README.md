# learn-go-with-tests

Complementary code related to the page with the same name (Lean Go with tests)

## Hello world

Even if the file is called `hello.go` it is necessary to set the package name as main. Calling the package `hello` will return an error `cannot run non-main package`.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```
