package gambas

import (
	"fmt"
)

type GnuplotOpt interface {
	createCmdString() string
}

// Refer to the page below for more detailed usage of gnuplot options.
// http://gnuplot.info/docs_5.5/loc9418.html
// Options not included: bind, clabel, colornames,

type angles struct {
	option string
	value  string
}

func (a angles) createCmdString() string {
	return fmt.Sprintf("set angles %s", a.value)
}

func SetAngles(value string) GnuplotOpt {
	a := angles{"angles", value}
	return a
}

type arrow struct {
	option string
	value  string
}

func (a arrow) createCmdString() string {
	return fmt.Sprintf("set arrow %s", a.value)
}

func Setarrow(value string) GnuplotOpt {
	a := arrow{"arrow", value}
	return a
}

type autoscale struct {
	option string
	value  string
}

func (a autoscale) createCmdString() string {
	return fmt.Sprintf("set autoscale %s", a.value)
}

func Setautoscale(value string) GnuplotOpt {
	a := autoscale{"autoscale", value}
	return a
}

type bmargin struct {
	option string
	value  string
}

func (b bmargin) createCmdString() string {
	return fmt.Sprintf("set bmargin %s", b.value)
}

func Setbmargin(value string) GnuplotOpt {
	b := bmargin{"bmargin", value}
	return b
}

type border struct {
	option string
	value  string
}

func (b border) createCmdString() string {
	return fmt.Sprintf("set border %s", b.value)
}

func Setborder(value string) GnuplotOpt {
	b := border{"border", value}
	return b
}

type boxwidth struct {
	option string
	value  string
}

func (b boxwidth) createCmdString() string {
	return fmt.Sprintf("set boxwidth %s", b.value)
}

func Setboxwidth(value string) GnuplotOpt {
	b := boxwidth{"boxwidth", value}
	return b
}

type boxdepth struct {
	option string
	value  string
}

func (b boxdepth) createCmdString() string {
	return fmt.Sprintf("set boxdepth %s", b.value)
}

func Setboxdepth(value string) GnuplotOpt {
	b := boxdepth{"boxdepth", value}
	return b
}

type color struct {
	option string
	value  string
}

func (c color) createCmdString() string {
	return fmt.Sprintf("set color %s", c.value)
}

func Setcolor() GnuplotOpt {
	c := color{"color", ""}
	return c
}

type colormap struct {
	option string
	value  string
}

func (c colormap) createCmdString() string {
	return fmt.Sprintf("set colormap %s", c.value)
}

func Setcolormap(value string) GnuplotOpt {
	c := colormap{"colormap", value}
	return c
}

type colorsequence struct {
	option string
	value  string
}

func (c colorsequence) createCmdString() string {
	return fmt.Sprintf("set colorsequence %s", c.value)
}

func Setcolorsequence(value string) GnuplotOpt {
	c := colorsequence{"colorsequence", value}
	return c
}

type clip struct {
	option string
	value  string
}

func (c clip) createCmdString() string {
	return fmt.Sprintf("set clip %s", c.value)
}

func Setclip(value string) GnuplotOpt {
	c := clip{"clip", value}
	return c
}

type cntrlabel struct {
	option string
	value  string
}

func (c cntrlabel) createCmdString() string {
	return fmt.Sprintf("set cntrlabel %s", c.value)
}

func Setcntrlabel(value string) GnuplotOpt {
	c := cntrlabel{"cntrlabel", value}
	return c
}

type cntrparam struct {
	option string
	value  string
}

func (c cntrparam) createCmdString() string {
	return fmt.Sprintf("set cntrparam %s", c.value)
}

func Setcntrparam(value string) GnuplotOpt {
	c := cntrparam{"cntrparam", value}
	return c
}

type colorbox struct {
	option string
	value  string
}

func (c colorbox) createCmdString() string {
	return fmt.Sprintf("set colorbox %s", c.value)
}

func Setcolorbox(value string) GnuplotOpt {
	c := colorbox{"colorbox", value}
	return c
}

type contour struct {
	option string
	value  string
}

func (c contour) createCmdString() string {
	return fmt.Sprintf("set contour %s", c.value)
}

func Setcontour(value string) GnuplotOpt {
	c := contour{"contour", value}
	return c
}

type cornerpoles struct {
	option string
	value  string
}

