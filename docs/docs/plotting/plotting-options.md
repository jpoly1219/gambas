# Plotting Options

`gambas` provides many plotting options, aiming for feature parity with `Gnuplot`'s options.

Options implement the `GnuplotOpt` interface and can be passed around in different plotting functions.

## Set

Refer to the [Gnuplot documentation](http://gnuplot.info/docs_5.5/loc9418.html) on `set` command for more info. 

### Setangles

```go
func Setangles(value string) GnuplotOpt
```

### Setarrow

```go
func Setarrow(value string) GnuplotOpt
```

### Setautoscale

```go
func Setautoscale(value string) GnuplotOpt
```

### Setbmargin

```go
func Setbmargin(value string) GnuplotOpt
```

### Setborder

```go
func Setborder(value string) GnuplotOpt
```

### Setboxwidth

```go
func Setboxwidth(value string) GnuplotOpt
```

### Setboxdepth

```go
func Setboxdepth(value string) GnuplotOpt
```

### Setcolor

```go
func Setcolor() GnuplotOpt
```

### Setcolormap

```go
func Setcolormap(value string) GnuplotOpt
```

### Setcolorsequence

```go
func Setcolorsequence(value string) GnuplotOpt
```

### Setclip

```go
func Setclip(value string) GnuplotOpt
```

### Setcntrlabel

```go
func Setcntrlabel(value string) GnuplotOpt
```

### Setcntrparam

```go
func Setcntrparam(value string) GnuplotOpt
```

### Setcolorbox

```go
func Setcolorbox(value string) GnuplotOpt
```

### Setcontour

```go
func Setcontour(value string) GnuplotOpt
```

### Setdashtype

```go
func Setdashtype(value string) GnuplotOpt
```

### Setdatafile

```go
func Setdatafile(value string) GnuplotOpt
```

### Setdecimalsign

```go
func Setdecimalsign(value string) GnuplotOpt
```

### Setdgrid3d

```go
func Setdgrid3d(value string) GnuplotOpt
```

### Setdummy

```go
func Setdummy(value string) GnuplotOpt
```

### Setencoding

```go
func Setencoding(value string) GnuplotOpt
```

### Seterrorbars

```go
func Seterrorbars(value string) GnuplotOpt
```

### Setfit

```go
func Setfit(value string) GnuplotOpt
```

### Setfontpath

```go
func Setfontpath(value string) GnuplotOpt
```

### Setformat

```go
func Setformat(value string) GnuplotOpt
```

### Setgrid

```go
func Setgrid(value string) GnuplotOpt
```

### Sethidden3d

```go
func Sethidden3d(value string) GnuplotOpt
```

### Sethistorysize

```go
func Sethistorysize(value string) GnuplotOpt
```

### Sethistory

```go
func Sethistory(value string) GnuplotOpt
```

### Setisosamples

```go
func Setisosamples(value string) GnuplotOpt
```

### Setisosurface

```go
func Setisosurface(value string) GnuplotOpt
```

### Setisotropic

```go
func Setisotropic() GnuplotOpt
```

### Setjitter

```go
func Setjitter(value string) GnuplotOpt
```

### Setkey

```go
func Setkey(value string) GnuplotOpt
```

### Setlabel

```go
func Setlabel(value string) GnuplotOpt
```

### Setlinetype

```go
func Setlinetype(value string) GnuplotOpt
```

### Setlink

```go
func Setlink(value string) GnuplotOpt
```

### Setlmargin

```go
func Setlmargin(value string) GnuplotOpt
```

### Setloadpath

```go
func Setloadpath(value string) GnuplotOpt
```

### Setlocale

```go
func Setlocale(value string) GnuplotOpt
```

### Setlogscale

```go
func Setlogscale(value string) GnuplotOpt
```

### Setmapping

```go
func Setmapping(value string) GnuplotOpt
```

### Setmicro

```go
func Setmicro(value string) GnuplotOpt
```

### Setminussign

```go
func Setminussign(value string) GnuplotOpt
```

### Setmonochrome

```go
func Setmonochrome(value string) GnuplotOpt
```

### Setmouse

```go
func Setmouse(value string) GnuplotOpt
```

### Setmttics

```go
func Setmttics(value string) GnuplotOpt
```

### Setmultiplot

```go
func Setmultiplot(value string) GnuplotOpt
```

### Setmx2tics

```go
func Setmx2tics(value string) GnuplotOpt
```

### Setmy2tics

```go
func Setmy2tics(value string) GnuplotOpt
```

### Setmytics

```go
func Setmytics(value string) GnuplotOpt
```

### Setmztics

```go
func Setmztics(value string) GnuplotOpt
```

### Setnonlinear

```go
func Setnonlinear(value string) GnuplotOpt
```

### Setobject

```go
func Setobject(value string) GnuplotOpt
```

### Setoffsets

```go
func Setoffsets(value string) GnuplotOpt
```

### Setorigin

```go
func Setorigin(value string) GnuplotOpt
```

### Setoutput

```go
func Setoutput(value string) GnuplotOpt
```

### Setoverflow

```go
func Setoverflow(value string) GnuplotOpt
```

### Setpalette

```go
func Setpalette(value string) GnuplotOpt
```

### Setparametric

```go
func Setparametric(value string) GnuplotOpt
```

### Setpaxis

```go
func Setpaxis(value string) GnuplotOpt
```

### Setpixmap

```go
func Setpixmap(value string) GnuplotOpt
```

### Setpm3d

```go
func Setpm3d(value string) GnuplotOpt
```

### Setpointintervalbox

```go
func Setpointintervalbox() GnuplotOpt
```

### Setpointsize

```go
func Setpointsize(value string) GnuplotOpt
```

### Setpolar

```go
func Setpolar() GnuplotOpt
```

### Setprint

```go
func Setprint(value string) GnuplotOpt
```

### Setpsdir

```go
func Setpsdir(value string) GnuplotOpt
```

### Setraxis

```go
func Setraxis() GnuplotOpt
```

### Setrgbmax

```go
func Setrgbmax(value string) GnuplotOpt
```

### Setrlabel

```go
func Setrlabel(value string) GnuplotOpt
```

### Setrmargin

```go
func Setrmargin(value string) GnuplotOpt
```

### Setrrange

```go
func Setrrange(value string) GnuplotOpt
```

### Setrtics

```go
func Setrtics(value string) GnuplotOpt
```

### Setsamples

```go
func Setsamples(value string) GnuplotOpt
```

### Setsize

```go
func Setsize(value string) GnuplotOpt
```

### Setspiderplot

```go
func Setspiderplot() GnuplotOpt
```

### Setstyle

```go
func Setstyle(value string) GnuplotOpt
```

### Setsurface

```go
func Setsurface(value string) GnuplotOpt
```

### Settable

```go
func Settable(value string) GnuplotOpt
```

### Setterminal

```go
func Setterminal(value string) GnuplotOpt
```

### Settermoption

```go
func Settermoption(value string) GnuplotOpt
```

### Settheta

```go
func Settheta(value string) GnuplotOpt
```

### Settics

```go
func Settics(value string) GnuplotOpt
```

### Settimestamp

```go
func Settimestamp(value string) GnuplotOpt
```

### Settimefmt

```go
func Settimefmt(value string) GnuplotOpt
```

### Settitle

```go
func Settitle(value string) GnuplotOpt
```

### Settmargin

```go
func Settmargin(value string) GnuplotOpt
```

### Settrange

```go
func Settrange(value string) GnuplotOpt
```

### Setttics

```go
func Setttics(value string) GnuplotOpt
```

### Seturange

```go
func Seturange(value string) GnuplotOpt
```

### Setvgrid

```go
func Setvgrid(value string) GnuplotOpt
```

### Setview

```go
func Setview(value string) GnuplotOpt
```

### Setvrange

```go
func Setvrange(value string) GnuplotOpt
```

### Setvxrange

```go
func Setvxrange(value string) GnuplotOpt
```

### Setvyrange

```go
func Setvyrange(value string) GnuplotOpt
```

### Setvzrange

```go
func Setvzrange(value string) GnuplotOpt
```

### Setwalls

```go
func Setwalls(value string) GnuplotOpt
```

### Setx2ata

```go
func Setx2ata(value string) GnuplotOpt
```

### Setx2dtics

```go
func Setx2dtics(value string) GnuplotOpt
```

### Setx2label

```go
func Setx2label(value string) GnuplotOpt
```

### Setx2mtics

```go
func Setx2mtics(value string) GnuplotOpt
```

### Setx2range

```go
func Setx2range(value string) GnuplotOpt
```

### Setx2tics

```go
func Setx2tics(value string) GnuplotOpt
```

### Setx2zeroaxis

```go
func Setx2zeroaxis(value string) GnuplotOpt
```

### Setxdata

```go
func Setxdata(value string) GnuplotOpt
```

### Setxdtics

```go
func Setxdtics(value string) GnuplotOpt
```

### Setxlabel

```go
func Setxlabel(value string) GnuplotOpt
```

### Setxmtics

```go
func Setxmtics(value string) GnuplotOpt
```

### Setxrange

```go
func Setxrange(value string) GnuplotOpt
```

### Setxtics

```go
func Setxtics(value string) GnuplotOpt
```

### Setxyplane

```go
func Setxyplane(value string) GnuplotOpt
```

### Setxzeroaxis

```go
func Setxzeroaxis(value string) GnuplotOpt
```

### Sety2data

```go
func Sety2data(value string) GnuplotOpt
```

### Sety2dtics

```go
func Sety2dtics(value string) GnuplotOpt
```

### Sety2label

```go
func Sety2label(value string) GnuplotOpt
```

### Sety2mtics

```go
func Sety2mtics(value string) GnuplotOpt
```

### Sety2range

```go
func Sety2range(value string) GnuplotOpt
```

### Sety2tics

```go
func Sety2tics(value string) GnuplotOpt
```

### Sety2zeroaxis

```go
func Sety2zeroaxis(value string) GnuplotOpt
```

### Setydata

```go
func Setydata(value string) GnuplotOpt
```

### Setydtics

```go
func Setydtics(value string) GnuplotOpt
```

### Setylabel

```go
func Setylabel(value string) GnuplotOpt
```

### Setymtics

```go
func Setymtics(value string) GnuplotOpt
```

### Setyrange

```go
func Setyrange(value string) GnuplotOpt
```

### Setytics

```go
func Setytics(value string) GnuplotOpt
```

### Setyzeroaxis

```go
func Setyzeroaxis(value string) GnuplotOpt
```

### Setzdata

```go
func Setzdata(value string) GnuplotOpt
```

### Setzdtics

```go
func Setzdtics(value string) GnuplotOpt
```

### Setzzeroaxis

```go
func Setzzeroaxis(value string) GnuplotOpt
```

### Setcbdata

```go
func Setcbdata(value string) GnuplotOpt
```

### Setcbdtics

```go
func Setcbdtics(value string) GnuplotOpt
```

### Setzero

```go
func Setzero(value string) GnuplotOpt
```

### Setzlabel

```go
func Setzlabel(value string) GnuplotOpt
```

### Setzmtics

```go
func Setzmtics(value string) GnuplotOpt
```

### Setzrange

```go
func Setzrange(value string) GnuplotOpt
```

### Setztics

```go
func Setztics(value string) GnuplotOpt
```

### Setcblabel

```go
func Setcblabel(value string) GnuplotOpt
```

### Setcbmtics

```go
func Setcbmtics(value string) GnuplotOpt
```

### Setcbrange

```go
func Setcbrange(value string) GnuplotOpt
```

### Setcbtics

```go
func Setcbtics(value string) GnuplotOpt
```

## Using

```go
func Using(value string) GnuplotOpt
```

Refer to the [Gnuplot documentation](http://gnuplot.info/docs_5.5/loc8480.html) on `using` command for more info.

## Via

```go
func Via(value string) GnuplotOpt
```

`Via` is generally used with `fit`.
Refer to the [Gnuplot documentation](http://gnuplot.info/docs_5.5/loc6410.html) on `fit` and `via` command for more info.

## With

```go
func With(value string) GnuplotOpt
```

Refer to the [Gnuplot documentation](http://gnuplot.info/docs_5.5/loc9051.html) on `with` command for more info.