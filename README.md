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

The code:

```go
func SumAll(numbersCollections ...[]int) []int {
	lengthOfNumbers := len(numbersCollections)
	sums := make([]int, lengthOfNumbers)

	for i, collection := range numbersCollections {
		sums[i] = Sum(collection)
	}

	return sums
}
```

is equivalent to:

```go
func SumAll(numbersCollections ...[]int) []int {
	var sums []int

	for _, collection := range numbersCollections {
		sums = append(sums, Sum(collection))
	}

	return sums
}
```

The function `reflect.DeepEqual` is not "type safe", so you need to take care when you use it. Wrapping it inside another function will make it typesafe.

```go
checkSums := func(t *testing.T, got, expected []int) {
    t.Helper()
    if !reflect.DeepEqual(got, expected) {
        t.Errorf("got %v want %v", got, expected)
    }
}
```