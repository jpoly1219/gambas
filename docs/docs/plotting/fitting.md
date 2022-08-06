# Fitting

## Fit

```go
func Fit(ff string, pd PlotData, viaOpts ...GnuplotOpt) error
```
`Fit` fits a user-defined function `ff` to data given in `PlotData` `pd`, and prints out the results.

Pass options such as `using` in `pd`, but `via` in `viaOpts`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "explike-data.csv"), []string{"index"})
if err != nil {
    fmt.Println(err)
}

pd := gambas.PlotData{
    Df:       &df,
    Columns:  []string{"index", "data", "error"},
    Function: "",
    Opts:     []gambas.GnuplotOpt{gambas.Using("0:1:2 yerrors")},
}

err = gambas.Fit("a*exp(-b*x)", pd, gambas.Via("a,b"))
if err != nil {
    fmt.Println(err)
}
```
```
// explike-data.csv
index,data,error
0,1.129778763,0.1
1,0.953320457,0.081873075
2,0.732501277,0.067032005
3,0.561141921,0.054881164
4,0.434710291,0.044932896
5,0.349139904,0.036787944
6,0.308242881,0.030119421
7,0.26402811,0.024659696
8,0.18692481,0.020189652
9,0.159507023,0.016529889
10,0.15148364,0.013533528
11,0.086788709,0.011080316
12,0.094842665,0.009071795
13,0.081597097,0.007427358
14,0.052315262,0.006081006
15,0.05010193,0.004978707
16,0.043722877,0.00407622
17,0.036272859,0.003337327
18,0.025103659,0.002732372
19,0.019053795,0.002237077
20,0.018988513,0.001831564
```
```
iter      chisq       delta/lim  lambda   a             b            
   0 3.2009094356e+06   0.00e+00  1.81e-01    1.000000e+00   1.000000e+00
   * -nan              -nan       1.81e+00   -3.218104e+00  -2.204204e+01
   * 7.4311004197e+27   1.00e+05  1.81e+01    1.213369e+00  -1.394988e+00
   1 3.2009089058e+06  -1.66e-02  1.81e+00    1.004569e+00   9.730191e-01
iter      chisq       delta/lim  lambda   a             b            

After 1 iterations the fit converged.
final sum of squares of residuals : 3.20091e+06
rel. change during last iteration : -1.65518e-07

degrees of freedom    (FIT_NDF)                        : 19
rms of residuals      (FIT_STDFIT) = sqrt(WSSR/ndf)    : 410.45
variance of residuals (reduced chisquare) = WSSR/ndf   : 168469
p-value of the Chisq distribution (FIT_P)              : 0

Final set of parameters            Asymptotic Standard Error
=======================            ==========================
a               = 1.00457          +/- 452.4        (4.504e+04%)
b               = 0.973019         +/- 673.1        (6.917e+04%)

correlation matrix of the fit parameters:
                a      b      
a               1.000 
b               0.411  1.000
```