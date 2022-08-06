# Introduction

Having a lot of data is awesome, but you only get to see the bigger picture once you start running calculations with them. `gambas` provides several statistics functions that you can use for your data.

Every statistics function are `StatsFunc`'s that accept a `[]interface{}` and returns a `StatsResult`.

```go
type StatsFunc func(dataset []interface{}) StatsResult

```

Here is the definition of `StatsResult`.

```go
type StatsResult struct {
    UsedFunc string
	Result   float64
	Err      error
}
```

`StatsResult` holds the results of calculation from a statistics function such as `Mean` or `Median`.
- `UsedFunc` denotes which function has been used.
- `Result` is the result of the calculation.
- `Err` holds any errors that have been returned during the calculation.