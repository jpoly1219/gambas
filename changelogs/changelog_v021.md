# Changelog for v0.2.1

## New features:
- More features for plotting!
  - `PlotData` type for storing plot data.
  ```go
  // A PlotData holds the data required for plotting.
  //
  // If you want to plot an arbitrary function, leave Df and Columns as nil.
  // Otherwise, populate Df and Columns, and leave Function as "".
  type PlotData struct {
    // Df is the DataFrame object you would like to plot.
    Df *DataFrame

    // Columns are the columns in Df that you want to plot. Usually, it's a pair of columns [xcol, ycol].
    // If you want to create a bar graph or a histogram, you can add more columns.
    Columns []string

    // Function is an arbitrary function such as sin(x) or an equation of the line of best fit.
    Function string

    // Opts are options such as `using` or `with`. `set` is passed in as an argument for other plotting functions.
    Opts []GnuplotOpt
  }
  ```
  - `PlotN` can be used to plot several graphs at once.
  - `Fit` can be used to fit an arbitrary function to your PlotData object.
  - Options added: `with`, `via`

## Changes
- `Plot` is no longer a struct method for DataFrame. Instead, it is a standalone function that accepts PlotData objects.

## Bug fixes
- Commands with `set` are properly parsed without omitting semicolons.
- Double quotes are no longer dropped unexpectedly.
