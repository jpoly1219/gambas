# gambas
---

**gambas** is a data analysis package for Go, inspired by the famous Python library **pandas**. It provides an intuitive way to malipulate tabular data.

gambas was created to serve the needs of Go developers. We can all agree that pandas is an amazing tool, and can be considered the industry standard when it comes to data analysis. However, some of us may not enjoy using Python very much. Go is a very attractive language with a very loyal userbase (including me!). You might be a Gopher, but you also might love the features Python provides. You may have jumped ship and used Go for some time now, but you also might miss some of the libraries you used when you were using Python. 

gambas aims to scratch that itch. You will be able to tap into the superpowers of pandas while using your favorite language Go. Gophers cannot into data analysis? No, of course we can! Hope you enjoy using this package.

### MVP Goals
---
- Provide basic features from the pandas tutorial.
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

- Extensive, easy to follow documentation