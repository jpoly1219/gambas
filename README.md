# gambas
---

**gambas** is a data analysis library for Go, inspired by the famous Python library **pandas**. It provides an intuitive way to malipulate tabular data.

gambas was created to serve the needs of Go developers. I thought it would be cool and helpful to have a data analysis library that is easy to use.

### MVP Goals
---
- Provide basic features, from the pandas tutorial.
  - [x] Providing `Series` and `DataFrame` data types
  - [ ] Reading and writing tabular data
    - [x] Reading CSV files
    - [ ] Writing to CSV files
    - [ ] Reading Excel files
    - [ ] Writing to Excel files
    - [ ] Reading JSON files
    - [ ] Writing to JSON files
  - [x] Selecting a subset of data
    - [x] At, IAt
    - [x] Loc, ILoc
  - [ ] Plotting
  - [ ] Creating new columns derived from existing columns
    - [x] Creating new columns
    - [ ] Applying operations to the new column
    - [x] Renaming columns
  - [x] Calculating summary statistics
    - [x] Mean, median, standard deviation
    - [x] Min, max, quartiles
    - [x] Count, describe
  - [ ] Reshaping the layout of tables
    - [x] Sorting by index
    - [x] Sorting by values
    - [x] Sorting by given index
    - [ ] Groupby
    - [ ] Pivot
    - [ ] PivotTable
    - [ ] Long to wide format
    - [ ] Wide to long format
  - [ ] Combining data from multiple tables
    - [ ] Concatenate
    - [ ] Merge
  - [ ] Handling time series data
    - [ ] Timestamp type
    - [ ] Timestamp type methods
    - [ ] ToDatetime
  - [ ] Manipulating textual data
  - [x] Multiindex
