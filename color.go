package color

import (
	"fmt"
	"io"
)

const (
	Black   = 100
	Red     = 110
	Green   = 120
	Yellow  = 130
	Blue    = 140
	Magenta = 150
	Cyan    = 160
	White   = 170
	Default = 190
)

const (
	Reset     = 9 // produces 0
	Bright    = 1
	Italic    = 3
	Underline = 4
	Blink     = 5
	Inverse   = 7
	Hide      = 8
)

const (
	clear = "\033[0m"
)

func code(flag int) string {
	color, effect := flag / 10 + 20, flag % 10
	if effect > 0 {
		if effect == Reset {
			effect = 0
		}
		return fmt.Sprintf("\033[%d;%dm", color, effect)
	}
	return fmt.Sprintf("\033[%dm", color)
}

func Sprint(flag int, a ...interface{}) string {
	return code(flag) + fmt.Sprint(a) + clear
}

func Sprintf(flag int, format string, a ...interface{}) string {
	return code(flag) + fmt.Sprintf(format, a...) + clear
}

func Sprintln(flag int, a ...interface{}) string {
	return code(flag) + fmt.Sprintln(a...) + clear
}

func Print(flag int, a ...interface{}) (n int, err error) {
	fmt.Print(code(flag))
	n, err = fmt.Print(a...)
	fmt.Print(clear)
	return
}

func Printf(flag int, format string, a ...interface{}) (int, error) {
	return fmt.Printf(code(flag) + format + clear, a...)
}

func Println(flag int, a ...interface{}) (n int, err error) {
	fmt.Print(code(flag))
	n, err = fmt.Println(a...)
	fmt.Print(clear)
	return
}

func Fprint(flag int, w io.Writer, a ...interface{}) (n int, err error) {
	fmt.Fprint(w, code(flag))
	n, err = fmt.Fprint(w, a...)
	fmt.Fprint(w, clear)
	return
}

func Fprintf(flag int, w io.Writer, format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(w, code(flag) + format + clear, a...)
}

func Fprintln(flag int, w io.Writer, a ...interface{}) (n int, err error) {
	fmt.Fprint(w, code(flag))
	n, err = fmt.Fprintln(w, a...)
	fmt.Fprint(w, clear)
	return
}
