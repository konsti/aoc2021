package color

import "fmt"

var (
	Green  = color("32")
	Yellow = color("33")
	Blue   = color("34")
	Purple = color("35")
	Teal   = color("36")
)

func color(colorString string) func(...interface{}) string {
	return func(args ...interface{}) string {
		return fmt.Sprintf("\x1b["+colorString+"m%s\x1b[0m", fmt.Sprint(args...))
	}
}
