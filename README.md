# Advent of Code

[Advent of Code](https://adventofcode.com/) is an annual set of Christmas-themed computer programming challenges.

The problems are solved here using Go.

## How to Run

Run the code that solves the problem(s) with these commands.
Solves and prints the problem year and day, part 1 and part 2 answers, and some timing information.

```shell
# all years, all days
go run main.go
# specific year, all days
go run main.go 2025
# specific year, specific day
go run main.go 2025 01
# alternatively, build the executable
go build
./advent
```

All days have tests for the sample and personal inputs with the correct answers.
There is a benchmark for more analysis.

```shell
# tests
go test ./...
# benchmarks
go test -bench=. ./...
# benchmark a specific day
go test -bench=. ./2025/day02/
```