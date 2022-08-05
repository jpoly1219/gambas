# Introduction

The `Series` and `DataFrame` objects rely on index to determine the order of their data. 

The main type of index in `gambas` is the `IndexData` type, defined as per below:

```go
type IndexData struct {
	index []Index
	names []string
}
```

- `index` holds the index values, which are of type `Index`.
- `names` are the names of each index column. For single index, `names` only has one element. For multiindex, `names` will have as many elements as necessary.

You can access the private fields with these accessors:

```go
func (id IndexData) Index() []Index {
	return id.index
}

func (id IndexData) Names() []string {
	return id.names
}
```

This is the definition of the `Index` type.

```go
type Index struct {
	id    int
	value []interface{}
}
```

- `id` is the ID of an index item.
- `value` is the value of an index item. For single index, `value` only has one element. For multiindex, `value` will have as many elements as necessary.

You can access the private fields with these accessors:

```go
func (i Index) Id() int {
	return i.id
}

func (i Index) Value() []interface{} {
	return i.value
}
```