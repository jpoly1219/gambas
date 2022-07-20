# gambas
---

`gambas` is a data analysis package for Go that provides an intuitive way to manipulate tabular data. The project is inspired by the famous Python library `pandas`.

## Installation
---
```
$ go get -u github.com/jpoly1219/gambas
```

## Documentation
---
The documentation can be found in our [pkg.go.dev page](https://pkg.go.dev/github.com/jpoly1219/gambas).

## Project Goals
---
- Provide basic features from the pandas tutorial.
  - [x] Providing `Series` and `DataFrame` data types
  - [x] Reading and writing tabular data
    - [x] Reading CSV files
    - [x] Writing to CSV files
    - [x] Reading Excel files
    - [x] Writing to Excel files
    - [x] Reading JSON files
    - [x] Writing to JSON files
  - [x] Selecting a subset of data
    - [x] At, IAt
    - [x] Loc, ILoc
  - [ ] Plotting
    - [x] Set
    - [x] Using
    - [x] Trendline (fit)
    - [x] Statistics
    - [ ] Categorical count
  - [x] Creating new columns derived from existing columns
    - [x] Creating new columns
    - [x] Applying operations to the new column
    - [x] Renaming columns
  - [x] Calculating summary statistics
    - [x] Mean, median, standard deviation
    - [x] Min, max, quartiles
    - [x] Count, describe
  - [x] Reshaping the layout of tables
    - [x] Sorting by index
    - [x] Sorting by values
    - [x] Sorting by given index
    - [x] Groupby
    - [x] Pivot (long to wide format)
    - [x] PivotTable (long to wide format)
    - [x] Melt (wide to long format)
  - [x] Combining data from multiple tables
    - [x] Concatenate
    - [x] Merge
  - [ ] Handling time series data
    - [ ] Timestamp type
    - [ ] Timestamp type methods
    - [ ] ToDatetime
  - [ ] Manipulating textual data
  - [x] Multiindex
- [x] pkg.go.dev page
- [ ] Documentation
- [ ] Project website
- [ ] Project logo

## Philosophy
---
`gambas` was created to serve the needs of Go developers who wanted a robust data analysis package. `pandas` is an amazing tool, and is considered the industry standard when it comes to data organization and manipulation.

We didn't have a solid alternative in the Go realm. According to the [Go Developer Survey 2021 Results](https://go.dev/blog/survey2021-results), **missing critical libraries** were one of **the most common barriers** to using Go. You may have used Go for some time now, but you might've missed some of the libraries you used when you were using Python. `gambas` aims to scratch that itch. You will be able to tap into the superpowers of `pandas` while using your favorite language Go.

Go is a very attractive language with a very loyal userbase. It provides a pleasant developer experience with its simple syntax and strong typing. However, Go currently tends to be skewed towards developing services. 49% of projects written in Go are API/RPC services, and another 10% are for web services. The ultimate goal for `gambas` is to allow the Go programming language to be a major player in the data analysis field.