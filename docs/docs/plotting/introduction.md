# Introduction

`gambas` provides tools for plotting your data. `gambas` uses `Gnuplot` as its plotting backend, so you need to have installed `Gnuplot` to use the plotting functions. It is advised that you be familiar with basic `Gnuplot` syntax.

Download `Gnuplot` [here](http://www.gnuplot.info/download.html).

Most plotting functions rely on the `PlotData` object to plot. Here is the definition of `PlotData` type:

```go
type PlotData struct {
	Df *DataFrame
	Columns []string
	Function string
	Opts []GnuplotOpt
}
```

- `Df` is the `DataFrame` object you would like to plot.
- `Columns` are the columns in `Df` that you want to plot. Usually, it's a pair of columns `[xcol, ycol]`. If you want to create a bar graph or a histogram, you can add more columns.
- `Function` is an arbitrary function such as `sin(x)` or an equation of the line of best fit.]
- `Opts` are options such as `using` or `with`. `set` is passed in as an argument for other plotting functions.

If you want to plot an arbitrary function, leave `Df` and `Columns` as `nil`. Otherwise, populate `Df` and `Columns`, and leave `Function` as `""`.