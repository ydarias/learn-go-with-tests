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

## Structs, methods & interfaces

Go doesn't allow different signatures for the same method :-( 

Go's interface resolution is implicit, if the Circle or Rectangle contains a method Area() then it matches Shape interface.

Go allows table driven tests very easily:

```go
func TestArea(t *testing.T) {

    areaTests := []struct {
        shape Shape
        want  float64
    }{
        {Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},
    }

    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }

}
```

## Pointers & errors

Like C you need to specify at the arguments of a method when they are passed by value or by reference, in Go that is done with `*`.

In Go the errors are unchecked so you need a linter to detect the places where errors are not managed.