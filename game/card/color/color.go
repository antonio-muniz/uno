package color

import (
	"io"

	"github.com/fatih/color"
)

type Color interface {
	Paint(string) string
	Paintf(string, ...interface{}) string
	String() string
}

type colorStruct struct {
	name          string
	colorFunction func(string, ...interface{}) string
}

func (c *colorStruct) Paint(text string) string {
	return c.colorFunction(text)
}

func (c *colorStruct) Paintf(text string, args ...interface{}) string {
	return c.colorFunction(text, args...)
}

func (c *colorStruct) String() string {
	return c.Paint(c.name)
}

var Red = &colorStruct{
	name:          "red",
	colorFunction: color.New(color.FgHiRed).SprintfFunc(),
}

var Yellow = &colorStruct{
	name:          "yellow",
	colorFunction: color.New(color.FgHiYellow).SprintfFunc(),
}

var Green = &colorStruct{
	name:          "green",
	colorFunction: color.New(color.FgHiGreen).SprintfFunc(),
}

var Blue = &colorStruct{
	name:          "blue",
	colorFunction: color.New(color.FgHiCyan).SprintfFunc(),
}

var Black = &colorStruct{
	name:          "black",
	colorFunction: color.New(color.FgHiWhite).SprintfFunc(),
}

var Stdout io.Writer = color.Output

var colors = map[string]Color{
	Red.name:    Red,
	Yellow.name: Yellow,
	Green.name:  Green,
	Blue.name:   Blue,
	Black.name:  Black,
}

func ByName(name string) Color {
	return colors[name]
}
