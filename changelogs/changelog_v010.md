# Changelog for v0.1.0

## New features:
- You can now merge DataFrame objects horizontally and vertically using `MergeDfsHorizontally` and `MergeDfsVertically`.
- Basic support for reading and writing Excel files has arrived! Use `ReadExcel` and `WriteExcel`.

## Bug fixes:
- `ReadJsonStream` now properly catches JSON tokens.
- `ReadJsonStream` now sorts the resulting DataFrame object `newDf` before returning.
- `NewCol`, `NewDerivedCol`, `RenameCol` no longer modifies the original DataFrame object.
- `ReadCsv` now uses RangeIndex instead of the first column of data.

## Optimization
- `Mean`, `Std`, `Min`, `Max` now run faster.
- All NaN values will now use `math.NaN` instead of mixing `math.NaN` and `"NaN"`.
