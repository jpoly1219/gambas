package gambas

import "fmt"

type GnuplotOpt interface {
	createCmdString() string
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
