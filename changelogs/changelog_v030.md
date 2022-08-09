# Changelog for v0.3.0

## New features:
- We have a new documentation page! I thought that pkg.go.dev page left a lot to be desired, so I made a separate documentation page for a better experience.
- `Series`, `DataFrame`, `IndexData`, and `Index` types now have a getter method that returns its private field values.
```go
type Series struct {
	data  []interface{}
	index IndexData
	name  string
	dtype string
}

func (s Series) Data() []interface{} {
	return s.data
}

func (s Series) Index() IndexData {
	return s.index
}

func (s Series) Name() string {
	return s.name
}

func (s Series) Dtype() string {
	return s.dtype
}
```
```go
type DataFrame struct {
	series  []Series
	index   IndexData
	columns []string
}

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
```go
type IndexData struct {
	index []Index
	names []string
}

func (id IndexData) Index() []Index {
	return id.index
}

func (id IndexData) Names() []string {
	return id.names
}
```
```go
type Index struct {
	id    int
	value []interface{}
}

func (i Index) Id() int {
	return i.id
}

func (i Index) Value() []interface{} {
	return i.value
}
```
- `df.LocCol` is a method that returns a column as a `Series` object. It's different from the preexisting `LocCols` which returns a `DataFrame` object.
```go
func (df *DataFrame) LocCol(col string) (Series, error)
```


## Changes
- `Describe` now returns a `[]StatsResult` instead of `[]float64`.
- `SortByGivenIndex` now accepts `withId` field for internal uses. For normal use, just pass `false`.

## Bug fixes
- `Fit` no longer bugs out while reading certain CSV files.