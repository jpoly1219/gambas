package gambas

import (
	"fmt"
)

type GnuplotOpt interface {
	createCmdString() string
}

// Refer to the page below for more detailed usage of gnuplot options.
// http://gnuplot.info/docs_5.5/loc9418.html
// Options not included: bind, clabel, colornames, macros

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

type grid struct {
	option string
	value  string
}

func (g grid) createCmdString() string {
	return fmt.Sprintf(`set grid %s`, g.value)
}

func Setgrid(value string) GnuplotOpt {
	g := grid{"grid", value}
	return g
}

type hidden3d struct {
	option string
	value  string
}

func (h hidden3d) createCmdString() string {
	return fmt.Sprintf(`set hidden3d %s`, h.value)
}

func Sethidden3d(value string) GnuplotOpt {
	h := hidden3d{"hidden3d", value}
	return h
}

type historysize struct {
	option string
	value  string
}

func (h historysize) createCmdString() string {
	return fmt.Sprintf(`set historysize %s`, h.value)
}

func Sethistorysize(value string) GnuplotOpt {
	h := historysize{"historysize", value}
	return h
}

type history struct {
	option string
	value  string
}

func (h history) createCmdString() string {
	return fmt.Sprintf(`set history %s`, h.value)
}

func Sethistory(value string) GnuplotOpt {
	h := history{"history", value}
	return h
}

type isosamples struct {
	option string
	value  string
}

func (i isosamples) createCmdString() string {
	return fmt.Sprintf(`set isosamples %s`, i.value)
}

func Setisosamples(value string) GnuplotOpt {
	i := isosamples{"isosamples", value}
	return i
}

type isosurface struct {
	option string
	value  string
}

func (i isosurface) createCmdString() string {
	return fmt.Sprintf(`set isosurface %s`, i.value)
}

func Setisosurface(value string) GnuplotOpt {
	i := isosurface{"isosurface", value}
	return i
}

type isotropic struct {
	option string
	value  string
}

func (i isotropic) createCmdString() string {
	return "set isotropic"
}

func Setisotropic() GnuplotOpt {
	i := isotropic{"isotropic", ""}
	return i
}

type jitter struct {
	option string
	value  string
}

func (j jitter) createCmdString() string {
	return fmt.Sprintf(`set jitter %s`, j.value)
}

func Setjitter(value string) GnuplotOpt {
	j := jitter{"jitter", value}
	return j
}

type key struct {
	option string
	value  string
}

func (k key) createCmdString() string {
	return fmt.Sprintf(`set key %s`, k.value)
}

func Setkey(value string) GnuplotOpt {
	k := key{"key", value}
	return k
}

type label struct {
	option string
	value  string
}

func (l label) createCmdString() string {
	return fmt.Sprintf(`set label %s`, l.value)
}

