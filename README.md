# Trimmed Mean

`trimmedmean` is a Go package for calculating symmetric and asymmetric trimmed means for integer and floating-point data.

A trimmed mean is calculated by sorting the observations, removing a proportion of observations from the lower and upper ends, and calculating the arithmetic mean of the remaining observations.

## Status

This package is under development as part of a graduate-level Go programming assignment.

Currently implemented:

- Symmetric trimming
- Asymmetric trimming
- Integer slices
- `float64` slices
- Input validation for trimming proportions
- Unit tests
- Protection against modification of the caller's input slice
- Rejection of non-finite observations
- Benchmarks for 100 integers and 100 floating-point values
- Execution-time and memory-allocation evidence

Profiling evidence and comparisons with R will be added as development continues.

## Requirements

- Go 1.27.1 or later

## Installation

From inside a Go module, run:

```bash
go get github.com/gummybeardream/trimmed-mean
```

Then import the package:

```go
import "github.com/gummybeardream/trimmed-mean"
```

The repository name contains a hyphen, but the Go package name used in code is `trimmedmean`.

## Symmetric trimming

Provide one trimming proportion to remove the same proportion from both ends of the sorted observations.

```go
package main

import (
	"fmt"
	"log"

	trimmedmean "github.com/gummybeardream/trimmed-mean"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result, err := trimmedmean.Mean(values, 0.10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
```

A trimming proportion of `0.10` removes 10% of the observations from each end. For this example, the package removes `1` and `10`, producing a trimmed mean of `5.5`.

## Asymmetric trimming

Provide two proportions to use different trimming amounts for the lower and upper ends:

```go
result, err := trimmedmean.Mean(values, 0.10, 0.20)
```

The first proportion is the lower trim, and the second is the upper trim. This example removes 10% from the lower end and 20% from the upper end.

## Trimming rules

All observations must be finite. The function returns an error if the input
contains `NaN`, positive infinity, or negative infinity. This validation applies
to the entire input slice, including observations that trimming might remove.

The number of observations removed from each end is calculated as:

```text
number of observations × trimming proportion
```

When this calculation is not a whole number, the fractional portion is discarded. For example:

```text
15 observations × 0.10 = 1.5
```

One observation is removed from each end. This behavior is intended to match the trimming-count rule used by the base R `mean` function.

Valid calls must provide one or two trimming proportions. Each proportion must be at least `0` and less than `1`, and the combined lower and upper proportions must be less than `1`.

## Supported data types

The current version supports:

- `[]int`
- `[]float64`
- Named types whose underlying type is `int` or `float64`

The returned mean has type `float64`.

## Testing

Run all unit tests:

```bash
go test ./...
```

Display each test and subtest:

```bash
go test -v ./...
```

Run Go's static analysis tool:

```bash
go vet ./...
```

Current tests cover:

- Symmetric integer trimming
- Symmetric floating-point trimming
- Asymmetric trimming
- Zero trimming
- Unsorted input
- Empty input
- Invalid trimming proportions
- Fractional trimming counts
- Preservation of the original input slice
- Rejection of `NaN` and infinite observations
- Duplicate observations

## Benchmarking

The package includes benchmarks using 100 integers and 100 floating-point
observations. Each benchmark measures the complete `Mean` operation, including
copying, sorting, trimming, and calculating the mean.

Run the benchmarks with:

```bash
go test -bench=. -benchmem
```

### Benchmark results

Five benchmark trials were run on an Apple M1 computer using Go 1.27.1 on
macOS with the ARM64 architecture.

| Benchmark | Average time | Memory | Allocations |
|---|---:|---:|---:|
| 100 integers | 868.4 ns/op | 896 B/op | 1 alloc/op |
| 100 floats | 1,070.2 ns/op | 896 B/op | 1 alloc/op |

Both benchmarks made one 896-byte allocation per operation. This allocation
comes from copying the input slice before sorting it, which prevents the
function from modifying the caller's data.

In this experiment, the floating-point benchmark took approximately 23% more
time than the integer benchmark. This result applies to this computer, dataset,
and implementation and should not be interpreted as proof that integer input
is always faster.

The unedited output from all five trials is available in
[`results/benchmarks.txt`](results/benchmarks.txt).

Run five benchmark trials with this command:

```bash
go test -bench=. -benchmem -count=5
```

## Known limitations

- Benchmarks have currently been run only on one Apple M1 computer using
  datasets of 100 observations.
- CPU and memory profiles have not yet been added.
- Validation against R using shared datasets of at least 100 integers and 100
  floating-point values has not yet been completed.
- Extremely large finite values could cause the running sum to overflow.

## GenAI Tools

Generative AI tools were used during development to:

- Break the assignment into smaller tasks
- Explain Go syntax and programming concepts
- Review the package design
- Develop and explain unit tests
- Review documentation

The package code and tests were reviewed and executed locally by the repository author. Conversation records or links will be added to the repository as required by the assignment.
