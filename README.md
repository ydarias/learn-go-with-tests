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

## Iteration

Benchmark tests are a first class citizen in Go.

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```

To run benchmarks the command is `go test -bench=.`.

## Arrays and slices

An array has the size coded into the type, so `[3]int{1, 2, 3}` is an array, but `[]int{1, 2, 3, 4` is a slice.

Test coverage can be checked with `go test -cover`.