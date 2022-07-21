package gambas

import (
	"fmt"
)

// A GnuplotOpt represents an option in gnuplot.
type GnuplotOpt interface {
	createCmdString() string
	getOption() string
}

// Refer to the page below for more detailed usage of gnuplot `set` options.
// http://gnuplot.info/docs_5.5/loc9418.html
// Options not included: bind, clabel, colornames, macros, ticslevel, ticscale, version

type angles struct {
	option string
	value  string
}

func (a angles) createCmdString() string {
	return fmt.Sprintf("set angles %s", a.value)
}

func (a angles) getOption() string {
	return a.option
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

func (a arrow) getOption() string {
	return a.option
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

func (a autoscale) getOption() string {
	return a.option
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

func (b bmargin) getOption() string {
	return b.option
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

func (b border) getOption() string {
	return b.option
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

func (b boxwidth) getOption() string {
	return b.option
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

func (b boxdepth) getOption() string {
	return b.option
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

func (c color) getOption() string {
	return c.option
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

func (c colormap) getOption() string {
	return c.option
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

func (c colorsequence) getOption() string {
	return c.option
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

func (c clip) getOption() string {
	return c.option
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

func (c cntrlabel) getOption() string {
	return c.option
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

func (c cntrparam) getOption() string {
	return c.option
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

func (c colorbox) getOption() string {
	return c.option
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

func (c contour) getOption() string {
	return c.option
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

func (c cornerpoles) getOption() string {
	return c.option
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

func (d dashtype) getOption() string {
	return d.option
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

func (d datafile) getOption() string {
	return d.option
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

func (d decimalsign) getOption() string {
	return d.option
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

func (d dgrid3d) getOption() string {
	return d.option
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

func (d dummy) getOption() string {
	return d.option
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

func (e encoding) getOption() string {
	return e.option
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

func (e errorbars) getOption() string {
	return e.option
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

func (f fit) getOption() string {
	return f.option
}

func Setfit(value string) GnuplotOpt {
	f := fit{"fit", value}
	return f
}

type fontpath struct {
	option string
	value  string
}

func (f fontpath) getOption() string {
	return f.option
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

func (f format) getOption() string {
	return f.option
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

func (g grid) getOption() string {
	return g.option
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

func (h hidden3d) getOption() string {
	return h.option
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

func (h historysize) getOption() string {
	return h.option
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

func (h history) getOption() string {
	return h.option
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

func (i isosamples) getOption() string {
	return i.option
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

func (i isosurface) getOption() string {
	return i.option
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

func (i isotropic) getOption() string {
	return i.option
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

func (j jitter) getOption() string {
	return j.option
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

func (k key) getOption() string {
	return k.option
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

func (l label) getOption() string {
	return l.option
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

func (l linetype) getOption() string {
	return l.option
}

func Setlinetype(value string) GnuplotOpt {
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

func (l link) getOption() string {
	return l.option
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

func (l lmargin) getOption() string {
	return l.option
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

func (l loadpath) getOption() string {
	return l.option
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

func (l locale) getOption() string {
	return l.option
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

func (l logscale) getOption() string {
	return l.option
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

func (m mapping) getOption() string {
	return m.option
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

func (m micro) getOption() string {
	return m.option
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

func (m minussign) getOption() string {
	return m.option
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

func (m monochrome) getOption() string {
	return m.option
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

func (m mouse) getOption() string {
	return m.option
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

func (m mttics) getOption() string {
	return m.option
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

func (m multiplot) getOption() string {
	return m.option
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

func (m mx2tics) getOption() string {
	return m.option
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

func (m my2tics) getOption() string {
	return m.option
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

func (m mytics) getOption() string {
	return m.option
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

func (m mztics) getOption() string {
	return m.option
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

func (n nonlinear) getOption() string {
	return n.option
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

func (o object) getOption() string {
	return o.option
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

func (o offsets) getOption() string {
	return o.option
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

func (o origin) getOption() string {
	return o.option
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

func (o output) getOption() string {
	return o.option
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

func (o overflow) getOption() string {
	return o.option
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

func (p palette) getOption() string {
	return p.option
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

func (p parametric) getOption() string {
	return p.option
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

func (p paxis) getOption() string {
	return p.option
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

func (p pixmap) getOption() string {
	return p.option
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

func (p pm3d) getOption() string {
	return p.option
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

func (p pointintervalbox) getOption() string {
	return p.option
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

func (p pointsize) getOption() string {
	return p.option
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

func (p polar) getOption() string {
	return p.option
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

func (p print) getOption() string {
	return p.option
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

func (p psdir) getOption() string {
	return p.option
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

func (r raxis) getOption() string {
	return r.option
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

func (r rgbmax) getOption() string {
	return r.option
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

func (r rlabel) getOption() string {
	return r.option
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

func (r rmargin) getOption() string {
	return r.option
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

func (r rrange) getOption() string {
	return r.option
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

func (r rtics) getOption() string {
	return r.option
}

func Setrtics(value string) GnuplotOpt {
	r := rtics{"rtics", value}
	return r
}

type samples struct {
	option string
	value  string
}

func (s samples) createCmdString() string {
	return fmt.Sprintf(`set samples "%s"`, s.value)
}

func (s samples) getOption() string {
	return s.option
}

func Setsamples(value string) GnuplotOpt {
	s := samples{"samples", value}
	return s
}

type size struct {
	option string
	value  string
}

func (s size) createCmdString() string {
	return fmt.Sprintf(`set size "%s"`, s.value)
}

func (s size) getOption() string {
	return s.option
}

func Setsize(value string) GnuplotOpt {
	s := size{"size", value}
	return s
}

type spiderplot struct {
	option string
	value  string
}

func (s spiderplot) createCmdString() string {
	return "set spiderplot"
}

func (s spiderplot) getOption() string {
	return s.option
}

func Setspiderplot() GnuplotOpt {
	s := spiderplot{"spiderplot", ""}
	return s
}

type style struct {
	option string
	value  string
}

func (s style) createCmdString() string {
	return fmt.Sprintf(`set style "%s"`, s.value)
}

func (s style) getOption() string {
	return s.option
}

func Setstyle(value string) GnuplotOpt {
	s := style{"style", value}
	return s
}

type surface struct {
	option string
	value  string
}

func (s surface) createCmdString() string {
	return fmt.Sprintf(`set surface "%s"`, s.value)
}

func (s surface) getOption() string {
	return s.option
}

func Setsurface(value string) GnuplotOpt {
	s := surface{"surface", value}
	return s
}

type table struct {
	option string
	value  string
}

func (t table) createCmdString() string {
	return fmt.Sprintf(`set table "%s"`, t.value)
}

func (t table) getOption() string {
	return t.option
}

func Settable(value string) GnuplotOpt {
	t := table{"table", value}
	return t
}

type terminal struct {
	option string
	value  string
}

func (t terminal) createCmdString() string {
	return fmt.Sprintf(`set terminal "%s"`, t.value)
}

func (t terminal) getOption() string {
	return t.option
}

func Setterminal(value string) GnuplotOpt {
	t := terminal{"terminal", value}
	return t
}

type termoption struct {
	option string
	value  string
}

func (t termoption) createCmdString() string {
	return fmt.Sprintf(`set termoption "%s"`, t.value)
}

func (t termoption) getOption() string {
	return t.option
}

func Settermoption(value string) GnuplotOpt {
	t := termoption{"termoption", value}
	return t
}

type theta struct {
	option string
	value  string
}

func (t theta) createCmdString() string {
	return fmt.Sprintf(`set theta "%s"`, t.value)
}

func (t theta) getOption() string {
	return t.option
}

func Settheta(value string) GnuplotOpt {
	t := theta{"theta", value}
	return t
}

type tics struct {
	option string
	value  string
}

func (t tics) createCmdString() string {
	return fmt.Sprintf(`set tics "%s"`, t.value)
}

func (t tics) getOption() string {
	return t.option
}

func Settics(value string) GnuplotOpt {
	t := tics{"tics", value}
	return t
}

type timestamp struct {
	option string
	value  string
}

func (t timestamp) createCmdString() string {
	return fmt.Sprintf(`set timestamp "%s"`, t.value)
}

func (t timestamp) getOption() string {
	return t.option
}

func Settimestamp(value string) GnuplotOpt {
	t := timestamp{"timestamp", value}
	return t
}

type timefmt struct {
	option string
	value  string
}

func (t timefmt) createCmdString() string {
	return fmt.Sprintf(`set timefmt "%s"`, t.value)
}

func (t timefmt) getOption() string {
	return t.option
}

func Settimefmt(value string) GnuplotOpt {
	t := timefmt{"timefmt", value}
	return t
}

type title struct {
	option string
	value  string
}

func (t title) createCmdString() string {
	return fmt.Sprintf(`set title "%s"`, t.value)
}

func (t title) getOption() string {
	return t.option
}

func Settitle(value string) GnuplotOpt {
	t := title{"title", value}
	return t
}

type tmargin struct {
	option string
	value  string
}

func (t tmargin) createCmdString() string {
	return fmt.Sprintf(`set tmargin "%s"`, t.value)
}

func (t tmargin) getOption() string {
	return t.option
}

func Settmargin(value string) GnuplotOpt {
	t := tmargin{"tmargin", value}
	return t
}

type trange struct {
	option string
	value  string
}

func (t trange) createCmdString() string {
	return fmt.Sprintf(`set trange "%s"`, t.value)
}

func (t trange) getOption() string {
	return t.option
}

func Settrange(value string) GnuplotOpt {
	t := trange{"trange", value}
	return t
}

type ttics struct {
	option string
	value  string
}

func (t ttics) createCmdString() string {
	return fmt.Sprintf(`set ttics "%s"`, t.value)
}

func (t ttics) getOption() string {
	return t.option
}

func Setttics(value string) GnuplotOpt {
	t := ttics{"ttics", value}
	return t
}

type urange struct {
	option string
	value  string
}

func (u urange) createCmdString() string {
	return fmt.Sprintf(`set urange "%s"`, u.value)
}

func (u urange) getOption() string {
	return u.option
}

func Seturange(value string) GnuplotOpt {
	u := urange{"urange", value}
	return u
}

type vgrid struct {
	option string
	value  string
}

func (v vgrid) createCmdString() string {
	return fmt.Sprintf(`set vgrid "%s"`, v.value)
}

func (v vgrid) getOption() string {
	return v.option
}

func Setvgrid(value string) GnuplotOpt {
	v := vgrid{"vgrid", value}
	return v
}

type view struct {
	option string
	value  string
}

func (v view) createCmdString() string {
	return fmt.Sprintf(`set view "%s"`, v.value)
}

func (v view) getOption() string {
	return v.option
}

func Setview(value string) GnuplotOpt {
	v := view{"view", value}
	return v
}

type vrange struct {
	option string
	value  string
}

func (v vrange) createCmdString() string {
	return fmt.Sprintf(`set vrange "%s"`, v.value)
}

func (v vrange) getOption() string {
	return v.option
}

func Setvrange(value string) GnuplotOpt {
	v := vrange{"vrange", value}
	return v
}

type vxrange struct {
	option string
	value  string
}

func (v vxrange) createCmdString() string {
	return fmt.Sprintf(`set vxrange "%s"`, v.value)
}

func (v vxrange) getOption() string {
	return v.option
}

func Setvxrange(value string) GnuplotOpt {
	v := vxrange{"vxrange", value}
	return v
}

type vyrange struct {
	option string
	value  string
}

func (v vyrange) createCmdString() string {
	return fmt.Sprintf(`set vyrange "%s"`, v.value)
}

func (v vyrange) getOption() string {
	return v.option
}

func Setvyrange(value string) GnuplotOpt {
	v := vyrange{"vyrange", value}
	return v
}

type vzrange struct {
	option string
	value  string
}

func (v vzrange) createCmdString() string {
	return fmt.Sprintf(`set vzrange "%s"`, v.value)
}

func (v vzrange) getOption() string {
	return v.option
}

func Setvzrange(value string) GnuplotOpt {
	v := vzrange{"vzrange", value}
	return v
}

type walls struct {
	option string
	value  string
}

func (w walls) createCmdString() string {
	return fmt.Sprintf(`set walls "%s"`, w.value)
}

func (w walls) getOption() string {
	return w.option
}

func Setwalls(value string) GnuplotOpt {
	w := walls{"walls", value}
	return w
}

type x2data struct {
	option string
	value  string
}

func (x x2data) createCmdString() string {
	return fmt.Sprintf("set x2ata %s", x.value)
}

func (x x2data) getOption() string {
	return x.option
}

func Setx2ata(value string) GnuplotOpt {
	x := x2data{"x2ata", value}
	return x
}

type x2dtics struct {
	option string
	value  string
}

func (x x2dtics) createCmdString() string {
	return fmt.Sprintf("set x2dtics %s", x.value)
}

func (x x2dtics) getOption() string {
	return x.option
}

func Setx2dtics(value string) GnuplotOpt {
	x := x2dtics{"x2dtics", value}
	return x
}

type x2label struct {
	option string
	value  string
}

func (x x2label) createCmdString() string {
	return fmt.Sprintf("set x2label %s", x.value)
}

func (x x2label) getOption() string {
	return x.option
}

func Setx2label(value string) GnuplotOpt {
	x := x2label{"x2label", value}
	return x
}

type x2mtics struct {
	option string
	value  string
}

func (x x2mtics) createCmdString() string {
	return fmt.Sprintf("set x2mtics %s", x.value)
}

func (x x2mtics) getOption() string {
	return x.option
}

func Setx2mtics(value string) GnuplotOpt {
	x := x2mtics{"x2mtics", value}
	return x
}

type x2range struct {
	option string
	value  string
}

func (x x2range) createCmdString() string {
	return fmt.Sprintf("set x2range %s", x.value)
}

func (x x2range) getOption() string {
	return x.option
}

func Setx2range(value string) GnuplotOpt {
	x := x2range{"x2range", value}
	return x
}

type x2tics struct {
	option string
	value  string
}

func (x x2tics) createCmdString() string {
	return fmt.Sprintf("set x2tics %s", x.value)
}

func (x x2tics) getOption() string {
	return x.option
}

func Setx2tics(value string) GnuplotOpt {
	x := x2tics{"x2tics", value}
	return x
}

type x2zeroaxis struct {
	option string
	value  string
}

func (x x2zeroaxis) createCmdString() string {
	return fmt.Sprintf("set x2zeroaxis %s", x.value)
}

func (x x2zeroaxis) getOption() string {
	return x.option
}

func Setx2zeroaxis(value string) GnuplotOpt {
	x := x2zeroaxis{"x2zeroaxis", value}
	return x
}

type xdata struct {
	option string
	value  string
}

func (x xdata) createCmdString() string {
	return fmt.Sprintf("set xdata %s", x.value)
}

func (x xdata) getOption() string {
	return x.option
}

func Setxdata(value string) GnuplotOpt {
	x := xdata{"xdata", value}
	return x
}

type xdtics struct {
	option string
	value  string
}

func (x xdtics) createCmdString() string {
	return fmt.Sprintf("set xdtics %s", x.value)
}

func (x xdtics) getOption() string {
	return x.option
}

func Setxdtics(value string) GnuplotOpt {
	x := xdtics{"xdtics", value}
	return x
}

type xlabel struct {
	option string
	value  string
}

func (x xlabel) createCmdString() string {
	return fmt.Sprintf("set xlabel %s", x.value)
}

func (x xlabel) getOption() string {
	return x.option
}

func Setxlabel(value string) GnuplotOpt {
	x := xlabel{"xlabel", value}
	return x
}

type xmtics struct {
	option string
	value  string
}

func (x xmtics) createCmdString() string {
	return fmt.Sprintf("set xmtics %s", x.value)
}

func (x xmtics) getOption() string {
	return x.option
}

func Setxmtics(value string) GnuplotOpt {
	x := xmtics{"xmtics", value}
	return x
}

type xrange struct {
	option string
	value  string
}

func (x xrange) createCmdString() string {
	return fmt.Sprintf("set xrange %s", x.value)
}

func (x xrange) getOption() string {
	return x.option
}

func Setxrange(value string) GnuplotOpt {
	x := xrange{"xrange", value}
	return x
}

type xtics struct {
	option string
	value  string
}

func (x xtics) createCmdString() string {
	return fmt.Sprintf("set xtics %s", x.value)
}

func (x xtics) getOption() string {
	return x.option
}

func Setxtics(value string) GnuplotOpt {
	x := xtics{"xtics", value}
	return x
}

type xyplane struct {
	option string
	value  string
}

func (x xyplane) createCmdString() string {
	return fmt.Sprintf("set xyplane %s", x.value)
}

func (x xyplane) getOption() string {
	return x.option
}

func Setxyplane(value string) GnuplotOpt {
	x := xyplane{"xyplane", value}
	return x
}

type xzeroaxis struct {
	option string
	value  string
}

func (x xzeroaxis) createCmdString() string {
	return fmt.Sprintf("set xzeroaxis %s", x.value)
}

func (x xzeroaxis) getOption() string {
	return x.option
}

func Setxzeroaxis(value string) GnuplotOpt {
	x := xzeroaxis{"xzeroaxis", value}
	return x
}

type y2data struct {
	option string
	value  string
}

func (y y2data) createCmdString() string {
	return fmt.Sprintf("set y2data%s", " "+y.value)
}

func (y y2data) getOption() string {
	return y.option
}

func Sety2data(value string) GnuplotOpt {
	y := y2data{"y2data", value}
	return y
}

type y2dtics struct {
	option string
	value  string
}

func (y y2dtics) createCmdString() string {
	return fmt.Sprintf("set y2dtics%s", " "+y.value)
}

func (y y2dtics) getOption() string {
	return y.option
}

func Sety2dtics(value string) GnuplotOpt {
	y := y2dtics{"y2dtics", value}
	return y
}

type y2label struct {
	option string
	value  string
}

func (y y2label) createCmdString() string {
	return fmt.Sprintf("set y2label%s", " "+y.value)
}

func (y y2label) getOption() string {
	return y.option
}

func Sety2label(value string) GnuplotOpt {
	y := y2label{"y2label", value}
	return y
}

type y2mtics struct {
	option string
	value  string
}

func (y y2mtics) createCmdString() string {
	return fmt.Sprintf("set y2mtics%s", " "+y.value)
}

func (y y2mtics) getOption() string {
	return y.option
}

func Sety2mtics(value string) GnuplotOpt {
	y := y2mtics{"y2mtics", value}
	return y
}

type y2range struct {
	option string
	value  string
}

func (y y2range) createCmdString() string {
	return fmt.Sprintf("set y2range%s", " "+y.value)
}

func (y y2range) getOption() string {
	return y.option
}

func Sety2range(value string) GnuplotOpt {
	y := y2range{"y2range", value}
	return y
}

type y2tics struct {
	option string
	value  string
}

func (y y2tics) createCmdString() string {
	return fmt.Sprintf("set y2tics%s", " "+y.value)
}

func (y y2tics) getOption() string {
	return y.option
}

func Sety2tics(value string) GnuplotOpt {
	y := y2tics{"y2tics", value}
	return y
}

type y2zeroaxis struct {
	option string
	value  string
}

func (y y2zeroaxis) createCmdString() string {
	return fmt.Sprintf("set y2zeroaxis%s", " "+y.value)
}

func (y y2zeroaxis) getOption() string {
	return y.option
}

func Sety2zeroaxis(value string) GnuplotOpt {
	y := y2zeroaxis{"y2zeroaxis", value}
	return y
}

type ydata struct {
	option string
	value  string
}

func (y ydata) createCmdString() string {
	return fmt.Sprintf("set ydata %s", y.value)
}

func (y ydata) getOption() string {
	return y.option
}

func Setydata(value string) GnuplotOpt {
	y := ydata{"ydata", value}
	return y
}

type ydtics struct {
	option string
	value  string
}

func (y ydtics) createCmdString() string {
	return fmt.Sprintf("set ydtics %s", y.value)
}

func (y ydtics) getOption() string {
	return y.option
}

func Setydtics(value string) GnuplotOpt {
	y := ydtics{"ydtics", value}
	return y
}

type ylabel struct {
	option string
	value  string
}

func (y ylabel) createCmdString() string {
	return fmt.Sprintf("set ylabel %s", y.value)
}

func (y ylabel) getOption() string {
	return y.option
}

func Setylabel(value string) GnuplotOpt {
	y := ylabel{"ylabel", value}
	return y
}

type ymtics struct {
	option string
	value  string
}

func (y ymtics) createCmdString() string {
	return fmt.Sprintf("set ymtics %s", y.value)
}

func (y ymtics) getOption() string {
	return y.option
}

func Setymtics(value string) GnuplotOpt {
	y := ymtics{"ymtics", value}
	return y
}

type yrange struct {
	option string
	value  string
}

func (y yrange) createCmdString() string {
	return fmt.Sprintf("set yrange %s", y.value)
}

func (y yrange) getOption() string {
	return y.option
}

func Setyrange(value string) GnuplotOpt {
	y := yrange{"yrange", value}
	return y
}

type ytics struct {
	option string
	value  string
}

func (y ytics) createCmdString() string {
	return fmt.Sprintf("set ytics %s", y.value)
}

func (y ytics) getOption() string {
	return y.option
}

func Setytics(value string) GnuplotOpt {
	y := ytics{"ytics", value}
	return y
}

type yzeroaxis struct {
	option string
	value  string
}

func (y yzeroaxis) createCmdString() string {
	return fmt.Sprintf("set yzeroaxis %s", y.value)
}

func (y yzeroaxis) getOption() string {
	return y.option
}

func Setyzeroaxis(value string) GnuplotOpt {
	y := yzeroaxis{"yzeroaxis", value}
	return y
}

type zdata struct {
	option string
	value  string
}

func (z zdata) createCmdString() string {
	return fmt.Sprintf("set zdata %s", z.value)
}

func (z zdata) getOption() string {
	return z.option
}

func Setzdata(value string) GnuplotOpt {
	z := zdata{"zdata", value}
	return z
}

type zdtics struct {
	option string
	value  string
}

func (z zdtics) createCmdString() string {
	return fmt.Sprintf("set zdtics %s", z.value)
}

func (z zdtics) getOption() string {
	return z.option
}

func Setzdtics(value string) GnuplotOpt {
	z := zdtics{"zdtics", value}
	return z
}

type zzeroaxis struct {
	option string
	value  string
}

func (z zzeroaxis) createCmdString() string {
	return fmt.Sprintf("set zzeroaxis %s", z.value)
}

func (z zzeroaxis) getOption() string {
	return z.option
}

func Setzzeroaxis(value string) GnuplotOpt {
	z := zzeroaxis{"zzeroaxis", value}
	return z
}

type cbdata struct {
	option string
	value  string
}

func (c cbdata) createCmdString() string {
	return fmt.Sprintf("set cbdata %s", c.value)
}

func (c cbdata) getOption() string {
	return c.option
}

func Setcbdata(value string) GnuplotOpt {
	c := cbdata{"cbdata", value}
	return c
}

type cbdtics struct {
	option string
	value  string
}

func (c cbdtics) createCmdString() string {
	return fmt.Sprintf("set cbdtics %s", c.value)
}

func (c cbdtics) getOption() string {
	return c.option
}

func Setcbdtics(value string) GnuplotOpt {
	c := cbdtics{"cbdtics", value}
	return c
}

type zero struct {
	option string
	value  string
}

func (z zero) createCmdString() string {
	return fmt.Sprintf("set zero %s", z.value)
}

func (z zero) getOption() string {
	return z.option
}

func Setzero(value string) GnuplotOpt {
	z := zero{"zero", value}
	return z
}

type zlabel struct {
	option string
	value  string
}

func (z zlabel) createCmdString() string {
	return fmt.Sprintf("set zlabel %s", z.value)
}

func (z zlabel) getOption() string {
	return z.option
}

func Setzlabel(value string) GnuplotOpt {
	z := zlabel{"zlabel", value}
	return z
}

type zmtics struct {
	option string
	value  string
}

func (z zmtics) createCmdString() string {
	return fmt.Sprintf("set zmtics %s", z.value)
}

func (z zmtics) getOption() string {
	return z.option
}

func Setzmtics(value string) GnuplotOpt {
	z := zmtics{"zmtics", value}
	return z
}

type zrange struct {
	option string
	value  string
}

func (z zrange) createCmdString() string {
	return fmt.Sprintf("set zrange %s", z.value)
}

func (z zrange) getOption() string {
	return z.option
}

func Setzrange(value string) GnuplotOpt {
	z := zrange{"zrange", value}
	return z
}

type ztics struct {
	option string
	value  string
}

func (z ztics) createCmdString() string {
	return fmt.Sprintf("set ztics %s", z.value)
}

func (z ztics) getOption() string {
	return z.option
}

func Setztics(value string) GnuplotOpt {
	z := ztics{"ztics", value}
	return z
}

type cblabel struct {
	option string
	value  string
}

func (c cblabel) createCmdString() string {
	return fmt.Sprintf("set cblabel %s", c.value)
}

func (c cblabel) getOption() string {
	return c.option
}

func Setcblabel(value string) GnuplotOpt {
	c := cblabel{"cblabel", value}
	return c
}

type cbmtics struct {
	option string
	value  string
}

func (c cbmtics) createCmdString() string {
	return fmt.Sprintf("set cbmtics %s", c.value)
}

func (c cbmtics) getOption() string {
	return c.option
}

func Setcbmtics(value string) GnuplotOpt {
	c := cbmtics{"cbmtics", value}
	return c
}

type cbrange struct {
	option string
	value  string
}

func (c cbrange) createCmdString() string {
	return fmt.Sprintf("set cbrange %s", c.value)
}

func (c cbrange) getOption() string {
	return c.option
}

func Setcbrange(value string) GnuplotOpt {
	c := cbrange{"cbrange", value}
	return c
}

type cbtics struct {
	option string
	value  string
}

func (c cbtics) createCmdString() string {
	return fmt.Sprintf("set cbtics %s", c.value)
}

func (c cbtics) getOption() string {
	return c.option
}

func Setcbtics(value string) GnuplotOpt {
	c := cbtics{"cbtics", value}
	return c
}

type using struct {
	option string
	value  string
}

func (u using) createCmdString() string {
	return fmt.Sprintf("using %s", u.value)
}

func (u using) getOption() string {
	return u.option
}

func Using(value string) GnuplotOpt {
	u := using{"using", value}
	return u
}

type via struct {
	option string
	value  string
}

func (v via) createCmdString() string {
	return fmt.Sprintf("via %s", v.value)
}

func (v via) getOption() string {
	return v.option
}

func Via(value string) GnuplotOpt {
	v := via{"via", value}
	return v
}

type with struct {
	option string
	value  string
}

func (w with) createCmdString() string {
	return fmt.Sprintf("with %s", w.value)
}

func (w with) getOption() string {
	return w.option
}

func With(value string) GnuplotOpt {
	w := with{"with", value}
	return w
}
