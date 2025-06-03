package console

import (
	"fmt"
	"io"
	"iter"
	"path/filepath"
	"runtime"
	"text/template"
)

type TerminalColor struct {
	name  string
	value string
}

func (c TerminalColor) Value() string {
	return c.value
}

func (c TerminalColor) Name() string {
	return c.name
}

func (c TerminalColor) String() string {
	return Colorize(c.Value(), c.Name())
}

var RED = TerminalColor{name: "Red", value: "31"}
var GREEN = TerminalColor{name: "Green", value: "32"}
var YELLOW = TerminalColor{name: "Yellow", value: "33"}
var BLUE = TerminalColor{name: "Blue", value: "34"}
var MAGENTA = TerminalColor{name: "Magenta", value: "35"}
var WHITE = TerminalColor{name: "White", value: "37"}

func RenderTemplate(name string, data any, writer io.Writer) {
	funcs := template.FuncMap{
		"colorize": Colorize,
		"colorizeBold": func(colorValue string, text string) string {
			return fmt.Sprintf("\033[1;%vm%s\033[0m", colorValue, text)
		},
		"dec":    func(i int) int { return i - 1 },
		"inc":    func(i int) int { return i + 1 },
		"mult":   func(a int, b int) int { return a * b },
		"repeat": repeat,
	}
	temp := template.Must(template.New(name).Funcs(funcs).ParseGlob(currentDir() + "/templates/*.tmpl"))

	io.WriteString(writer, "\033c")
	err := temp.ExecuteTemplate(writer, name, data)
	if err != nil {
		panic(err)
	}
}

func colorize(colorValue string, text string, bold bool) string {
	boldCode := ""
	if bold {
		boldCode = "1;"
	}
	return fmt.Sprintf("\033[%s%vm%s\033[0m", boldCode, colorValue, text)
}

func Colorize(colorValue string, text string) string {
	return colorize(colorValue, text, false)
}

func ColorizeBold(colorValue string, text string) string {
	return colorize(colorValue, text, true)
}

func currentDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error determining template directory")
	}
	return filepath.Dir(filename)
}

func repeat(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		list := make([]int, n)
		for i := range list {
			yield(i)
		}
	}
}