func Setlabel(value string) GnuplotOpt {
	l := label{"label", value}
	return l
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

type link struct {
	option string
	value  string
}

func (l link) createCmdString() string {
	return fmt.Sprintf("set link %s", l.value)
}

func Setlink(value string) GnuplotOpt {
	l := link{"link", value}
	return l
}

type lmargin struct {
	option string
	value  string
}

func (l lmargin) createCmdString() string {
	return fmt.Sprintf("set lmargin %s", l.value)
}

func Setlmargin(value string) GnuplotOpt {
	l := lmargin{"lmargin", value}
	return l
}

type loadpath struct {
	option string
	value  string
}

func (l loadpath) createCmdString() string {
	return fmt.Sprintf("set loadpath %s", l.value)
}

func Setloadpath(value string) GnuplotOpt {
	l := loadpath{"loadpath", value}
	return l
}

type locale struct {
	option string
	value  string
}

func (l locale) createCmdString() string {
	return fmt.Sprintf("set locale %s", l.value)
}

func Setlocale(value string) GnuplotOpt {
	l := locale{"locale", value}
	return l
}

type logscale struct {
	option string
	value  string
}

func (l logscale) createCmdString() string {
	return fmt.Sprintf("set logscale %s", l.value)
}

func Setlogscale(value string) GnuplotOpt {
	l := logscale{"logscale", value}
	return l
}

type mapping struct {
	option string
	value  string
}

func (m mapping) createCmdString() string {
	return fmt.Sprintf("set mapping %s", m.value)
}

func Setmapping(value string) GnuplotOpt {
	m := mapping{"mapping", value}
	return m
}

type micro struct {
	option string
	value  string
}

func (m micro) createCmdString() string {
	return fmt.Sprintf("set micro %s", m.value)
}

func Setmicro(value string) GnuplotOpt {
	m := micro{"micro", value}
	return m
}

type minussign struct {
	option string
	value  string
}

func (m minussign) createCmdString() string {
	return fmt.Sprintf("set minussign %s", m.value)
}

func Setminussign(value string) GnuplotOpt {
	m := minussign{"minussign", value}
	return m
}

type monochrome struct {
	option string
	value  string
}

func (m monochrome) createCmdString() string {
	return fmt.Sprintf("set monochrome %s", m.value)
}

func Setmonochrome(value string) GnuplotOpt {
	m := monochrome{"monochrome", value}
	return m
}

type mouse struct {
	option string
	value  string
}

func (m mouse) createCmdString() string {
	return fmt.Sprintf("set mouse %s", m.value)
}

func Setmouse(value string) GnuplotOpt {
	m := mouse{"mouse", value}
	return m
}

type mttics struct {
	option string
	value  string
}

func (m mttics) createCmdString() string {
	return fmt.Sprintf("set mttics %s", m.value)
}

func Setmttics(value string) GnuplotOpt {
	m := mttics{"mttics", value}
	return m
}

type multiplot struct {
	option string
	value  string
}

func (m multiplot) createCmdString() string {
	return fmt.Sprintf("set multiplot %s", m.value)
}

func Setmultiplot(value string) GnuplotOpt {
	m := multiplot{"multiplot", value}
	return m
}

type mx2tics struct {
	option string
	value  string
}

func (m mx2tics) createCmdString() string {
	return fmt.Sprintf("set mx2tics %s", m.value)
}

func Setmx2tics(value string) GnuplotOpt {
	m := mx2tics{"mx2tics", value}
	return m
}

type my2tics struct {
	option string
	value  string
}

func (m my2tics) createCmdString() string {
	return fmt.Sprintf("set my2tics %s", m.value)
}

func Setmy2tics(value string) GnuplotOpt {
	m := my2tics{"my2tics", value}
	return m
}

type mytics struct {
	option string
	value  string
}

func (m mytics) createCmdString() string {
	return fmt.Sprintf("set mytics %s", m.value)
}

func Setmytics(value string) GnuplotOpt {
	m := mytics{"mytics", value}
	return m
}

type mztics struct {
	option string
	value  string
}

func (m mztics) createCmdString() string {
	return fmt.Sprintf("set mztics %s", m.value)
}

func Setmztics(value string) GnuplotOpt {
	m := mztics{"mztics", value}
	return m
}

type nonlinear struct {
	option string
	value  string
}

func (n nonlinear) createCmdString() string {
	return fmt.Sprintf("set nonlinear %s", n.value)
}

func Setnonlinear(value string) GnuplotOpt {
	n := nonlinear{"nonlinear", value}
	return n
}

type object struct {
	option string
	value  string
}

func (o object) createCmdString() string {
	return fmt.Sprintf("set object %s", o.value)
}

func Setobject(value string) GnuplotOpt {
	o := object{"object", value}
	return o
}

type offsets struct {
	option string
	value  string
}

func (o offsets) createCmdString() string {
	return fmt.Sprintf("set offsets %s", o.value)
}

func Setoffsets(value string) GnuplotOpt {
	o := offsets{"offsets", value}
	return o
}

type origin struct {
	option string
	value  string
}

func (o origin) createCmdString() string {
	return fmt.Sprintf("set origin %s", o.value)
}

func Setorigin(value string) GnuplotOpt {
	o := origin{"origin", value}
	return o
}

type output struct {
	option string
	value  string
}

func (o output) createCmdString() string {
	return fmt.Sprintf("set output %s", o.value)
}

func Setoutput(value string) GnuplotOpt {
	o := output{"output", value}
	return o
}

type overflow struct {
	option string
	value  string
}

func (o overflow) createCmdString() string {
	return fmt.Sprintf("set overflow %s", o.value)
}

func Setoverflow(value string) GnuplotOpt {
	o := overflow{"overflow", value}
	return o
}

type palette struct {
	option string
	value  string
}

func (p palette) createCmdString() string {
	return fmt.Sprintf("set palette %s", p.value)
}

func Setpalette(value string) GnuplotOpt {
	p := palette{"palette", value}
	return p
}

type parametric struct {
	option string
	value  string
}

func (p parametric) createCmdString() string {
	return fmt.Sprintf("set parametric %s", p.value)
}

func Setparametric(value string) GnuplotOpt {
	p := parametric{"parametric", value}
	return p
}

type paxis struct {
	option string
	value  string
}

func (p paxis) createCmdString() string {
	return fmt.Sprintf("set paxis %s", p.value)
}

func Setpaxis(value string) GnuplotOpt {
	p := paxis{"paxis", value}
	return p
}

type pixmap struct {
	option string
	value  string
}

func (p pixmap) createCmdString() string {
	return fmt.Sprintf("set pixmap %s", p.value)
}

func Setpixmap(value string) GnuplotOpt {
	p := pixmap{"pixmap", value}
	return p
}

type pm3d struct {
	option string
	value  string
}

func (p pm3d) createCmdString() string {
	return fmt.Sprintf("set pm3d %s", p.value)
}

func Setpm3d(value string) GnuplotOpt {
	p := pm3d{"pm3d", value}
	return p
}

type pointintervalbox struct {
	option string
	value  string
}

func (p pointintervalbox) createCmdString() string {
	return "set pointintervalbox"
}

func Setpointintervalbox() GnuplotOpt {
	p := pointintervalbox{"pointintervalbox", ""}
	return p
}

type pointsize struct {
	option string
	value  string
}

func (p pointsize) createCmdString() string {
	return fmt.Sprintf("set pointsize %s", p.value)
}

func Setpointsize(value string) GnuplotOpt {
	p := pointsize{"pointsize", value}
	return p
}

type polar struct {
	option string
	value  string
}

func (p polar) createCmdString() string {
	return "set polar"
}

func Setpolar() GnuplotOpt {
	p := polar{"polar", ""}
	return p
}

type print struct {
	option string
	value  string
}

func (p print) createCmdString() string {
	return fmt.Sprintf("set print %s", p.value)
}

func Setprint(value string) GnuplotOpt {
	p := print{"print", value}
	return p
}

type psdir struct {
	option string
	value  string
}

func (p psdir) createCmdString() string {
	return fmt.Sprintf("set psdir %s", p.value)
}

func Setpsdir(value string) GnuplotOpt {
	p := psdir{"psdir", value}
	return p
}

type raxis struct {
	option string
	value  string
}

func (r raxis) createCmdString() string {
	return "set raxis"
}

func Setraxis() GnuplotOpt {
	r := raxis{"raxis", ""}
	return r
}

type rgbmax struct {
	option string
	value  string
}

func (r rgbmax) createCmdString() string {
	return fmt.Sprintf("set rgbmax %s", r.value)
}

func Setrgbmax(value string) GnuplotOpt {
	r := rgbmax{"rgbmax", value}
	return r
}

type rlabel struct {
	option string
	value  string
}

func (r rlabel) createCmdString() string {
	return fmt.Sprintf("set rlabel %s", r.value)
}

func Setrlabel(value string) GnuplotOpt {
	r := rlabel{"rlabel", value}
	return r
}

type rmargin struct {
	option string
	value  string
}

func (r rmargin) createCmdString() string {
	return fmt.Sprintf("set rmargin %s", r.value)
}

func Setrmargin(value string) GnuplotOpt {
	r := rmargin{"rmargin", value}
	return r
}

type rrange struct {
	option string
	value  string
}

func (r rrange) createCmdString() string {
	return fmt.Sprintf("set rrange %s", r.value)
}

func Setrrange(value string) GnuplotOpt {
	r := rrange{"rrange", value}
	return r
}

type rtics struct {
	option string
	value  string
}

func (r rtics) createCmdString() string {
	return fmt.Sprintf("set rtics %s", r.value)
}

func Setrtics(value string) GnuplotOpt {
	r := rtics{"rtics", value}
	return r
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
