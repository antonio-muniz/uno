package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/antonio-muniz/uno/game/card/color"
)

func Printfln(format string, args ...interface{}) {
	Println(fmt.Sprintf(format, args...))
}

func Printlns(lines []string) {
	Println(strings.Join(lines, "\n"))
}

func Println(args ...interface{}) {
	fmt.Fprintln(color.Stdout, args...)
	time.Sleep(2 * time.Second)
}
