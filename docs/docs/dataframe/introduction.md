# Introduction

A `DataFrame` is one of the fundamental types used in `gambas`. It represents a tabular 2D data. `DataFrame` is a column-based data, meaning that it is made of several `Series`'s.

The `DataFrame` type is defined as per below:

```go
type DataFrame struct {
	series  []Series
	index   IndexData
	columns []string
}

```

- `series` holds the `Series` objects that make up the `DataFrame` object.
- `index` holds the index of the `DataFrame` object.
- `columns` is the names of columns in the `DataFrame` object.

The fields are private, but `gambas` provides accessors to get these fields.

```go
func (df DataFrame) Series() []Series {
	return df.series
}

func (df DataFrame) Index() IndexData {
	return df.index
}
func (df DataFrame) Columns() []string {
	return df.columns
}
```