func (c cornerpoles) createCmdString() string {
	return "unset cornerpoles"
}

func Unsetcornerpoles() GnuplotOpt {
	c := cornerpoles{"cornerpoles", ""}
	return c
}

type dashtype struct {
	option string
	value  string
}

func (d dashtype) createCmdString() string {
	return fmt.Sprintf("set dashtype %s", d.value)
}

func Setdashtype(value string) GnuplotOpt {
	d := dashtype{"dashtype", value}
	return d
}

type datafile struct {
	option string
	value  string
}

func (d datafile) createCmdString() string {
	return fmt.Sprintf("set datafile %s", d.value)
}

func Setdatafile(value string) GnuplotOpt {
	d := datafile{"datafile", value}
	return d
}

type decimalsign struct {
	option string
	value  string
}

func (d decimalsign) createCmdString() string {
	return fmt.Sprintf("set decimalsign %s", d.value)
}

func Setdecimalsign(value string) GnuplotOpt {
	d := decimalsign{"decimalsign", value}
	return d
}

type dgrid3d struct {
	option string
	value  string
}

func (d dgrid3d) createCmdString() string {
	return fmt.Sprintf("set dgrid3d %s", d.value)
}

func Setdgrid3d(value string) GnuplotOpt {
	d := dgrid3d{"dgrid3d", value}
	return d
}

type dummy struct {
	option string
	value  string
}

func (d dummy) createCmdString() string {
	return fmt.Sprintf("set dummy %s", d.value)
}

func Setdummy(value string) GnuplotOpt {
	d := dummy{"dummy", value}
	return d
}

type encoding struct {
	option string
	value  string
}

func (e encoding) createCmdString() string {
	return fmt.Sprintf("set encoding %s", e.value)
}

func Setencoding(value string) GnuplotOpt {
	e := encoding{"encoding", value}
	return e
}

type errorbars struct {
	option string
	value  string
}

func (e errorbars) createCmdString() string {
	return fmt.Sprintf("set errorbars %s", e.value)
}

func Seterrorbars(value string) GnuplotOpt {
	e := errorbars{"errorbars", value}
	return e
}

type fit struct {
	option string
	value  string
}

func (f fit) createCmdString() string {
	return fmt.Sprintf("set fit %s", f.value)
}

func Setfit(value string) GnuplotOpt {
	f := fit{"fit", value}
	return f
}

type fontpath struct {
	option string
	value  string
}

func (f fontpath) createCmdString() string {
	return fmt.Sprintf(`set fontpath "%s"`, f.value)
}

func Setfontpath(value string) GnuplotOpt {
	f := fontpath{"fontpath", value}
	return f
}

type format struct {
	option string
	value  string
}

func (f format) createCmdString() string {
	return fmt.Sprintf(`set format %s`, f.value)
}

func Setformat(value string) GnuplotOpt {
	f := format{"format", value}
	return f
}

type linetype struct {
	option string
	value  string
}

func (l linetype) createCmdString() string {
	return fmt.Sprintf("set linetype %s", l.value)
}

func SetLinetype(value string) GnuplotOpt {
	l := linetype{"linetype", value}
	return l
}

type style struct {
	option string
	value  string
}

func (s style) createCmdString() string {
	return fmt.Sprintf(`set style "%s"`, s.value)
}

func Setstyle(value string) GnuplotOpt {
	s := style{"style", value}
	return s
}

type timefmt struct {
	option string
	value  string
}

func (t timefmt) createCmdString() string {
	return fmt.Sprintf(`set timefmt "%s"`, t.value)
}

func SetTimefmt(value string) GnuplotOpt {
	t := timefmt{"timefmt", value}
	return t
}

type xdata struct {
	option string
	value  string
}

func (x xdata) createCmdString() string {
	if x.value == "" {
		return "set xdata"
	}
	return fmt.Sprintf("set xdata %s", x.value)
}

func SetXdata(value string) GnuplotOpt {
	x := xdata{"xdata", value}
	return x
}

type ydata struct {
	option string
	value  string
}

func (y ydata) createCmdString() string {
	if y.value == "" {
		return "set ydata"
	}
	return fmt.Sprintf("set ydata%s", " "+y.value)
}

func SetYdata(value string) GnuplotOpt {
	y := ydata{"ydata", value}
	return y
}
