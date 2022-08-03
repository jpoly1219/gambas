# Introduction

A `Series` is one of the fundamental types used in `gambas`. It represents a column of data. 

The `Series` type is defined as per below:

```go
type Series struct {
	data  []interface{}
	index IndexData
	name  string
	dtype string
}
```

- `data` holds the column data.
- `index` holds the index of the `Series` object.
- `name` is the name of the `Series` object.
- `dtype` is the data type of the data in the `Series` object.

The fields are private, but `gambas` provides accessors to get these fields.

```go
func (s Series) Data() []interface{}

func (s Series) Index() IndexData

func (s Series) Name() string

func (s Series) Dtype() string
```

`Series`'s make up a `DataFrame`.