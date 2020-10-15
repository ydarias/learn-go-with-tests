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

## Integers

Using examples you can improve your documentation, like:

```go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```

Comments will be automatically added to the Go Doc.